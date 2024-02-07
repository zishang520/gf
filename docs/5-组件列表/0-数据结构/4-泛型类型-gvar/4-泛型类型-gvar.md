---
title: '泛型类型-gvar'
sidebar_position: 4
---

![](/markdown/cd9ed75865d6b5efe704f58156a42fa4.png)

## 基本介绍

`gvar` 是一种 **运行时泛型** 实现，以较小的运行时开销提高开发便捷性以及研发效率，支持各种内置的数据类型转换，可以作为 `interface{}` 类型的替代数据类型，并且该类型支持并发安全开关。

框架同时提供了 `g.Var` 的数据类型，其实也是 `gvar.Var` 数据类型的别名。

**使用场景**：

使用 `interface{}` 的场景，各种不固定数据类型格式，或者需要对变量进行频繁的数据类型转换的场景。

**使用方式**：

```  go
import "github.com/gogf/gf/v2/container/gvar"
```

**接口文档**：

[https://pkg.go.dev/github.com/gogf/gf/v2/container/gvar](https://pkg.go.dev/github.com/gogf/gf/v2/container/gvar)

## 相关文档

- [泛型类型-基本使用](/docs/组件列表/数据结构/泛型类型-gvar/泛型类型-基本使用)
- [泛型类型-方法介绍](/docs/组件列表/数据结构/泛型类型-gvar/泛型类型-方法介绍)
- [泛型类型-注意事项](/docs/组件列表/数据结构/泛型类型-gvar/泛型类型-注意事项)