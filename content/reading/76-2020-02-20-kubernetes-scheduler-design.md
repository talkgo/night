---
desc: Go 夜读之 Kubernetes Scheduler 设计与实现
title: 第 76 期 Kubernetes Scheduler 设计与实现
date: 2020-02-20T21:00:00+08:00
author: Draven
---

## 【Go 夜读】#76 Kubernetes Scheduler 设计与实现

谈谈 Kubernetes 的架构设计以及 Kubernetes Scheduler 怎么把你的 Pod 调度到某个 Node 上的。

## 大纲

+ Kubernetes 架构设计
+ Kubernetes 调度器历史与设计演变
    + 基于谓词与优先级的调度器
    + 基于调度框架的调度器
+ Kubernetes 调度器实现分析
+ 番外篇（Optional）
    + 反调度器
    + 批处理调度器

## 分享者自我介绍

Draven，[面向信仰编程](https://draveness.me/) 作者，Kubernetes 搬砖工，~负责给 issue 贴 Label~。

## 分享时间

2020-02-20 21:00:00 UTC+8 (真是个好日子)

## Slides

https://docs.google.com/presentation/d/1zVftY8VhOfTqGYvQogFMUvB8WLTCIPJXUmeZnDOGzNE/edit?usp=sharing

## 参考资料

+ [调度系统设计精要](https://draveness.me/system-design-scheduler)
+ [Go 语言调度器的实现原理](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-goroutine/)
+ [谈 Kubernetes 的架构设计与实现原理](https://draveness.me/understanding-kubernetes)

----

## 观看视频

{{< youtube id="1cQt2bXJtME" >}}
