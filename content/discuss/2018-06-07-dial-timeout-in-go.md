---
title: 2018-06-07 Dial 超时
---
来源：《Go 夜读》微信群

时间：2018-06-07

----

## 问题1

```go
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error) {
func DialTimeout(network, address string, timeout time.Duration) (Conn, error) { 
```

这两个方法返回的对象类型不一样，想得到TCPConn对象，同时可以设置连接超时，应该怎么写？有谁可以帮帮忙吗？

回答：

Use [`net.Dialer`](https://godoc.org/net#Dialer) with either the [Timeout](https://godoc.org/net#Dialer.Timeout) or [Deadline](https://godoc.org/net#Dialer.Deadline) fields set.

```go
d := net.Dialer{Timeout: timeout}
conn, err := d.Dial("tcp", addr)
if err != nil {
   // handle error
}
```

A variation is to call [`Dialer.DialContext`](https://godoc.org/net#Dialer.DialContext) with a [deadline](https://godoc.org/context#WithDeadline) or [timeout](https://godoc.org/context#WithTimeout) applied to the context.

Type assert to `*net.TCPConn` if you specifically need that type instead of a `net.Conn`:

```go
tcpConn, ok := conn.(*net.TCPConn)
```

## 参考资料

1. [https://stackoverflow.com/questions/47117850/how-to-set-timeout-while-doing-a-net-dialtcp-in-golang](https://stackoverflow.com/questions/47117850/how-to-set-timeout-while-doing-a-net-dialtcp-in-golang)
