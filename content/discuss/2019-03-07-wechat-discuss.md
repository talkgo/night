---
title: "2019-03-07 微信讨论"
date: 2019-03-07T10:51:00+08:00
---

来源：《Go 夜读》微信群

Question: 下面的代码，为何输出顺序不确定？
```
package main

import "fmt"

func main() {
    println("hello")
    println("world")
    fmt.Print("go\n")
}

输出结果：
hello
world
go

或是：
go
hello
world

```

Answer:

由于println输出流是stderr，而fmt.Print输出流是stdout。stderr和stdout默认都是向屏幕输出，但是stdout输出方式是行缓冲，即输出的字符先存放到缓冲区，再输出到屏幕；而stderr则是不带缓冲的，直接输出到屏幕。因此，这两类函数使用的是不同的输出流，因此输出到屏幕上的顺序也无法保证。

