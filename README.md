# Go 夜读

[![Build Status](https://travis-ci.org/developer-learning/night-reading-go.svg?branch=master)](https://travis-ci.org/developer-learning/night-reading-go) [![Go Report Card](https://goreportcard.com/badge/github.com/developer-learning/night-reading-go)](https://goreportcard.com/report/github.com/developer-learning/night-reading-go) [![GitHub stars](https://img.shields.io/github/stars/developer-learning/night-reading-go.svg?label=Stars)](https://github.com/developer-learning/night-reading-go) [![GitHub forks](https://img.shields.io/github/forks/developer-learning/night-reading-go.svg?label=Fork)](https://github.com/developer-learning/night-reading-go) [![All Contributors](https://img.shields.io/badge/all_contributors-18-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/developer-learning/night-reading-go?status.svg)](http://godoc.org/github.com/developer-learning/night-reading-go) [![Coverage Status](https://coveralls.io/repos/github/developer-learning/night-reading-go/badge.svg?branch=master)](https://coveralls.io/github/developer-learning/night-reading-go?branch=master) [![GitHub issues](https://img.shields.io/github/issues/developer-learning/night-reading-go.svg?label=Issue)](https://github.com/developer-learning/night-reading-go/issues) [![license](https://img.shields.io/github/license/developer-learning/night-reading-go.svg)](https://github.com/developer-learning/night-reading-go/blob/master/LICENSE)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/developer-learning/night-reading-go.svg)](https://starcharts.herokuapp.com/developer-learning/night-reading-go)

[night-reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/?u=developer-learning&r=night-reading-go)

每周约定一个晚上进行 Go 源码阅读，Go 新手可以先去这里看看 **[Go 学习之路](https://github.com/developer-learning/learning-golang)**。

>阅读计划：Go 标准包、开源项目源代码。

----

## 0. 预告

暂无。

## 1. 往期回顾

| 源码总结 | YouTube |
| :---- | :---- |
| Kubernetes 入门指南 | [https://youtu.be/DJgYlmGCmDA](https://youtu.be/DJgYlmGCmDA) |
| [2018-08-02-(runtime goroutine调度实现)](./reading/20180802/README.md) | [https://youtu.be/98pIzaOeD2k](https://youtu.be/98pIzaOeD2k) |
| [2018-07-26 Golang 代码质量持续检测实践](./articles/sonarqube-for-golang/2018-07-22-sonarqube-for-golang.md) | [https://youtu.be/d95PZDAabqQ](https://youtu.be/d95PZDAabqQ) |
| [2018-06-28-(net/http/server.go、request.go和net/textproto/reader.go)](./reading/20180628/README.md) | [https://youtu.be/xodlVBWxTYM](https://youtu.be/xodlVBWxTYM) |
| [2018-06-14-(net/http/server.go 和 h2_bundle.go 系列三)](./reading/20180614/README.md)                             | |
| [2018-05-31-(net/http/server.go 系列二)](./reading/20180531/README.md) | [https://youtu.be/U84dn76gixQ](https://youtu.be/U84dn76gixQ) |
| [2018-05-24-(net/http/server.go 系列一)](./reading/20180524/README.md) | [https://youtu.be/H3oXjpiOReQ](https://youtu.be/H3oXjpiOReQ) |
| [2018-05-17-(strings/strings.go 系列二)](./reading/20180517/README.md) | |
| [2018-05-10-(strings/strings.go 系列一)](./reading/20180510/README.md) | |
| [2018-04-25-(strings/replace.go)、strings/search.go](./reading/20180425/README.md) | |
| [2018-04-18-(strings/builder.go、strings/compare.go、strings/reader.go)](./reading/20180418/README.md) | |
| [2018-04-11-(telport、tp-micro、ants)](./reading/20180411/README.md) | |
| [2018-03-21-(cannot take address of temporary variables、telport、goutil、neochain)](./reading/20180321/README.md) | |

>[查看更多](./reading/README.md)

## 2. 日常讨论总结

- [2018-08-15 一篇 zap 日志库引发的激烈讨论，以及 sync.Pool 到底是用来干嘛的？](./discuss/2018-08-15-wechat-discuss.md)
- [2018-08-14 做实时语音流，用什么来做比较好？rtmp？还是ws？Go 从 1.5 开始默认使用了 CPU 核数？etcd相关讨论](./discuss/2018-08-14-wechat-discuss.md)
- [2018-08-09 Go 语言终端日志着色](./discuss/2018-08-09-log-color-in-go.md)
- [2018-08-02 APNS 批量发送推送通知](./discuss/2018-08-02-apns-push-notification.md)
- [2018-08-02 go 调用 shell 脚本如何传递参数讨论](./discuss/2018-08-02-go-shell.md)
- [2018-07-31 println与fmt.Println有何猫腻？](./discuss/2018-07-31-println-Println-add_context.md)
- [2018-07-14 包版本管理？](./discuss/2018-07-14-version-gopath-go-command.md)
- [2018-07-11 Go在32位系统中使用64位原子操作的坑](./discuss/2018-07-11-using_64bit_atomic_in_32bit_system.md)
- [2018-07-09 Go 中 make 和 new 的用法讨论](./discuss/2018-07-09-make-new-in-go.md)
- [2018-07-04 关于包命名的讨论](./discuss/2018-07-04-package-names.md)
- [2018-07-02 关于C1000k问题的讨论](./discuss/2018-07-02-c1000k-on-linux.md)
- [2018-06-07 tcp连接设置timeout的问题讨论](./discuss/2018-06-07-dial-timeout-in-go.md)
- [2018-05-28 Go 语言中 CPU 和内存分析的问题讨论](./discuss/2018-05-28-pprof-in-go.md)
- [2018-05-23 锁失效和RPC框架选择的问题讨论](./discuss/2018-05-23-wechat-discuss.md)
- [2018-05-22 字符串转字节切片的问题讨论](./discuss/2018-05-22-go-string-to-byte-slice.md)
- [2018-05-21 在循环迭代值在 goroutine 中的使用等问题讨论](./discuss/2018-05-21-using-goroutines-on-loop-iterator-variables.md)
- [2018-05-18 bitset 实现和循环导入问题讨论](./discuss/2018-05-18-bitset-and-import-cycle-not-allowed.md)
- [2018-05-13 变量的作用域是贯穿整个 if-else-if 的](./discuss/2018-05-13-declaring-variables-on-if-else.md)
- [Go Vendor && GOPATH 讨论](./discuss/2018-05-10-which-vendor-tool.md)
- [2018-05-09 调试-开发工具-赋值:=和=的差别-变量作用域等等](./discuss/2018-05-09-wechat-discuss.md)
- [2018-05-08 Go 语言中下划线的分析总结](./discuss/2018-05-08-anlayze-underscore-in-go.md)

>[查看更多](./discuss/README.md)

## 3. 技术分析总结

- [2018-07-26 Sony gobreaker容断器源码分析](./articles/sony-gobreaker/README.md)
- [2018-07-22 Golang 代码持续检测实践](./articles/sonarqube-for-golang/2018-07-22-sonarqube-for-golang.md)
- [2018-06-12 sync.Map源码分析](./articles/sync/sync_Map_source_code_analysis.md)
- [2018-06-12 sync.RWMutex源码分析](./articles/sync/sync_rwmutex_source_code_analysis.md)
- [2018-06-12 sync.WaitGroup源码分析](./articles/sync/sync_waitgroup_source_code_analysis.md)
- [2018-06-13 sync.Cond源码分析](./articles/sync/sync_cond_source_code_analysis.md)
- [2018-06-13 sync.Once源码分析](./articles/sync/sync_once_source_code_analysis.md)
- [2018-05-31 批量删除redis key](./articles/2018-05-31_batch-del-redis-key.md)

>[查看更多](./articles/README.md)

## 4. 深度剖析

1. 深度剖析 Boyer-Moore 和 Rabin-Karp 等字符串搜索算法。
2. 深度剖析 bitset 。

## 5. 大咖技术分享

有兴趣的可以联系我，并且提供你要分享的话题。

[其他更多](./other/README.md)

## 6. 面试题专区

- [笔试题](./interview/interview-pen.md)
- [Golang语言](./interview/interview-golang-language.md)
- [数据库](./interview/interview-database.md)
- [网络](./interview/interview-network.md)
- [操作系统](./interview/interview-os.md)
- [数据结构](./interview/interview-data-structure.md)
- [架构](./interview/interview-architecture.md)
- [设计题](./interview/interview-design.md)
- [算法题](./interview/interview-algorithm.md)

[其他更多](./interview/README.md)

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

----

## 如何参与贡献？

```sh
├── reading   // Go 源码阅读
├── discuss   // 日常微信群讨论的总结
├── articles  // 个人原创的技术文章
└── other     // 其他
```

[如何参与贡献](./CONTRIBUTING.md)

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
| [<img src="https://avatars3.githubusercontent.com/u/10513552?v=4" width="100px;"/><br /><sub><b>The notes of SQL optimize </b></sub>](https://github.com/zhongxuan123)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=zhongxuan123 "Code") | [<img src="https://avatars2.githubusercontent.com/u/29008269?v=4" width="100px;"/><br /><sub><b>zhouxinxin19920802</b></sub>](https://github.com/zhouxinxin19920802)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=zhouxinxin19920802 "Code") | [<img src="https://avatars2.githubusercontent.com/u/20811449?v=4" width="100px;"/><br /><sub><b>macaria</b></sub>](https://github.com/macaria)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=macaria "Code") | [<img src="https://avatars3.githubusercontent.com/u/15226239?v=4" width="100px;"/><br /><sub><b>Dennis</b></sub>](http://github.com/DennisMao)<br />[💻](https://github.com/developer-learning/night-reading-go/commits?author=DennisMao "Code") |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!