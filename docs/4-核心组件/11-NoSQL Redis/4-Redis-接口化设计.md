---
title: 'Redis-接口化设计'
sidebar_position: 4
---

`gredis` 采用接口化设计，具有强大的灵活性和扩展性。

## 接口定义

[https://pkg.go.dev/github.com/gogf/gf/v2/database/gredis#Adapter](https://pkg.go.dev/github.com/gogf/gf/v2/database/gredis#Adapter)

## 相关方法

```
// SetAdapter sets custom adapter for current redis client.
func (r *Redis) SetAdapter(adapter Adapter)

// GetAdapter returns the adapter that is set in current redis client.
func (r *Redis) GetAdapter() Adapter
```

## 更进一步

由于 `gredis` 组件的接口实现是高阶功能，一般来说开发者也无需替换 `Redis` 接口实现。感兴趣的朋友可以自行研究。