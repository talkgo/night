---
desc: Go 夜读之 Paper Reading CSP 理解顺序进程间通信
title: 第 66 期 Paper Reading CSP 理解顺序进程间通信
date: 2019-11-07T21:00:00+08:00
author: 欧长坤@慕尼黑大学
---

## Go 夜读第 66 期 Paper Reading CSP 理解顺序进程间通信

本期 Go 夜读是由 Go 夜读 SIG 核心小组成员欧长坤给大家带来的经典论文 CSP 的 Paper Reading。

## CSP 是什么？

我们常常在讨论中提及 CSP，但鲜有人能真正说清楚 CSP 的演进历史，及其最核心的基本思想。我们已经对 Go 提供的并发原语足够熟悉了，是时候深入理解其背后的基础理论 —— 顺序进程间通信（Communicating Sequential Processes, CSP）了。本次分享我们针对 [Hoare 1978] 探讨 CSP 理论的原始设计（CSP 1978），主要围绕以下几个问题展开：

Tony Hoare 提出 CSP 的时代背景是什么？
- CSP 1978 理论到底有哪些值得我们研究的地方？
- CSP 1978 理论是否真的就是我们目前熟知的基于通道的同步方式？
- CSP 1978 理论的早期设计存在什么样的缺陷？

## 大纲

- CSP 1978 的诞生背景
- CSP 1978 的主要内容及其结论
- CSP 1978 理论中存在的设计缺陷
- 讨论与反思

### 分享 Slides 

- https://docs.google.com/presentation/d/1N5skL6vR9Wxk-I82AYs3dlsOsJkUAGJCsb5NGpXWpqo/edit?usp=sharing

## 回看视频

- https://www.bilibili.com/video/av74891823/
- https://youtu.be/Z8ZpWVuEx8c

## 参考资料

- [Hoare 1978] [Hoare, C. A. R. (1978). Communicating sequential processes. Communications of the ACM, 21(8), 666–677.](https://spinroot.com/courses/summer/Papers/hoare_1978.pdf)
- [Ou 2019a] [CSP1978 的 Go 语言实现](https://github.com/changkun/gobase/blob/master/csp/csp.go)
- [Ou 2019b] [第 56 期 channel & select 源码分析](https://github.com/talk-go/night/issues/450)
- [Ou 2019c] [第 59 期 Real-world Go Concurrency Bugs](https://github.com/talk-go/night/issues/464)


---

## 观看视频

{{< youtube id="Z8ZpWVuEx8c" >}}
