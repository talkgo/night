---
desc: Go 夜读之 channel & select 源码分析
title: 第 56 期 channel & select 源码分析
date: 2019-08-22T21:00:00+08:00
author: 欧长坤
---

## Go 夜读第 56 期 channel & select 源码分析

内容简介

Go 语言除了提供传统的互斥量、同步组等同步原语之外，还受 CSP 理论的影响，提供了 channel 这一强有力的同步原语。本次分享将讨论 channel 及其相关的 select 语句的源码，并简要讨论 Go 语言同步原语中的性能差异与反思。

内容大纲
- 同步原语概述
- channel/select 回顾
- channel 的结构设计
- channel 的初始化行为
- channel 的发送与接收过程及其性能优化
- channel 的回收
- select 的本质及其相关编译器优化

## 分享地址
2019.08.22, 20:30 ~ 21:30, UTC+8

https://zoom.us/j/6923842137

## 进一步阅读的材料

[Ou, 2019] channel & select 源码分析
分享内容的 PPT
[Ou, 2018] Go 源码研究
分享者写的一本 Go 源码分析
[Mullender and Cox, 2008] S. Mullender and R. Cox, Semaphores in Plan 9
影响 Go 语言信号量设计的一篇文章
[Drepper, 2003] U. Drepper, Futexes are Tricky
第一篇正确实现 Linux Futex 机制的文章
[Vyukov, 2014a] D. Vyukov, Go channels on steroids, January 2014
无锁式 channel 的设计提案
[Vyukov, 2014b] D. Vyukov, runtime: lock-free channels, October 2014
关于无锁式 channel 的讨论、早期实现与现状
[Hoare, 2015] C. A. R. Hoare, Communicating Sequential Processes. May 18, 2015
有关 CSP 理论的一切，与早期 1978 年的版本相比更为完善和严谨
[Creager, 2016] D. Creager, An oversimplified history of CSP, 2016
CSP 理论的极简史

更多见：https://github.com/developer-learning/reading-go/issues/450

## 观看视频

{{< youtube id="O_FJgYKOBYQ" >}}
