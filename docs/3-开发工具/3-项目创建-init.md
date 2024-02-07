---
title: '项目创建-init'
sidebar_position: 3
---

从 `v2` 版本开始，项目的创建不再依赖远端获取，仓库模板已经通过 [资源管理](/docs/核心组件/资源管理) 的方式内置到了工具二进制文件中，因此项目创建速度非常迅速。

## 使用方式

```
$ gf init -h
USAGE
    gf init ARGUMENT [OPTION]

ARGUMENT
    NAME    name for the project. It will create a folder with NAME in current directory.The NAME will also be the
            module name for the project.

OPTION
    -m, --mono    initialize a mono-repo instead a single-repo
    -h, --help    more information about this command

EXAMPLE
    gf init my-project
    gf init my-mono-repo -m
```

我们可以使用 `init` 命令在当前目录生成一个示例的 `GoFrame` 空框架项目，并可给定项目名称参数。生成的项目目录结构仅供参考，根据业务项目具体情况可自行调整。生成的目录结构请参考 [代码分层设计](/docs/框架设计/工程开发设计/代码分层设计) 章节。

`GoFrame` 框架开发推荐统一使用官方的 `go module` 特性进行依赖包管理，因此空项目根目录下也有一个 `go.mod` 文件。

工程目录采用了通用化的设计，实际项目中可以根据项目需要适当增减模板给定的目录。例如，没有 `kubernetes` 部署需求的场景，直接删除对应 `deploy` 目录即可。

## 使用示例

### 在当前目录下初始化项目

```
$ gf init .
initializing...
initialization done!
you can now run 'gf run main.go' to start your journey, enjoy!
```

### 创建一个指定名称的项目

```
$ gf init myapp
initializing...
initialization done!
you can now run 'cd myapp && gf run main.go' to start your journey, enjoy!
```

### 创建一个 `MonoRepo` 项目

默认情况下创建的是 `SingleRepo` 项目，若有需要也可以创建一个 `MonoRepo`（大仓）项目，通过使用 `-m` 选项即可。

```
$ gf init mymono -m
initializing...
initialization done!
```

关于大仓的介绍请参考章节： [微服务大仓管理模式](/docs/框架设计/工程开发设计/微服务大仓管理模式)