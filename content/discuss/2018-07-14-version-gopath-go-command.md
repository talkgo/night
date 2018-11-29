---
title: 2018-07-14 包版本管理
---
来源：《Go 夜读》微信群

时间：2018-07-14

----

## 问题1：包版本管理？

govendor,vgo?

vgo 是 Go Module 的前身。

## 问题2：为什么要搞一个 GOPATH ？

GOPATH 是 Go 最初设计的产物，在 Go 语言快速发展的今天，人们日益发现 GOPATH 似乎不那么重要了，尤其是在引入 vendor 以及诸多包管理工具后。并且 GOPATH 的设置还会让 Go 语言新手感到些许困惑，提高了入门的门槛。Go core team 也一直在寻求 “去 GOPATH” 的方案，当然这一过程是循序渐进的。Go 1.8 版本中，如果开发者没有显式设置 GOPATH，Go 会赋予 GOPATH 一个默认值（在 linux 上为 $HOME/go ）。虽说不用再设置 GOPATH，但 GOPATH 还是事实存在的，它在 go toolchain 中依旧发挥着至关重要的作用。

## 问题3：go get 命令内部应该也是用 git clone 命令吧？

https://github.com/golang/go/wiki/GoGetTools
https://github.com/hyper0x/go_command_tutorial/blob/master/0.3.md

### 参考链接

1. https://github.com/golang/go/wiki/PackageManagementTools
2. [初窥 Go Module](https://mp.weixin.qq.com/s/ris9hYqRMKMX-HCZMpNMkg)
3. https://git-scm.com/docs/revisions
4. https://dave.cheney.net/2018/07/14/taking-go-modules-for-a-spin

