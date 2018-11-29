---
title: 2018-06-28 线下活动
---
>参与人数: 10 人

### Go 标准包阅读

- Go版本：`go 1.10.2`

### net包

- http/server.go
- http/request.go
- textproto/reader.go

### 读取位置

- textproto/reader.go(`140行`)

### 问题

> **1.各个系统的回车换行符区别**                

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/reading/20180628/images/20180628-1.jpeg)

- 注意:`10.13及其以上是macOS系统`

> **2.URI，URL和URN的区别**                   

- [查看详情](http://www.cnblogs.com/hust-ghtao/p/4724885.html)

> **3.HTTP CONNECT方法介绍**                  

**会议讨论小结**

```
	可以建立一个代理服务器到目标服务器的透明通道（tcp连接通道），中间完全不会对数据做任何处理，直接转发（支持https，一种翻墙的手段，专线独享）
```

- [HTTP代理协议 HTTP/1.1的CONNECT方法](https://www.web-tinker.com/article/20055.html)

> **4.peek读取字节内部实现**                  

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/reading/20180628/images/20180628-4.jpeg)

- 这里先peek获取流数据(注意：`这里没有对Peek的错误进行处理，而是根据是否Buffered读取到数据来判断错误`)
- 为什么没有对Peek的错误进行处理呢？`主要是因Peek失败了也有可能不会返回错误`

```
	golang读取字节表现形式是阻塞式的，但其实底层是用了非阻塞式的NIO，如果没有读取到数据会定时轮询读取
```

> **5.http header尾部的符号什么情况下会存在\n\n的情况？(`待解决，欢迎在下面评论`)**             

看源码发现hearder结尾会存在`\r\n\r\n`和`\n\n`两种字符情况

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/reading/20180628/images/20180628-2.jpeg)

网络上查资料发现只会存在`\r\n\r\n`

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/reading/20180628/images/20180628-3.jpeg)

- TODO                 


### 相关链接

- [uri和url的详细规范](https://tools.ietf.org/html/rfc3986)
- [扒一扒HTTP的构成](http://mrpeak.cn/blog/http-constitution/)
- [20180628直播视频](https://www.youtube.com/watch?v=xodlVBWxTYM)
