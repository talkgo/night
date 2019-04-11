# [Go 夜读](https://reading.developerlearning.cn/)
[![Build Status](https://travis-ci.org/developer-learning/reading-go.svg?branch=master)](https://travis-ci.org/developer-learning/reading-go) [![Go Report Card](https://goreportcard.com/badge/github.com/developer-learning/reading-go)](https://goreportcard.com/report/github.com/developer-learning/reading-go) [![GitHub stars](https://img.shields.io/github/stars/developer-learning/reading-go.svg?label=Stars)](https://github.com/developer-learning/reading-go) [![GitHub forks](https://img.shields.io/github/forks/developer-learning/reading-go.svg?label=Fork)](https://github.com/developer-learning/reading-go) [![All Contributors](https://img.shields.io/badge/all_contributors-0-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/developer-learning/reading-go?status.svg)](http://godoc.org/github.com/developer-learning/reading-go) [![Coverage Status](https://coveralls.io/repos/github/developer-learning/reading-go/badge.svg?branch=master)](https://coveralls.io/github/developer-learning/reading-go?branch=master) [![GitHub issues](https://img.shields.io/github/issues/developer-learning/reading-go.svg?label=Issue)](https://github.com/developer-learning/reading-go/issues) [![license](https://img.shields.io/github/license/developer-learning/reading-go.svg)](https://github.com/developer-learning/reading-go/blob/master/LICENSE)

<img src="https://raw.githubusercontent.com/developer-learning/reading-go/master/static/images/2018-12-11-night-reading-go.jpg" width="400px;"/>

*其他语言版本: [Deutsch](README_DE.md), [English](README_EN.md), [简体中文](README.md).*

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/developer-learning/reading-go.svg)](https://starcharts.herokuapp.com/developer-learning/reading-go)

[reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/?u=developer-learning&r=reading-go)

Go 学习与分享：

- [Go 夜读](https://github.com/developer-learning/reading-go/labels/Go%20%E5%A4%9C%E8%AF%BB)

*每周约定一个晚上进行 Go 源码阅读，Go 新手可以先去这里看看 **[Go 学习之路](https://github.com/developer-learning/learning-golang)**。*
>阅读范畴：Go 标准包、开源项目。

- [每日阅读](https://github.com/developer-learning/reading-go/labels/%E6%AF%8F%E6%97%A5%E9%98%85%E8%AF%BB)

*你愿意来，那说明你想改变，你也保证自己能做到，那你就得理解一点：在这里，我们是共同付出，你不是吸收者，我也不是分享者，我们可以讨论，但是我们不能帮你坚持或者教你，带你。*

- [每日一问](https://github.com/developer-learning/reading-go/labels/%E6%AF%8F%E6%97%A5%E4%B8%80%E9%97%AE)

- [Go 项目实践](https://github.com/developer-learning/reading-go/labels/Go%20%E9%A1%B9%E7%9B%AE%E5%AE%9E%E8%B7%B5)
- [Gin 开发](https://github.com/developer-learning/reading-go/labels/Gin%20%E5%BC%80%E5%8F%91)

----

<br>

|![notification](/static/images/bell-outline-badged.svg)预告|
|:------------------:|
| 《k8s 源代码解析 - 调度器》 4月13日 John|
| github.com/golang/sync -> errgroup、syncmap 等 源码阅读 |

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
- [ ] golang/sync
- [ ] kubernetes

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

<img src="/static/images/wechat_reading_go.jpg" width="400px;"/>

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

- [如何参与贡献](https://github.com/developer-learning/reading-go/blob/master/CONTRIBUTING.md)
- *[Git Commit Message Conventions](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit?pref=2&pli=1#)*

## Contributors

我非常重视每一个对这个项目的贡献者，我会将贡献者列表更新到这里，目前只有提交 Pull Request 的小伙伴，但是贡献不仅仅如此，还可以包括提交 Issue 以及在社群中有所贡献的人。

贡献者自己可以提 PR ，方法如下：

- 安装 `npm install -g --save-dev all-contributors-cli`
- `sh gen_contributors.sh`

贡献类型有多种，比如："code", "ideas","review","talk","tutorial"，你可以在 `.all-contributorsrc` 中修改。

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!
