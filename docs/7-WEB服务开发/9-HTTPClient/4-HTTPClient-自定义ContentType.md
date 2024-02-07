---
title: 'HTTPClient-自定义ContentType'
sidebar_position: 4
---

## 示例1，提交 `Json` 请求

```
g.Client().ContentJson().PostContent(ctx, "http://order.svc/v1/order", g.Map{
    "uid"         : 1,
    "sku_id"      : 10000,
    "amount"      : 19.99,
    "create_time" : "2020-03-26 12:00:00",
})
```

该请求将会将 `Content-Type` 设置为 `application/json`，并且将提交参数自动编码为 `Json`:

```
{"uid":1,"sku_id":10000,"amount":19.99,"create_time":"2020-03-26 12:00:00"}
```

## 示例2，提交 `Xml` 请求

```
g.Client().ContentXml().PostContent(ctx, "http://order.svc/v1/order", g.Map{
    "uid"         : 1,
    "sku_id"      : 10000,
    "amount"      : 19.99,
    "create_time" : "2020-03-26 12:00:00",
})
```

该请求将会将 `Content-Type` 设置为 `application/xml`，并且将提交参数自动编码为 `Xml`:

```
<doc><amount>19.99</amount><create_time>2020-03-26 12:00:00</create_time><sku_id>10000</sku_id><uid>1</uid></doc>
```

## 示例3，自定义 `ContentType`

我们可以通过 `ContentType` 方法自定义客户端请求的 `ContentType`。例如：

```
g.Client().ContentType("application/json").PostContent(ctx, "http://order.svc/v1/order", g.Map{
    "uid"         : 1,
    "sku_id"      : 10000,
    "amount"      : 19.99,
    "create_time" : "2020-03-26 12:00:00",
})
```