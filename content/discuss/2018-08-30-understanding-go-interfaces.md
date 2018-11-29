---
title: 2018-08-30 理解 Go interface
---
来源：《Go 夜读》微信群

时间：2018-08-30

### Go 语言不同接口、声明了同名方法，怎么解决问题？

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/2018-08-30-interface.png)
![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/2018-08-30-interface2.png)

1. 防止被其他对象误实现接口。

----

>由此，我陷入了一些思考，所以也引申出来了对 Go 接口实现或者是否这个是一个推荐的做法呢？

### 网络上的一些讨论[为什么我不喜欢Go语言式的接口（即Structural Typing）](http://blog.zhaojie.me/2013/04/why-i-dont-like-go-style-interface-or-structural-typing.html)

[Structural_type_system 是什么？摘自于 wiki](https://en.wikipedia.org/wiki/Structural_type_system)

[Duck_typing](http://en.wikipedia.org/wiki/Duck_typing)

### 一些讨论。。。

#### Wuvist 对 Go 语言接口的个人观点

Go里面类适用于哪些接口，不也是程序员显示指定的吗？

区别在于是实现者显示指定，还是调用者显示指定。

我也可以说作为调用者，你非要乱使用一个你不了解的类，那也没人可以拦着你。

但抬杠毫无意义。

值得讨论的，是接口调用者还是接口实现者责任的问题；这才是“Go语言式接口”与常见语言接口的本质不同。

#### 科技球对 Go 语言接口的个人观点

我对“自由严格”之争的基本观点：

第一，我不认为用interface去检查一切能够检查的东西，换句话说，用interface去表达一切library实现者对使用者的约定就是最好的design。我们需要分清楚什么东西是适合用程序语言表达的，什么东西是适合用自然语言表达的，甚至什么东西是适合用数学语言表达的（O(1), O(n)什么的不正是数学语言吗？）。程序语言不适合表达所有的东西，不然我们干嘛需要看编程书，算法书呢？

第二，什么东西适合用程序语言去表达？我觉得程序语言归根结底是人与机器交流的平台，而不是人与人交流的平台。我们之所以提倡静态语言，是为了给计算机表达出足够的信息让计算机能够根据这些信息进行优化，同时利用计算机在编译的时候做一些类似于拼写检查的东西避免人类常犯的错误。但程序不适合用来做人与人交流的平台，哪怕是程序员与程序员交流的平台，与其设计一种语言去表达什么O(1),O(n)，thread-safety之类的信息，何不更简单地在注释或者文档里写清楚？同理，技术上的检验也不是越多越好，拼写错误这些是人类常犯的错误，但是high-level的概念上的错误，比如送画画的小明去决斗却是程序员本身就不应该犯的错误。用interface这种程序语言去表达编程的思想，最终的结果就是编程的思想禁锢在程序语言设计者制定的牢笼里。

...

## 参考资料

1. [A Tour of the Go Programming Language with Russ Cox](http://www.youtube.com/watch?v=MzYZhh6gpI0)
2. [深入理解 interface](https://zhuanlan.zhihu.com/p/32926119)
3. [Understanding Go Interfaces - YouTube](https://www.youtube.com/watch?v=F4wUrj6pmSI)
4. [Understanding Go Interfaces - Slide](https://speakerdeck.com/campoy/understanding-the-interface)
5. [Go语言中隐式接口的冲突问题](https://my.oschina.net/chai2010/blog/416679)
6. [《Go语言高级编程》- 接口](https://github.com/chai2010/advanced-go-programming-book/blob/master/ch1-basic/ch1-04-func-method-interface.md#143-%E6%8E%A5%E5%8F%A3)
7. [理解 Go interface 的 5 个关键点](https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/)