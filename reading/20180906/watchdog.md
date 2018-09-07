# watchdog

**监视器**

监视器提供了一个外部世界和函数之间的非托管的通用接口。它的工作是收集从API网关来的HTTP请求，然后调用程序。监视器是一个小型的Golang服务——下图展示了它是如何工作的：



![](https://ws1.sinaimg.cn/large/006tNbRwgy1fuzw3rkue0j30k00bamy2.jpg)



> 上图：一个小型的web服务，可以为每个传入的HTTP请求分配所需要的进程。

每个函数都需要嵌入这个二进制文件并将其作为`ENTRYPOINT` 或 `CMD`，实际上是把它作为容器的初始化进程。一旦你的进程被创建分支，监视器就会通过`stdin` 传递HTTP请求并从`stdout`中读取HTTP响应。这意味着你的程序无需知道web和HTTP的任何信息。

## **轻松创建新函数**

**从CLI创建一个函数**

创建函数最简单的方法是使用FaaS CLI和模板。CLI抽象了所有Docker的知识，使得你只需要编写所支持语言的handler文件即可。

- [你的第一个使用OpenFaaS的无服务器Python函数](https://link.zhihu.com/?target=https%3A//blog.alexellis.io/first-faas-python-function/)
- [阅读有关FaaS CLI的教程](https://link.zhihu.com/?target=https%3A//github.com/openfaas/faas-cli)

## **深入研究**

**Package your function打包你的函数**

如果你不想使用CLI或者现有的二进制文件或镜像，可以使用下面的方法去打包函数：

- 使用一个现有的或者一个新的Docker镜像作为基础镜像 `FROM`
- 通过`curl` 或 `ADD https://`从 [Releases 页面](https://link.zhihu.com/?target=https%3A//github.com/openfaas/faas/releases) 添加fwatchdog二进制文件
- 为每个你要运行的函数设置 `fprocess`(函数进程) 环境变量
- Expose port 8080
- 暴露端口8080
- Set the `CMD` to `fwatchdog`
- 设置 `CMD`为`fwatchdog`

一个`echo`函数的示例Dockerfile：

```text
FROM alpine:3.7

ADD https://github.com/openfaas/faas/releases/download/0.8.0/fwatchdog /usr/bin
RUN chmod +x /usr/bin/fwatchdog

# Define your binary here
ENV fprocess="/bin/cat"

CMD ["fwatchdog"]
```

**Implementing a Docker healthcheck实现一个Docker健康检查**

Docke的健康检查不是必需的，但是它是最佳实践。这会确保监视器已经在API网关转发请求之前准备好接收请求。如果函数或者监视器遇到一个不可恢复的问题，Swarm也会重启容器。

Here is an example of the `echo` function implementing a healthcheck with a 5-second checking interval.

下面是实现了一个具有5秒间隔的健康检查的`echo`函数示例：

```text
FROM functions/alpine

ENV fprocess="cat /etc/hostname"

HEALTHCHECK --interval=5s CMD [ -e /tmp/.lock ] || exit 1
```

监视器进程早启动内部Golang HTTP服务的时候会在 `/tmp/`下面创建一个.lock文件。`[ -e file_name ]`shell命令可以检查文件是否存在。在Windows容器中，这是一个不合法的路径，所以你可能需要设置`suppress_lock` 环境变量。

有关健康检查，请阅读我的Docker Swarm教程：

- [10分钟内试用Docker的健康检查](https://link.zhihu.com/?target=http%3A//blog.alexellis.io/test-drive-healthcheck/)

**环境变量重载:**

监视器可以通过环境变量来配置，你必须始终指定一个`fprocess` 变量

## **高级/调整**

## **(新)——子监视器和HTTP模式**

- 部分的监视器

为每个请求创建一个新的进程分支具有进程隔离，可移植和简单的优点。任何进程都可以在没有任何附加代码的情况下变成一个函数。of-watchdog可和HTTP模式是一种优化，这样就可以在所有请求之间维护一个单一的进程。

新版本的监视器正在[openfaas-incubator/of-watchdog](https://link.zhihu.com/?target=https%3A//github.com/openfaas-incubator/of-watchdog)上测试。

这种重写主要是生成一个可以持续维护的结构。它将会替代现有的监视器，也会有二进制的释放版。

## **使用HTTP头**

HTTP的头和其他请求信息以下面的格式注入到环境变量中：

```
X-Forwarded-By`头变成了`Http_X_Forwarded_By
```

- `Http_Method` - GET/POST etc
- `Http_Method` - GET/POST 等等
- `Http_Query` - QueryString value
- `Http_Query` - 查询字符串的值
- `Http_ContentLength` - gives the total content-length of the incoming HTTP request received by the watchdog.
- `Http_ContentLength` - 监视器收到的HTTP请求的内容长度。

> 默认情况下，通过`cgi_headers` 环境变量启用该行为。

以下是带有附加头和查询字符串的POST请求的示例：

```text
$ cgi_headers=true fprocess=env ./watchdog &
2017/06/23 17:02:58 Writing lock-file to: /tmp/.lock

$ curl "localhost:8080?q=serverless&page=1" -X POST -H X-Forwarded-By:http://my.vpn.com
```

如果你再Linux系统下设置了`fprocess` 到 `env`中，会看到如下结果：

```text
Http_User_Agent=curl/7.43.0
Http_Accept=*/*
Http_X_Forwarded_By=http://my.vpn.com
Http_Method=POST
Http_Query=q=serverless&page=1
```

也可以使用`GET`请求：

```text
$ curl "localhost:8080?action=quote&qty=1&productId=105"
```

监视器的输出如下：

```text
Http_User_Agent=curl/7.43.0
Http_Accept=*/*
Http_Method=GET
Http_Query=action=quote&qty=1&productId=105
```

现在就可以在程序中使用HTTP状态来做决策了。

## **HTTP方法**

监视器支持的HTTP方法有：

带有请求体的：

- POST, PUT, DELETE, UPDATE

不带请求体的：

- GET

> API网关现在支持函数的POST路由。

## **请求响应的内容类型**

默认情况下，监视器会匹配客户端的"Content-Type"。

- 如果客户端发送Content-Type 为 `application/json` 的json形式的post请求，将会在响应的时候自动匹配。
- 如果客户端发送Content-Type 为 `text/plain` 的json形式的post请求，响应也会自动匹配。

若要重载所有响应的Content-Type ，需要设置`content_type` 环境变量。

## **I don't want to use the watchdog**

## **我不想使用监视器**

这种案例是OpenFaaS所不支持的，但是如果你的容器符合以下要求，那么OpenFaaS的网关和其他工具也会管理和伸缩服务。

你需要提供一个锁文件 `/tmp/.lock`，以便业务流程系统可以在容器中运行健康检查。如果你正在使用swarm，那么请确保在Dockerfile中提供`HEALTHCHECK`指令——在 `faas`存储库中有示例。

- 在HTTP之上暴露TCP端口8080
- 创建`/tmp/.lock` 文件，或者在响应操作tempdir系统调用的任何位置。

## **调整自动伸缩**

自动伸缩式从1个副本开始，以5个位一个单位进行升级：

- 1->5
- 5->10
- 10->15
- 15->20

你可以通过标签来覆盖一个函数minimum 和 maximum 。

如果要在2到15之间的话，请在部署的时候配置以下标签：

```text
com.openfaas.scale.min: "2"
com.openfaas.scale.max: "15"
```

这些标签是可选的

**禁用自动伸缩**

如果要禁用某个函数的自动伸缩，将最小和最大的副本数设置为相同的值，即“1”。

同样也可以删除AlertManager。