---
desc: Go 夜读 && TiDB 源码阅读之 Executor
title: 第 47 期 TiDB 源码阅读之 Executor
date: 2019-06-12T21:10:00+08:00
author: chenshuang
---

## TiDB Executor 内容介绍

本次分享主要讲 TiDB 中 insert/update/delete/select, 以及 DDL 等是如何执行的，以及涉及到相关模块。大概会涉及以下模块：

* executor
* distsql
* ddl

![](https://mmbiz.qpic.cn/mmbiz_png/2jnWxKdgFb8uuoI7WAicOIPWnheB8ovPRXtaF9Lyq1pj52DGCfvxg7hI6pamSc9fiaNTf3vfdoibWZRibibKmoal2xw/640?wx_fmt=png&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

PPT: [TiDB Executor 源码阅读.pdf](https://github.com/talkgo/night/files/3281080/TiDB.Executor.pdf)

## 推荐阅读

* [Select 语句概览](https://pingcap.com/blog-cn/tidb-source-code-reading-6/)
* [INSERT 语句详解](https://pingcap.com/blog-cn/tidb-source-code-reading-16/)
* [DDL 源码解析](https://pingcap.com/blog-cn/tidb-source-code-reading-17/)

## 视频回看

1. [TiDB 源码学习之 Executor - YouTube](https://youtu.be/Rcrm4w7sqbM)
2. [TiDB 源码学习之 Executor - Bilibili](https://www.bilibili.com/video/av55403428/)

PPT: https://github.com/talkgo/night/files/3281080/TiDB.Executor.pdf

## 问题

- 表的信息是怎么存的呢 
- id的生成规则是什么 
- 如果索引里面不保存handle_id，那怎么根据索引找到这行数据呢 
- 索引字段很大会不会有问题，作为id的一部分的话
- 单条6m的限制是怎么计算出来的？还是压力测出来的？
- ddl时，job放到tikv的队列，tikv是分布式的，job具体是放到哪个tikv上的呢？
- 并行ddl 如何跑
- tikv整体上可以看成一个kv store
- region这部分概念可以配合hbase去看看能更好的理解
- 难道own tidb server要遍历所有的tikv server上的queue，去取ddl的job？
- tidb 的统计信息也是放一个表里面，每次parse 都会去拿这个信息，这样的话请求到一个region,这个表是不是很容易成为热点

## 观看视频

{{< youtube id="Rcrm4w7sqbM" >}}
