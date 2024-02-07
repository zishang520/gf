---
title: 'ORM使用配置'
sidebar_position: 0
---

## 配置文件

我们推荐使用配置组件来管理数据库配置，并使用 `g` 对象管理模块中的 `g.DB("数据库分组名称")` 方法获取数据库操作对象，数据库对象将会自动读取配置组件中的相应配置项，并自动初始化该数据库操作的单例对象。数据库配置管理功能使用的是配置管理组件实现（配置组件采用接口化设计默认使用文件系统实现），同样支持多种数据格式如： `toml/yaml/json/xml/ini/properties`。默认并且推荐的配置文件数据格式为 `yaml`。

### 简单配置

从 `v2.2.0` 版本开始，使用 `link` 进行数据库配置时，数据库组件统一了不同数据库类型的配置格式，以简化配置管理。

简化配置通过配置项 `link` 指定，格式如下：

```
type:username:password@protocol(address)[/dbname][?param1=value1&...&paramN=valueN]
```

即：

```
类型:账号:密码@协议(地址)/数据库名称?特性配置
```

其中：

- **数据库名称** 及 **特性配置** 为非必须参数，其他参数为必须参数。
- **协议** 可选配置为： `tcp/udp/file`，常见配置为 `tcp`
- **特性配置** 根据不同的数据库类型，由其底层实现的第三方驱动定义，具体需要参考第三方驱动官网。例如，针对 `mysql` 驱动而言，使用的第三方驱动为： [https://github.com/go-sql-driver/mysql](https://github.com/go-sql-driver/mysql) 支持的特性配置如 `multiStatements` 和 `loc` 等。

示例：

```
database:
  default:
    link:  "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
  user:
    link:  "sqlite::@file(/var/data/db.sqlite3)"
```

不同数据类型对应的 `link` 示例如下:

| 类型 | link示例 | extra参数 |
| --- | --- | --- |
| `mysql` | `mysql:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true` | [mysql](https://github.com/go-sql-driver/mysql) |
| `mariadb` | `mariadb:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true` | [mysql](https://github.com/go-sql-driver/mysql) |
| `tidb` | `tidb:root:12345678@tcp(127.0.0.1:3306)/test?loc=Local&parseTime=true` | [mysql](https://github.com/go-sql-driver/mysql) |
| `pgsql` | `pgsql:root:12345678@tcp(127.0.0.1:5432)/test` | [pq](https://github.com/lib/pq) |
| `mssql` | `mssql:root:12345678@tcp(127.0.0.1:1433)/test?encrypt=disable` | [go-mssqldb](https://github.com/denisenkom/go-mssqldb) |
| `sqlite` | `sqlite::@file(/var/data/db.sqlite3)  (可以使用相对路径，如: db.sqlite3)` | [go-sqlite3](https://github.com/mattn/go-sqlite3) |
| `oracle` | `oracle:root:12345678@tcp(127.0.0.1:5432)/test` | [go-oci8](https://github.com/mattn/go-oci8) |
| `clickhouse` | `clickhouse:root:12345678@tcp(127.0.0.1:9000)/test` | [clickhouse-go](https://github.com/ClickHouse/clickhouse-go) |
| `dm` | `dm:root:12345678@tcp(127.0.0.1:5236)/test` | [dm](https://gitee.com/chunanyong/dm) |

更多框架支持的数据库类型请参考： [https://github.com/gogf/gf/tree/master/contrib/drivers](https://github.com/gogf/gf/tree/master/contrib/drivers)

### 完整配置

完整的 `config.yaml` 数据库配置项的数据格式形如下：

```
database:
  分组名称:
    host:                  "地址"
    port:                  "端口"
    user:                  "账号"
    pass:                  "密码"
    name:                  "数据库名称"
    type:                  "数据库类型(如：mariadb/tidb/mysql/pgsql/mssql/sqlite/oracle/clickhouse/dm)"
    link:                  "(可选)自定义数据库链接信息，当该字段被设置值时，以上链接字段(Host,Port,User,Pass,Name)将失效，但是type必须有值"
    extra:                 "(可选)不同数据库的额外特性配置，由底层数据库driver定义"
    role:                  "(可选)数据库主从角色(master/slave)，不使用应用层的主从机制请均设置为master"
    debug:                 "(可选)开启调试模式"
    prefix:                "(可选)表名前缀"
    dryRun:                "(可选)ORM空跑(只读不写)"
    charset:               "(可选)数据库编码(如: utf8/gbk/gb2312)，一般设置为utf8"
	protocol:              "(可选)数据库连接协议，默认为TCP"
    weight:                "(可选)负载均衡权重，用于负载均衡控制，不使用应用层的负载均衡机制请置空"
    timezone:              "(可选)时区配置，例如:local"
    namespace:             "(可选)用以支持个别数据库服务Catalog&Schema区分的问题，原有的Schema代表数据库名称，而NameSpace代表个别数据库服务的Schema"
    maxIdle:               "(可选)连接池最大闲置的连接数(默认10)"
    maxOpen:               "(可选)连接池最大打开的连接数(默认无限制)"
    maxLifetime:           "(可选)连接对象可重复使用的时间长度(默认30秒)"
	queryTimeout:          "(可选)查询语句超时时长(默认无限制，注意ctx的超时时间设置)"
	execTimeout:           "(可选)写入语句超时时长(默认无限制，注意ctx的超时时间设置)"
	tranTimeout:           "(可选)事务处理超时时长(默认无限制，注意ctx的超时时间设置)"
	prepareTimeout:        "(可选)预准备SQL语句执行超时时长(默认无限制，注意ctx的超时时间设置)""
    createdAt:             "(可选)自动创建时间字段名称"
    updatedAt:             "(可选)自动更新时间字段名称"
    deletedAt:             "(可选)软删除时间字段名称"
    timeMaintainDisabled:  "(可选)是否完全关闭时间更新特性，true时CreatedAt/UpdatedAt/DeletedAt都将失效"
```

完整的数据库配置项示例( `YAML`)：

```
database:
  default:
    host:          "127.0.0.1"
    port:          "3306"
    user:          "root"
    pass:          "12345678"
    name:          "test"
    type:          "mysql"
    extra:         "local=Local&parseTime=true"
    role:          "master"
    debug:         "true"
    dryrun:        0
    weight:        "100"
    prefix:        "gf_"
    charset:       "utf8"
    timezone:      "local"
    maxIdle:       "10"
    maxOpen:       "100"
    maxLifetime:   "30s"
 	protocol
```

使用该配置方式时， **为保证数据库安全，默认底层不支持多行 `SQL` 语句执行**。为了得到更多配置项控制，请参考推荐的简化配置，同时建议您务必了解清楚简化配置项中每个连接参数的功能作用。

### 集群模式

`gdb` 的配置支持集群模式，数据库配置中每一项分组配置均可以是多个节点，支持负载均衡权重策略，例如：

```
database:
  default:
  - link: "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
    role: "master"
  - link: "mysql:root:12345678@tcp(127.0.0.1:3306)/test"
    role: "slave"

  user:
  - link: "mysql:root:12345678@tcp(127.0.0.1:3306)/user"
    role: "master"
  - link: "mysql:root:12345678@tcp(127.0.0.1:3306)/user"
    role: "slave"
  - link: "mysql:root:12345678@tcp(127.0.0.1:3306)/user"
    role: "slave"
```

以上数据库配置示例中包含两个数据库分组 `default` 和 `user`，其中 `default` 分组包含一主一从， `user` 分组包含一主两从。在代码中可以通过 `g.DB()` 和 `g.DB("user")` 获取对应的数据库连接对象。

### 日志配置

`gdb` 支持日志输出，内部使用的是 `glog.Logger` 对象实现日志管理，并且可以通过配置文件对日志对象进行配置。默认情况下 `gdb` 关闭了 `DEBUG` 日志输出，如果需要打开 `DEBUG` 信息需要将数据库的 `debug` 参数设置为 `true`。以下是为一个配置文件示例：

```
database:
  logger:
    path:    "/var/log/gf-app/sql"
    level:   "all"
    stdout:  true
  default:
    link:    "mysql:root:12345678@tcp(127.0.0.1:3306)/user_center"
    debug:   true
```

其中 `database.logger` 即为 `gdb` 的日志配置，当该配置不存在时，将会使用日志组件的默认配置，具体请参考 [日志组件-配置管理](/docs/核心组件/日志组件/日志组件-配置管理) 章节。

需要注意哦：由于 `ORM` 底层都是采用安全的预处理执行方式，提交到底层的 `SQL` 与参数其实是分开的，因此日志中记录的完整 `SQL` 仅作参考方便人工阅读，并不是真正提交到底层的 `SQL` 语句。

## 原生配置(高阶，可选)

以下为数据库底层管理配置介绍，如果您对数据库的底层配置管理比较感兴趣，可继续阅读后续章节。

### 数据结构

`gdb` 数据库管理模块的内部配置管理数据结构如下：

`ConfigNode` 用于存储一个数据库节点信息； `ConfigGroup` 用于管理多个数据库节点组成的配置分组(一般一个分组对应一个业务数据库集群)； `Config` 用于管理多个 `ConfigGroup` 配置分组。

**配置管理特点：**

1. 支持多节点数据库集群管理；
2. 每个节点可以单独配置连接属性；
3. 采用单例模式管理数据库实例化对象；
4. 支持对数据库集群分组管理，按照分组名称获取实例化的数据库操作对象；
5. 支持多种关系型数据库管理，可通过 `ConfigNode.Type` 属性进行配置；
6. 支持 `Master-Slave` 读写分离，可通过 `ConfigNode.Role` 属性进行配置；
7. 支持客户端的负载均衡管理，可通过 `ConfigNode.Weight` 属性进行配置，值越大，优先级越高；

```
type Config      map[string]ConfigGroup // 数据库配置对象
type ConfigGroup []ConfigNode           // 数据库分组配置
// 数据库配置项(一个分组配置对应多个配置项)
type ConfigNode  struct {
    Host             string        // 地址
    Port             string        // 端口
    User             string        // 账号
    Pass             string        // 密码
    Name             string        // 数据库名称
    Type             string        // 数据库类型：mysql, sqlite, mssql, pgsql, oracle
	Link             string        // (可选)自定义链接信息，当该字段被设置值时，以上链接字段(Host,Port,User,Pass,Name)将失效(该字段是一个扩展功能)
    Extra            string        // (可选)不同数据库的额外特性配置，由底层数据库driver定义
    Role             string        // (可选，默认为master)数据库的角色，用于主从操作分离，至少需要有一个master，参数值：master, slave
    Debug            bool          // (可选)开启调试模式
    Charset          string        // (可选，默认为 utf8)编码，默认为 utf8
    Prefix           string        // (可选)表名前缀
    Weight           int           // (可选)用于负载均衡的权重计算，当集群中只有一个节点时，权重没有任何意义
    MaxIdleConnCount int           // (可选)连接池最大闲置的连接数
    MaxOpenConnCount int           // (可选)连接池最大打开的连接数
    MaxConnLifetime  time.Duration // (可选，单位秒)连接对象可重复使用的时间长度
}
```

特别说明， `gdb` 的配置管理最大的 **特点** 是，（同一进程中）所有的数据库集群信息都使用同一个配置管理模块进行统一维护， **不同业务的数据库集群配置使用不同的分组名称** 进行配置和获取。

### 配置方法

这是原生调用 `gdb` 模块来配置管理数据库。如果开发者想要自行控制数据库配置管理可以参考以下方法。若无需要可忽略该章节。

接口文档： [https://pkg.go.dev/github.com/gogf/gf/v2/database/gdb](https://pkg.go.dev/github.com/gogf/gf/v2/database/gdb)

```
// 添加一个数据库节点到指定的分组中
func AddConfigNode(group string, node ConfigNode)
// 添加一个配置分组到数据库配置管理中(同名覆盖)
func AddConfigGroup(group string, nodes ConfigGroup)

// 添加一个数据库节点到默认的分组中(默认为default，可修改)
func AddDefaultConfigNode(node ConfigNode)
// 添加一个配置分组到数据库配置管理中(默认分组为default，可修改)
func AddDefaultConfigGroup(nodes ConfigGroup)

// 设置默认的分组名称，获取默认数据库对象时将会自动读取该分组配置
func SetDefaultGroup(groupName string)

// 设置数据库配置为定义的配置信息，会将原有配置覆盖
func SetConfig(c Config)
```

默认分组表示，如果获取数据库对象时不指定配置分组名称，那么 `gdb` 默认读取的配置分组。例如： `gdb.NewByGroup()` 可获取一个默认分组的数据库对象。简单的做法，我们可以通过 `gdb` 包的 `SetConfig` 配置管理方法进行自定义的数据库全局配置，例如：

```
gdb.SetConfig(gdb.Config {
    "default" : gdb.ConfigGroup {
        gdb.ConfigNode {
            Host     : "192.168.1.100",
            Port     : "3306",
            User     : "root",
            Pass     : "123456",
            Name     : "test",
            Type     : "mysql",
            Role     : "master",
            Weight   : 100,
        },
        gdb.ConfigNode {
            Host     : "192.168.1.101",
            Port     : "3306",
            User     : "root",
            Pass     : "123456",
            Name     : "test",
            Type     : "mysql",
            Role     : "slave",
            Weight   : 100,
        },
    },
    "user-center" : gdb.ConfigGroup {
        gdb.ConfigNode {
            Host     : "192.168.1.110",
            Port     : "3306",
            User     : "root",
            Pass     : "123456",
            Name     : "test",
            Type     : "mysql",
            Role     : "master",
            Weight   : 100,
        },
    },
})
```

随后，我们可以使用 `gdb.NewByGroup("数据库分组名称")` 来获取一个数据库操作对象。该对象用于后续的数据库一系列方法/链式操作。

## 常见问题

### 如何实现数据库账号密码在配置文件中加密

在某些场景下，数据库的账号密码无法明文配置到配置文件中，需要进行一定的加密。在连接数据库的时候，再对配置文件中加密的字段进行解密处理。这种需求可以通过自定义 `Driver` 来实现（关于 `Driver` 的详细介绍请参考章节： [ORM接口开发](/docs/核心组件/数据库ORM/ORM接口开发)）。以 `mysql` 为例，我们可以自己编写一个 `Driver`，包裹框架社区组件中的 `mysql driver`，并且覆盖它的 `Open` 方法即可。代码示例：

```
import (
	"database/sql"

	"github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
)

type MyBizDriver struct {
	mysql.Driver
}

// Open creates and returns an underlying sql.DB object for mysql.
// Note that it converts time.Time argument to local timezone in default.
func (d *MyBizDriver) Open(config *gdb.ConfigNode) (db *sql.DB, err error) {
	config.User = d.decode(config.User)
	config.Pass = d.decode(config.Pass)
	return d.Driver.Open(config)
}

func (d *MyBizDriver) decode(s string) string {
	// 执行字段解密处理逻辑
	// ...
	return s
}
```