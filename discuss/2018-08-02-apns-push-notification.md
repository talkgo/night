## 2018-08-02

## Push Notification (apns)

来源：《Go 夜读》微信群
时间：2018-08-02

----
 
## 问题: 如何批量发送apns消息推送


Install apns2
```go
    go get -u github.com/sideshow/apns2
```


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

#### tls

而tls是14年的时候,苹果官方发布更换了所有的通讯协议:

"鉴于SSL 3.0最新发现的漏洞，为了保护用户，APNS决定在下周三也就是10月29号起开始停止对SSL 3.0的支持。所有仅支持SSL 3.0的推送服务需要更换为TLS以确保推送服务能够正常运行，同时支持了SSL 3.0和TLS的服务不会受到此次更新的影响。"

#### 使用tls解决了http1的两个问题:
1.安全性无法保证
2.单次请求获取一个文档的方式，不满足如今流式传输体验所要求的性能

所以http2中的tls具有流式传输功能的,相比http1.他可以首次建立连接请求,响应返回之后并不会断开连接.他不用向HTTP1那样,每次发送请求都需要建立一次.更重要的是他一次建立连接后,通过这个连接的所有请求都不用等待上一次请求的响应返回,而是可以同时发送请求.

####  解决批量的方案是:

服务端接收到发送推送消息的请求后,不是马上发送,而是等待100ms,收集所有这段时间的所有推送请求(设置一个上限),一次向一个tls中去推送.

先看下我们的推送架构,基本是根据现有的架构自我调整:

|  消息队列  |
        ↓
| 消息收集器 |
        ↓
 |  消息合并  |
        ↓
|  过 滤 器  |
        ↓        
  推送模块
  
  
  1.最上层就不多加说明了,根据自己公司的消息队列获取到要推送的消息.
  
  2.消息收集器是一次获取批量的推送消息,基本公司使用像成熟的消息队列也可以不需要这一层,如果类似使用的redis队列,一次只能获取一个推送消息,就需要加这一层,获取一段时间或者一定量的推送消息,再推给下一层.
  ```go
  for   {
    //检验是否跳出循环,将数据给到下层
    nowTime := time.Now().UnixNano()
    if nowTime - beginTime > 150 * millisecond {
        if len(p_items) != 0 || len(g_items) != 0 || len(c_items) != 0 || len(s_items) != 0  {
            break
        }
        beginTime = time.Now().UnixNano()
    }else if len(p_items) >= 100 {
        break
    }else if len(g_items) >= 500 {
        break
    }else if len(c_items) >= 100 {
        break
    }else if len(s_items) >= 100 {
        break
    }
    
    /*从队列获取推送消息*/
    
 }
 //推给下一层消息收集器
 if len(p_items) != 0 {
 p_chan<-p_items
 }
 
 if len(g_items) != 0 {
 g_chan<-g_items
 }
 
 if len(c_items) != 0 {
 c_chan<-c_items
 }
 
 if len(s_items) != 0 {
 s_chan<-s_items
 }
 
```
  
  
  3.消息合并是用来合并群消息,系统消息.适用于同样的内容,推给不同的对象.
```go

for msgs := range g_chan {   //这里要小心,不要用到了defer,不然会一直都不释放.
    
    // 合并相同的群消息
    var msgs_dict = make(map[string]*group_msg)
    /*略*/
    for _, gmsg := range group_msgs {
    
        obj, ok := msgs_dict[gmsg.uuid]
        if ok {
        receivers := obj.receivers
        for i := 0; i < len(gmsg.receivers); i++ {
            //合并
            receivers = append(receivers, gmsg.receivers[i])
        }
        //保存在map
        obj.receivers = receivers
        msgs_dict[gmsg.uuid] = obj
    
        } else {
        msgs_dict[gmsg.uuid] = gmsg
        }
    }
    /*略*/
}
  
```
  4.过滤器是过滤一些敏感的词语,过滤一些特殊的消息体.
  
  5.推送模块,我们要讲的就是apns2.
  
  获取certificate文件:
  ```go
  func get_apns_cert(P12 string,P12_SECRET string) *tls.Certificate {
    cert, err := certificate.FromP12File(P12, P12_SECRET)
    if err != nil {
        log.Fatal("Cert Error:", err)
        return nil
    }
    return &cert
  }
```
  
  
  获取apns:
  ```go
  func (ap *Apns_push)get_apns_client(appid int64,P12 string,P12_SECRET string) *apns2.Client {
  
    if cer,ok := ap.Apns_cert_dict[appid];ok {
        client := ap.Client_manager.Get(cer) //Client_manager -> *apns2.ClientManager
        return client
    }
    cert := get_apns_cert(P12 ,P12_SECRET)
    if cert == nil {
        return nil
    }
    ap.Lock()
    ap.Apns_cert_dict[appid] = *cert
    ap.Unlock()
    client := apns2.NewClient(*cert)
    ap.Client_manager.Add(client) //Client_manager -> *apns2.ClientManager
    return client
  }
```
  
  

编辑apns消息体:
  ```go
func(ap *Apns_push)Apns_push(ams []*APNs_msg)  {
    var notifications []*apns2.Notification
    for _,am := range ams {

        p := payload.NewPayload()
        p.Badge(am.Badge)
        p.Alert(am.Content)
        if am.Sound == "" {
            am.Sound = "default"
        }
        p.ThreadID(am.Uuid)
        p.Sound(am.Sound)
        if am.Custom !=nil {
            for k,v := range am.Custom {
                p.Custom(k,v)
            }
        }

        notification := &apns2.Notification{
            DeviceToken:am.DeviceToken,
            Topic:am.Bundle_id,
            Payload:p,
        }
        notifications = append(notifications, notification)
    }
    ap.NotificationsCh<-notifications
}
```

推送:
```go
func send(client *apns2.Client,notification *apns2.Notification,isRepeat bool,seq int64)  {
    res, err := client.Development().Push(notification)
    if err != nil {
        log.Warnf("Push fail seq:%d -- %v %v %v   token:%s \n",seq, res.StatusCode, res.ApnsID, res.Reason ,notification.DeviceToken)
        if isRepeat {
            seq++
            send(client,notification,false,seq)
        }
        return
    }
    log.Infof("Push success seq:%d -- %v %v %v   token:%s \n",seq, res.StatusCode, res.ApnsID, res.Reason ,notification.DeviceToken)
}
```

## 参考资料

 [go-apns2](https://github.com/sideshow/apns2)
 
 [网络协议之TLS](https://www.cnblogs.com/syfwhu/p/5219814.html)
