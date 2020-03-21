---
title: "2018-12-07 微信讨论"
date: 2018-12-07T23:13:46+08:00
---

来源：『Go 夜读』微信群

>应该还是负数问题，你实验一下

下面这种情况的负数报错是在 build 阶段报的：

```golang
errMake := make([]byte, 64-22*5)  // negative len argument in make([]byte)
```

下面这种情况的负数，编译时不知道，是在运行时的报错：

```golang
errMake := make([]byte, 64-len(testStr)*5)  // runtime error: makeslice: len out of range
```
