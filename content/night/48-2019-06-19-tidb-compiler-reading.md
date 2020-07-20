---
desc: Go 夜读 && TiDB 源码阅读之 Compiler
title: 第 48 期 TiDB 源码阅读之 Compiler
date: 2019-06-19T21:10:00+08:00
author: wangcong
---

## TiDB Compiler 内容介绍

本次分享主要讲 TiDB 的优化器框架以及具体的 SQL 执行优化原理 。主要涉及 TiDB 的 planner 模块。欢迎大家参加！

PPT: [TiDB Compiler.pdf](https://github.com/talkgo/night/files/3305279/TiDB.Compiler.pdf)


## 推荐阅读

* [TiDB 源码阅读系列文章（七）基于规则的优化](https://pingcap.com/blog-cn/tidb-source-code-reading-7)
* [TiDB 源码阅读系列文章（八）基于代价的优化](https://pingcap.com/blog-cn/tidb-source-code-reading-8/)
* [TiDB 源码阅读系列文章（二十一）基于规则的优化 II](https://pingcap.com/blog-cn/tidb-source-code-reading-21/)

## 视频回看

1. [TiDB 源码学习之 Executor - YouTube](https://youtu.be/4mgx8bq_fcQ)
2. [TiDB 源码学习之 Executor - Bilibili](https://www.bilibili.com/video/av56138440/)

## 问题

22:13:46	 From mai yang : rule 怎么对照文档帮助理解呢？
22:16:31	 From dqyuan : 怎么快速找到代码对应的pr？
22:18:43	 From Wei Yao : git blame
22:23:04	 From kzl : Order by 是会下推到tikv吗？
22:25:02	 From Wei Yao : 除非 order by 带了 limit，要不然推下去没意义
22:25:25	 From Wei Yao : 有一些情况，如果是 order by 一个索引，那就直接消除掉这个 排序操作了
22:26:07	 From zhao : 有意义吧，推了之后 tidb端可以直接stream merge，不知道实现了没有
22:27:38	 From Wei Yao : 是可以 stream merge, 但是现在 tidb 还没实现这个，因为优先级不是太高
22:30:42	 From Wei Yao : stream merge 主要是可以节省一些内存，避免 order by 太多导致 tidb oom
22:30:56	 From Heng Long : 嗯，会让 tikv 的压力变大
22:34:28	 From hezhiyong : limit  offset   分页性能不好
22:34:53	 From hezhiyong : 有好变通改写方法没
22:35:52	 From hezhiyong : limit offset 会有下推到tikv么
22:36:19	 From Wei Yao : limit offset 没办法的，这个是全局的 offset
22:36:40	 From Wei Yao : tikv 并不知道自己的 offset 在全局的 offset 是多少
22:37:05	 From Wei Yao : 这个其他数据库其实也一样
22:37:37	 From hezhiyong : 那就是这个数据就是要全拿到tidb层在来过滤
22:37:38	 From hezhiyong : 是吧
22:39:27	 From Hao’s iPad : 喝口水吧
22:45:35	 From zhao : 这个skyline prune有相关的资料吗
22:45:41	 From zhao : paper之类的
22:47:09	 From Wei Yao : 我记得暂时还没有 public
22:47:20	 From Wei Yao : skyline pruning 就是消除一些路径
23:04:22	 From mai yang : 怎么快速找到代码对应的pr？git blame 这个可以演示一下吗？
23:06:42	 From tangenta : github 上面看文件的时候有个选项是 blame，那里应该比较清晰
23:09:44	 From mai yang : github 上面看文件的时候有个选项是 blame，那里应该比较清晰 ——  这个不错，看到了。

## 观看视频

{{< youtube id="4mgx8bq_fcQ" >}}
