---
title: 第 1 期每日阅读特训营
date: 2019-03-14T09:10:00+08:00
---

# 阅读清单

| 标题 | 阅读者 |
|----|----|
|[漫话：如何给女朋友解释为什么有些网站域名不以 www 开头](https://mp.weixin.qq.com/s?__biz=Mzg3MjA4MTExMw==&mid=2247484994&idx=1&sn=e5cbc3175ef0dd88e76aa7b69e31a82b&chksm=cef5f5f4f9827ce2d91c11f62219d60ccaa09bf09cadc8bae4a786da4b8a3880055582ceff9b&token=79184148&lang=zh_CN#rd) | mai |
| [LeetCode 的刷题利器（伪装到老板都无法 diss 你没有工作）](https://github.com/jdneo/vscode-leetcode/blob/master/docs/README_zh-CN.md) | mai |
|[深入 GO 语言文本类型](https://vonng.com/blog/go-text-types/) | ch |
| [k8s cpu 资源限制](https://mp.weixin.qq.com/s/yLQrBPl729yQD26YBSZ50A) [k8s 内存资源限制](https://mp.weixin.qq.com/s?__biz=MzIzNzU5NTYzMA==&mid=2247486237&idx=1&sn=640b7ad99e3ddf144027f113cccfa728&chksm=e8c7759cdfb0fc8aeaac1b76c019c796aa702307bbaf648ccdb7eb0873e0c453fe93611a978a&scene=21#wechat_redirect) | Jason |
| [Golang 之轻松化解 defer 的温柔陷阱](https://mp.weixin.qq.com/s/txj7jQNki_8zIArb9kSHeg?scene=25#wechat_redirect) [关于 go 语言中的延迟执行函数](https://www.jianshu.com/p/441c016f527e) | Littlesqx |


## “漫话”的阅读笔记

1. 域名的一个重要功能——为数字化的互联网资源提供易于记忆的名称。
2. 域名具有唯一性。
3. www，其实是 World Wide Web 的缩写，中文翻译为万维网
4. 互联网并不等同万维网（WWW），万维网只是一个基于超文本相互链接而成的全球性系统，且是互联网所能提供的服务其中之一。
5. 为了区分互联网中的各种应用，就有了不同的子域名，比如互联网就以 www 作为子域名，文件传输以 ftp 作为子域名，电子邮件以 mail 作为子域名。
6. 正是因为万维网是互联网中最重要的一部分，很多域名的最主要用途也是搭建 web 网站，所以，会有很多公司直接忽略 www。
7. 通用顶级域（英语：Generic top-level domain，缩写为 gTLD），是互联网名称与数字地址分配机构（IANA）管理的顶级域（TLD）之一。该机构专门负责互联网的域名系统。
8. 域名支持中文，并且域名中也已经支持颜文字了。“👀.我爱你”

## “Go文本类型”的阅读笔记

1. 为什么字符串直接取下标，s[0] 类型是 byte，range 时 类型为 rune？
2. k,v := range string，k 是 UTF-8 编码字节的下标，v 是 Unicode。[For statements with range clause](https://golang.org/ref/spec#For_statements)
3. 扩展 [fmt 格式化](https://golang.org/pkg/fmt/)

## k8s 资源限制

1. pod 调度是按 container 中 request 资源总和调度的；
2. limit 帮助 kubelet 约束本节点上容器资源使用最大份额；
3. 针对 内存 这种不可压缩资源，超额将导致 container 中程序 OOM 退出；而 CPU 只是影响了程序的等待时长；
4. 内存资源：cgroup 中 `memory.limit_in_bytes` (最大内存 limit 份额)、`memory.soft_limit_in_bytes` (request 内存份额)
5. CPU 资源，分两部分：
    * request cpu 份额: cgroup 中 `cpu.shares`，其中记录的数值是 CPU 分片数，一个 CPU 分为 1024 片（k8s 分成 1000 片），request.memory: 500m 表示申请 500/1000 个 cpu; 它保证程序能接收到申请的份数 CPU 片，但是如果程序没完全使用，其他程序是可以占用的；
    * limit cpu 份额：`cpu.cfs_period_us` 表示一个 CPU 执行时间周期的时间，通常为 100ms（带宽控制系统定义了一个通常是 1/10 秒的周期）;`cpu.cfs_quota_us` 表示一个周期内分配的 CPU 时间; cpu.cfs_period_us/cpu.cfs_quota_us 合起来表示 limit 的 CPU 份额；
6. 查看 cgroup 命令：

```
1. 找到容器中程序在宿主机上进程 ID
$ cat /proc/${PID}/cgroup
...
8:cpuacct,cpu:/kubepods/burstable/podxxx/dockerid
7:memory:/kubepods/burstable/podxxx/dockerid
...
2. ls /sys/fs/cgroup/cpuacct,cpu/kubepods/burstable/podxxx/dockerid/ 
3. ls /sys/fs/cgroup/memory/kubepods/burstable/podxxx/dockerid
```

## “defer” 笔记

1. 每次 defer 语句执行的时候，会把函数“压栈”，**函数参数会被拷贝下来**；当在当前函数执行完毕后（包括通过 return 正常结束或者**panic 导致的异常结束**），defer 函数按照定义的逆序执行；**如果 defer 执行的函数为 nil, 那么会在最终调用函数的产生 panic**。
2. Golang 内置的带返回值的函数无法进行延迟调用（调用的结果不可以抛弃，copy 和 recover 例外），可通过放入到匿名函数中再 defer。
3. defer 常用场景：控制资源（文件、数据库、锁等）的申请和释放，使代码简洁；配合 recover 实现异常恢复。
4. defer 副作用（坑）：延迟执行的机制损耗性能；延迟意味着资源占用，需要时刻警惕尽可能早地释放。

## 观看视频

{{< youtube id="" >}}
