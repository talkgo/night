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

## 直播中的文字交流内容

<details>
<summary>点击展开</summary>

```
00:06:46	shouge:	你好
00:37:28	tangyinpeng:	盛神好
00:37:31	tangyinpeng:	紧张不
00:37:48	盛傲飞:	可紧张了，我怎么连接不到语音啊
00:38:09	tangyinpeng:	进来的时候有电脑音频的
00:38:39	tangyinpeng:	左下角
00:38:43	tangyinpeng:	你看看
00:38:55	tangyinpeng:	解除静音
00:39:37	煎鱼:	11
00:39:39	煎鱼:	哟了
00:39:41	煎鱼:	有了
00:39:41	shouge:	可以
00:39:42	tangyinpeng:	有了
00:40:00	煎鱼:	11
00:40:11	cibudayi:	有了
00:40:17	煎鱼:	有的
00:40:30	煎鱼:	嗯嗯，等多一会吧。
00:40:47	shouge:	第一次来现场  以前都是看的回放 
00:40:52	mai yang:	再等一下开始
00:40:59	煎鱼:	哈哈淡定，等等到点再开始。
00:42:53	tangyinpeng:	有点
00:42:54	煎鱼:	是有点杂。
00:42:59	Kevin Bai:	有点
00:43:15	tangyinpeng:	有线耳机有木有
00:43:18	tangyinpeng:	苹果的还行
00:43:25	tangyinpeng:	好man
00:43:30	Kevin Bai:	嗡嗡的
00:43:31	cibudayi:	好多了
00:43:32	liuxuan:	有环境音效了
00:44:40	iPhone:	不开视频吗
00:44:53	煎鱼:	还行，我都听得清楚。
00:44:58	tangyinpeng:	还行
00:45:02	iPhone:	性感傲飞 在线直播 哈哈
00:45:24	蘇醒若蘅:	有点断断续续
00:45:34	Kevin Bai:	声音够大 但是有点嗡嗡的感觉
00:45:46	mai yang:	可以开始了
00:46:34	煎鱼:	可以开始了。
00:46:44	煎鱼:	不慌。
00:48:00	shouge:	我开始以为你是七牛的
00:48:02	shouge:	哈哈哈
00:49:20	hanyouqing:	没有声音？
00:49:46	tangyinpeng:	有
00:50:01	hanyouqing:	呃，我退出重进下试试，谢谢
00:50:27	热心市民孙先生:	改一下声音输出设备
00:55:21	cluas:	声音无了 是我卡了吗
00:55:27	cluas:	又有了
00:55:28	Wang:	1
00:55:35	Lx:	有吗？
00:55:35	hanyu:	1
00:55:39	煎鱼:	1
00:55:43	锦锐:	1
00:56:36	hanyouqing:	我网络不好····，请问后面会分享PPT和视频么～
00:56:55	shouge:	有
00:57:11	hanyouqing:	通过什么途径发出来
00:57:18	Lx:	会有录播,ppt 在 google 上。。另外 你们真的有声音吗？
00:57:19	shouge:	b站 youtube都有回放
00:57:36	Quintin Zhang:	有声音啊
00:57:41	shouge:	有声音的啊 
00:57:42	Quintin Zhang:	你用什么客户端的
00:57:47	cluas:	有声音
00:57:50	Quintin Zhang:	https://meetzoom.net/client/latest/ZoomInstaller.exe
00:57:59	Quintin Zhang:	下载这个
00:58:05	Quintin Zhang:	windows的
00:58:14	cluas:	就是信号一般 跟电话信号不好的时候差不多
00:58:38	Quintin Zhang:	不用爬，
00:58:48	Quintin Zhang:	直接下载这个包就可以中国使用
00:58:49	Kevin Bai:	不好意思发错了 quintin
00:59:01	Kevin Bai:	估计是翻墙的，网络不好
00:59:11	Quintin Zhang:	我刚才发的是国内版本
00:59:27	Quintin Zhang:	可以在国内使用 把 vpn 关掉
00:59:39	cluas:	比看录像好多了 录像模糊到怀疑人生
00:59:43	煎鱼:	666
01:00:31	Quintin Zhang:	本地录制 效果取决于 本人的网络。。
01:00:52	Quintin Zhang:	https://meetzoom.net/download
01:01:10	oliverch:	回看录像 体验很糟糕 上一期的音画不同步很严重
01:01:25	Quintin Zhang:	用梯子的 质量不好的 可以关闭 梯子，使用我发的link 下载client
01:01:27	Quintin Zhang:	https://meetzoom.net/download
01:03:26	cluas:	mac 可以吗
01:03:40	cluas:	不开梯子
01:03:41	Quintin Zhang:	https://meetzoom.net/download
01:03:44	Quintin Zhang:	可以的
01:03:50	Quintin Zhang:	所有平台都可以了
01:04:02	cluas:	这里跟.us域名下下的是同一个apk
01:04:05	cluas:	pkg
01:04:26	cluas:	我关梯子试试
01:04:31	Quintin Zhang:	不一样的 国内要求实名制
01:04:55	tangyinpeng:	没有
01:04:58	煎鱼:	111
01:04:59	郭源:	没
01:05:12	煎鱼:	我们这还看得到。
01:05:17	tangyinpeng:	不开梯子进不来的
01:05:25	Ricardo:	能看到
01:05:25	JACK:	可以看到
01:05:42	Quintin Zhang:	https://meetzoom.net/download
01:05:49	Quintin Zhang:	用这个里面的客户端
01:05:52	faultzone:	开始了？
01:05:55	Quintin Zhang:	绝对可以
01:06:02	faultzone:	好难呀
01:06:08	cluas:	我梯子关了
01:06:08	cluas:	进来了 音效好多了 没卡顿了
01:06:17	Quintin Zhang:	肯定的
01:06:48	Quintin Zhang:	sorry 最近我们正在努力解决 zoom 中国的问题
01:14:46	煎鱼:	111
01:14:47	煎鱼:	在
01:14:54	shouge:	没问题 
01:26:36	JACK:	PPT会分享吗？
01:27:59	煎鱼:	分享结束后会公开PPT。
01:28:40	oliverch:	sum.golang.google.cn 这个sumdb 的可以用
01:31:21	tangyinpeng:	没声音了
01:31:24	tangyinpeng:	我的问题吗
01:31:36	佳飞 邱:	我是正常
01:31:39	煎鱼:	我有
01:31:41	JACK:	我有
01:32:03	mai yang:	你自己检查一下网络
01:33:03	“Joey”的 iPhone:	不好意思，晚到了和goproxy.io有什么区别
01:35:32	cluas:	可能代码里带了宣传视频
01:35:54	oliverch:	/强
01:38:01	mai yang:	👍七牛云和傲飞
01:38:34	煎鱼:	11
01:38:41	tangyinpeng:	字体小
01:38:49	haohongfan:	确实有点小
01:38:54	haohongfan:	1
01:38:54	Kevin Bai:	太小了
01:38:55	tangyinpeng:	再大点
01:38:57	cjh:	再大点.
01:39:03	tangyinpeng:	大大大
01:39:22	tangyinpeng:	差不多了
01:39:25	tangyinpeng:	小
01:39:27	tangyinpeng:	网页
01:48:14	cluas:	同vim党  你的tmux加vim不卡吗…
01:48:19	煎鱼:	6666
01:48:27	shouge:	不卡啊
01:48:30	alphababy:	确实快
01:48:36	cluas:	终端用的哪个
01:48:45	weiluoliang:	vim大佬
01:48:45	Quintin Zhang:	都是大佬。。这样会带坏小朋友的。。
01:48:51	alphababy:	Iterm2 吧
01:49:04	weiluoliang:	我也用iTerm2
01:49:06	shouge:	还好吧 可能安装了什么插件
01:49:12	cluas:	我用自带的 好卡 尤其是向下滚屏的时候
01:49:23	weiluoliang:	自带的不好用
01:49:47	shouge:	Neovim 
01:50:05	cluas:	我是nvim
01:50:19	cluas:	在tmux里会卡 在终端直接开vim不会
02:03:58	shouge:	不是会删除 cache嘛？
02:04:04	cluas:	来讲下终端呀
02:04:09	cluas:	我的终端巨卡
02:04:12	tangyinpeng:	大家提问
02:04:14	煎鱼:	大家有问题的可以说。
02:04:15	tangyinpeng:	或者语音
02:04:23	houruxin:	有个问题想问问
02:04:49	alphababy:	大佬 分享下你的go多版本管理呀
02:04:52	tangyinpeng:	听得见
02:05:10	shouge:	go  最新的版本 
02:06:49	hanyu:	也可以发微信群里
02:06:51	644262163:	是不是公司的gitlab版本比较低
02:06:52	煎鱼:	有截图么，报错信息什么的。
02:06:55	tangyinpeng:	聊天框打字啥的
02:07:53	cluas:	dotfile链接可以发下吗
02:08:09	houruxin:	go get gitlab.51y5.net/houruxin/EtcdClient
go: finding gitlab.51y5.net/houruxin/EtcdClient latest
go get gitlab.51y5.net/houruxin/EtcdClient: git ls-remote -q https://gitlab.51y5.net/houruxin.git in /Users/houruxin/Documents/golang/pkg/mod/cache/vcs/1987238e5f8450887f471b9feb5b0cad955ee995627ce1eed61f207f66d0e356: exit status 128:
	GitLab: The project you were looking for could not be found.
	fatal: Could not read from remote repository.

	Please make sure you have the correct access rights
	and the repository exists.
这里EtcdClient是一个我自己打包的内部repo，然后可以看到这个go get已经找到了这个repo，但是在后面go get的时候它会先去git ls-remote -q https://gitlab.51y5.net/houruxin.git， 但是houruxin这个只是内部gitlab中的一个域名，并不是一个项目，所以这个肯定会报错。但是关闭GO111MODULE，通过老的方式做go get，又不会出错 
02:08:15	王克亚的 iPhone:	我想问下go语言一般怎么打包发布的
02:08:46	644262163:	gitlab是subgroup有问题 11.7修复的
02:09:45	oliverch:	看错误信息 是 权限的问题吧
02:10:48	oliverch:	同时依赖一个依赖库的不同小版本 怎么搞
02:12:01	煎鱼:	example.com/apple v0.1.2 h1:哈希值
example.com/apple v0.1.2/go.mod h1:哈希值

刚刚有提到 go.sum 下，apple v0.1.2 h1 和 apple v0.1.2/go.mod h1 不一定是成对出现，那什么情况下有前者，又或是只有后者，又或是什么情况下两个会同时都出现？
02:12:03	oliverch:	了解了
02:13:21	煎鱼:	了解。
02:13:43	煎鱼:	Go Modules 目前来看，官方还有其他规划么。
02:13:52	佳飞 邱:	go get 的类库会缓存在本地一份吗？
02:15:33	cjh:	如果我的GOPATH=/aaa:/bbb 这个时候是存储在/aaa/pkg目录下面吗?
02:16:26	alphababy:	啥网站呀？可以发个域名？
02:17:02	oliverch:	https://gfw.go101.org/article/101.html go101
02:17:06	shouge:	https://go101.org/
02:17:09	alphababy:	thx
02:17:19	王克亚的 iPhone:	你们都刷leetcode算法吗，感觉好难
02:17:54	alphababy:	你去耍耍acm，就会发现lc简单好多
02:17:59	shouge:	有时间会刷 
02:18:27	oliverch:	算法 好浪费时间啊
02:18:31	alphababy:	Acm 太累了，扛不住
02:18:36	oliverch:	顶不住
02:18:45	alphababy:	算法才是灵魂，个人看法
02:18:49	王克亚的 iPhone:	哈哈
02:19:58	shouge:	哈哈哈
02:20:01	alphababy:	我也太菜了
02:20:12	weiluoliang:	docker
02:20:26	oliverch:	太过分了/哈哈
02:20:35	oliverch:	k8s/docker etcd
02:21:35	alphababy:	遇到什么学什么，我室友就是这样
02:21:41	IPhone7:	Golang源码有深入阅读吗
02:22:16	王克亚的 iPhone:	对GPM有深入了解过没
02:22:32	shouge:	at 错了 我是想发公屏的 
02:22:38	alphababy:	太深入就是理论了，了解了解进行了，
02:22:43	Quintin Zhang:	兴趣引发需求 需求驱动兴趣
02:22:55	cluas:	tmux
02:22:57	cluas:	哈哈哈哈
02:22:58	alphababy:	删掉抖音
02:23:26	cluas:	alfred
02:23:28	oliverch:	用 Google 代替 Baidu
02:23:42	alphababy:	安利一个 软件 叫 Mos，上github上可以搜到
02:23:45	Kevin Bai:	Dash 还是不错的，除了升级大版本时
02:23:47	cluas:	离线文档
02:24:57	alphababy:	fzf 安利
02:25:00	hawken:	zsh
02:25:01	shouge:	自由开发者 
02:25:05	shouge:	好舒服 
02:25:18	oliverch:	vscode 的远程开发 感觉okey
02:25:21	cluas:	linux系统上的vim体验 优于mac
02:25:25	cluas:	亲身体验
02:25:46	shouge:	zsh tmux vim
02:25:51	shouge:	三剑客
02:27:07	cluas:	而且Alfred用的就是它的缓存
02:27:16	cluas:	不过可定制化强一点
02:27:35	shouge:	 感谢分享 收拾收拾先下班了 
02:27:50	cluas:	辛苦
02:27:55	alphababy:	真实，996实锤了
02:28:06	cluas:	我们7点下班
02:28:12	cluas:	还是比较养老
02:28:13	煎鱼:	辛苦了
02:28:17	shouge:	羡慕自由开发者啊
02:28:24	mai yang:	辛苦了
02:28:26	alphababy:	多版本的管理切换时怎么样子的
02:28:27	IPhone7:	怎么进群
02:28:31	alphababy:	分享下被
02:28:45	煎鱼:	内推你免试啊哈哈哈。
02:28:54	shouge:	饶哥 
02:28:58	mai yang:	mai_yang 微信
02:29:01	shouge:	经常看你文章
02:29:11	王克亚的 iPhone:	多谢
02:29:22	shouge:	我也姓饶  哈哈哈 
02:29:28	煎鱼:	感谢。
02:29:33	cluas:	谢谢
02:29:34	全成 饶:	大家好多问题都是从夜读 github 上看到的
02:29:35	佳飞 邱:	感谢分享
02:29:38	tangyinpeng:	感谢
02:29:40	cjh:	谢谢分享.
02:29:44	alphababy:	感谢
02:29:45	tangyinpeng:	辛苦了
02:29:50	oliverch:	辛苦了
02:29:50	shouge:	感谢分享 
02:29:52	Quintin Zhang:	感谢
02:29:52	gopher:	谢谢 辛苦了
02:29:55	煎鱼:	Gl.
02:29:55	weiluoliang:	感谢分享
02:30:00	chrisSun:	感谢
02:30:00	bocai:	感谢分享
02:30:01	煎鱼:	对
```

