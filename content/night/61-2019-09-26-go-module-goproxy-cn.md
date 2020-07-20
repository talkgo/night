---
desc: Go 夜读 之 Go Modules、Go Module Proxy 和 goproxy.cn #468
title: 第 61 期 Go Modules、Go Module Proxy 和 goproxy.cn #468
date: 2019-09-26T21:00:00+08:00
author: 盛傲飞
---

## Go 夜读第 61 期 Go Modules、Go Module Proxy 和 goproxy.cn #468

## 内容简介

Go 1.11 推出的[模块（Modules）](https://github.com/golang/go/wiki/Modules)为 Go 语言开发者打开了一扇新的大门，理想化的依赖管理解决方案使得 Go 语言朝着计算机编程史上的第一个依赖乌托邦（Deptopia）迈进。随着模块一起推出的还有[模块代理协议（Module proxy protocol）](https://golang.org/cmd/go/#hdr-Module_proxy_protocol)，通过这个协议我们可以实现 Go 模块代理（Go module proxy），aka 依赖镜像。Go 1.13 的发布为模块带来了大量的改进，所以模块的扶正就是这次 Go 1.13 发布中开发者能直接感觉到的最大变化。Go 1.13 中的 `GOPROXY` 环境变量拥有了一个在中国大陆无法访问到的默认值 [proxy.golang.org](https://proxy.golang.org)，经过大家在 https://github.com/golang/go/issues/31755 中激烈的讨论（有些人甚至将话提上升到了“自由世界”的层次），最终 Go 核心团队仍然无法为中国开发者提供一个可在中国大陆访问的官方模块代理。为了今后中国的 Go 语言开发者能更好地进行开发，七牛云推出了非营利性项目 [goproxy.cn](https://goproxy.cn)，其目标是为中国和世界上其他地方的 Gopher 们提供一个免费的、可靠的、持续在线的且经过 CDN 加速的模块代理。可以预见未来是属于模块化的，所以 Go 语言开发者能越早切入模块就能越早进入未来。如果说 Go 1.11 和 Go 1.12 时由于模块的不完善你不愿意切入，那么 Go 1.13 你则可以大胆地开始放心使用。本次分享将讨论如何使用模块和模块代理，以及在它们的使用中会常遇见的坑，还会讲解如何快速搭建自己的私有模块代理，并简单地介绍一下七牛云推出的 [goproxy.cn](https://goproxy.cn) 以及它的出现对于中国 Go 语言开发者来说重要在何处。

## 内容大纲

* Go Modules 简介
* 快速迁移项目至 Go Modules
* 使用 Go Modules 时常遇见的坑
	* 坑 1：判断项目是否启用了 Go Modules
	* 坑 2：管理 Go 的环境变量
	* 坑 3：从 dep、glide 等迁移至 Go Modules
	* 坑 4：拉取私有模块
	* 坑 5：更新现有的模块
	* 坑 6：主版本号
* Go Module Proxy 简介
* 快速搭建私有 Go Module Proxy
* Goproxy 中国（goproxy.cn）

## 分享者自我介绍

盛傲飞，两个多月前刚本科毕业，目前刚毕业旅行归来还未找工作，从事 Go 语言开发 4 年左右，开源爱好者，[goproxy.cn](https://goproxy.cn) 的作者。

## 分享时间

2019-09-26 21:00 UTC+8

## 分享地址

https://zoom.us/j/6923842137

## 录播地址

Bilibili（音视频不同步，正在修复……）：https://www.bilibili.com/video/av69111199

YouTube：https://youtu.be/H3LVVwZ9zNY

## Slides

https://docs.google.com/presentation/d/1LRs_D-IlrSU-ZrJN7KiBhNmCY5tWXH1b2CyDF7jvClY/edit?usp=sharing

## 参考资料

* https://github.com/golang/go/wiki/Modules
* https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more
* https://blog.golang.org/versioning-proposal
* https://go.googlesource.com/proposal/+/master/design/25530-sumdb.md
* https://youtu.be/F8nrpe0XWRg
* https://golang.org/doc/go1.11#modules
* https://blog.golang.org/modules2019
* https://golang.org/doc/go1.12#modules
* https://blog.golang.org/using-go-modules
* https://blog.golang.org/migrating-to-go-modules
* https://blog.golang.org/module-mirror-launch
* https://golang.org/doc/go1.13#modules
* https://studygolang.com/topics/9994
* https://studygolang.com/topics/10014

## 直播中的 Q&A 环节内容

**问：如何解决 Go 1.13 在从 GitLab 拉取模块版本时遇到的，Go 错误地按照非期望值的路径寻找目标模块版本结果致使最终目标模块拉取失败的问题？**

答：GitLab 中配合 `go get` 而设置的 `<meta>` 存在些许问题，导致 Go 1.13 错误地识别了模块的具体路径，这是个 Bug，据说在 GitLab 的新版本中已经被修复了，详细内容可以看 https://github.com/golang/go/issues/34094 这个 Issue。然后目前的解决办法的话除了升级 GitLab 的版本外，还可以参考 https://github.com/talkgo/night/issues/468#issuecomment-535850154 这条回复。

**问：使用 Go modules 时可以同时依赖同一个模块的不同的两个或者多个小版本（修订版本号不同）吗？**

答：不可以的，Go modules 只可以同时依赖一个模块的不同的两个或者多个大版本（主版本号不同）。比如可以同时依赖 `example.com/foobar@v1.2.3` 和 `example.com/foobar/v2@v2.3.4`，因为他们的模块路径（module path）不同，Go modules 规定主版本号不是 `v0` 或者 `v1` 时，那么主版本号必须显式地出现在模块路径的尾部。但是，同时依赖两个或者多个小版本是不支持的。比如如果模块 A 同时直接依赖了模块 B 和模块 C，且模块 A 直接依赖的是模块 C 的 `v1.0.0` 版本，然后模块 B 直接依赖的是模块 C 的 `v1.0.1` 版本，那么最终 Go modules 会为模块 A 选用模块 C 的 `v1.0.1` 版本而不是模块 A 的 `go.mod` 文件中指明的 `v1.0.0` 版本。这是因为，Go modules 认为只要主版本号不变，那么剩下的都可以直接升级采用最新的。但是如果采用了最新的结果导致项目 Break 掉了，那么 Go modules 就会 Fallback 到上一个老的版本，比如在前面的例子中就会 Fallback 到 `v1.0.0` 版本。

**问：在 `go.sum` 文件中的一个模块版本的 Hash 校验数据什么情况下会成对出现，什么情况下只会存在一行？**

答：通常情况下，在 `go.sum` 文件中的一个模块版本的 Hash 校验数据会有两行，前一行是该模块的 ZIP 文件的 Hash 校验数据，后一行是该模块的 `go.mod` 文件的 Hash 校验数据。但是也有些情况下只会出现一行该模块的 `go.mod` 文件的 Hash 校验数据，而不包含该模块的 ZIP 文件本身的 Hash 校验数据，这个情况发生在 Go modules 判定为你当前这个项目完全用不到该模块，根本也不会下载该模块的 ZIP 文件，所以就没必要对其作出 Hash 校验保证，只需要对该模块的 `go.mod` 文件作出 Hash 校验保证即可，因为 `go.mod` 文件是用得着的，在深入挖取项目依赖的时候要用。

**问：能不能更详细地讲解一下 `go.mod` 文件中的 `replace` 动词的行为以及用法？**

答：这个 `replace` 动词的作用是把一个“模块版本”替换为另外一个“模块版本”，这是“模块版本”和“模块版本（module path）”之间的替换，“=>”标识符前面的内容是待替换的“模块版本”的“模块路径”，后面的内容是要替换的目标“模块版本”的所在地，即路径，这个路径可以是一个本地磁盘的相对路径，也可以是一个本地磁盘的绝对路径，还可以是一个网络路径，但是这个目标路径并不会在今后你的项目代码中作为你“导入路径（import path）”出现，代码里的“导入路径”还是得以你替换成的这个目标“模块版本”的“模块路径”作为前缀。注意，Go modules 是不支持在“导入路径”里写相对路径的。举个例子，如果项目 A 依赖了模块 B，比如模块 B 的“模块路径”是 `example.com/b`，然后它在的磁盘路径是 `~/b`，在项目 A 里的 `go.mod` 文件中你有一行 `replace example.com/b => ~/b`，然后在项目 A 里的代码中的“导入路径”就是 `import "example.com/b"`，而不是 `import "~/b"`，剩下的工作是 Go modules 帮你自动完成了的。然后就是我在分享中也提到了，`exclude` 和 `replace` 这两个动词只作用于当前主模块，也就是当前项目，它所依赖的那些其他模块版本中如果出现了你待替换的那个模块版本的话，Go modules 还是会为你依赖的那个模块版本去拉取你的这个待替换的模块版本。比如项目 A 直接依赖了模块 B 和模块 C，然后模块 B 也直接依赖了模块 C，那么你在项目 A 中的 `go.mod` 文件里的 `replace c => ~/some/path/c` 是只会影响项目 A 里写的代码中，而模块 B 所用到的还是你 `replace` 之前的那个 `c`，并不是你替换成的 `~/some/path/c` 这个。


---

## 观看视频

{{< youtube id="H3LVVwZ9zNY" >}}
