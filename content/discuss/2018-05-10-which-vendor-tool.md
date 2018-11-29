---
title: 2018-05-10 用什么工具解决项目依赖
---
来源：《Go 夜读》微信群
时间：2018-05-10 13:30

## 1. 你们是用什么工具解决项目依赖的呢？

- [dep](https://github.com/golang/dep)
- govendor

govendor 最简洁，glide最方便，不过一般团队用啥就用啥。

## 2. 平时你们的工作目录结果是怎样的？

- *一个大的 GOPATH 下，放所有项目*【比较多是这种方式】
- 一个项目一个 GOPATH
- 有一个默认 GOPATH，不同项目有不同的 GOPATH
	
	VisualStudio Code 的设置： `"go.gopath": "${workspaceRoot}:/Users/username/gopath"`
	公共第三方包是公用的，每个有自己的 vendor，编译打包什么的都是各自的 makefile（很多Github上的大项目，比如docker，都是自己写 Makefile ，整理项目依赖，没有一个统一的形式）

Rust 的项目结构比 Go 好，Rust 的 cargo 还不错。

- 可以给 GOPATH 写一个脚本

```shell
gop() {
	if [ "$1" = "" ]; then
	elif [ "$1" = "d" ]; then
		export GOPATH=`echo $DEF_GOPATH`
	elif [ "$1" = "a" ]; then
		export GOPATH=`echo $DEF_GOPATH`:`pwd`
	elif [ "$1" = "f" ]; then
		export GOPATH=`pwd`
	fi

	echo "current GOPATH = "$GOPATH
}
```

## 参考资料

1. [Go Vendoring Tools 使用总结](http://researchlab.github.io/2016/05/24/comparison-of-Go-Vendoring-Tools/)
2. [Golang 代码规范](https://sheepbao.github.io/post/golang_code_specification/)