</details>

## 直播中的 Q&A 环节内容

**问：如何解决 Go 1.13 在从 GitLab 拉取模块版本时遇到的，Go 错误地按照非期望值的路径寻找目标模块版本结果致使最终目标模块拉取失败的问题？**

答：GitLab 中配合 `go get` 而设置的 `<meta>` 存在些许问题，导致 Go 1.13 错误地识别了模块的具体路径，这是个 Bug，据说在 GitLab 的新版本中已经被修复了，详细内容可以看 https://github.com/golang/go/issues/34094 这个 Issue。然后目前的解决办法的话除了升级 GitLab 的版本外，还可以参考 https://github.com/developer-learning/night-reading-go/issues/468#issuecomment-535850154 这条回复。

**问：使用 Go modules 时可以同时依赖同一个模块的不同的两个或者多个小版本（修订版本号不同）吗？**

答：不可以的，Go modules 只可以同时依赖一个模块的不同的两个或者多个大版本（主版本号不同）。比如可以同时依赖 `example.com/foobar@v1.2.3` 和 `example.com/foobar/v2@v2.3.4`，因为他们的模块路径（module path）不同，Go modules 规定主版本号不是 `v0` 或者 `v1` 时，那么主版本号必须显式地出现在模块路径的尾部。但是，同时依赖两个或者多个小版本是不支持的。比如如果模块 A 同时直接依赖了模块 B 和模块 C，且模块 A 直接依赖的是模块 C 的 `v1.0.0` 版本，然后模块 B 直接依赖的是模块 C 的 `v1.0.1` 版本，那么最终 Go modules 会为模块 A 选用模块 C 的 `v1.0.1` 版本而不是模块 A 的 `go.mod` 文件中指明的 `v1.0.0` 版本。这是因为，Go modules 认为只要主版本号不变，那么剩下的都可以直接升级采用最新的。但是如果采用了最新的结果导致项目 Break 掉了，那么 Go modules 就会 Fallback 到上一个老的版本，比如在前面的例子中就会 Fallback 到 `v1.0.0` 版本。

