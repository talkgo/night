---
title: Go 快速入门
---
创建一个新函数

```
faas-cli new --lang node hell-node
```

构建函数

```
faas-cli build -f hello-node.yml
```

推送函数到docker仓库

```
faas-cli push -f hello-node.yml
```

部署函数

```
faas-cli deploy -f hello-node.yml
```

稍等几秒钟，等待部署，然后就可以从postman发送get或者post请求。

![img](https://ws3.sinaimg.cn/large/006tNbRwgy1fusv6plgoxj31kw0vidom.jpg)

在rancher中的状态

![](https://ws2.sinaimg.cn/large/006tNbRwgy1fuzw7a38gbj31kw0gfaei.jpg)

函数的状态

![image-20180906162055229](/var/folders/c_/jg_300b169qggntnwc88n1g80000gn/T/abnerworks.Typora/image-20180906162055229.png)



