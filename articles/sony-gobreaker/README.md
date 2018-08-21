---
title: "Sony gobreaker容断器源码分析]"
date: 2018-07-26
author: HuangChuanTonG@WPS.cn
---

最近看了一下go-kit，发现这个微服务框架的容断器，也是使用sony开源的作为基础。
[sony开源在 github 的容断器](https://github.com/sony/gobreaker)
在源代头注释中发现，原来sony实现的是微软2015时公布的CircuitBreaker标准，果然微软才开源界的大神。


## 1）微软定义的 Circuit breaker

微软的原文件在此：https://msdn.microsoft.com/en-us/library/dn589784.aspx
名不知道怎么正确翻译，直观翻译，可能叫：环形容断器（或叫：循环状态自动切换中断器）。
因为它是在下面3个状态循环切换  ：
```
         Closed 
         /    \
 Half-Open <--> Open

初始状态是：Closed，指容断器放行所有请求。
达到一定数量的错误计数，进入Open 状态，指容断发生，下游出现错误，不能再放行请求。
经过一段Interval时间后，自动进入Half-Open状态，然后开始尝试对成功请求计数。
进入Half-Open后，根据成功/失败计数情况，会自动进入Closed或Open。

```
## 2）sony开源的go实现
```go
// 从定义的错误来看，sony的应该增加了对连接数进行了限制 。
 var (
	// ErrTooManyRequests is returned when the CB state is half open and the requests count is over the cb maxRequests
	ErrTooManyRequests = errors.New("too many requests")
	// ErrOpenState is returned when the CB state is open
	ErrOpenState = errors.New("circuit breaker is open")
)
```
### 2.1） 通过Settings的实现，了解可配置功能：
```go
type Settings struct {
	Name          string
	MaxRequests   uint32        // 半开状态期最大允许放行请求：即进入Half-Open状态时，一个时间周期内允许最大同时请求数（如果还达不到切回closed状态条件，则不能再放行请求）。
	Interval      time.Duration // closed状态时，重置计数的时间周期；如果配为0，切入Open后永不切回Closed--有点暴力。
	Timeout       time.Duration // 进入Open状态后，多长时间会自动切成 Half-open，默认60s，不能配为0。

    // ReadyToTrip回调函数：进入Open状态的条件，比如默认是连接5次出错，即进入Open状态，即可对容断条件进行配置。在fail计数发生后，回调一次。
	ReadyToTrip   func(counts Counts) bool 

	// 状态切换时的容断器
	OnStateChange func(name string, from State, to State)
}
```
### 2.2）核心的*执行函数*实现

要把容断器使用到工程中，只需要，实例化一个gobreaker，再使用这个Execute包一下原来的请求函数。
```go
func (cb *CircuitBreaker) Execute(req func() (interface{}, error)) (interface{}, error) {
	generation, err := cb.beforeRequest() // 
	if err != nil {
		return nil, err
	}

	defer func() {
		e := recover()
		if e != nil {
			cb.afterRequest(generation, false)
			panic(e) // 如果代码发生了panic，继续panic给上层调用者去recover。
		}
	}()

	result, err := req()
	cb.afterRequest(generation, err == nil)
	return result, err
}
```

### 2.2 关键  func beforeRequest()

函数做了几件事：

 0. 函数的核心功能：判断是否放行请求，计数或达到切换新条件刚切换。
 1. 判断是否Closed，如是，放行所有请求。
	 - 并且判断时间是否达到Interval周期，从而清空计数，进入新周期，调用toNewGeneration()	 
 2. 如果是Open状态，返回ErrOpenState，---不放行所有请求。
	- 同样判断周期时间，到达则 同样调用 toNewGeneration(){清空计数}	
 3. 如果是half-open状态，则判断是否已放行MaxRequests个请求，如未达到刚放行；否则返回:ErrTooManyRequests。
 4. 此函数一旦放行请求，就会对请求计数加1（conut.onRequest())，请求后到另一个关键函数 : afterRequest()。

### 2.3 关键  func afterRequest()
 1. 函数核心内容很简单，就对成功/失败进行计数，达到条件则切换状态。
 2. 与beforeRequest一样，会调用公共函数 currentState(now) 
	 - currentState(now) 先判断是否进入一个先的计数时间周期(Interval), 是则重置计数，改变容断器状态，并返回新一代。
	 - 如果request耗时大于Interval, 几本每次都会进入新的计数周期，容断器就没什么意义了。

## 代码的核心内容

 1. 使用了一个generation的概念，每一个时间周期(Interval)的计数(count)状态称为一个generation。
 2. 在before/after的两个函数中，实现了两个状态自动切换的机制：
	 - 在同一个generation(即时间）周期内，计数满足状态切换条件，即自动切换；
	 - 超过一个generation时间周期的也会自动切换；
 3. 没有使用定时器，只在请求调用时，去检测当时状态与时间间隔。
 
 
