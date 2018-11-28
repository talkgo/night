# queue-worker源码分析

## **异步函数和同步函数**

在OpenFaaS中同步调用函数时，将会连接到网关，直到函数成功返回才会关闭连接。同步调用是阻塞的。

- 网关的路由是：`/function/<function_name>`
- 必须等待
- 在结束的时候得到结果
- 明确知道是成功还是失败

异步函数会有一些差异：

- 网关的路由是：`/async-function/<function_name>`
- 客户端获得202的即时响应码
- 从queue-worker中调用函数
- 默认情况下，结果是被丢弃的。

## **查看queue-worker的日志**

```text
docker service logs -f func_queue-worker
```

## **利用requestbin和X-Callback-Url获取异步函数的结果**

如果需要获得异步函数的结果，有两个方法：

- 更改代码，将结果返回给端点或者消息系统
- 利用内置的回调
  内置的回调将会允许函数提供一个url，queue-worker会报告函数的成功或失败。
  requestbin会创建一个新的bin，这是互联网的一个url地址，可以从这里获取函数的结果。



![](https://ws4.sinaimg.cn/large/006tNbRwgy1fuzvzbcz2vj30k00b3aar.jpg)





![](https://ws3.sinaimg.cn/large/006tNbRwgy1fuzvztfu66j30k00b2757.jpg)





![](https://ws4.sinaimg.cn/large/006tNbRwgy1fuzw08o8w3j30k00b1aay.jpg)



## **源码分析**

## **依赖项**

```go
github.com/nats-io/go-nats-streaming
github.com/nats-io/go-nats

github.com/openfaas/faas
```

go-nats和go-nats-streaming是nats和nats-streaming的go版本的客户端。

faas这个依赖其实是只用到了queue包下面的types.go文件。这个文件是定义了异步请求的Request结构体和一个CanQueueRequests接口。如下所示：

```go
package queue

import "net/url"
import "net/http"

// Request for asynchronous processing
type Request struct {
    Header      http.Header
    Body        []byte
    Method      string
    QueryString string
    Function    string
    CallbackURL *url.URL `json:"CallbackUrl"`
}

// CanQueueRequests can take on asynchronous requests
type CanQueueRequests interface {
    Queue(req *Request) error
}
```

从这里我们就可以明白作者的设计思路，只要是实现了这个CanQueueRequests接口，就可以作为一个queue-worker。

## **接口实现类NatsQueue**

接口的实现类NatsQueue是在handler包里。它的属性都是nats中常用到的，包括clientId，clusterId，url，连接，主题等，如下所示：

```go
// NatsQueue queue for work
type NatsQueue struct {
    nc        stan.Conn    // nats的连接
    ClientID  string       // nats的clientId
    ClusterID string       // nats的clusterId
    NATSURL   string       // nats的URL
    Topic     string       // 主题
}
```

它的queue方法也很简单，主要做了两件事儿：

1. 解析传入的Request对象，并转为json对象out
2. 将out发布到队列里

```go
// Queue request for processing
func (q *NatsQueue) Queue(req *queue.Request) error {
    var err error

    fmt.Printf("NatsQueue - submitting request: %s.\n", req.Function)

    out, err := json.Marshal(req)
    if err != nil {
        log.Println(err)
    }

    err = q.nc.Publish(q.Topic, out)

    return err
}
```

go语言没有构造方法，所以NatsQueue还用于创建NatsQueue的实例的方法，这里就成为工厂方法。这个工厂方法主要就是从配置文件中读取环境变量的值，然后创建一个nats的连接，相当于给NatsQueue的对象的每个属性进行赋值。

```go
func CreateNatsQueue(address string, port int, clientConfig NatsConfig) (*NatsQueue, error) {
    queue1 := NatsQueue{}
    var err error
    natsURL := fmt.Sprintf("nats://%s:%d", address, port)
    log.Printf("Opening connection to %s\n", natsURL)

    clientID := clientConfig.GetClientID()
    clusterID := "faas-cluster"

    nc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
    queue1.nc = nc

    return &queue1, err
}
```

这个CreateNatsQueue方法是Gateway项目中进行调用，我们可以在Gateway项目的main.go中找到，如果Gateway的配置开启了异步函数支持，就会调用该方法，创建一个NatsQueue对象，然后把函数放到队列中，这里就不深入讲解：

```go
if config.UseNATS() {
        log.Println("Async enabled: Using NATS Streaming.")
        natsQueue, queueErr := natsHandler.CreateNatsQueue(*config.NATSAddress, *config.NATSPort, natsHandler.DefaultNatsConfig{})
        if queueErr != nil {
            log.Fatalln(queueErr)
        }

        faasHandlers.QueuedProxy = handlers.MakeQueuedProxy(metricsOptions, true, natsQueue)
        faasHandlers.AsyncReport = handlers.MakeAsyncReport(metricsOptions)
}
```

到这里，我相信读者也了解到，Gateway其实就是一个发布者，将异步请求扔到队列里。接下来肯定要有一个订阅者将请求消费处理。

## **订阅者处理**

我们都知道，nats streaming的订阅者订阅到消息之后，会把消息扔给一个回调函数去处理。queue-worker的订阅者实现也是这样，它的实现并不复杂，所有逻辑都在main.go的中。

我们先看回调函数mcb都做了什么：

1. 首先当然是将消息体反序列化成上面说到的用于异步处理的Request对象。
2. 构造http请求的url和querystring，url的格式如下：
   functionURL := fmt.Sprintf("http://%s%s:8080/%s", req.Function, config.FunctionSuffix, queryString)
3. 设置http的header，并以post的形式向functionURL发起请求。
4. 如果请求失败，设置返回状态码为`http.StatusServiceUnavailable`，并分别处理CallbackURL是否存在的情况。
5. 如果请求成功，同样也是要分别处理CallbackURL是否存在的情况。

当然在这个callback中会根据一些环境变量的存在，选择是否打印日志出来。

```go
mcb := func(msg *stan.Msg) {
        i++

        printMsg(msg, i)

        started := time.Now()

        req := queue.Request{}
        unmarshalErr := json.Unmarshal(msg.Data, &req)

        if unmarshalErr != nil {
            log.Printf("Unmarshal error: %s with data %s", unmarshalErr, msg.Data)
            return
        }

        fmt.Printf("Request for %s.\n", req.Function)

        if config.DebugPrintBody {
            fmt.Println(string(req.Body))
        }

        queryString := ""
        if len(req.QueryString) > 0 {
            queryString = fmt.Sprintf("?%s", strings.TrimLeft(req.QueryString, "?"))
        }

        functionURL := fmt.Sprintf("http://%s%s:8080/%s", req.Function, config.FunctionSuffix, queryString)

        request, err := http.NewRequest(http.MethodPost, functionURL, bytes.NewReader(req.Body))
        defer request.Body.Close()

        copyHeaders(request.Header, &req.Header)

        res, err := client.Do(request)
        var status int
        var functionResult []byte

        if err != nil {
            status = http.StatusServiceUnavailable

            log.Println(err)
            timeTaken := time.Since(started).Seconds()

            if req.CallbackURL != nil {
                log.Printf("Callback to: %s\n", req.CallbackURL.String())

                resultStatusCode, resultErr := postResult(&client, res, functionResult, req.CallbackURL.String())
                if resultErr != nil {
                    log.Println(resultErr)
                } else {
                    log.Printf("Posted result: %d", resultStatusCode)
                }
            }

            statusCode, reportErr := postReport(&client, req.Function, status, timeTaken, config.GatewayAddress)
            if reportErr != nil {
                log.Println(reportErr)
            } else {
                log.Printf("Posting report - %d\n", statusCode)
            }
            return
        }

        if res.Body != nil {
            defer res.Body.Close()

            resData, err := ioutil.ReadAll(res.Body)
            functionResult = resData

            if err != nil {
                log.Println(err)
            }

            if config.WriteDebug {
                fmt.Println(string(functionResult))
            } else {
                fmt.Printf("Wrote %d Bytes\n", len(string(functionResult)))
            }
        }

        timeTaken := time.Since(started).Seconds()

        fmt.Println(res.Status)

        if req.CallbackURL != nil {
            log.Printf("Callback to: %s\n", req.CallbackURL.String())
            resultStatusCode, resultErr := postResult(&client, res, functionResult, req.CallbackURL.String())
            if resultErr != nil {
                log.Println(resultErr)
            } else {
                log.Printf("Posted result: %d", resultStatusCode)
            }
        }

        statusCode, reportErr := postReport(&client, req.Function, res.StatusCode, timeTaken, config.GatewayAddress)

        if reportErr != nil {
            log.Println(reportErr)
        } else {
            log.Printf("Posting report - %d\n", statusCode)
        }
}
```

`postResult`函数是用来处理callbackURL存在的情况，在这个函数中将结果，以post请求调用callbackURL发送出去。

`postReport`函数用来处理callbackURL不存在的情况，这里是将结果发到Gateway网关的`"http://" + gatewayAddress + ":8088/system/async-report"`中，我们之后就可以从这个url里查询异步函数的执行结果了。

## **总结**

本文主要分析了NATS Streaming版本的queue worker的实现，通过分析源码我们可以看到OpenFaaS在架构的设计很有考究，充分的考虑到了可扩展性，通过定义接口规范，使得开发者很容易实现自定义。