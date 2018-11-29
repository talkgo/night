---
title: gateway-reading
date: 2018-08-01 09:15:35
tags:
---

OpenFaaS的Gateway是一个golang实现的请求转发的网关，在这个网关服务中，主要有以下几个功能：

- UI
- 部署函数
- 监控
- 自动伸缩

## 架构分析

![图：Kubernetes作为Provider的架构图](https://ws1.sinaimg.cn/large/b831e4c7gy1fttxzimdp3j20sm0hqabx.jpg)

从图中可以发现，当Gateway作为一个入口，当CLI或者web页面发来要部署或者调用一个函数的时候，Gateway会将请求转发给Provider，同时会将监控指标发给Prometheus。AlterManager会根据需求，调用API自动伸缩函数。

## 源码分析

### 依赖

```go
github.com/gorilla/mux

github.com/nats-io/go-nats-streaming
github.com/nats-io/go-nats
github.com/openfaas/nats-queue-worker

github.com/prometheus/client_golang
```

mux 是一个用来执行http请求的路由和分发的第三方扩展包。

go-nats-streaming，go-nats，nats-queue-worker这三个依赖是异步函数的时候才会用到，在分析queue-worker的时候有说到Gateway也是一个发布者。

client_golang是Prometheus的客户端。

### 项目结构

```bash
├── Dockerfile
├── Dockerfile.arm64
├── Dockerfile.armhf
├── Gopkg.lock
├── Gopkg.toml
├── README.md
├── assets
├── build.sh
├── handlers
│   ├── alerthandler.go
│   ├── alerthandler_test.go
│   ├── asyncreport.go
│   ├── baseurlresolver_test.go
│   ├── basic_auth.go
│   ├── basic_auth_test.go
│   ├── callid_middleware.go
│   ├── cors.go
│   ├── cors_test.go
│   ├── forwarding_proxy.go
│   ├── forwarding_proxy_test.go
│   ├── function_cache.go
│   ├── function_cache_test.go
│   ├── infohandler.go
│   ├── metrics.go
│   ├── queueproxy.go
│   ├── scaling.go
│   └── service_query.go
├── metrics
│   ├── add_metrics.go
│   ├── add_metrics_test.go
│   ├── externalwatcher.go
│   ├── metrics.go
│   └── prometheus_query.go
├── plugin
│   ├── external.go
│   └── external_test.go
├── queue
│   └── types.go
├── requests
│   ├── forward_request.go
│   ├── forward_request_test.go
│   ├── prometheus.go
│   ├── prometheus_test.go
│   └── requests.go
├── server.go
├── tests
│   └── integration
├── types
│   ├── handler_set.go
│   ├── inforequest.go
│   ├── load_credentials.go
│   ├── proxy_client.go
│   ├── readconfig.go
│   └── readconfig_test.go
├── vendor
│   └── github.com
└── version
    └── version.go
```

Gateway的目录明显多了很多，看源码的时候，首先要找到的是main包，从main函数看起，就能很容易分析出来项目是如何运行的。

从server.go的main函数中我们可以看到，其实有如下几个模块：

- 基本的安全验证
- 和函数相关的代理转发
  - 同步函数
    - 列出函数
    - 部署函数
    - 删除函数
    - 更新函数
  - 异步函数
- Prometheus的监控
- ui
- 自动伸缩

### 基本的安全验证

如果配置了开启基本安全验证，会从磁盘中读取密钥：

```go
var credentials *types.BasicAuthCredentials

if config.UseBasicAuth {
	var readErr error
	reader := types.ReadBasicAuthFromDisk{
		SecretMountPath: config.SecretMountPath,
	}
	credentials, readErr = reader.Read()

	if readErr != nil {
		log.Panicf(readErr.Error())
	}
}
```

在Gateway的配置相关的，都会有一个read()方法，进行初始化赋值。

如果credentials被赋值之后，就会对一些要加密的API handler进行一个修饰，被修饰的API有：

- UpdateFunction
- DeleteFunction
- DeployFunction
- ListFunctions
- ScaleFunction

```go
if credentials != nil {
	faasHandlers.UpdateFunction =
			handlers.DecorateWithBasicAuth(faasHandlers.UpdateFunction, credentials)
	faasHandlers.DeleteFunction =
			handlers.DecorateWithBasicAuth(faasHandlers.DeleteFunction, credentials)
	faasHandlers.DeployFunction =
			handlers.DecorateWithBasicAuth(faasHandlers.DeployFunction, credentials)
	faasHandlers.ListFunctions =
			handlers.DecorateWithBasicAuth(faasHandlers.ListFunctions, credentials)
	faasHandlers.ScaleFunction =
			handlers.DecorateWithBasicAuth(faasHandlers.ScaleFunction, credentials)
}
```

这个DecorateWithBasicAuth()方法是一个路由中间件：

1. 调用mux路由的BasicAuth()，从http的header中取到用户名和密码
2. 然后给请求头上设置一个字段`WWW-Authenticate`，值为`Basic realm="Restricted"`
3. 如果校验失败，则返回错误，成功的话调用next方法继续进入下一个handler。

```go
// DecorateWithBasicAuth enforces basic auth as a middleware with given credentials
func DecorateWithBasicAuth(next http.HandlerFunc, credentials *types.BasicAuthCredentials) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, password, ok := r.BasicAuth()
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

		if !ok || !(credentials.Password == password && user == credentials.User) {

			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("invalid credentials"))
			return
		}
		next.ServeHTTP(w, r)
	}
}
```

### 代理转发

Gateway本身不做任何和部署发布函数的事情，它只是作为一个代理，把请求转发给相应的Provider去处理，所有的请求都要通过这个网关。

#### 同步函数转发

主要转发的API有：

- RoutelessProxy
- ListFunctions
- DeployFunction
- DeleteFunction
- UpdateFunction

```go
faasHandlers.RoutelessProxy = handlers.MakeForwardingProxyHandler(reverseProxy, forwardingNotifiers, urlResolver)
	faasHandlers.ListFunctions = handlers.MakeForwardingProxyHandler(reverseProxy, forwardingNotifiers, urlResolver)
	faasHandlers.DeployFunction = handlers.MakeForwardingProxyHandler(reverseProxy, forwardingNotifiers, urlResolver)
	faasHandlers.DeleteFunction = handlers.MakeForwardingProxyHandler(reverseProxy, forwardingNotifiers, urlResolver)
	faasHandlers.UpdateFunction = handlers.MakeForwardingProxyHandler(reverseProxy, forwardingNotifiers, urlResolver)
```

MakeForwardingProxyHandler()有三个参数：

- proxy

  这是一个http的客户端，作者把这个客户端抽成一个类，然后使用该类的NewHTTPClientReverseProxy方法创建实例，这样就简化了代码，不用每次都得写一堆相同的配置。

- notifiers

  这个其实是要打印的日志，这里是一个HTTPNotifier的接口。而在这个MakeForwardingProxyHandler中其实有两个实现类，一个是LoggingNotifier，一个是PrometheusFunctionNotifier，分别用来打印和函数http请求相关的日志以及和Prometheus监控相关的日志。

- baseURLResolver

  这个就是Provider的url地址。

在这个MakeForwardingProxyHandler中主要做了三件事儿：

1. 解析要转发的url

2. 调用forwardRequest方法转发请求，

   forwardRequest方法的逻辑比较简单，只是把请求发出去。这里就不深入分析了。

3. 打印日志

```go
// MakeForwardingProxyHandler create a handler which forwards HTTP requests
func MakeForwardingProxyHandler(proxy *types.HTTPClientReverseProxy, notifiers []HTTPNotifier, baseURLResolver BaseURLResolver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		baseURL := baseURLResolver.Resolve(r)
		requestURL := r.URL.Path
		start := time.Now()
		statusCode, err := forwardRequest(w, r, proxy.Client, baseURL, requestURL, proxy.Timeout)
		seconds := time.Since(start)
		if err != nil {
			log.Printf("error with upstream request to: %s, %s\n", requestURL, err.Error())
		}
		for _, notifier := range notifiers {
			notifier.Notify(r.Method, requestURL, statusCode, seconds)
		}
	}
}
```

#### 异步函数转发

前面说过，如果是异步函数，Gateway就作为一个发布者，将函数放到队列里。MakeQueuedProxy方法就是做这件事儿的：

1. 读取请求体
2. 将`X-Callback-Url`参数从参数中http的header中读出来
3. 实例化用于异步处理的Request对象
4. 调用canQueueRequests.Queue(req)，将请求发布到队列中

```go
// MakeQueuedProxy accepts work onto a queue
func MakeQueuedProxy(metrics metrics.MetricOptions, wildcard bool, canQueueRequests queue.CanQueueRequests) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		// 省略错误处理代码
		vars := mux.Vars(r)
		name := vars["name"]

		callbackURLHeader := r.Header.Get("X-Callback-Url")
		var callbackURL *url.URL
		if len(callbackURLHeader) > 0 {
			urlVal, urlErr := url.Parse(callbackURLHeader)
			// 省略错误处理代码
			callbackURL = urlVal
		}
		req := &queue.Request{
			Function:    name,
			Body:        body,
			Method:      r.Method,
			QueryString: r.URL.RawQuery,
			Header:      r.Header,
			CallbackURL: callbackURL,
		}
		err = canQueueRequests.Queue(req)
		// 省略错误处理代码
		w.WriteHeader(http.StatusAccepted)
	}
}
```

#### 自动伸缩

伸缩性其实有两种，一种是可以通过调用API接口，来将函数进行缩放。另外一种就是通过AlertHandler。

自动伸缩是OpenFaaS的一大特点，触发自动伸缩主要是根据不同的指标需求。

- 根据每秒请求数来做伸缩

  OpenFaaS附带了一个自动伸缩的规则，这个规则是在AlertManager配置文件中定义。AlertManager从Prometheus中读取使用情况（每秒请求数），然后在满足一定条件时向Gateway发送警报。

  可以通过删除AlertManager，或者将部署扩展的环境变量设置为0，来禁用此方式。

- 最小/最大副本数

  通过向函数添加标签, 可以在部署时设置最小 (初始) 和最大副本数。

  - `com.openfaas.scale.min` 默认是 `1`
  - `com.openfaas.scale.max` 默认是 `20` 
  - `com.openfaas.scale.factor` 默认是 `20%` ，在0-100之间，这是每次扩容的时候，新增实例的百分比，若是100的话，会瞬间飙升到副本数的最大值。

  `com.openfaas.scale.min` 和 `com.openfaas.scale.max`值一样的时候，可以关闭自动伸缩。

  `com.openfaas.scale.factor`是0时，也会关闭自动伸缩。

- 通过内存和CPU的使用量。

  使用k8s内置的HPA，也可以触发AlertManager。

##### 手动指定伸缩的值

可以从这句代码中发现，调用这个路由，转发给了provider处理。

```
r.HandleFunc("/system/scale-function/{name:[-a-zA-Z_0-9]+}", faasHandlers.ScaleFunction).Methods(http.MethodPost)
```
##### 处理AlertManager的伸缩请求

Prometheus将监控指标发给AlertManager之后，会触发AlterManager调用`/system/alert`接口，这个接口的handler是由`handlers.MakeAlertHandler`方法生成。

MakeAlertHandler方法接收的参数是ServiceQuery。ServiceQuery是一个接口，它有两个函数，用来get或者ser最大的副本数。Gateway中实现这个接口的类是ExternalServiceQuery，这个实现类是在plugin包中，我们也可以直接定制这个实现类，用来实现满足特定条件。

```go
// ServiceQuery provides interface for replica querying/setting
type ServiceQuery interface {
	GetReplicas(service string) (response ServiceQueryResponse, err error)
	SetReplicas(service string, count uint64) error
}

// ExternalServiceQuery proxies service queries to external plugin via HTTP
type ExternalServiceQuery struct {
	URL         url.URL
	ProxyClient http.Client
}
```

这个ExternalServiceQuery有一个`NewExternalServiceQuery`方法，这个方法也是一个工厂方法，用来创建实例。这个url其实就是provider的url，proxyClient就是一个http的client对象。

- `GetReplicas`方法

  从`system/function/:name`接口获取到函数的信息，组装一个`ServiceQueryResponse`对象即可。

- `SetReplicas`方法

  调用`system/scale-function/:name`接口，设置副本数。

MakeAlertHandler的函数主要是从`http.Request`中读取body，然后反序列化成`PrometheusAlert`对象：

```go
// PrometheusAlert as produced by AlertManager
type PrometheusAlert struct {
	Status   string                 `json:"status"`
	Receiver string                 `json:"receiver"`
	Alerts   []PrometheusInnerAlert `json:"alerts"`
}
```

可以发现，这个Alerts是一个数组对象，所以可以是对多个函数进行缩放。反序列化之后，调用`handleAlerts`方法，而`handleAlerts`对Alerts进行遍历，针对每个Alerts调用了`scaleService`方法。`scaleService`才是真正处理伸缩服务的函数。

```go
func scaleService(alert requests.PrometheusInnerAlert, service ServiceQuery) error {
	var err error
	serviceName := alert.Labels.FunctionName

	if len(serviceName) > 0 {
		queryResponse, getErr := service.GetReplicas(serviceName)
		if getErr == nil {
			status := alert.Status

			newReplicas := CalculateReplicas(status, queryResponse.Replicas, uint64(queryResponse.MaxReplicas), queryResponse.MinReplicas, queryResponse.ScalingFactor)

			log.Printf("[Scale] function=%s %d => %d.\n", serviceName, queryResponse.Replicas, newReplicas)
			if newReplicas == queryResponse.Replicas {
				return nil
			}

			updateErr := service.SetReplicas(serviceName, newReplicas)
			if updateErr != nil {
				err = updateErr
			}
		}
	}
	return err
}
```

从代码总就可以看到，scaleService做了三件事儿：

- 获取现在的副本数

- 计算新的副本数

  新副本数的计算方法是根据`com.openfaas.scale.factor`计算步长：

  ```go
  step := uint64((float64(maxReplicas) / 100) * float64(scalingFactor))
  ```

- 设置为新的副本数

##### 从0增加副本到的最小值

我们在调用函数的时候，用的路由是：`/function/:name`。如果环境变量里有配置`scale_from_zero`为true，先用`MakeScalingHandler()`方法对proxyHandler进行一次包装。

`MakeScalingHandler`接受参数主要是：

- next：就是下一个httpHandlerFunc，中间件都会有这样一个参数

- config：`ScalingConfig`的对象：

  ```go
  // ScalingConfig for scaling behaviours
  type ScalingConfig struct {
  	MaxPollCount         uint              // 查到的最大数量
  	FunctionPollInterval time.Duration     // 函数调用时间间隔
  	CacheExpiry          time.Duration     // 缓存过期时间
  	ServiceQuery         ServiceQuery      // 外部服务调用的一个接口
  }
  ```

这个`MakeScalingHandler`中间件主要做了如下的事情：

   - 先从FunctionCache缓存中获取该函数的基本信息，从这个缓存可以拿到每个函数的副本数量。
   - 为了加快函数的启动速度，如果缓存中可以获该得函数，且函数的副本数大于0，满足条件，return即可。
   - 如果不满足上一步，就会调用`SetReplicas`方法设置副本数，并更新FunctionCache的缓存。

```go
// MakeScalingHandler creates handler which can scale a function from
// zero to 1 replica(s).
func MakeScalingHandler(next http.HandlerFunc, upstream http.HandlerFunc, config ScalingConfig) http.HandlerFunc {
	cache := FunctionCache{
		Cache:  make(map[string]*FunctionMeta),
		Expiry: config.CacheExpiry,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		functionName := getServiceName(r.URL.String())
		if serviceQueryResponse, hit := cache.Get(functionName); hit && serviceQueryResponse.AvailableReplicas > 0 {
			next.ServeHTTP(w, r)
			return
		} 
		queryResponse, err := config.ServiceQuery.GetReplicas(functionName)
		cache.Set(functionName, queryResponse)
        // 省略错误处理
		if queryResponse.AvailableReplicas == 0 {
			minReplicas := uint64(1)
			if queryResponse.MinReplicas > 0 {
				minReplicas = queryResponse.MinReplicas
			}
			err := config.ServiceQuery.SetReplicas(functionName, minReplicas)
			// 省略错误处理代码
			for i := 0; i < int(config.MaxPollCount); i++ {
				queryResponse, err := config.ServiceQuery.GetReplicas(functionName)
				cache.Set(functionName, queryResponse)
				// 省略错误处理
				time.Sleep(config.FunctionPollInterval)
			}
		}
		next.ServeHTTP(w, r)
	}
}
```

### 监控

监控是一个定时任务，开启了一个新协程，利用go的ticker.C的间隔不停的去调用`/system/functions`接口。反序列化到MetricOptions对象中。

```go
func AttachExternalWatcher(endpointURL url.URL, metricsOptions MetricOptions, label string, interval time.Duration) {
	ticker := time.NewTicker(interval)
	quit := make(chan struct{})
	proxyClient := // 省略创建一个http.Client对象

	go func() {
		for {
			select {
			case <-ticker.C:
				get, _ := http.NewRequest(http.MethodGet, endpointURL.String()+"system/functions", nil)
				services := []requests.Function{}
				res, err := proxyClient.Do(get)
				// 省略反序列的代码
				for _, service := range services {
					metricsOptions.ServiceReplicasCounter.
						WithLabelValues(service.Name).
						Set(float64(service.Replicas))
				}
				break
			case <-quit:
				return
			}
		}
	}()
}
```

### UI

UI的代码很简单，主要就是一些前端的代码，调用上面的讲的一些API接口即可，这里就略去不表。

## 总结

Gateway是OpenFaaS最为重要的一个组件。回过头看整个项目的结构，Gateway就是一个rest转发服务，一个一个的handler，每个模块之间的耦合性不是很高，可以很容易的去拆卸，自定义实现相应的模块。