---
desc: Go 夜读之 Go database/sql 数据库连接池分析
title: 第 67 期 Go database/sql 数据库连接池分析
date: 2019-11-14T21:00:00+08:00
author: 邹文通@POP
---

## Go 夜读第 67 期 Go database/sql 数据库连接池分析

本期 Go 夜读是由 POP 后端团队的邹文通给大家带来的 Go 标准包 database/sql 数据库连接池源码剖析。

### 大纲

- sql 连接池简介
- 连接池的工作原理
- sql 包连接池源码分析
- 连接池使用 tips

## Slides

- https://docs.google.com/presentation/d/10kGjeHGbB0h0Cz8f58reXOyCdyWSOSKrr2160IFNla4/edit?usp=sharing

## 回看视频

- https://www.bilibili.com/video/av75690189/
- https://youtu.be/JKJ8ehtiqUM

----

## QA

### 1. database/sql 中 MaxIdleConns 和 MaxOpenConns 应该怎么设置才是相对合理的，在选择设置具体的值时，他们又受什么因素影响呢？

- 关于这个问题，可以参考这篇文章 [Production-ready Database Connection Pooling in Go](https://making.pusher.com/production-ready-connection-pooling-in-go/). 文章的建议是 MaxOpenConns 应该和实际的打开的连接数的监测值相关。然后按照 MaxOpenConns 的一定比值设置 MaxIdleConns，比方说 50%，这个值取决于你对业务的预估。每维持一个闲散连接，会造成 1MB 左右的客户端内存开销和 2MB 左右的数据库内存开销，CPU 开销相对小一点。文章还给出了一些 benchmark 的测试，在默认 MaxIdleConns 和  MaxIdleConns = 50% * MaxOpenConns 情况下的一个性能的对比，可以参考一下。

## 参考资料

- [Go组件学习——database/sql数据库连接池你用对了吗](https://juejin.im/post/5d624abde51d45621655352c)
- [Go组件学习——手写连接池并没有那么简单](https://mp.weixin.qq.com/s/-2T9BovG8TG32DQKn93LaA)
- [Chapter 8 Connection Pooling with Connector/J](https://dev.mysql.com/doc/connector-j/8.0/en/connector-j-usagenotes-j2ee-concepts-connection-pooling.html)
- [彻底弄懂mysql（二）--连接方式](https://blog.csdn.net/LYue123/article/details/89285157)


---

## 观看视频

{{< youtube id="JKJ8ehtiqUM" >}}
