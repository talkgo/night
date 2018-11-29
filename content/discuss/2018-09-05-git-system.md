---
title: 2018-09-05 微信讨论
---
来源：《Go 夜读》微信群

### [xxxxx [1 2 3 4]] 如何做到输出为[ssss 1 2 3 4]

```golang
func main() {
	MutilParam("ssss", 1, 2, 3, 4) //[ssss 1 2 3 4]
	iis := []int{1, 2, 3, 4}
	MutilParam("xxxxx", iis) //MutilParam= [xxxxx [1 2 3 4]] 如何做到输出为[xxxx 1 2 3 4]
}

func MutilParam(p ...interface{}) {
	fmt.Println("MutilParam=", p)
}
```
考点：**函数变参**
这样的情况会在开源类库如xorm升级版本后出现Exce函数不兼容的问题。
解决方式有两个：
```go
//方法一：interface[]
tmpParams := make([]interface{}, 0, len(iis)+1)
tmpParams = append(tmpParams, "ssss")
for _, ii := range iis {
    tmpParams = append(tmpParams, ii)
}
MutilParam(tmpParams...)
//方法二:反射
f := MutilParam
value := reflect.ValueOf(f)
pps := make([]reflect.Value, 0, len(iis)+1)
pps = append(pps, reflect.ValueOf("ssss"))
for _, ii := range iis {
    pps = append(pps, reflect.ValueOf(ii))
}
value.Call(pps)
```
实际上都是一样原理的，也很简单的方式处理了。

### git 系统平台的选择？

git 系统都有哪一些？

- Github
- Gitlab
- gogs
- Gitea
- Bitbucket
- 码云
- Coding
- ？

更多讨论：

- Gitlab 资源占用太夸张了，但是确实 Gitlab 做的比较好。

>云服务器至少得 4核8G

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/2018-09-05-gitlab.png)

- gogs 有一个超级大的痛点：CodeReview，但是优点也很突出：速度快，资源占用少。

>无法选中一行代码发表评论/注释。

- Coding 企业应该都不太敢使用吧。

### 时序序列告警框架

- bosun 是 Time Series Alerting Framework，比较轻量级，但是文档严重滞后。

- Prometheus 是用 opentsdb，有点重，还要配上 hbase，玩不转啊。

>由于我们是物联网的项目，监控的指标比较多样化，tag的不确定性，应该会导致Prometheus内存耗费很大，而且prometheus采用pull的模式，虽然有pushgateway，还是挺担心的，用了prometheus后，发现业务层的监控代码也会要写的很复杂。

## 参考资料

- [Enum value with index 0 not shown #3808](https://github.com/protocolbuffers/protobuf/issues/3808)
- [proto3#enum](https://developers.google.com/protocol-buffers/docs/proto3#enum)
- https://github.com/bosun-monitor/bosun/issues/2301
