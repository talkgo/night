---
title: "Go 源码阅读之 flag 包"
date: 2020-02-08T17:54:04+08:00
---

- [简介](#%e7%ae%80%e4%bb%8b)
- [文件结构](#%e6%96%87%e4%bb%b6%e7%bb%93%e6%9e%84)
- [运行测试](#%e8%bf%90%e8%a1%8c%e6%b5%8b%e8%af%95)
- [总结](#%e6%80%bb%e7%bb%93)
	- [接口转换能实现类似 C++ 中模板的功能](#%e6%8e%a5%e5%8f%a3%e8%bd%ac%e6%8d%a2%e8%83%bd%e5%ae%9e%e7%8e%b0%e7%b1%bb%e4%bc%bc-c-%e4%b8%ad%e6%a8%a1%e6%9d%bf%e7%9a%84%e5%8a%9f%e8%83%bd)
	- [函数 vs 方法](#%e5%87%bd%e6%95%b0-vs-%e6%96%b9%e6%b3%95)
	- [`new` vs `make`](#new-vs-make)
	- [指针赋值给接口变量](#%e6%8c%87%e9%92%88%e8%b5%8b%e5%80%bc%e7%bb%99%e6%8e%a5%e5%8f%a3%e5%8f%98%e9%87%8f)
	- [flag文件夹中有`flag_test`包](#flag%e6%96%87%e4%bb%b6%e5%a4%b9%e4%b8%ad%e6%9c%89flagtest%e5%8c%85)
	- [作用域](#%e4%bd%9c%e7%94%a8%e5%9f%9f)
- [参考文献](#%e5%8f%82%e8%80%83%e6%96%87%e7%8c%ae)

## 简介
flag 包是 Go 里用于解析命令行参数的包。为什么选择它作为第一个阅读的包，因为它的代码量少。其核心代码只有一个 1000 不到的 flag.go 文件。

## 文件结构
flag 包的文件结构很简单，就一层。一个文件夹里放了 5 个文件，其文件及其作用如下：

* flag.go 
  
  flag 的核心包，实现了命令行参数解析的所有功能
* export_test.go
  
  测试的实用工具，定义了所有测试需要的基础变量和函数
* flag_test.go 
  
  flag 的测试文件，包含了 17 个测试单元
* example_test.go
  
  flag 的样例文件，介绍了 flag 包的三种常用的用法样例
* example_value_test.go 
  
  flag 的样例文件，介绍了一个更复杂的样例

## 运行测试
我先介绍一下 Go 的运行环境。
```bash
# 通过 brew install go 安装，源码位置为 $GOROOT/src
GOROOT=/usr/local/opt/go/libexec
# 阅读的源码通过 go get -v -d github.com/haojunyu/go 下载，源码位置为 $GOPATH/src/github.com
GOPATH=$HOME/go
```

单独测试 flag 包踩过的坑：
1. 无法针对单个文件进行测试，需要针对包。

这里重点说一下 export_test.go 文件，它是flag包的一部分`package flag`，但是它确实专门为测试而存在的，说白了也就一个`ResetForTesting`方法，用来清除所有命令参数状态并且直接设置Usage函数。该方法会在测试用例中被频繁使用。所以单独运行以下命令会报错"flag_test.go:30:2: undefined: ResetForTesting"
```bash
# 测试当前目录（报错）
go test -v .
# 测试包
go test -v flag
```

2. `go test -v flag` 测试的源码是 `$GOROOT/src` 下的（以我当前的测试环境）

指定 flag 包后，实际运行的源码是 `$GOROOT` 下的，这个应该和我的安装方式有关系。

## 总结
### 接口转换能实现类似 C++ 中模板的功能
flag 包中定义了一个结构体类型叫 `Flag`，它用来存放一个命令参数，其定义如下。
```go
// A Flag represents the state of a flag.
// 结构体Flag表示一个参数的所有信息，包括名称，帮助信息，实际值和默认值
type Flag struct {
	Name     string // name as it appears on command line名称
	Usage    string // help message帮助信息
	Value    Value  // value as set实现了取值/赋值方法的接口
	DefValue string // default value (as text); for usage message默认值
}
```
其中命令参数的值是一个 `Value` 接口类型，其定义如下：
```go
// Set is called once, in command line order, for each flag present.
// The flag package may call the String method with a zero-valued receiver,
// such as a nil pointer.
// 接口Value是个接口，在结构体Flag中用来存储每个参数的动态值（参数类型格式各样）
type Value interface {
	String() string   // 取值方法
	Set(string) error // 赋值方法
}
```
为什么这么做？因为这样做能够实现类似模板的功能。任何一个类型 `T` 只要实现了 `Value` 接口里的 `String` 和 `Set` 方法，那么该类型 `T` 的变量 `v` 就可以转换成 `Value` 接口类型，并使用 `String` 来取值，使用 `Set` 来赋值。这样就能完美的解决不同类型使用相同的代码操作目的，和 C++ 中的模板有相同的功效。


### 函数 vs 方法
函数和方法都是一组一起执行一个任务的语句，二者的区别在于调用者不同，函数的调用者是包 package，而方法的调用者是接受者 receiver。在 flag 的源码中，有太多的函数里面只有一行，就是用包里的变量 `CommandLine` 调用同名方法。
```go
// Parsed reports whether f.Parse has been called.
// Parsed方法： 命令行参数是否已经解析
func (f *FlagSet) Parsed() bool {
	return f.parsed
}

// Parsed reports whether the command-line flags have been parsed.
func Parsed() bool {
	return CommandLine.Parsed()
}
```

### `new` vs `make`
`new` 和 `make` 是 Go 语言中两种内存分配原语。二者所做的事情和针对的类型都不一样。
`new` 和其他编程语言中的关键字功能类似，都是向系统申请一段内存空间来存储对应类型的数据，但又有些区别，区别在于它会将该片空间置零。也就是说 `new(T)` 会根据类型 `T` 在堆上 申请一片置零的内存空间，并返回指针 `*T`。
`make` 只针对切片，映射和信道三种数据类型 `T` 的构建，并返回类型为 `T` 的一个已经初始化（而非零）的值。原因是这三种数据类型都是引用数据类型，在使用前必须初始化。就像切片是一个具有三项内容的描述符，包含一个指向数组的指针，长度和容量。通过 `make` 创建对应类型的变量过程是先分配一段空间，接着根据对应的描述符来创建对应的类型变量。关于 `make` 的细节可以看 draveness 写的 [Go语言设计与实现](book_golang)。
```go
// Bool defines a bool flag with specified name, default value, and usage string.
// The return value is the address of a bool variable that stores the value of the flag.
func (f *FlagSet) Bool(name string, value bool, usage string) *bool {
	p := new(bool)
	f.BoolVar(p, name, value, usage)
	return p
}


// sortFlags returns the flags as a slice in lexicographical sorted order.
// sortFlags函数：按字典顺序排序命令参数，并返回Flag的切片
func sortFlags(flags map[string]*Flag) []*Flag {
	result := make([]*Flag, len(flags))
	i := 0
	for _, f := range flags {
		result[i] = f
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}
```


### 指针赋值给接口变量
Go 中的接口有两层含义，第一层是一组方法（不是函数）的签名，它需要接受者（具体类型 `T` 或具体类型指针 `*T` ）来实现细节；另一层是一个类型，而该类型能接受所有现实该接受的接受者。深入理解接口的概念可以细读 [Go语言设计与实现之接口](book_golang_interface)。在 flag 包中的 `StringVar` 方法中`newStringValue(value, p)`返回的是 `*stringValue` 类型，而该类型（接受者）实现了 `Value` 接口（ `String` 和 `Set` 方法），此时该类型就可以赋值给 `Value` 接口变量。
```go
// StringVar defines a string flag with specified name, default value, and usage string.
// The argument p points to a string variable in which to store the value of the flag.
// StringVar方法：将命令行参数的默认值value赋值给变量*p,并生成结构Flag并置于接受者中f.formal
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.Var(newStringValue(value, p), name, usage) // newStringValue返回值是*stringValue类型，之所以能赋值给Value接口是因为newStringValue实现Value接口时定义的接受者为*stringValue
}
```


### flag文件夹中有`flag_test`包
flag 文件夹下有 `flag_test` 包，是因为该文件夹下包含了核心代码 flag.go 和测试代码 *_test.go 。这两部分代码并没有通过文件夹来区分。所以该 `flag_test` 包存在的意义是将测试代码与核心代码区分出来。而该包被引用时只会使用到核心代码。
```go
// example_test.go
package flag_test
```

### 作用域 
关于作用域 [Golang变量作用域](blog_varScope) 和 [GO语言圣经中关于作用域](go_bible) 都有了详细的介绍，前者更通俗易懂些，后者更专业些。在 flag 包的 `TestUsage` 测试样例中，因为 `func(){called=true}` 是在函数 `TestUsage` 中定义函数，并且直接作为形参传递给 `ResetForTesting` 函数，所以该函数是和局部变量 `called` 是同级的，当然在该函数中给该变量赋值也是合理的。
```go
//  called变量的作用域
func TestUsage(t *testing.T) {
	called := false
	// 变量called的作用域
	ResetForTesting(func() { called = true })
	if CommandLine.Parse([]string{"-x"}) == nil {
		t.Error("parse did not fail for unknown flag")
	} else {
		t.Error("hahahh")
	}
	if !called {
		t.Error("did not call Usage for unknown flag")
	}
}
```


## 参考文献
1. [Go 夜读之 flag 包视频](video_nightreading)
2. [实效 Go 编程之内存分配](book_effectGo)
3. [Go 语言设计与实现之 make 和 new](book_golang_makeNew)
4. [菜鸟教程之 Go 语言变量作用域](runoob_varScope)
5. [Go 语言圣经中关于作用域](go_bible)
6. [Go 语言中值 receiver 和指针 receiver 的对比](blog_receiver)
7. [Go CodeReviewComments](gowiki_receiver)
8. [Golang 变量作用域](blog_varScope)
9. [Go 语言圣经中关于作用域](go_bible)
10. [Go 语言设计与实现之接口](book_golang_interface)

[book_effectGo]: https://go-zh.org/doc/effective_go.html#new%E5%88%86%E9%85%8D
[book_golang_makeNew]:https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/
[github_goSource]: https://github.com/haojunyu/go
[runoob_varScope]: https://www.runoob.com/go/go-scope-rules.html
[book_goBible]:https://docs.hacknode.org/gopl-zh/ch2/ch2-07.html
[blog_reflect]:https://juejin.im/post/5a75a4fb5188257a82110544
[book_golang_interface]:https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/#42-
[video_nightreading]: https://www.bilibili.com/video/av45158627

[blog_receiver]: https://maiyang.me/post/2018-12-12-values-receiver-vs-pointer-receiver-in-golang
[gowiki_receiver]: https://github.com/golang/go/wiki/CodeReviewComments#receiver-type
