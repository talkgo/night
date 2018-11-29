---
title:  2018-05-23 Get passes lock by value
---
来源：《Go 夜读》微信群

时间：2018-05-23

----

## 1. Get passes lock by value

```go
type ExecuterList struct {
	sync.Map
	length int
}

func (e ExecuterList) Get(key string) IExecuter {
	value, ok := e.Load(key)
	if !ok {
		return nil
	}
	if value == nil {
		return nil
	}
	res, _ := value.(IExecuter)
	return res
}
```

使用 `go tool vet` ，出现“Get passes lock by value: ExecuterList contains sync.Map contains sync.Mutex”， 解决方案有两种：。

1，sync.Map用指针

```go
type X struct {
	*sync.Map
}
```

2， 也可以用 `(e *ExecutorList)` ,避免锁的复制。

为什么会失效呢？

>因为你每次操作都复制了一遍整个struct，当然也复制了Map里面的Mutex，多线程同时读写时Map里面的锁相当于失效了。理解这个你需要知道的知识点有两个，一是go的参数都是值传递，二是只有用同一把锁才能对某个资源边界进行锁与解锁的操作。

## 2. RPC 微服务框架

rpc 可以封装程序间网络通信层，服务间调用只需要关注目标服务是什么，不用关心我到底用什么协议和数据格式了。

[为什么我推荐避免使用 go-kit 库？](https://gist.github.com/posener/330c2b08aaefdea6f900ff0543773b2e)
