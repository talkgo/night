---
date: 2018-11-08T21:07:13+01:00
title: Go 夜读
type: index
weight: 0
---

[![Build Status](https://travis-ci.org/developer-learning/reading-go.svg?branch=master)](https://travis-ci.org/developer-learning/reading-go) [![Go Report Card](https://goreportcard.com/badge/github.com/developer-learning/reading-go)](https://goreportcard.com/report/github.com/developer-learning/reading-go) [![GitHub stars](https://img.shields.io/github/stars/developer-learning/reading-go.svg?label=Stars)](https://github.com/developer-learning/reading-go) [![GitHub forks](https://img.shields.io/github/forks/developer-learning/reading-go.svg?label=Fork)](https://github.com/developer-learning/reading-go) [![All Contributors](https://img.shields.io/badge/all_contributors-27-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/developer-learning/reading-go?status.svg)](http://godoc.org/github.com/developer-learning/reading-go) [![Coverage Status](https://coveralls.io/repos/github/developer-learning/reading-go/badge.svg?branch=master)](https://coveralls.io/github/developer-learning/reading-go?branch=master) [![GitHub issues](https://img.shields.io/github/issues/developer-learning/reading-go.svg?label=Issue)](https://github.com/developer-learning/reading-go/issues) [![license](https://img.shields.io/github/license/developer-learning/reading-go.svg)](https://github.com/developer-learning/reading-go/blob/master/LICENSE)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/developer-learning/reading-go.svg)](https://starcharts.herokuapp.com/developer-learning/reading-go)

[reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/?u=developer-learning&r=reading-go)

Go 学习与分享：

- [Go 夜读](https://github.com/developer-learning/reading-go/labels/Go%20%E5%A4%9C%E8%AF%BB)

*根据[【草案】Go 夜读重大调整(请每个人都来说说你的看法和意见）](https://github.com/developer-learning/reading-go/issues/348)，我们将按计划进行 Go 源码阅读或者 Go 项目实践，你如果是 Go 新手可以先去这里看看 **[Go 学习之路](https://github.com/developer-learning/learning-golang)**。*
>范畴：Go 标准包、开源项目、Go 项目最佳实践等。

- [每日一问](https://github.com/developer-learning/reading-go/labels/%E6%AF%8F%E6%97%A5%E4%B8%80%E9%97%AE)

- [Go 项目实践](https://github.com/developer-learning/reading-go/labels/Go%20%E9%A1%B9%E7%9B%AE%E5%AE%9E%E8%B7%B5)
	- [Gin 开发](https://github.com/developer-learning/reading-go/labels/Gin%20%E5%BC%80%E5%8F%91)
	- [goim 开发(todo)]()
- **[TiDB 源码学习](https://github.com/developer-learning/reading-go/issues/382)（从6月5日开始，每周三晚上21点 [zoom live](https://zoom.us/j/6923842137) ）**：
	- TiDB 源码概览；
	- 执行引擎；
	- 优化器；
	- 事务；
	- 其他（工程化，测试等）；

----

## 阅读清单

- [x] strings
- [x] strconv
- [x] testing
- [x] net/http
- [x] sync
- [x] flag
- [x] etcd/raft
- [x] defer
- [x] context
- [x] golang/sync
- [x] kubernetes

## 回看地址

- [Go 夜读(YouTuBe)](https://www.youtube.com/channel/UCZwrjDu5Rf6O_CX2CVx7n8Q?sub_confirmation=1)
- [Go 夜读(B 站)](https://space.bilibili.com/326749661)

## 我们的目标

我们希望可以推进大家深入了解 Go ，快速成长为资深的 Gopher 。我们希望每次来了的人和没来的人都能够有收获，成长。

## 我们的方式

由一个主讲人带着大家一起去阅读 Go 源代码，一起去啃那些难啃的算法、学习代码里面的奇淫技巧，遇到问题或者有疑惑了，我们可以一起去检索，解答这些问题。我们可以一起学习，共同成长。

>阅读规则：选取 package 包，然后从上往下开始读 xxx.go 文件，每个文件从上往下读导出的函数（一步一步跟逻辑，如果逻辑跳出这个 package 则不做深入探究）。

## 我们的精神

开源！开源！开源！重要的事，一定要说三遍。

>希望有兴趣的小伙伴们一起加入，让我们一起把 《Go 夜读》建立成一个对大家都有帮助的开源社区。

## 怎么加入

如果你想加入微信群，请搜索 `mai_yang` ，然后备注你的姓名、公司、工作岗位和职责，备注来源：Github。

有同学想要用 Slack 交流，我开放了一个：[reading-go Slack](https://join.slack.com/t/reading-go/shared_invite/enQtMjgwNTU5MTE5NjgxLTA5NDQwYzE4NGNhNDI3N2E0ZmYwOGM2MWNjMDUyNjczY2I0OThiNzA5ZTk0MTc1MGYyYzk0NTA0MjM4OTZhYWE)

----

## 如何参与贡献？

```sh
├── reading   // Go 源码阅读
├── discuss   // 日常微信群讨论的总结
├── articles  // 个人原创的技术文章
├── interview // Go 面试专区
└── other     // 其他
```

[如何参与贡献](other/contributing/)

## Contributors

我非常重视每一个对这个项目的贡献者，我会将贡献者列表更新到这里，目前只有提交 Pull Request 的小伙伴，但是贡献不仅仅如此，还可以包括提交 Issue 以及在社群中有所贡献的人。

贡献者自己可以提 PR ，方法如下：

- 安装 `npm install -g --save-dev all-contributors-cli`
- `sh gen_contributors.sh`

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
| [<img src="https://avatars3.githubusercontent.com/u/1710912?v=4" width="100px;"/><br /><sub><b>maiyang</b></sub>](https://github.com/yangwenmai)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=yangwenmai "Code") [🤔](#ideas-yangwenmai "Ideas, Planning, & Feedback") [👀](#review-yangwenmai "Reviewed Pull Requests") [📢](#talk-yangwenmai "Talks") [✅](#tutorial-yangwenmai "Tutorials") | [<img src="https://avatars1.githubusercontent.com/u/16773339?v=4" width="100px;"/><br /><sub><b>Simple Min</b></sub>](https://github.com/mougeCM)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=mougeCM "Code") | [<img src="https://avatars3.githubusercontent.com/u/35653599?v=4" width="100px;"/><br /><sub><b>kenny</b></sub>](https://github.com/yuhao5)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=yuhao5 "Code") | [<img src="https://avatars2.githubusercontent.com/u/13843868?v=4" width="100px;"/><br /><sub><b>charnlsxy</b></sub>](https://github.com/charnlsxy)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=charnlsxy "Code") | [<img src="https://avatars3.githubusercontent.com/u/11901298?v=4" width="100px;"/><br /><sub><b>AceDarkknight</b></sub>](https://github.com/AceDarkknight)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=AceDarkknight "Code") | [<img src="https://avatars2.githubusercontent.com/u/3014297?v=4" width="100px;"/><br /><sub><b>Data</b></sub>](https://github.com/gnuos)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=gnuos "Code") | [<img src="https://avatars0.githubusercontent.com/u/2876745?v=4" width="100px;"/><br /><sub><b>侯名</b></sub>](https://github.com/KISSMonX)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=KISSMonX "Code") |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| [<img src="https://avatars0.githubusercontent.com/u/12060175?v=4" width="100px;"/><br /><sub><b>dumliu01</b></sub>](https://github.com/dumliu01)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=dumliu01 "Code") | [<img src="https://avatars0.githubusercontent.com/u/1411282?v=4" width="100px;"/><br /><sub><b>hlily2005</b></sub>](https://github.com/hlily2005)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=hlily2005 "Code") | [<img src="https://avatars0.githubusercontent.com/u/16982786?v=4" width="100px;"/><br /><sub><b>fenggolang</b></sub>](https://github.com/fenggolang)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=fenggolang "Code") | [<img src="https://avatars3.githubusercontent.com/u/10174178?v=4" width="100px;"/><br /><sub><b>henrylee2cn</b></sub>](https://github.com/henrylee2cn)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=henrylee2cn "Code") | [<img src="https://avatars0.githubusercontent.com/u/1336914?v=4" width="100px;"/><br /><sub><b>shaqsnake</b></sub>](https://github.com/shaqsnake)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=shaqsnake "Code") | [<img src="https://avatars0.githubusercontent.com/u/5728787?v=4" width="100px;"/><br /><sub><b>tbwisk</b></sub>](https://github.com/TBWISK)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=TBWISK "Code") | [<img src="https://avatars3.githubusercontent.com/u/416141?v=4" width="100px;"/><br /><sub><b>Huang ChuanTong</b></sub>](https://github.com/toontong)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=toontong "Code") |
| [<img src="https://avatars3.githubusercontent.com/u/10513552?v=4" width="100px;"/><br /><sub><b>The notes of SQL optimize </b></sub>](https://github.com/zhongxuan123)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=zhongxuan123 "Code") | [<img src="https://avatars2.githubusercontent.com/u/29008269?v=4" width="100px;"/><br /><sub><b>zhouxinxin19920802</b></sub>](https://github.com/zhouxinxin19920802)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=zhouxinxin19920802 "Code") | [<img src="https://avatars2.githubusercontent.com/u/20811449?v=4" width="100px;"/><br /><sub><b>macaria</b></sub>](https://github.com/macaria)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=macaria "Code") | [<img src="https://avatars3.githubusercontent.com/u/15226239?v=4" width="100px;"/><br /><sub><b>Dennis</b></sub>](http://github.com/DennisMao)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=DennisMao "Code") | [<img src="https://avatars1.githubusercontent.com/u/2696746?v=4" width="100px;"/><br /><sub><b>orangleliu</b></sub>](http://blog.csdn.net/orangleliu)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=orangle "Code") | [<img src="https://avatars1.githubusercontent.com/u/21693162?v=4" width="100px;"/><br /><sub><b>HarbinZhang</b></sub>](https://github.com/HarbinZhang)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=HarbinZhang "Code") | [<img src="https://avatars1.githubusercontent.com/u/7344921?v=4" width="100px;"/><br /><sub><b>LiMingji</b></sub>](https://github.com/SwanSpouse)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=SwanSpouse "Code") |
| [<img src="https://avatars0.githubusercontent.com/u/22164927?v=4" width="100px;"/><br /><sub><b>wintersnow</b></sub>](https://mickey0524.github.io/)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=mickey0524 "Code") | [<img src="https://avatars2.githubusercontent.com/u/44076738?v=4" width="100px;"/><br /><sub><b>zhuzhenfeng</b></sub>](https://github.com/zhuzhenfeng-finogeeks)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=zhuzhenfeng-finogeeks "Code") | [<img src="https://avatars2.githubusercontent.com/u/6065007?v=4" width="100px;"/><br /><sub><b>徐佳军</b></sub>](https://xujiajun.cn)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=xujiajun "Code") | [<img src="https://avatars0.githubusercontent.com/u/6884499?v=4" width="100px;"/><br /><sub><b>nicho</b></sub>](https://github.com/NichoZhang)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=NichoZhang "Code") | [<img src="https://avatars1.githubusercontent.com/u/17244565?v=4" width="100px;"/><br /><sub><b>Weifeng Wang</b></sub>](https://www.btxiaowei.net)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=qclaogui "Code") | [<img src="https://avatars3.githubusercontent.com/u/6748475?v=4" width="100px;"/><br /><sub><b>John Deng</b></sub>](https://hiboot.hidevops.io)<br />[💻](https://github.com/developer-learning/reading-go/commits?author=john-deng "Code") |
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!
