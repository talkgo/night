---
date: 2018-11-08T21:07:13+01:00
title: Go 夜读
type: index
weight: 0
---

[![Build Status](https://travis-ci.org/developer-learning/night-reading-go.svg?branch=master)](https://travis-ci.org/developer-learning/night-reading-go) [![Go Report Card](https://goreportcard.com/badge/github.com/developer-learning/night-reading-go)](https://goreportcard.com/report/github.com/developer-learning/night-reading-go) [![GitHub stars](https://img.shields.io/github/stars/developer-learning/night-reading-go.svg?label=Stars)](https://github.com/developer-learning/night-reading-go) [![GitHub forks](https://img.shields.io/github/forks/developer-learning/night-reading-go.svg?label=Fork)](https://github.com/developer-learning/night-reading-go) [![All Contributors](https://img.shields.io/badge/all_contributors-24-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/developer-learning/night-reading-go?status.svg)](http://godoc.org/github.com/developer-learning/night-reading-go) [![Coverage Status](https://coveralls.io/repos/github/developer-learning/night-reading-go/badge.svg?branch=master)](https://coveralls.io/github/developer-learning/night-reading-go?branch=master) [![GitHub issues](https://img.shields.io/github/issues/developer-learning/night-reading-go.svg?label=Issue)](https://github.com/developer-learning/night-reading-go/issues) [![license](https://img.shields.io/github/license/developer-learning/night-reading-go.svg)](https://github.com/developer-learning/night-reading-go/blob/master/LICENSE)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/developer-learning/night-reading-go.svg)](https://starcharts.herokuapp.com/developer-learning/night-reading-go)

[night-reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/?u=developer-learning&r=night-reading-go)

每周约定一个晚上进行 Go 源码阅读，Go 新手可以先去这里看看 **[Go 学习之路](https://github.com/developer-learning/learning-golang)**。

>阅读计划：Go 标准包、开源项目源代码。

----
<br>

|![notification](./images/bell-outline-badged.svg)预告|
|:------------------:|
|2018年11月29号晚上8点开始，讨论主题：《Go errors 处理》|

## 1. 往期回顾

| 期数 | 分析总结 | YouTube |
| :---- | :---- | :---- |
| 第 20 期 | [go test 及测试覆盖率] | [https://youtu.be/qkFFIIaTgHM](https://youtu.be/qkFFIIaTgHM) |
| 第 19 期 | [如何开发一个简单高性能的http router及gorouter源码分析 - xujiajun](https://github.com/xujiajun/gorouter) | [https://youtu.be/3BoStxKECL0](https://youtu.be/3BoStxKECL0) |
| 第 18 期 | 去中心化加密通信框架 [CovenantSQL](https://github.com/CovenantSQL/CovenantSQL)/DH-RPC的设计 | [https://youtu.be/bAfiKsLbDeE](https://youtu.be/bAfiKsLbDeE) |
| 第 17 期 | [2018-09-20 grpc 开发及 [grpcp](https://github.com/0x5010/grpcp) 的源码分析] | [https://youtu.be/sMBgYYEgm3c](https://youtu.be/sMBgYYEgm3c) |
| 第 16 期 | [2018-09-06 OpenFaas 介绍及源码分析](./reading/20180906/2018-09-06-openfaas-guide/) | [https://youtu.be/bZtgrAVR9HQ](https://youtu.be/bZtgrAVR9HQ) |
| 第 15 期 | [2018-08-23 多路复用资源池组件剖析](./reading/20180823/2018-08-23-pool-workshop-in-go/) | [https://youtu.be/CDfrRzgmR4E](https://youtu.be/CDfrRzgmR4E) |
| 第 14 期 | [2018-08-17 sync.Pool 源码分析及适用场景](./reading/20180817/2018-08-17-sync-pool-reading.pdf) | [https://youtu.be/jaepwn2PWPk](https://youtu.be/jaepwn2PWPk) |
| 第 13 期 | Kubernetes 入门指南 | [https://youtu.be/DJgYlmGCmDA](https://youtu.be/DJgYlmGCmDA) |
| 第 12 期 | [2018-08-02-(runtime goroutine调度实现)](./reading/20180802/README/) | [https://youtu.be/98pIzaOeD2k](https://youtu.be/98pIzaOeD2k) |
| 第 11 期 | [2018-07-26 Golang 代码质量持续检测实践](./articles/sonarqube-for-golang/2018-07-22-sonarqube-for-golang/) | [https://youtu.be/d95PZDAabqQ](https://youtu.be/d95PZDAabqQ) |
| 第 10 期 | [2018-06-28-(net/http/server.go、request.go和net/textproto/reader.go)](./reading/20180628/README/) | [https://youtu.be/xodlVBWxTYM](https://youtu.be/xodlVBWxTYM) |
| 第 9 期 | [2018-06-14-(net/http/server.go 和 h2_bundle.go 系列三)](./reading/20180614/README/)                             | |
| 第 8 期 | [2018-05-31-(net/http/server.go 系列二)](./reading/20180531/README/) | [https://youtu.be/U84dn76gixQ](https://youtu.be/U84dn76gixQ) |
| 第 7 期 | [2018-05-24-(net/http/server.go 系列一)](./reading/20180524/README/) | [https://youtu.be/H3oXjpiOReQ](https://youtu.be/H3oXjpiOReQ) |
| 第 6 期 | [2018-05-17-(strings/strings.go 系列二)](./reading/20180517/README/) | |
| 第 5 期 | [2018-05-10-(strings/strings.go 系列一)](./reading/20180510/README/) | |
| 第 4 期 | [2018-04-25-(strings/replace.go)、strings/search.go](./reading/20180425/README/) | |
| 第 3 期 | [2018-04-18-(strings/builder.go、strings/compare.go、strings/reader.go)](./reading/20180418/README/) | |
| 第 2 期 | [2018-04-11-(telport、tp-micro、ants)](./reading/20180411/README/) | |
| 第 1 期 | [2018-03-21-(cannot take address of temporary variables、telport、goutil、neochain)](./reading/20180321/README/) | |

>[查看更多](./reading/README/)

## 2. 日常讨论总结

- [2018-11-09 强制使用字段命名方式初始化结构体](./discuss/2018-11-09-force-to-use-keyed-struct-literals/)
- [2018-11-08 aws的ec2虚拟机示例创建后无法ssh连接登录问题](./discuss/2018-11-08-aws-ec2-ssh-login-problem/)
- [2018-10-18 cannot take address of temporary variables](./discuss/2018-11-08-address_operators/)
- [2018-10-18 encOp 怎么实现的呢？](./discuss/2018-10-18-encOp/)
- [2018-09-28 goroutine 中怎么得到返回值？](./discuss/2018-09-28-return-value-in-waitgroup/)
- [2018-09-19 Producer Consumer 丢包问题](./discuss/2018-09-19-wechat-discuss/)
- [2018-09-18 HTTP 压力测试工具](./discuss/2018-09-18-benchmark-tools/)
- [2018-09-14 VSCode Go 代码自动补全和自动导入包](./discuss/2018-09-14-tips-in-vscode/)
- [2018-09-05 浅谈 Git 系统平台](./discuss/2018-09-05-git-system/)
- [2018-09-04 protobuffer 3 enum && Body bind](./discuss/2018-09-04-wechat-discuss/)
- [2018-08-30 深入理解 Go Interfaces](./discuss/2018-08-30-understanding-go-interfaces/)
- [2018-08-24 数据库管理工具 GUI&CLI](./discuss/2018-08-24-wechat-discuss/)
- [2018-08-23 博客平台的选择 && kafka && nats && spinning threads讨论](./discuss/2018-08-23-wechat-discuss/)
- [2018-08-15 一篇 zap 日志库引发的激烈讨论，以及 sync.Pool 到底是用来干嘛的？](./discuss/2018-08-15-wechat-discuss/)
- [2018-08-14 做实时语音流，用什么来做比较好？rtmp？还是ws？Go 从 1.5 开始默认使用了 CPU 核数？etcd相关讨论](./discuss/2018-08-14-wechat-discuss/)
- [2018-08-09 Go 语言终端日志着色](./discuss/2018-08-09-log-color-in-go/)
- [2018-08-02 APNS 批量发送推送通知](./discuss/2018-08-02-apns-push-notification/)
- [2018-08-02 go 调用 shell 脚本如何传递参数讨论](./discuss/2018-08-02-go-shell/)
- [2018-07-31 println与fmt.Println有何猫腻？](./discuss/2018-07-31-println-Println-add_context/)
- [2018-07-14 包版本管理？](./discuss/2018-07-14-version-gopath-go-command/)
- [2018-07-11 Go在32位系统中使用64位原子操作的坑](./discuss/2018-07-11-using_64bit_atomic_in_32bit_system/)
- [2018-07-09 Go 中 make 和 new 的用法讨论](./discuss/2018-07-09-make-new-in-go/)
- [2018-07-04 关于包命名的讨论](./discuss/2018-07-04-package-names/)
- [2018-07-02 关于C1000k问题的讨论](./discuss/2018-07-02-c1000k-on-linux/)
- [2018-06-07 tcp连接设置timeout的问题讨论](./discuss/2018-06-07-dial-timeout-in-go/)
- [2018-05-28 Go 语言中 CPU 和内存分析的问题讨论](./discuss/2018-05-28-pprof-in-go/)
- [2018-05-23 锁失效和RPC框架选择的问题讨论](./discuss/2018-05-23-wechat-discuss/)
- [2018-05-22 字符串转字节切片的问题讨论](./discuss/2018-05-22-go-string-to-byte-slice/)
- [2018-05-21 在循环迭代值在 goroutine 中的使用等问题讨论](./discuss/2018-05-21-using-goroutines-on-loop-iterator-variables/)
- [2018-05-18 bitset 实现和循环导入问题讨论](./discuss/2018-05-18-bitset-and-import-cycle-not-allowed/)
- [2018-05-13 变量的作用域是贯穿整个 if-else-if 的](./discuss/2018-05-13-declaring-variables-on-if-else/)
- [Go Vendor && GOPATH 讨论](./discuss/2018-05-10-which-vendor-tool/)
- [2018-05-09 调试-开发工具-赋值:=和=的差别-变量作用域等等](./discuss/2018-05-09-wechat-discuss/)
- [2018-05-08 Go 语言中下划线的分析总结](./discuss/2018-05-08-anlayze-underscore-in-go/)

>[查看更多](./discuss/README/)

## 3. 技术分析总结

- [2018-11-11 golang 文件锁操作](./articles/2018-11-11-golang-file-lock/)
- [2018-11-02 sync.RWMutex 源码分析](./articles/sync/sync_rwmutex_source_code_analysis/)
- [2018-10-31 sync.Mutex 源码分析](./articles/sync/sync_mutex_source_code_analysis/)
- [2018-07-26 Sony gobreaker 熔断器源码分析](./articles/sony-gobreaker/README/)
- [2018-07-22 Golang 代码持续检测实践](./articles/sonarqube-for-golang/2018-07-22-sonarqube-for-golang/)
- [2018-06-12 sync.Map源码分析](./articles/sync/sync_Map_source_code_analysis/)
- [2018-06-12 sync.RWMutex源码分析](./articles/sync/sync_rwmutex_source_code_analysis/)
- [2018-06-12 sync.WaitGroup源码分析](./articles/sync/sync_waitgroup_source_code_analysis/)
- [2018-06-13 sync.Cond源码分析](./articles/sync/sync_cond_source_code_analysis/)
- [2018-06-13 sync.Once源码分析](./articles/sync/sync_once_source_code_analysis/)
- [2018-05-31 批量删除redis key](./articles/2018-05-31-batch-del-redis-key/)

>[查看更多](./articles/README/)

## 4. 深度剖析

1. 深度剖析 Boyer-Moore 和 Rabin-Karp 等字符串搜索算法。
2. 深度剖析 bitset 。

## 5. 大咖技术分享

有兴趣的可以联系我，并且提供你要分享的话题。

[其他更多](./other/README/)

## 6. 面试题专区

- [笔试题](./interview/interview-pen/)
- [Golang语言](./interview/interview-golang-language/)
- [数据库](./interview/interview-database/)
- [网络](./interview/interview-network/)
- [操作系统](./interview/interview-os/)
- [数据结构](./interview/interview-data-structure/)
- [架构](./interview/interview-architecture/)
- [设计题](./interview/interview-design/)
- [算法题](./interview/interview-algorithm/)

[其他更多](./interview/README/)

## 7. 其他

- [数据库工具推荐](./other/db-client-tools/)
- [压力测试工具](./other/benchmark-tools/)

----

## 我们的目标

我们希望可以推进大家深入了解 Go ，快速成长为资深的 Gopher 。我们希望每次来了的人和没来的人都能够有收获，成长。

## 我们的方式

由一个主讲人带着大家一起去阅读 Go 源代码，一起去啃那些难啃的算法、学习代码里面的奇淫技巧，遇到问题或者有疑惑了，我们可以一起去检索，解答这些问题。我们可以一起学习，共同成长。

>阅读规则：选取 package 包，然后从上往下开始读 xxx.go 文件，每个文件从上往下读导出的函数（一步一步跟逻辑，如果逻辑跳出这个 package 则不做深入探究）。

## 我们的精神

开源！开源！开源！重要的事，一定要说三遍。

>希望有兴趣的小伙伴们一起加入，让我们一起把 《Go 夜读》建立成一个对大家都有帮助的开源社区。

## 怎么加入

目前微信群已经超过 `100` 人，请微信搜索 `mai_yang` ，然后备注你的姓名、公司、工作岗位和职责，备注来源：Github。

## Discord

Discord 入群链接：https://discord.gg/Q2FDua9

----

## 如何参与贡献？

```sh
├── reading   // Go 源码阅读
├── discuss   // 日常微信群讨论的总结
├── articles  // 个人原创的技术文章
└── other     // 其他
```

[如何参与贡献](./CONTRIBUTING/)

## Contributors

我非常重视每一个对这个项目的贡献者，我会将贡献者列表更新到这里，目前只有提交 Pull Request 的小伙伴，但是贡献不仅仅如此，还可以包括提交 Issue 以及在社群中有所贡献的人。

贡献者自己可以提 PR ，方法如下：

- 安装 `npm install -g --save-dev all-contributors-cli`
- `sh gen_contributors.sh`

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars3.githubusercontent.com/u/1710912?v=4" width="100px;"/><br /><sub><b>maiyang</b></sub>](https://github.com/yangwenmai)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=yangwenmai "Code") [🤔](#ideas-yangwenmai "Ideas, Planning, & Feedback") [👀](#review-yangwenmai "Reviewed Pull Requests") [📢](#talk-yangwenmai "Talks") [✅](#tutorial-yangwenmai "Tutorials") | [<img src="https://avatars1.githubusercontent.com/u/16773339?v=4" width="100px;"/><br /><sub><b>Simple Min</b></sub>](https://github.com/mougeCM)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=mougeCM "Code") | [<img src="https://avatars3.githubusercontent.com/u/35653599?v=4" width="100px;"/><br /><sub><b>kenny</b></sub>](https://github.com/yuhao5)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=yuhao5 "Code") | [<img src="https://avatars2.githubusercontent.com/u/13843868?v=4" width="100px;"/><br /><sub><b>charnlsxy</b></sub>](https://github.com/charnlsxy)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=charnlsxy "Code") | [<img src="https://avatars3.githubusercontent.com/u/11901298?v=4" width="100px;"/><br /><sub><b>AceDarkknight</b></sub>](https://github.com/AceDarkknight)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=AceDarkknight "Code") | [<img src="https://avatars2.githubusercontent.com/u/3014297?v=4" width="100px;"/><br /><sub><b>Data</b></sub>](https://github.com/gnuos)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=gnuos "Code") | [<img src="https://avatars0.githubusercontent.com/u/2876745?v=4" width="100px;"/><br /><sub><b>侯名</b></sub>](https://github.com/KISSMonX)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=KISSMonX "Code") |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars0.githubusercontent.com/u/12060175?v=4" width="100px;"/><br /><sub><b>dumliu01</b></sub>](https://github.com/dumliu01)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=dumliu01 "Code") | [<img src="https://avatars0.githubusercontent.com/u/1411282?v=4" width="100px;"/><br /><sub><b>hlily2005</b></sub>](https://github.com/hlily2005)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=hlily2005 "Code") | [<img src="https://avatars0.githubusercontent.com/u/16982786?v=4" width="100px;"/><br /><sub><b>fenggolang</b></sub>](https://github.com/fenggolang)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=fenggolang "Code") | [<img src="https://avatars3.githubusercontent.com/u/10174178?v=4" width="100px;"/><br /><sub><b>henrylee2cn</b></sub>](https://github.com/henrylee2cn)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=henrylee2cn "Code") | [<img src="https://avatars0.githubusercontent.com/u/1336914?v=4" width="100px;"/><br /><sub><b>shaqsnake</b></sub>](https://github.com/shaqsnake)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=shaqsnake "Code") | [<img src="https://avatars0.githubusercontent.com/u/5728787?v=4" width="100px;"/><br /><sub><b>tbwisk</b></sub>](https://github.com/TBWISK)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=TBWISK "Code") | [<img src="https://avatars3.githubusercontent.com/u/416141?v=4" width="100px;"/><br /><sub><b>Huang ChuanTong</b></sub>](https://github.com/toontong)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=toontong "Code") |
| [<img src="https://avatars3.githubusercontent.com/u/10513552?v=4" width="100px;"/><br /><sub><b>The notes of SQL optimize </b></sub>](https://github.com/zhongxuan123)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=zhongxuan123 "Code") | [<img src="https://avatars2.githubusercontent.com/u/29008269?v=4" width="100px;"/><br /><sub><b>zhouxinxin19920802</b></sub>](https://github.com/zhouxinxin19920802)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=zhouxinxin19920802 "Code") | [<img src="https://avatars2.githubusercontent.com/u/20811449?v=4" width="100px;"/><br /><sub><b>macaria</b></sub>](https://github.com/macaria)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=macaria "Code") | [<img src="https://avatars3.githubusercontent.com/u/15226239?v=4" width="100px;"/><br /><sub><b>Dennis</b></sub>](http://github.com/DennisMao)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=DennisMao "Code") | [<img src="https://avatars1.githubusercontent.com/u/2696746?v=4" width="100px;"/><br /><sub><b>orangleliu</b></sub>](http://blog.csdn.net/orangleliu)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=orangle "Code") | [<img src="https://avatars1.githubusercontent.com/u/21693162?v=4" width="100px;"/><br /><sub><b>HarbinZhang</b></sub>](https://github.com/HarbinZhang)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=HarbinZhang "Code") | [<img src="https://avatars1.githubusercontent.com/u/7344921?v=4" width="100px;"/><br /><sub><b>LiMingji</b></sub>](https://github.com/SwanSpouse)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=SwanSpouse "Code") |
| [<img src="https://avatars0.githubusercontent.com/u/22164927?v=4" width="100px;"/><br /><sub><b>wintersnow</b></sub>](https://mickey0524.github.io/)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=mickey0524 "Code") | [<img src="https://avatars2.githubusercontent.com/u/44076738?v=4" width="100px;"/><br /><sub><b>zhuzhenfeng</b></sub>](https://github.com/zhuzhenfeng-finogeeks)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=zhuzhenfeng-finogeeks "Code") | [<img src="https://avatars2.githubusercontent.com/u/6065007?v=4" width="100px;"/><br /><sub><b>徐佳军</b></sub>](https://xujiajun.cn)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=xujiajun "Code") |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!
