---
title: 2018-06-14 线下活动
---
>参与人数: 12 人

*Go 标准包阅读*

Go 版本：go 1.10.2

### net/http

- server.go
- h2_bundle.go

### 问题

1. 	WriteHeader(statusCode int)

- 要先调用 header.set()
- 再调用 WriteHeader() 
- 然后调用 Write()
	- 如果在调用 Write() 之后，还有比较多的逻辑要处理，则一定要紧跟着马上调一下 Flush()
- 然后调用 Flush()

2. HTTP2 不支持 Hijacker

3. 使用了 Hijacker 之后不能再使用 Request.Body

```go
type Hijacker interface {
	// After a call to Hijack, the original Request.Body must not be used.
	Hijack() (net.Conn, *bufio.ReadWriter, error)
}
```

4. The returned bufio.Reader may contain unprocessed buffered data from the client.

5. CloseNotifier 主要用于 HTTP2

6. CloseNotify may wait to notify until Request.Body has been fully read.

7. HTTP2 中是如何在 net/http/server.go 中调用 serve() 触发的呢？

```go
if proto := c.tlsState.NegotiatedProtocol; validNPN(proto) {
	if fn := c.server.TLSNextProto[proto]; fn != nil {
		h := initNPNRequest{tlsConn, serverHandler{c.server}}
		fn(c.server, tlsConn, h)
	}
	return
}
```

主要是 TLSNextProto ，然后查询得到 onceSetNextProtoDefaults() 调用。

```go
// onceSetNextProtoDefaults configures HTTP/2, if the user hasn't
// configured otherwise. (by setting srv.TLSNextProto non-nil)
// It must only be called via srv.nextProtoOnce (use srv.setupHTTP2_*).
func (srv *Server) onceSetNextProtoDefaults() {
	if strings.Contains(os.Getenv("GODEBUG"), "http2server=0") {
		return
	}
	// Enable HTTP/2 by default if the user hasn't otherwise
	// configured their TLSNextProto map.
	if srv.TLSNextProto == nil {
		conf := &http2Server{
			NewWriteScheduler: func() http2WriteScheduler { return http2NewPriorityWriteScheduler(nil) },
		}
		srv.nextProtoErr = http2ConfigureServer(srv, conf)
	}
}
```

然后是调用如下代码：

```go
#h2_bundle.go
...
// ConfigureServer adds HTTP/2 support to a net/http Server.
//
// The configuration conf may be nil.
//
// ConfigureServer must be called before s begins serving.
func http2ConfigureServer(s *Server, conf *http2Server) error {
...
}
```

8. HTTP1 流水线，一条连接一个并发；HTTP2 是每个连接一个并发，每处理一个请求又是一个并发。

## 延伸阅读

1. HTTP2 协议
2. 多路复用
