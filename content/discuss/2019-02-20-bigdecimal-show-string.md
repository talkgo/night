---
title: 2019-02-20 浮点数如何输出
date: 2019-02-20T00:00:00+08:00
---

来源: Wechat discuss


### 将数据序列化为json的时候，怎么让序列化后的 json 里面不要使用科学计数法
### json.Marshal 转出来的json怎么能不让转成科学计数法 大家有解决办法吗？

```golang
dec := decimal.NewFromFloat(0.000001)
fmt.Println(dec.String())
```

## 参考资料

1. [shopspring/decimal](https://github.com/shopspring/decimal)
