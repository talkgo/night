---
desc: Go 夜读之 Go&WebAssembly 简介
title: 第 55 期 Go&WebAssembly 简介
date: 2019-08-15T21:00:00+08:00
author: 柴大
---

## Go 夜读第 55 期 Go&WebAssembly 简介

WebAssembly 简介
WebAssembly 是一种新兴的网页虚拟机标准，它的设计目标包括：高可移植性、高安全性、高效率（包括载入效率和运行效率）、尽可能小的程序体积。

根据 Ending 定律：⼀切可被编译为 WebAssembly 的，终将被编译为 WebAssembly。

本次分享 Go&WebAssembly 相关的用法。

## 分享时间

2019-08-15 21:00:00

## 分享平台

[zoom 在线直播 - https://zoom.us/j/6923842137](https://zoom.us/j/6923842137)

## 更多讨论

FelixSeptem：补充一下 go 官方给的 wiki https://github.com/golang/go/wiki/WebAssembly 以及
WebAssembly 官网 https://webassembly.org/
个人比较倾向于对于 https://github.com/gopherjs/gopherjs 来比较理解，相对于 go->js (包括 react 等等) 的方案，WebAssembly 带来的异同是什么？

chai2010 ：@FelixSeptem wasm 和 gopherjs 最大的差异：wasm 是官方支持，同时 wasm 是国际标准是其它语言认可的中间格式。

以前虽然很多工具输出 js，那是因为没有 wasm 可以选择。
现在有了 wasm，大家肯定只支持 wasm 而逐渐弱化 js 的支持。
毕竟 wasm 虚拟机实现比 v8 简单多了，性能又可以秒杀 js。

wasm 最大的潜力是在浏览器之外，甚至可以想象成一个轻量化的 Docker 环境。
我觉得这个才是 wasm 真正有意思的地方，wasm 对于 js 完全属于降维打击。

changkun：还没有在生产环境使用过 wasm。从给的马里奥的例子来看，go wasm 本质上是分发由 Go 编译好 .wasm，而 Go 端的本质就是提供了一些能够解释为 wasm 的 utils。不太清楚会不会在分享中提及这一本质。

长远来看，这个 .wasm 文件在特性支持的情况下最终会包含完整的 Go 运行时，
但 go wasm 并没有明确在 web 场景下为什么一定需要它，当然不可否认它的确为兼容并移植 Go 代码来发展 web 应用带来了便捷，但前提是我们必须有足够多基础设施是基于 Go 的，但游戏并没有，非常希望看到一些能够说服用 Go 写 wasm 而不是其他语言（C/C++ 有着丰富的图形资产，而 Go 在这方面的积累为 0，甚至连马里奥的例子都是依赖一个 cgo 对 c sdl2 renderer 的封装）编译 wasm 的论点。

## 参考资料

- [《Go&WebAssembly 简介》PPT ](https://talks.godoc.org/github.com/chai2010/awesome-go-zh/chai2010/chai2010-golang-wasm.slide)
- [《WebAssembly 标准入门》图书](https://github.com/chai2010/awesome-wasm-zh/blob/master/webassembly-primer.md)

## 观看视频

{{< youtube id="O_FJgYKOBYQ" >}}
