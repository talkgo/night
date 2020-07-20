# [Go 夜讀](https://talkgo.org/)
[![Build Status](https://travis-ci.org/talkgo/night.svg?branch=master)](https://travis-ci.org/talkgo/night) [![Go Report Card](https://goreportcard.com/badge/github.com/talkgo/night)](https://goreportcard.com/report/github.com/talkgo/night) [![GitHub stars](https://img.shields.io/github/stars/talkgo/night.svg?label=Stars)](https://github.com/talkgo/night) [![GitHub forks](https://img.shields.io/github/forks/talkgo/night.svg?label=Fork)](https://github.com/talkgo/night) [![All Contributors](https://img.shields.io/badge/all_contributors-0-orange.svg?style=flat-square)](#contributors) [![Documentation](https://godoc.org/github.com/talkgo/night?status.svg)](http://godoc.org/github.com/talkgo/night) [![Coverage Status](https://coveralls.io/repos/github/talkgo/night/badge.svg?branch=master)](https://coveralls.io/github/talkgo/night?branch=master) [![GitHub issues](https://img.shields.io/github/issues/talkgo/night.svg?label=Issue)](https://github.com/talkgo/night/issues) [![license](https://img.shields.io/github/license/talkgo/night.svg)](https://github.com/talkgo/night/blob/master/LICENSE)

<img src="https://raw.githubusercontent.com/talkgo/night/master/static/images/2018-12-11-night-reading-go.jpg" width="400px;"/>

#### *以[其他語言](Translations.md)閱讀*

[🇨🇳](README.md)
[🇭🇰](README.cht.md)
[🇺🇸](README.en.md)
[🇩🇪](README.de.md)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/talkgo/night.svg)](https://starcharts.herokuapp.com/talkgo/night)

[reading-go Star History and Stats](https://seladb.github.io/StarTrack-js/#/preload?r=developer-learning,night-reading-go)

Go 學習與分享：

- [Go 夜讀](https://github.com/talkgo/night/labels/Go%20%E5%A4%9C%E8%AF%BB)

*每週約定一個晚上進行 Go 源碼閱讀，Go 新手可以先去這裡看看 **[Go 學習之路](https://github.com/developer-learning/learning-golang)**。 *
>閱讀範疇：Go 標準包、開源項目。

- [每日閱讀](https://github.com/talkgo/night/labels/%E6%AF%8F%E6%97%A5%E9%98%85%E8%AF%BB)

*你願意來，那說明你想改變，你也保證自己能做到，那你就得理解一點：在這裡，我們是共同付出，你不是吸收者，我也不是分享者，我們可以討論，但是我們不能幫你堅持或者教你，帶你。*

- [每日一問](https://github.com/talkgo/night/labels/%E6%AF%8F%E6%97%A5%E4%B8%80%E9%97%AE)

- [Go 項目實踐](https://github.com/talkgo/night/labels/Go%20%E9%A1%B9%E7%9B%AE%E5%AE%9E%E8%B7%B5)
- [Gin 開發](https://github.com/talkgo/night/labels/Gin%20%E5%BC%80%E5%8F%91)

----

<br>

|![notification](/static/images/bell-outline-badged.svg)預告|
|:------------------:|
| k8s 源代碼解析 - 調度器 4月13日 John|
| github.com/golang/sync -> errgroup、syncmap 等 源碼閱讀 |

----

## 閱讀清單

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

- [Go 夜讀(YouTuBe)](https://youtube.com/c/talkgo_night)
- [Go 夜讀(Bilibili)](https://space.bilibili.com/326749661)
- [Go 夜讀(Tencent)](https://v.qq.com/vplus/e05f55c8ca5e36f8e370ba49c9e883e0)

## 我們的目標

我們希望可以推進大家深入了解 Go ，快速成長為資深的 Gopher 。我們希望每次來了的人和沒來的人都能夠有收穫，成長。

## 我们的方式

由一個主講人帶著大家一起去閱讀 Go 源代碼，一起去啃那些難啃的算法、學習代碼裡面的奇淫技巧，遇到問題或者有疑惑了，我們可以一起去檢索，解答這些問題。我們可以一起學習，共同成長。

>閱讀規則：選取 package 包，然後從上往下開始讀 xxx.go 文件，每個文件從上往下讀導出的函數（一步一步跟邏輯，如果邏輯跳出這個 package 則不做深入探究）。

## 我們的精神

開源！開源！開源！重要的事，一定要說三遍。

>希望有興趣的小伙伴們一起加入，讓我們一起把 『Go 夜讀』建立成一個對大家都有幫助的開源社區。

## 怎麼加入

<img src="/static/images/wechat_reading_go.jpg" width="400px;"/>

如果你想加入微信群，請搜索 `night_reading_go` ，然後備註你的姓名、公司、工作崗位和職責，備註來源：Github。

有同學想要用 Slack 交流，我開放了一個：[Talk Go Slack](https://join.slack.com/t/talkgo/shared_invite/zt-89zh1000-KX2tZ6l~FSNP14Oy2B~onQ)

## Night Reading Go SIG members

>Thanks for you effort.

- [杨文 yangwenmai](https://github.com/yangwenmai)
- [欧长坤 changkun](https://github.com/changkun)
- [FelixSeptem](https://github.com/FelixSeptem)
- [饶全成 qcrao](https://github.com/qcrao)
- [煎鱼 - EDDYCJY](https://github.com/EDDYCJY)
- [盛傲飞 - aofei](https://github.com/aofei)

----

## 如何參與貢獻？

```sh
├── reading   // Go 源碼閱讀
├── discuss   // 日常微信群討論的總結
├── articles  // 個人原創的技術文章
├── interview // Go 面試專區
└── other     // 其他
```

- [如何參與貢獻](https://github.com/talkgo/night/blob/master/CONTRIBUTING.md)
- *[Git Commit 規範指南](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit?pref=2&pli=1#)*

## Contributors

我非常重視每一個對這個項目的貢獻者，我會將貢獻者列表更新到這裡，目前只有提交 Pull Request 的小夥伴，但是貢獻不僅僅如此，還可以包括提交 Issue 以及在社群中有所貢獻的人。

貢獻者自己可以提PR，方法如下：

- 安裝 `npm install -g --save-dev all-contributors-cli`
- `sh gen_contributors.sh`

貢獻類型有多種，比如："code", "ideas","review","talk","tutorial"，你可以在 `.all-contributorsrc` 中修改。

Thanks goes to these wonderful people ([emoji key](https://github.com/kentcdodds/all-contributors#emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore -->
<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/kentcdodds/all-contributors) specification. Contributions of any kind welcome!
