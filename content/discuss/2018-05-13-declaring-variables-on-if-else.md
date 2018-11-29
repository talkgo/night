---
title:  关于变量在 if-else 条件表达式里的作用域范围
---
## 一. 如何发现这个问题?

我是在浏览 `Twitter` 的时候, 发现博主 [David Crawshaw](https://twitter.com/davidcrawshaw) 分享了一段代码,

推文地址:
[https://twitter.com/davidcrawshaw/status/994614426064928770](https://twitter.com/davidcrawshaw/status/994614426064928770)

![ifelsetwitter.png](/images/ifelsetwitter.png)

原博文代码如下, [点击在  play 里运行](https://play.golang.org/p/1tD1C6sOxcV):

```go
package main

import "fmt"

func main() {
	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		fmt.Println(a, b, c) // 结果: 1 2 3
	}
}

```

所以, 就扩展一下如下这种方式:

```go
package main

import "fmt"

func main() {
	if a := 1; false {
		aa := 11
	} else if b := 2; false {
		bb := 22
	} else if c := 3; false {
		cc := 33
	} else {
		fmt.Println(a, b, c)
		fmt.Println(aa, bb, cc) // 结果: undefined aa bb cc
	}
}

```

再尝试如下这种方式:

```go
package main

import "fmt"

func main() {
	if x := 10; x < 9 {
		fmt.Println(x)
	} else {
		fmt.Println("not x")
	}
	fmt.Println(x) // 结果: other if-else.go:11:14: undefined: x
}

```

脑洞再大点:

```go
package main

import "fmt"

func main() {
	if a := 1; false {
		fmt.Println(a, b, c) // undefined: b, c
	} else if b := 2; false {
		fmt.Println(a, b, c) // undefined: c
	} else if c := 3; false {
		fmt.Println(a, b, c)
	} else {
		fmt.Println(a, b, c)
	}
}

```

合理的猜测: 
```go
package main

import "fmt"

func main() {
	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		if a := 11; true {
			fmt.Println(a, b, c) // 11 2 3
		}
		fmt.Println(a, b, c) // 1 2 3
	}
}

```

所以可以得出一个初步的结论:

**只有 if/else-if 条件表达式里的变量声明作用域才会向下到达最后 `else` 内部, 显而易见的不能向上作用, 作用域范围仅限于本级, 不影响正常的变量作用范围以及屏蔽作用.**

## 二. 这么做有什么好处和坏处?

### A. 好处: 
> 1. 如下面 `Chris Hines` [评论列出的代码](https://twitter.com/chris_csguy/status/994627365576806401) , 在处理 sql 任务的时候, 对返回值和错误做处理, 逻辑上看, 比较流畅. 

```go
if result, err := db.Exec(updateSql, ...); err != nil {
	return err
} else if count, err := result.RowsAffected(); err != nil {
	return err
} else if count != 1 {
	return ErrNotUpdated
}
```

> 2. 另一个[评论的代码](https://twitter.com/davidcrawshaw/status/994621058702499840), 
emmmmm, 怎么说呢? 这么看代码确实比较简洁优美一些, 不过对于不了解这个特性的人来说, 可能有坑, 增加阅读难度.

```go
if got, err := f(); err != nil {
	// 巴拉巴拉巴拉
} else if want := ...; got != want {
    t.Errorf(..., got, want)
}
```

> 好在我们可以加注释. 改成如下, 在可读性上会好点:

```go
// NOTE: golang 特性, got 的作用域贯穿整个 if-else...
if got, err := f(); err != nil {
	// 巴拉巴拉巴拉
} else if want := ...; got != want {
    t.Errorf(..., got, want)
}
```

> 3. 语言的设计就是如此: 
参考设计文档: [https://golang.org/ref/spec#If_statements](https://golang.org/ref/spec#If_statements), 官方并没有明确的说这么做的好处. 不过确实给书写代码带来了方便. 

### B. 坏处:
> 1. emmmmm, 就是对于不知道这个特性的人来说比较奇怪. 不是很理解, 不符合编程的核心原则: 代码首先是给人看的, 顺便给机器运行. 