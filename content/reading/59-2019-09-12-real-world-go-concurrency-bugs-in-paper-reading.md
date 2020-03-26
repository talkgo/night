---
desc: Go 夜读之 Real-world Go Concurrency Bugs
title: 第 59 期 Real-world Go Concurrency Bugs
date: 2019-09-12T21:00:00+08:00
author: Go 夜读 SIG 小组
---

## Go 夜读第 59 期 Real-world Go Concurrency Bugs

### 内容简介

Go 语言鼓励其用户多使用基于消息传递的同步原语 channel，但也不排斥其用户使用基于内存共享的同步原语，提供了诸如 sync.Mutex 等互斥原语。在过去十年的时间里，Go 的实践者不断思考着这些问题：哪种同步原语更加优秀？究竟什么场景下应该使用何种同步原语？哪类同步原语能够更好的保证数据操纵的正确性？哪类同步原语对程序员的心智负担较大？何种同步原语更容易产生程序 Bug？channel 是一种反模式吗？什么类型的 Bug 能够更好的被 Bug 检测器发现？……

[Tu et al., 2019] 调研了包括 Docker, Kubernetes, gRPC 等六款主流大型 Go 程序在演进过程中出现的 171 个与同步原语相关的 Bug，并给出了一些有趣的见解。本次分享将讨论 [Tu et al. 2019] 的研究论文。

### 内容大纲

- Go 常见的并发模式与论文的研究背景
- 论文的研究方法
- Go 并发 Bug 的分类及部分主要结论
	- 阻塞式 Bug
	- 非阻塞式 Bug
- Go 运行时死锁、数据竞争检测器对 Bug 的检测能力与算法原理（如果时间允许）
- 论文的结论、争议与我们的反思

## 分享地址

2019-09-12, 21:00 ~ 22:10, UTC+8

https://zoom.us/j/6923842137

## 进一步阅读的材料

- [Ou, 2019a] Real-world Go Concurrency Bugs PPT 
  - 本次分享的 PPT: [这里](https://docs.google.com/presentation/d/1clppbBqjxzPrj-26d_zVeJK2fFiXCsNVXYhKPjEZ4Tc/edit?usp=sharing)
- [Pike, 2012] [*Go Concurrency Patterns*](https://talks.golang.org/2012/concurrency.slide)
  - Rob Pike 关于 「Go 并发模式」的 PPT
- [Gerrand, 2013] [*Advanced Go Concurrency Patterns*](http://talks.golang.org/2013/advconc.slide)
  - Sameer Ajmani 关于 「Go 并发模式进阶」的 PPT
- [Tu et al., 2019] [Understanding Real-World Concurrency Bugs in Go](https://songlh.github.io/paper/go-study.pdf)
  - 本次分享讨论的论文
  - [论文作者的 PPT](https://slideplayer.com/slide/17049966/)
  - [Bug Table](https://github.com/talk-go/night/files/3587505/bug.table.xlsx)
  - 论文中对应的 GitHub 仓库：https://github.com/system-pclub/go-concurrency-bugs
  - 与论文作者的一次面谈记录：https://www.jexia.com/en/blog/golang-error-proneness-message-passing/
  - Hacker News 对本文的讨论 https://news.ycombinator.com/item?id=19280927
- [Utahn, 2019] [Go channels are bad and you should feel bad](https://www.jtolio.com/2016/03/go-channels-are-bad-and-you-should-feel-bad/)
  - Reddit 对本文的讨论：https://www.reddit.com/r/golang/comments/48mnrp/go_channels_are_bad_and_you_should_feel_bad/
- [Ou, 2019b] [Go 夜读 第 56 期：channel & select 源码分析](https://github.com/talk-go/night/issues/450)


请点击：https://github.com/talk-go/night/issues/464

----

## QA

**Q: 可以再详细说一下 lift 指标吗？**

A: 可以从两种不同的角度来思考这个指标。

1. 借鉴 person 相关性系数（余弦相似性) |X·Y|/(|X|*|Y|)
2. 借鉴 Bayes 公式 P(B|A) = P(AB)/P(A) 
  - Lift(cause, fix) = 导致阻塞的 cause 且使用了 fix 进行修复的概率 除以 cause 的概率乘以 fix 的概率 = P(cause, fix) / (P(cause)P(fix)) = P(cause|fix)/P(cause)
  - 接近 1 时，说明 fix 导致 cause 的概率接近 cause 自己的概率，即 P(cause|fix) 约等于 P(cause) 于是 fix 和 cause 独立
  - 大于 1 时，说明 fix 导致 cause 的概率比 cause 自己的概率大，即 P(cause|fix) > P(cause) => P(fix | cause) > P(fix)，即 cause 下 fix 的概率比 fix 本身的概率大，正相关
  - 小于 1 时，同理，负相关

**Q: 可以贴一下提到的两篇相关文献吗？**

A: 论文引用了两篇很硬核的形式化验证的论文：

- Julien Lange, Nicholas Ng, Bernardo Toninho, and Nobuko Yoshida. Fencing off go: Liveness and safety for channel-based programming. In Proceedings ofthe 44th ACMSIGPLANSymposium on Principles of Programming Languages (POPL ’17), Paris, France, January 2017.

- Julien Lange, Nicholas Ng, Bernardo Toninho, and Nobuko Yoshida. A static verification framework for message passing in go using be- havioural types. In IEEE/ACM40th International Conference on Software Engineering (ICSE ’18), Gothenburg, Sweden, June 2018.

**Q: 作者还分享了其他语言的关于并发 Bug 的论文，比如 Rust。**

A: 地址在[这里](https://arxiv.org/pdf/1902.01906.pdf)，但是思路完全一致，可以直接扫一眼结论。

Q: 能否将 CSP 和 Actor 模型进行一下简单比较？

A: CSP 和 Actor 的本质区别在于如何对消息传递进行建模。Actor 模型强调每个通信双方都有一个“信箱”，传递消息的发送方对将这个消息发给谁是很明确的，这种情况下 Actor 的信箱必须能够容纳不同类型的消息，而且理论上这个信箱的大小必须无穷大，很像你想要送一件礼物给别人，你会直接把礼物递给这个人，如果这个人不在，你就扔到他家的信箱里；CSP 需要有一个信道，因此对发送方而言，其实它并不知道这个消息的接收方是谁，更像是你朝大海扔了一个漂流瓶，大海这个信道根据洋流将这个漂流瓶传递给了其他正在观察监听大海的人。

**Q: 读论文的目标是什么？**

A: 我读论文主要有两个目标：1. 了解论文的研究方法，因为研究方法可能可以用在我未来的研究中；2. 了解论文的整体思路，因为论文很多，思路远比它们的结果对我未来自己的研究更重要。

**Q: 去哪儿找这类论文？**

A: 我们这次讨论的论文是我偶然在 Go 语言 GitHub 仓库的 Wiki 上看到的；一般情况下我会订阅 ArXiv，然后定期浏览新发出来的文章。


---


## 观看视频

{{< youtube id="WZUii-Czaps" >}}
