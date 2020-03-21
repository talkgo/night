---
title: go mod下出现looping trying to add package
date: 2019-10-13T12:12:00+08:00
---
来源：『Go 夜读』微信群

时间：2019-10-13

---

## 问题

忽然之间出现looping trying to add package

意思就是循环引用某个类库，提示如下：

```text

go: github.com/sirupsen/logrus imports
	golang.org/x/sys/unix: looping trying to add package

```

## 解决方法

删除go mod下的缓存，重新引用。

```shell

sudo rm -fr "$(go env GOPATH)/pkg/mod"

```

[参考](https://github.com/rancher/k3s/issues/315)

```text

sudo rm -fr "$(go env GOPATH)/pkg/mod”
go get github.com/rancher/k3s@v.14.1-k3s.1
After going through a screenful of dependencies, the program aborts with following error.
....
....
go: downloading github.com/rancher/k3s v1.14.1-k3s.1
go: finding github.com/rancher/k3s v0.3.0
go: downloading github.com/rancher/k3s v0.3.0
go: import "github.com/rancher/k3s": looping trying to add package

```
