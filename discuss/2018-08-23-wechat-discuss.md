## 2018-08-23

来源：《Go 夜读》微信群

时间：2018-08-23

### 有什么好的博客平台吗？（除了简书）

- Github pages
- hugo + Github
- 自建
- jekyll
- Ghost
- hexo + Github
- hexo + Coding pages

图片CDN：七牛

### 请问nats有类似[kafka manager](https://github.com/yahoo/kafka-manager)这样的管理后台吗？

- kafka manager 的管理后台：kafkatool

### 对kafka的抽象

![](../images/2018-08-23-kafka-producer-consumer.png)

kafka聚焦于数据管道，nats聚焦于message bus

### for-select ?

下面这段代码为什么运行一会儿就停止了呢?

![](../images/2018-08-23-for-select.png)

用 waitgroup 肯定是正常的，研究一个 go sceduler 的问题，所以才故意这样写的。

for {}这种死循环（现实中应该用不着，反正我没用过这种需求），编译的时候不会被go插入抢占的代码（morestack 函数），导致调度切换不出去。这个问题go团队已经在着手解决了 [golang/go/issues#24543](https://github.com/golang/go/issues/24543)，不理解的话，可以看看:[](https://tonybai.com/2017/11/23/the-simple-analysis-of-goroutine-schedule-examples/)

加打印会有系统调用，就会插入调度抢占代码，抢占调度的前提是需要插入morestackt代码，和编译器和调度机制都有关系。。

go 中加入 spinning threads 的目的是啥，有相关资料吗？
>让 M 工作。

![](../images/2018-08-23-spinning.png)

spin 和 unspin 对应就是 M 运行和休眠的状态，也就是线程运行和休眠的状态，M 只有实在找不多 G 来做的时候才会休眠。

## 参考

1. [系统设计入门](https://github.com/donnemartin/system-design-primer/blob/master/README-zh-Hans.md)
2. [https://stackshare.io/stackups/kafka-vs-nsq-vs-rabbitmq](https://stackshare.io/stackups/kafka-vs-nsq-vs-rabbitmq)
3. [https://github.com/golang/go/issues/15442](https://github.com/golang/go/issues/15442)
4. [https://rakyll.org/scheduler/](https://rakyll.org/scheduler/)

