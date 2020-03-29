---
desc: Go opentracing jaeger 集成及源码分析
title: 第 29 期 Go opentracing jaeger 集成及源码分析
date: 2019-01-23T21:00:00+08:00
author: jukylin
---

# Go opentracing jaeger 集成及源码分析

## 一、分布式追踪论文

论文地址：http://bigbully.github.io/Dapper-translation/ 


### 为什么要用分布式追踪

> 当代的互联网的服务，通常都是用复杂的、大规模分布式集群来实现的。
互联网应用构建在不同的软件模块集上，这些软件模块，有可能是由不同的团队开发、
可能使用不同的编程语言来实现、有可能布在了几千台服务器，横跨多个不同的数据中心。
因此，就需要一些可以帮助理解系统行为、用于分析性能问题的工具。

### 分布式系统调用过程

![image](https://bigbully.github.io/Dapper-translation/images/img1.png)

  
###  使用分布式追踪要留意哪些问题

* 低损耗

	> 跟踪系统对在线服务的影响应该做到足够小。

* 应用透明

	>  对于应用的程序员来说，是不需要知道有跟踪系统这回事的。
	
	
## 二、Opentracing简介

### Opentracing的作用

* OpenTracing通过提供平台无关、厂商无关的API，使得开发人员能够方便的添加（或更换）追踪系统的实现。
 
* 可以很自由的在不同的分布式追
     踪系统中切换

* 不负责具体实现

![image](https://camo.githubusercontent.com/e5b7d545b447ac93dfdbac415f4180a7c1644369/68747470733a2f2f75706c6f61642e63632f692f6156364e7a432e706e67)


### Opentracing主要组成

* 一个Trace
	> 一个trace代表了一个事务或者流程在（分布式）系统中的执行过程

* Span
	> 记录Trace在执行过程中的信息

* 无限极分类
	> 服务与服务之间使用无限极分类的方式，通过HTTP头部或者请求地址传输到最低层，从而把整个调用链串起来。

![image](https://camo.githubusercontent.com/57a991f13b85b69442aa728fb92253391c309ea0/68747470733a2f2f75706c6f61642e63632f692f4f68736a41302e6a7067)

### Jaeger-client的实现

#### Jaeger-client源码

##### 提取

* 为什么要提取
> 主要作用是为了找到父亲

* 从哪里提取
	> 进程内，不同进程之间各自约定
	> 粟子：github.com/opentracing-contrib/go-stdlib/nethttp/server.go	 P86

* 提取什么
	> 	traceid:spanid:parentid:是否采集
	> uber-trace-id=157b74261b51d917:157b74261b51d917:0:1
	> github.com/jaegertracing/jaeger-client-go/propagation.go P124


##### 注入

* 为什么要注入
	> 主要为了让孩子能找到爸爸

* 注入到哪里
	> 和提取相对
	> github.com/jaegertracing/jaeger-client-go/propagation_test.go

* 注入了什么
	> github.com/jaegertracing/jaeger-client-go/propagation.go P103


##### 异步report
* Span.finish
	> github.com/jaegertracing/jaeger-client-go/span.go P177

* 把Span放入队列
	> github.com/jaegertracing/jaeger-client-go/reporter.go P219

* 从队列取出，生成thrift，放入spanBuffer
	> github.com/jaegertracing/jaeger-client-go/reporter.go P253

* Flush到远程
	> github.com/jaegertracing/jaeger-client-go/transport_udp.go P113


#### 低消耗

* 消耗在哪里	
	> Jaeger-client作用于应用层，提取、注入、生成span、序列化成Thrift、发送到远程等，一系列操作这些都会带来性能上的损耗。

* 如何处理
	> 选择合适采集策略：
    1. Constant
    2. Probabilistic
    3. Rate Limiting
    4. Remote


#### 应用透明

* 如何做到让业务开发人员无感知
	1. Golang：
		约定第一个参数为ctx，把parentSpan放入ctx
		github.com/opentracing/opentracing-go/gocontext.go
	2. PHP：
		使用全局变量 


## 三、Jaeger服务端源码阅读

### 服务端组件职责 
> 各组件按照微服务架构风格设计，职责单一


![image](https://camo.githubusercontent.com/e877b9ef989f6ca60f4cce8bfe39350237a92d6a/687474703a2f2f6a61656765722e72656164746865646f63732e696f2f656e2f6c61746573742f696d616765732f6172636869746563747572652e706e67)


* Jaeger-agent负责上报数据的整理

* Jaeger-collector负责数据保存

* Jaeger-query负责数据查询

* Jaeger-agent和Jaeger-collector使用基于TCP协议实现的RPC进行通讯


![image](https://camo.githubusercontent.com/efed552ff18aa3f8583d0b4af0bd5a35bea67bd9/68747470733a2f2f75706c6f61642e63632f692f324a41516b702e706e67)

### Jaeger-agent 源码阅读
* 监听3个UDP端口   
	> github.com/jaegertracing/jaeger/cmd/agent/app/flags.go P35
	> github.com/jaegertracing/jaeger/cmd/agent/app/servers/thriftudp/transport.go P73

* 接收Jaeger-client的数据，放入队列dataChan
	> github.com/jaegertracing/jaeger/cmd/agent/app/servers/tbuffered_server.go #80 
* 从队列dataChan获取数据，进行校验
	> github.com/jaegertracing/jaeger/cmd/agent/app/processors/thrift_processor.go P108

* 提交数据
	> github.com/jaegertracing/jaeger/thrift-gen/jaeger/tchan-jaeger.go #39

### Jaeger-collector 源码阅读

* 协程池
	> github.com/jaegertracing/jaeger/pkg/queue/bounded_queue.go

* 接收jaeger-agent数据
	> github.com/jaegertracing/jaeger/cmd/collector/app/span_handler.go P69

* 放入队列
	> github.com/jaegertracing/jaeger/cmd/collector/app/span_processor.go P112

* 从队列拿出来，写入数据库 	 
    > github.com/jaegertracing/jaeger/cmd/collector/app/span_processor.go p54 	
    > github.com/jaegertracing/jaeger/plugin/storage/cassandra/spanstore/writer.go P136

## 四、Jaeger使用经验

### 监听指标

* Jaeger-client 监听 reporter_spans

* Jaeger-agent 监听 thrift.udp.server.packets.dropped 

* Jaeger-collector 监听 spans.dropped

http://localhost:16686/metrics

### 测试环境debug
> 测试环境记录执行mysql语句，redis命令，RPC参数、结果
可以很方便定位问题

### 性能调优
> 观察Jaeger-ui，对线上接口，mysql执行时间进行监控调优

## 观看视频

{{< youtube id="ub7jtN13KHA" >}}
