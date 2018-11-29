---
title: 2018-07-04 文件夹命名
---
来源：《Go 夜读》微信群

时间：2018-07-04

----

## 问题1. 话说大家 go 文件夹命名 是什么规范。下划线 横线 驼峰？有标准么？

有标准 [https://blog.golang.org/package-names](https://blog.golang.org/package-names)

## 问题2. 包的命名多个单词的情况呢？

包名和文件名用小写,使用短命名,尽量和标准库不要冲突。

文件名或者文件夹不要用大写，切记!!!

>系统有些区别大小写，有些不区分，比如 Mac 不区分大小写，而 linux 区分。

>制作规范的时候参考的是 effective go。

----

## 参考

* [Golang 代码规范](https://sheepbao.github.io/post/golang_code_specification/)
