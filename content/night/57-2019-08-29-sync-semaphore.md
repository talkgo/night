---
desc: Go 夜读之 sync/semaphore 源码浅析
title: 第 57 期 sync/semaphore 源码浅析
date: 2019-08-29T21:00:00+08:00
author: Felix
---

## Go 夜读第 57 期 sync/semaphore 源码浅析

内容简介

主要分析 golang.org/x/sync/semaphore 相关代码和 semaphore 部分使用场景。

内容大纲
- semaphore 定义
- 源码分析
- Q&A

## 分享地址

2019.08.29, 21:00 ~ 21:40, UTC+8

https://zoom.us/j/6923842137

## 进一步阅读的材料

- [semaphore 定义](https://en.wikipedia.org/wiki/Semaphore_(programming))
- [源码](https://github.com/golang/sync/blob/master/semaphore/semaphore.go)
- [分享 PPT](https://docs.google.com/presentation/d/17Moou4_Z5kD9xuvCIFUT4d7KbkyS73DCQmdOPlJ5P2U/edit?usp=sharing)

## 补充资料
- [同步原语](https://draveness.me/golang/concurrency/golang-sync-primitives.html)
- [结合 errgroup 使用](https://github.com/golang/go/issues/27837#issuecomment-513443404)
- [关于是否应该支持 resize 的讨论](https://github.com/golang/go/issues/29721)
- [semaphore 实现的 taskpool](https://github.com/eleniums/async/blob/master/pool.go)

更多见：https://github.com/talkgo/night/issues/456


---


## 观看视频

{{< youtube id="VEtdLBFY_y4" >}}
