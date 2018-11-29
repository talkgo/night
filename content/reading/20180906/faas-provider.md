---
title: faas-provider
date: 2018-08-01 19:53:23
tags:
---

faas-provider是一个模板，只要实现了这个模板的接口，就可以自定义实现自己的provider。

## faas-provider

OpenFaaS官方提供了两套后台provider：

- Docker Swarm
- Kubernetes

这两者在部署和调用函数的时候流程图如下：

部署一个函数

![](https://ws1.sinaimg.cn/large/b831e4c7gy1ftuggjuhtpj20xc07k0x7.jpg)

调用一个函数

![](https://ws1.sinaimg.cn/large/b831e4c7gy1ftuggyto07j20xc071n1c.jpg)

provider要提供的一些API有：

- List / Create / Delete 一个函数

`/system/functions`

方法: GET / POST / DELETE

- 获取一个函数

`/system/function/{name:[-a-zA-Z_0-9]+}`

方法: GET

- 伸缩一个函数

`/system/scale-function/{name:[-a-zA-Z_0-9]+}`

方法: POST

- 调用一个函数

`/function/{name:[-a-zA-Z_0-9]+}`

方法: POST

在provider的server.go的serve方法，可以看到这个serve方法创建了几个路由，接受一个FaaSHandler对象。

```go
// Serve load your handlers into the correct OpenFaaS route spec. This function is blocking.
func Serve(handlers *types.FaaSHandlers, config *types.FaaSConfig) {
	r.HandleFunc("/system/functions", handlers.FunctionReader).Methods("GET")
	r.HandleFunc("/system/functions", handlers.DeployHandler).Methods("POST")
	r.HandleFunc("/system/functions", handlers.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/system/functions", handlers.UpdateHandler).Methods("PUT")

	r.HandleFunc("/system/function/{name:[-a-zA-Z_0-9]+}", handlers.ReplicaReader).Methods("GET")
	r.HandleFunc("/system/scale-function/{name:[-a-zA-Z_0-9]+}", handlers.ReplicaUpdater).Methods("POST")

	r.HandleFunc("/function/{name:[-a-zA-Z_0-9]+}", handlers.FunctionProxy)
	r.HandleFunc("/function/{name:[-a-zA-Z_0-9]+}/", handlers.FunctionProxy)

	r.HandleFunc("/system/info", handlers.InfoHandler).Methods("GET")

	if config.EnableHealth {
		r.HandleFunc("/healthz", handlers.Health).Methods("GET")
	}
	// 省略
}
```

因此在自定义的provider，只需实现FaaSHandlers中的几个路由处理函数即可。这几个handler是：

```go
// FaaSHandlers provide handlers for OpenFaaS
type FaaSHandlers struct {
	FunctionReader http.HandlerFunc
	DeployHandler  http.HandlerFunc
	DeleteHandler  http.HandlerFunc
	ReplicaReader  http.HandlerFunc
	FunctionProxy  http.HandlerFunc
	ReplicaUpdater http.HandlerFunc

	// Optional: Update an existing function
	UpdateHandler http.HandlerFunc
	Health        http.HandlerFunc
	InfoHandler   http.HandlerFunc
}
```

我们以官方实现的faas-netes为例，讲解一下这几个hander的实现过程。

## faas-netes

我们看下在faas-netes的中的FaaSHandlers实现：

```go
bootstrapHandlers := bootTypes.FaaSHandlers{
	FunctionProxy:  handlers.MakeProxy(functionNamespace, cfg.ReadTimeout),
	DeleteHandler:  handlers.MakeDeleteHandler(functionNamespace, clientset),
	DeployHandler:  handlers.MakeDeployHandler(functionNamespace, clientset, deployConfig),
	FunctionReader: handlers.MakeFunctionReader(functionNamespace, clientset),
	ReplicaReader:  handlers.MakeReplicaReader(functionNamespace, clientset),
	ReplicaUpdater: handlers.MakeReplicaUpdater(functionNamespace, clientset),
	UpdateHandler:  handlers.MakeUpdateHandler(functionNamespace, clientset),
		Health:         handlers.MakeHealthHandler(),
		InfoHandler:    handlers.MakeInfoHandler(version.BuildVersion(), version.GitCommit),
}
```

因为是Kubernetes上的provider实现，所以这些函数都带有一个namespace的参数。

### FunctionProxy

这里最重要的就是FunctionProxy，它主要负责调用函数。这个handler其实也是起到了一个代理转发的作用，在这个函数中，只接受get和post。调用函数只接受post和get请求

1. 创建一个http的client对象

2. 只处理get和post请求。

3. 组装代理转发的watchdog的地址

   ```go
   url := forwardReq.ToURL(fmt.Sprintf("%s.%s", service, functionNamespace), watchdogPort)
   ```

   所以最后请求的格式就会形如：

   ```
   http://函数名.namespace:监视器的端口/路径
   ```

4. 将请求发出去

5. 设置http响应的头

### ReplicaReader和ReplicaUpdater

这两个是和副本数相关的，所以放在一起对比讲解。这两个的实现依赖于Kubernetes的客户端，获取代码如下：

```go
clientset, err := kubernetes.NewForConfig(config)
```

这个config主要满足以下几个条件就行：

```go
Config{
		// TODO: switch to using cluster DNS.
		Host:            "https://" + net.JoinHostPort(host, port),
		BearerToken:     string(token),
		TLSClientConfig: tlsClientConfig,
}
```

Kubernetes的所有操作都可以通过rest api来完成，这两个handler也是通过调用Kubernetes的api来做的。

#### ReplicaReader

`MakeReplicaReader`函数是获取当前的副本数：

1. 通过mux从路由中获取到name参数

2. 调用getService方法获取副本数，getService的核心代码就一句：

   ```go
   item, err := clientset.ExtensionsV1beta1().Deployments(functionNamespace).Get(functionName, getOpts)
   ```

3. 序列化之后，把结果返回

#### ReplicaUpdater

`MakeReplicaUpdater`是解析从gateway传过来的post请求，调用k8s的API设置副本数。

1. 从请求中取出body

2. 首先获取该函数的已部署的deployment对象

3. 然后将deployment的副本数量设置为应设数量，这样做的目的是为了仅仅修改副本数，而不修改别的属性。

   ```go
   _, err = clientset.ExtensionsV1beta1().Deployments(functionNamespace).Update(deployment)
   ```

> 注：mux做路由的时候，如果成功的时候不对w做任何处理，是会默认状态码为200，空字符串。

#### DeleteHandler，DeployHandler，FunctionReader和UpdateHandler

这几个都是对函数的操作，其实就是调用一下Kubernetes的API进行操作。

这几个是核心的几句代码：

```go
clientset.ExtensionsV1beta1().Deployments(functionNamespace).Delete(request.FunctionName, opts)

deploy := clientset.Extensions().Deployments(functionNamespace)

res, err := clientset.ExtensionsV1beta1().Deployments(functionNamespace).List(listOpts)

_, updateErr := clientset.CoreV1().Services(functionNamespace).Update(service)
```

## 总结

官方还提供了一个faas-swarm，其实现思路也是这样，操作swarm的api来做对容器的操作。至于如何调用一个函数，都是在函数的watchdog中实现。