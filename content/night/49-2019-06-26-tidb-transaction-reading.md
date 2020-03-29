---
desc: Go 夜读 && TiDB 源码阅读之 Transaction
title: 第 49 期 TiDB 源码阅读之 Transaction
date: 2019-06-26T21:10:00+08:00
author: zimulala
---

## TiDB Transaction 内容介绍

本次分享主要讲 TiDB 的事务执行过程和一些异常处理，涉及 TiDB 的 session 和 tikv 部分模块。

PDF: [Source code reading of TiDB Transaction .pdf](https://github.com/talk-go/night/files/3329306/Source.code.reading.of.TiDB.Transaction.pdf)

## 推荐阅读

* [TiDB 源码阅读系列文章（十九）tikv-client（下）](https://pingcap.com/blog-cn/tidb-source-code-reading-19/)
* [三篇文章了解 TiDB 技术内幕 - 说存储](https://pingcap.com/blog-cn/tidb-internal-1/)
* [Transaction in TiDB](https://andremouche.github.io/tidb/transaction_in_tidb.html)
* [Coprocessor in TiKV](https://andremouche.github.io/tidb/coprocessor_in_tikv.html)


## 视频回看

1. [TiDB 源码学习之 Executor - YouTube](https://youtu.be/A46VE3aUTKo)
2. [TiDB 源码学习之 Executor - Bilibili](https://www.bilibili.com/video/av56945776/)

## 问题

21:17:34	 From zq : 分享妹子用的是什么IDE
21:17:44	 From mrj :   goland
21:17:45	 From Pure White : 左上角，goland
21:17:45	 From tangyinpeng : goland
21:17:46	 From Heng Long : goland
21:17:57	 From zq : goland现在做得这么好看啦
21:18:05	 From Heng Long : Meterial theme
21:18:07	 From mrj : 下来主题
21:18:22	 From lk : 有什么比较不错的主题吗？
21:18:31	 From Pure White : darcula
21:18:35	 From mrj : 默认的就挺好的
21:20:04	 From mai yang : 明天晚上将由 GoLand 布道师给我们分享 GoLand 的使用及技巧实践分享。
21:28:23	 From HAITAO的 iPhone : 点查不带timestamp，直接读最新稳定版本么？
21:28:32	 From Wei Yao : 对
21:28:52	 From liber xue : 双击shift 直接search
21:28:55	 From Wei Yao : 最新 commited 版本
21:35:50	 From HAITAO的 iPhone : 点查，实际会默认给一个当前最新的timestamp,根据这个ts，kv返回对应的版本值?还是不带任何ts，发给kv ？
21:36:29	 From Wei Yao : 用 maxTs
21:50:04	 From openinx : A very nice talk.
22:05:22	 From kzl : 获取完成之后，region扩容了，数据迁移走了怎么办？
22:06:16	 From jeff : 是说 region 分裂了吧。
22:06:33	 From kzl : 对的
22:08:46	 From ruiayLin : region信息就会过期
22:11:05	 From jeff : 那提交的时候会重试吧
22:11:39	 From hezhiyong : tidb不断缓存region 的信息会不会占用很大的内存
22:13:23	 From jeff : 唔，这里应该只缓冲曾经用到的 region ，并不是集群中所有 region
22:13:52	 From jeff : s/ 缓冲 / 缓存 /g
22:14:44	 From jeff : 貌似讲到刚才数据 region 分裂后的场景了。
22:14:56	 From Wei Yao : 会重试
22:26:54	 From fj : 大神 tidb的事物隔离级别 能介绍下吗？- ̗̀(๑ᵔ⌔ᵔ๑)
22:29:35	 From Tengjin Xie : snapshot isolation?
22:31:48	 From Wei Yao : 比 mysql 的 rr 稍微高一点
22:33:41	 From fj : 刚才 讲的tidb的隔离级别是？
22:34:03	 From Wei Yao : 你可以认为是 可重复读
22:34:11	 From Wei Yao : 其实这是快照隔离级别

## 观看视频

{{< youtube id="A46VE3aUTKo" >}}
