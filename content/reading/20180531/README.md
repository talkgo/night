---
title: 2018-05-31 线下活动 - Go 标准包阅读
---
>参与人数: 10 人

*Go 标准包阅读*

Go 版本：go 1.10.2

### net/http

- server.go

### 问题

1. 

```go
func (s *Server) doKeepAlives() bool {
	return atomic.LoadInt32(&s.disableKeepAlives) == 0 && !s.shuttingDown()
}
```

为什么要用 `atomic.LoadInt32(&s.disableKeepAlives) == 0` ？

原子操作比用锁更节约一点性能。

2. server.go#Shutdown 不保险

3. 	panicChan := make(chan interface{}, 1)

```go
	panicChan := make(chan interface{}, 1)
	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		h.handler.ServeHTTP(tw, r)
		close(done)
	}()
	select {
	case p := <-panicChan:
		panic(p)
		...
```

外部处理就不能按照你的意愿去处理了，如果不拿出来，那么进程就挂掉了。

4. // Deprecated: ErrWriteAfterFlush is no longer used.
	ErrWriteAfterFlush = errors.New("unused")

5. Header() Header 注释引发的Trailer的思考？

![](/images/2018-05-31-night-reading-go-01.jpeg)
![](/images/2018-05-31-night-reading-go-03.jpeg)
![](/images/2018-05-31-night-reading-go-02.jpeg)

## 延伸阅读

1. [HTTP Chunked Body/Trailer编码](http://www.unclekevin.org/?p=203)
2. [example_ResponseWriter_trailers](https://golang.org/pkg/net/http/#example_ResponseWriter_trailers)
3. [HTTP Header Trailer](https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Headers/Trailer)
