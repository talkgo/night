## 2018-08-14

来源：《Go 夜读》微信群

时间：2018-08-14

### 问题：做实时语音流，用什么来做比较好？rtmp？还是ws？

- 第三方：即构（推荐）、声网（自己做不稳定，还得搞流媒体服务器）
- 自己做一般都是rtmp

>实时？需要有多实时？直播还是点到点通信？
>A：有一个实现是在 boya 的 rtmp 协议基础上，完善了 golang 的 rtmp 流媒体服务器，最近已经上网运行了，支持查询回源，定点回源等，资源消耗很低。也可参考项目[livego](https://github.com/gwuhaolin/livego)

### Go 默认使用 CPU 核数？

![](../images/2018-08-14-discuss-01.png)

具体代码：`proc.go#schedinit()`。

文档说明：[runtime-GOMAXPROCS](https://golang.org/doc/go1.5#runtime)

>Another potentially breaking change is that the runtime now sets the default number of threads to run simultaneously, defined by GOMAXPROCS, to the number of cores available on the CPU. In prior releases the default was 1. Programs that do not expect to run with multiple cores may break inadvertently. They can be updated by removing the restriction or by setting GOMAXPROCS explicitly. For a more detailed discussion of this change, see [the design document - Russ Cox](https://docs.google.com/document/d/1At2Ls5_fhJQ59kDK2DFVhFu3g5mATSXqqV5QrxinasI/edit).

[discuss on golang-dev](https://groups.google.com/forum/#!msg/golang-dev/POSw7qrelso/dI3YPTeGbkMJ)

### etcd

etcd 心跳超时是 1s ；
etcd 用了[raft 分布式一致性算法](http://thesecretlivesofdata.com/raft)。
其他的内容有待后续整理。。。

## 参考

1. [Raft一致性算法论文](https://github.com/maemual/raft-zh_cn/blob/master/raft-zh_cn.md)
2. [分布式一致性算法：可能比你想象得更复杂](https://mp.weixin.qq.com/s/ohTXhFFywGHGDOkzO45aaQ)