**问：在 `go.sum` 文件中的一个模块版本的 Hash 校验数据什么情况下会成对出现，什么情况下只会存在一行？**

答：通常情况下，在 `go.sum` 文件中的一个模块版本的 Hash 校验数据会有两行，前一行是该模块的 ZIP 文件的 Hash 校验数据，后一行是该模块的 `go.mod` 文件的 Hash 校验数据。但是也有些情况下只会出现一行该模块的 `go.mod` 文件的 Hash 校验数据，而不包含该模块的 ZIP 文件本身的 Hash 校验数据，这个情况发生在 Go modules 判定为你当前这个项目完全用不到该模块，根本也不会下载该模块的 ZIP 文件，所以就没必要对其作出 Hash 校验保证，只需要对该模块的 `go.mod` 文件作出 Hash 校验保证即可，因为 `go.mod` 文件是用得着的，在深入挖取项目依赖的时候要用。

**问：能不能更详细地讲解一下 `go.mod` 文件中的 `replace` 动词的行为以及用法？**

答：这个 `replace` 动词的作用是把一个“模块版本”替换为另外一个“模块版本”，这是“模块版本”和“模块版本（module path）”之间的替换，“=>”标识符前面的内容是待替换的“模块版本”的“模块路径”，后面的内容是要替换的目标“模块版本”的所在地，即路径，这个路径可以是一个本地磁盘的相对路径，也可以是一个本地磁盘的绝对路径，还可以是一个网络路径，但是这个目标路径并不会在今后你的项目代码中作为你“导入路径（import path）”出现，代码里的“导入路径”还是得以你替换成的这个目标“模块版本”的“模块路径”作为前缀。注意，Go modules 是不支持在“导入路径”里写相对路径的。举个例子，如果项目 A 依赖了模块 B，比如模块 B 的“模块路径”是 `example.com/b`，然后它在的磁盘路径是 `~/b`，在项目 A 里的 `go.mod` 文件中你有一行 `replace example.com/b => ~/b`，然后在项目 A 里的代码中的“导入路径”就是 `import "example.com/b"`，而不是 `import "~/b"`，剩下的工作是 Go modules 帮你自动完成了的。然后就是我在分享中也提到了，`exclude` 和 `replace` 这两个动词只作用于当前主模块，也就是当前项目，它所依赖的那些其他模块版本中如果出现了你待替换的那个模块版本的话，Go modules 还是会为你依赖的那个模块版本去拉取你的这个待替换的模块版本。比如项目 A 直接依赖了模块 B 和模块 C，然后模块 B 也直接依赖了模块 C，那么你在项目 A 中的 `go.mod` 文件里的 `replace c => ~/some/path/c` 是只会影响项目 A 里写的代码中，而模块 B 所用到的还是你 `replace` 之前的那个 `c`，并不是你替换成的 `~/some/path/c` 这个。


---

## 观看视频

{{< youtube id="H3LVVwZ9zNY" >}}