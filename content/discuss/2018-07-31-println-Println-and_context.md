---
title: 2018-07-31 println与fmt.Println有何猫腻
---

来源：《Go 夜读》微信群

时间：2018-07-31

----

## 问题1：println与fmt.Println有何猫腻？
  下面2个输出语句代码，哪一行先输出
```go
package main

import "fmt"

func main() {
	println("11111")
	fmt.Println("22222")
}
```
本人的是debian linux系统下，go version为1.9.1，goland IDE,反复执行几次之后发现始终"反着"输出
```
22222
11111
```
按照一般人的思维惯性(或者说下意识)，应该先输出11111，再输出22222,然后向Go夜读微信群发起提问，结果群友测试的结果就是顺序输出，他的测试环境是macbook, go version1.10.3
```go
package main

import "fmt"

func main() {
	println("aaaaa")
	fmt.Println("bbbbb")
}
```
输入
```bash
aaaaa
bbbbb
```
难道是因为内容不同就不一样？难道是因为和系统有关？难道是和go版本有关？

我们看下println源代码，如下：
```go
// 大概意思是：println内置函数将其参数格式化为特定实现方式，然后将结果写入标准错误流
// 而Println函数对于我们开发程序的时候调试输出很有用,且不能保证以后会留在语言中

// The println built-in function formats its arguments in an
// implementation-specific way and writes the result to standard error.
// Spaces are always added between arguments and a newline is appended.
// Println is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func println(args ...Type)
```
 我们再看看Println源代码,如下：
```go
// 大概意思是：Println函数使用默认格式化并将结果写入标准输出流
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```
看了2个函数的实现之后，总结如下：
1. println是输出到stderr
2. fmt.Println是输出到stdout

fmt.Println writes go standard output(stdout) and println writes to standard error (stderr),two different,unsynchronized files

结论：linux中，标准输入，标准输出，标准错误分别对应三个管道，编号分别是0,1,2，不同的管道输出是并发，不存在不同的管道之间相互等待，也就不存在先后顺序，上面两个函数是输出到不同的输出管道，就有可能导致顺序不一致，一个终端总是会存在标准输出和标准错误输出。所以最开始的那个问题是随机输出，没有固定的顺序!,写个大点的循环就能容易发现是随机，不存在系统不一致，或者go版本不一致导致顺序可能不同。

## 问题2：context.Context普通用，是不是主要通过with value？
问题： context.Context普通用，是不是主要通过with value,可以传递部分参数而不用传递全部参数，然后调用的函数可以把对应的参数取出来？对这个问题，各位有什么看法？
答：参考链接：http://www.flysnow.org/2017/05/12/go-in-action-go-context.html
