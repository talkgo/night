---
title: 2018-05-24 线下活动 - Go 标准包阅读
---
>参与人数: 10 人

*Go 标准包阅读*

Go 版本：go 1.10.1

### net/http

- server.go

### 问题

1. Next Protocol Negotiation = NPN
2. Expect 100 Continue support

>见参考资料

3. header提到了：Expect和host
4. 判断了 header里面的HOST，但是后面又删除，为什么？

server.go#L980

```go	
delete(req.Header, "Host")
```

5. 判断是否支持 HTTP2 （isH2Upgrade）

```go
// isH2Upgrade reports whether r represents the http2 "client preface"
// magic string.
func (r *Request) isH2Upgrade() bool {
	return r.Method == "PRI" && len(r.Header) == 0 && r.URL.Path == "*" && r.Proto == "HTTP/2.0"
}
```

```go
调用：ProtoAtLeast(1, 1)
...
// ProtoAtLeast reports whether the HTTP protocol used
// in the request is at least major.minor.
func (r *Request) ProtoAtLeast(major, minor int) bool {
	return r.ProtoMajor > major ||
		r.ProtoMajor == major && r.ProtoMinor >= minor
}
```

待补充。。。

## 延伸阅读

1. https://github.com/golang/go/issues/22128
2. https://tools.ietf.org/html/draft-ietf-httpbis-p2-semantics-26#section-6.2.1
3. https://www.cnblogs.com/tekkaman/archive/2013/04/03/2997781.html
4. https://benramsey.com/blog/2008/04/http-status-100-continue/
5. http://www.ituring.com.cn/article/130844

