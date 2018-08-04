## 2018-08-02
##Push Notification (apns)
来源：《Go 夜读》微信群
时间：2018-08-02

----

## 问题: 如何批量发送apns消息推送

先看下如何发送一条的apns.
```go
func push(){
    cert, err := certificate.FromP12File("../cert.p12", "")
    if err != nil {
        log.Fatal("Cert Error:", err)
    }

    notification := &apns2.Notification{}
    notification.DeviceToken = "11aa01229f15f0f0c52029d8cf8cd0aeaf2365fe4cebc4af26cd6d76b7919ef7"
    notification.Topic = "com.sideshow.Apns2"
    notification.Payload = []byte(`{"aps":{"alert":"Hello!"}}`) // See Payload section below

    client := apns2.NewClient(cert).Production()
    res, err := client.Push(notification)

    if err != nil {
        log.Fatal("Error:", err)
    }
    fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

```

一开始会想到,当需要推送大量消息的时候,会想到每发一条推送的时候,开一个协程(当然也要限制协程的数量),像这样:
```go
    go push()
```

但是一个问题,apns我记得是最多只能挂载15个tls.也就是最多同时发送15条推送消息,远远不能满足像IM需要大量同时推送的需求.

14年的时候,苹果官方发布更换了所有的通讯协议:
"鉴于SSL 3.0最新发现的漏洞，为了保护用户，APNS决定在下周三也就是10月29号起开始停止对SSL 3.0的支持。所有仅支持SSL 3.0的推送服务需要更换为TLS以确保推送服务能够正常运行，同时支持了SSL 3.0和TLS的服务不会受到此次更新的影响。"

使用tls解决了http1的两个问题:
1.安全性无法保证
2.单次请求获取一个文档的方式，不满足如今流式传输体验所要求的性能

所以http2中的tls具有流式传输功能的,相比http1.他可以首次建立连接请求,响应返回之后并不会断开连接.他不用向HTTP1那样,每次发送请求都需要建立一次.更重要的是他一次建立连接后,通过这个连接的所有请求都不用等待上一次请求的响应返回,而是可以同时发送请求.

解决批量的方案是:

服务端接收到发送推送消息的请求后,不是马上发送,而是等待100ms,收集所有这段时间的所有推送请求(设置一个上限),一次向一个tls中去推送.

```go
cert, err := certificate.FromP12File(*certPath, "")
if err != nil {
    log.Fatal("Cert Error:", err)
}

notifications := make(chan *apns2.Notification, 100)
responses := make(chan *apns2.Response, *count)

client := apns2.NewClient(cert).Production()

for i := 0; i < 50; i++ {
    go worker(client, notifications, responses)
}

for i := 0; i < *count; i++ {
    n := &apns2.Notification{
        DeviceToken: *token,
        Topic:       *topic,
        Payload:     payload.NewPayload().Alert(fmt.Sprintf("Hello! %v", i)),
        }
        notifications <- n
}

for i := 0; i < *count; i++ {
    res := <-responses
    fmt.Printf("%v %v %v\n", res.StatusCode, res.ApnsID, res.Reason)
}

```

```go
func worker(client *apns2.Client, notifications <-chan *apns2.Notification, responses chan<- *apns2.Response) {
    for n := range notifications {
    res, err := client.Push(n)
    if err != nil {
        log.Fatal("Push Error:", err)
        }
        responses <- res
    }
}
```

## 参考资料

 [go-apns2](https://github.com/sideshow/apns2)
 [网络协议之TLS](https://www.cnblogs.com/syfwhu/p/5219814.html)
