---
title: "2019-03-07 微信讨论"
date: 2019-03-07T10:51:00+08:00
---

来源：『Go 夜读』微信群

## 1 Goland中println和fmt.Print乱序原因

**Question**: 下面的代码，在goland上运行后，为何输出顺序不确定？但是使用命令行go run xxx运行，输出顺序却是确定的？

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

**Answer**:

println输出流是stderr，而fmt.Print输出流是stdout。而goland应该是区分stderr和stdout的（输出结果中stderr为红色），这种顺序不确定应该是goland自身处理stderr/stdout输出的时候导致的。因此通过非goland（比如console）方式运行该代码，顺序是确定的。



reference：

- [golang i use fmt.Println() after println() but](https://stackoverflow.com/questions/35931166/golang-i-use-fmt-println-after-println-but)
- [为什么Go自带的日志默认输出到os.Stderr？](https://www.zhihu.com/question/67629357)



## 2 Go 终端输出彩色文字方法

### 2.1 色彩编码

ANSI转义序列是一种带内信号的转义序列标准，用于控制视频文本终端上的光标位置/颜色和其他选项。在文本中嵌入确定的字节序列，大部分以`ESC转义`字符和`[`字符开始，终端会把这些字节序列解释为相应的指令，而不是普通的字符编码。

其中有专门控制字符颜色的控制符。一般由`ESC[`开始，中间包含若干个（包括0个）参数字节，以及一个最终字节组成。例如：`\x1b[37;44;4;1m hello go \x1b[0m`，代表输出的文字`hello go`的格式是：蓝色背景（44），灰色字体（37），带下滑下（4）并且加粗（1），go 语言代码如下：

```
package main

import "fmt"

func main() {
    fmt.Printf("\x1b[37;44;4;1m hello go \x1b[0m")
}
```

其中：

```
- \x1b 标志字符，代表转义序列开始，其实就是0x1B
- [ 转义序列的开始符
- 以 ; 分割的数字是控制字符
- m 代表结束控制字符序列
```

常用的文本样式控制符如下：

编码  | 说明
------------- | -------------
0  | 重置/清除样式
1  | 加粗
3  | 斜体
4  | 下划线
5  | 闪烁
8  | 隐藏
30～37 | 前景色，参考下文『1位颜色编码』
38 | 设置前景色，后跟 5;n 代表使用8位256颜色码，后跟 2;r;g;b代表24位RGB颜色码
40~47 | 背景色，参考下文『1位颜色编码』
48  | 设置背景色，后跟 5;n代表使用8位256颜色码，后跟 2;r;g;b代表24位RGB颜色码
90～97 | 亮色前景色，参考下文 『1 位颜色编码』
100～107 | 亮色背景，参考下文 『1 位颜色编码』

1位颜色编码


颜色  | 前景色编码 | 背景色编码
------------- | ------------- | -------------
黑色  | 30  | 40
红色  | 31  | 41
绿色  | 32  | 42
黄色  | 33  | 43
蓝色  | 34  | 44
品红色  | 35  | 45
青色  | 36  | 46
白色（灰）  | 37  | 47
亮黑色（灰）  | 90  | 100
亮红色  | 91  | 101
亮绿色  | 92  | 102
亮黄色  | 93  | 103
亮蓝色  | 94  | 104
亮品红色  | 95  | 105
亮青色  | 96  | 106
亮白色  | 97  | 107


### 2.2 例子

通过颜色控制，我们可以输出带颜色的log，例子如下：

```
package main

import (
	"fmt"
	"time"
)

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta

	succ = "[succ]"
	error = "[error]"
	warn = "[warn]"
	info = "[info]"
	debug = "[debug]"
)

func red(s string) string {
	return fmt.Sprint("\x1b[%dm%s\x1b[0m", color_red, s)
}

func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}

func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}

func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}

func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

func Success(format string, a ... interface{}) {
	prefix := green(succ)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a ...))
}

func Error(format string, a ... interface{}) {
	prefix := red(error)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Warning(format string, a ... interface{}) {
	prefix := magenta(warn)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Info(format string, a ... interface{}) {
	prefix := blue(info)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func Debug(format string, a ... interface{}) {
	prefix := yellow(debug)
	fmt.Println(formatLog(prefix), fmt.Sprintf(format, a...))
}

func formatLog(prefix string) string {
	return time.Now().Format("2006/01/02 15:04:05") + "" + prefix + ""
}

func main() {
	Debug("%s", "hello, go")
}
```

### 2.3 参考
- [ANSI转义序列](https://zh.wikipedia.org/wiki/ANSI%E8%BD%AC%E4%B9%89%E5%BA%8F%E5%88%97)
- [教你写一个color日志库](https://toutiao.io/posts/2889gp/preview)
- [在终端中输出彩色文字](https://segmentfault.com/a/1190000012666612)
- [IntelliJ IDEA 安装 Grep Console 自定义控制台输出多颜色格式](http://www.ibloger.net/article/2975.html)
