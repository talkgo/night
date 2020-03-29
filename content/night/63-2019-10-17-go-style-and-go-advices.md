---
desc: Go 夜读 之 Go 编码风格阅读与讨论
title: 第 63 期 Go 编码风格阅读与讨论
date: 2019-10-17T21:00:00+08:00
author: mai
---

## Go 夜读第 63 期 Go 编码风格阅读与讨论

## 内容简介

本期主要是针对近期 uber-go/guide style 和 go-advices 的解读以及开发者讨论。

## 内容大纲

- Go CodeReview Comments
- Uber-go/style
- Go-advices

## 分享地址

2019-10-17 21:00:00 ~ 22:10:00, UTC+8

https://zoom.us/j/6923842137

## 分享 Slides

https://docs.google.com/presentation/d/1MlzZJBK0Zq0VzJVC_AqSWmmlS4Of-8xY6NGZmfhKQXI/edit?usp=sharing

## 进一步阅读的材料

- [Go CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
- [uber-go/guide style](https://github.com/uber-go/guide/blob/master/style.md)
- [go-advices](https://github.com/cristaloleg/go-advices)

## Go CodeReviewComments 翻译

- [Go Code Review Comments 译文（截止 2018 年 7 月 27 日）](https://www.zybuluo.com/wddpct/note/1264988)

Go 官方的建议已经涉及到非常方面：

- Gofmt
- Comment Sentences
- Contexts
- Copying
- Crypto Rand
- Declaring Empty Slices
- Doc Comments
- Don't Panic
- Error Strings
- Examples
- Goroutine Lifetimes
- Handle Errors
- Imports
- Import Dot
- In-Band Errors
- Indent Error Flow
- Initialisms
- Interfaces
- Line Length
- Mixed Caps
- Named Result Parameters
- Naked Returns
- Package Comments
- Package Names
- Pass Values
- Receiver Names
- Receiver Type
- Synchronous Functions
- Useful Test Failures
- Variable Names

### gofmt

不管你是用什么开发工具，都推荐一定要配置 goimports。

### Context

- Context 应该在函数的第一个参数；
- 不要将 Context 加到结构体中，而应该加一个 ctx 参数；
- 不要创建自定义的 Context 类型；
- Context 是不可变的，所以可以将相同的 ctx 传递给调用共享相同截止日期，取消信号，凭证，父跟踪等；

虽然官方已经说明了，但是也还是有不少公司或者开源项目有自己的设计和实现。

### Declaring Empty Slices

`var t[]string` 比 `t:= []string{}` 更好

### Imports

应该按系统库、内部库、第三方库分层分隔。

### Indent Error Flow

```go
if err != nil {
	// error handling...
	return // or continue, etc.
}
// other code
```

```go
x, err := f()
if err != nil {
	// error handling...
	return
}
// use x...
```

### Variable Names

- 局部变量应该越精简越好；
- 不通用的或者全局变量，应该描述更清楚的命名；

## Uber Go 风格指南翻译

- [Uber Go 风格指南 (译) by 徐旭](https://note.mogutou.xyz/articles/2019/10/13/1570978862812.html)
- [Uber Go 语言编程规范 by legendtkl](https://mp.weixin.qq.com/s/SNmq0llxuu8NUkhwenegRg)
- [【重磅】Uber Go 语言代码风格指南 by Go 中国](https://mp.weixin.qq.com/s/cu6IZl_BhWokJxMXYmSytg)
- [Uber Go 语言编码规范 by TonyBai](https://mp.weixin.qq.com/s/LYLLghOjevBDieAM_LKrjA)

上周刚出来，过了2天，就出现大量的翻译文章，也能够看出来 Go 语言虽然官方有 gofmt，以及 go vet 静态代码检测工具，但是也抵挡不住大家对于代码风格的热衷。
>也说明大家还希望将代码风格更统一，追求更好的代码。

## Go-advices

- Code
- Concurrency
- Performance
- Modules
- Build
- Testing
- Tools
- Misc

### 代码方面

- `var foo time.Duration` 比 `var fooMillis int64` 更好
- 检查 `defer` 中的 error
- 用 `%+v` 打印足够详细信息
- 小心 range
- map 中读取一个不存在的 key 不会 panic，建议：`value, ok := map[“no_key”]`
- 将 defer 移到顶部

## Dave Cheney

- [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- 中文翻译版（2018 年）：https://www.flysnow.org/2018/12/04/golang-the-go-best-presentations.html

----

## Q&A 总结

1. 下划线开头 来声明 全局变量
>很少见，也不太适用。
2. 为什么建议channel尽量不加buffer? 
>按需分配。
3. go.uber.org/atomic 
>库非常好，原子操作很方便。新增多种数据类型。
4. package_test 比 package 好？
>比较清晰，但是也有局限性，测试不了内部逻辑，类似于外部包调用。
5. go test 指定 -count 可以消除偶然因素导致的不稳定结果。
>-count=1 也可以消除 cache。
6. 有没有认证或授权库的推荐？
>github 上搜索即可，一般可以根据 star 数量和活跃情况来评判。
也可以去 godoc.org 搜索，查看 imports 引入数据来评判。
https://github.com/dgrijalva/jwt-go


---

## 观看视频

{{< youtube id="91YbbwlKZ2k" >}}
