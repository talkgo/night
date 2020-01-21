# goroutine和通道
<!-- TOC -->

- [goroutine和通道](#goroutine%e5%92%8c%e9%80%9a%e9%81%93)
	- [goroutine和通道](#goroutine%e5%92%8c%e9%80%9a%e9%81%93-1)
		- [CSP并发模型](#csp%e5%b9%b6%e5%8f%91%e6%a8%a1%e5%9e%8b)
		- [goroutine的调度模型](#goroutine%e7%9a%84%e8%b0%83%e5%ba%a6%e6%a8%a1%e5%9e%8b)
		- [Go协程和主线程](#go%e5%8d%8f%e7%a8%8b%e5%92%8c%e4%b8%bb%e7%ba%bf%e7%a8%8b)
		- [goroutine入门](#goroutine%e5%85%a5%e9%97%a8)
			- [设置Golang运行的CPU数](#%e8%ae%be%e7%bd%aegolang%e8%bf%90%e8%a1%8c%e7%9a%84cpu%e6%95%b0)
		- [管道](#%e7%ae%a1%e9%81%93)
			- [加锁](#%e5%8a%a0%e9%94%81)
			- [引入管道](#%e5%bc%95%e5%85%a5%e7%ae%a1%e9%81%93)
			- [channel介绍](#channel%e4%bb%8b%e7%bb%8d)
		- [协程和管道](#%e5%8d%8f%e7%a8%8b%e5%92%8c%e7%ae%a1%e9%81%93)
		- [代码效率](#%e4%bb%a3%e7%a0%81%e6%95%88%e7%8e%87)
	- [golang管道细节总结](#golang%e7%ae%a1%e9%81%93%e7%bb%86%e8%8a%82%e6%80%bb%e7%bb%93)
		- [细节1](#%e7%bb%86%e8%8a%821)
		- [细节2](#%e7%bb%86%e8%8a%822)
		- [细节3](#%e7%bb%86%e8%8a%823)

<!-- /TOC -->

## goroutine和通道

### CSP并发模型

用于描述两个独立的并发实体通过共享的通讯 channel(管道)进行通信的并发模型

Golang 就是借用CSP模型的一些概念为之实现并发进行理论支持

process是在go语言上的表现就是 goroutine 是实际并发执行的实体，每个实体之间是通过channel通讯来实现数据共享。


### goroutine的调度模型

MPG模型

M: 操作系统的主线程

P: 协程执行所需要的上下文

G:协程


[具体了解可以点击这里](https://www.jianshu.com/p/36e246c6153d)

### Go协程和主线程

1、Go主线程（线程或者叫进程）：一个Go主线程，可以起多个协程，协程就是轻量级的线程

2、Go协程的特点

>有独立的栈空间

>共享程序堆空间

>调度由用户控制

>协程是轻量级的线程【编译器做优化】

**问题：**

**`什么是栈空间&堆空间？`**

> 栈空间？
> **编译器自动分配释放**，存放函数的参数值，局部变量的值等，其操作方式类似于数据结构的栈。
>
>堆空间？
>**一般是由程序员分配释放**，若程序员不释放的话，程序结束时可能由OS回收，值得注意的是他与数据结构的堆是两回事，分配方式倒是类似于数据结构的链表

**`怎么理解这段话？`**

注意我们此处谈到的堆和栈是对操作系统中的，这个和数据结构中的堆和栈还是又一定区别的。

栈: 可以简单得理解成**一次函数调用内部申请到的内存，它们会随着函数的返回把内存还给系统**。

```go
func F() {
    temp := make([]int, 0, 20)
    ...
}
```
类似于上面代码里面的temp变量，只是内函数内部申请的临时变量，并不会作为返回值返回，它就是被编译器申请到栈里面。

申请到 栈内存 好处：**函数返回直接释放，不会引起垃圾回收，对性能没有影响。**


再来看看堆得情况之一如下代码：
```go
func F() []int{
    a := make([]int, 0, 20)
    return a
}
```
而上面这段代码，申请的代码一模一样，但是申请后作为返回值返回了，编译器会认为变量之后还会被使用，当函数返回之后并不会将其内存归还，那么它就会被申请到 堆 上面了。

申请到堆上面的内存才会引起垃圾回收，如果这个过程（特指垃圾回收不断被触发）过于高频就会导致 gc 压力过大，程序性能出问题。

参考文献：

[Golang内存分配逃逸分析](https://driverzhang.github.io/post/golang%E5%86%85%E5%AD%98%E5%88%86%E9%85%8D%E9%80%83%E9%80%B8%E5%88%86%E6%9E%90/)

[Go的变量到底在堆还是栈中分配](http://www.zenlife.tk/go-allocated-on-heap-or-stack.md)

后面我会单独出一章介绍Golang 堆空间&栈空间理解


### goroutine入门

例子：
```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test () hello world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	// test()
	go test() //开启一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println("main () hello golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}

```

运行结果：
```
main () hello golang1
test () hello world1
test () hello world2
main () hello golang2
test () hello world3
main () hello golang3
main () hello golang4
test () hello world4
test () hello world5
main () hello golang5
main () hello golang6
test () hello world6
test () hello world7
main () hello golang7
main () hello golang8
test () hello world8
main () hello golang9
test () hello world9
main () hello golang10
test () hello world10
```

运行结果： **说明main这个主线程和test协程同时运行**


可以画个逻辑图来说明这个情况：

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/逻辑图.png)



逻辑图讲解：

1、主线程是一个物理线程、直接作用在CPU上、是重量级的，非常耗费CPU资源

2、协程是主线程开启的，是轻量级的线程，是逻辑态，对资源消耗相对小

3、Golang的协程机制是重要的特点，可以轻松开启上万个协程

其他编程语言的开发机制一般基于线程，开启过多的线程，资源耗费大

这里就凸显了golang在并发上的优势了



#### 设置Golang运行的CPU数

注意:
>1、Go1.8之前 要进行设置下 可以更高效的利用CPU

>2、GO1.8之后 默认让程序运行在多个核上 可以不用设置

这里使用的是**go version go1.13.1** 

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpunum:", cpuNum)

	//可以自己设置使用多个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}

```



### 管道

看一个例子来解释为什么要用到管道这个技术？

```go
package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i

	}

	myMap[n] = res

}

func main() {

	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 10)

	//遍历结果
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}

}

```

运行结果：
```
map[76]=0
map[81]=0
map[104]=0
map[117]=0
map[118]=0
map[124]=0
map[139]=0
map[153]=0
map[162]=0
map[2]=2
map[16]=20922789888000
....

```


发现的问题：

**`多个协程 同时写 会出现资源竞争`**

<br>
解决思路：

####  加锁
全局变量加锁同步

没有对全局变量加锁，会出现资源竞争问题，代码会报错： concurrent map writes

加入互斥锁

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	myMap = make(map[int]int, 10)

	//声明全局互斥锁
	//lock 是一个全局互斥锁
	//sync 表示同步
	//Mutex 表示互斥
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i

	}

	//加锁
	lock.Lock()
	myMap[n] = res

	//解锁
	lock.Unlock()

}

func main() {

	for i := 1; i <= 200; i++ {
		go test(i)
	}

	//休眠几秒合适？
	time.Sleep(time.Second * 10)

	//遍历结果
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()

}

```

遍历结果也要加入锁机制， 原因：

程序从设计上可以指定10秒执行了所有协程，但是主线程并不知道，因此底层可能仍然出现资源争夺


#### 引入管道
前面使用全局变量加锁解决 但不完美：

主要有三个地方：

>1）主线程在等待所有gorouting全部完成的时间很难确定，这里设置了10秒，仅仅是估算

>2）如果主线程休眠时间长了，会加长等待时间
    如果等待时间短了，可能还有goroutine处于工作状态，
    这时会随着主线程的退出而销毁

>3）通过全局变量加锁，也并不利用协程对全局变量的读写操作（不知道在哪里加锁、释放锁）


#### channel介绍

1.主要有下面几个特点：

>1.Channel本质就是一个数据结构 -队列

>2.数据是先进先出

>3.线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全

>4.channel是有类型的，一个string的channel只能存放string类型数据


2.基本使用：
```
定义 /声明 channel
var 变量 chan  数据类型
var intChan  chan int

说明：
1）channel是引用类型
2）channel必须初始化才能写入数据、即make后才能使用

```

3.例子

```go
package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan的值=%v\n", intChan) //intChan的值=0xc00001a100

}
```


4.管道写入

例子1：
```go
package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println()

	//管道写入
	intChan <- 10
	num := 211
	intChan <- num

	//管道长度和容量
	fmt.Printf("channel len=%v cap=%v", len(intChan), cap(intChan))

}

```

例子2：

```go
package main

import (
    "fmt"
)

func main() {
    var intChan chan int
    intChan = make(chan int, 3)

    fmt.Println()

    //管道写入
    intChan <- 10
    num := 211
    intChan <- num

    //当写入数据不能超过容量,超过报错
    intChan <- 50
    intChan <- 98

    //管道长度和容量
    fmt.Printf("channel len=%v cap=%v", len(intChan), cap(intChan))

}
```

例子3：

```go
package main

import (
    "fmt"
)

func main() {
    var intChan chan int
    intChan = make(chan int, 3)

    fmt.Println()

    //管道写入
    intChan <- 10
    num := 211
    intChan <- num

    //当写入数据不能超过容量
    intChan <- 50

    //管道长度和容量
    fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

    //读数据
    var num2 int
    num2 = <-intChan
    fmt.Println("num2=", num2)
    fmt.Printf("channel len=%v cap=%v\n", len(intChan), cap(intChan))

    //在没有使用协程的情况下，管道数据已经全部取出,再取就会报错deadlock
    num3 := <-intChan
    num4 := <-intChan
    num5 := <-intChan

    fmt.Println("num3=", num3, "num4=", num4, "num5=", num5)

}
```

5.管道细节总结：

>1.channel只能存放指定的数据类型

>2.channel的数据放满后，就不能再放入了

>3.如果从channel取出数据后，可以继续放入

>4.在没有使用协程的情况下，如果channel数据取完了再取， 就会报deadlock



6.channel的关闭

使用内置函数close可以关闭channel，当channel关闭后 就不能再向channel写数据

但是可以从channel读取数据

```go
package main

func main() {
	intChan := make(chan int, 3)

	intChan <- 100
	intChan <- 200
	close(intChan)
	intChan <- 300  //panic: send on closed channel
}

```


7.channel的遍历

支持for-range的方式来遍历：

1.在遍历时，如果channel没有关闭，则出现deadlock

2.在遍历时，如果channel已经关闭，会正常遍历数据，遍历完后会退出遍历


```go
package main

import "fmt"

func main() {
	intChan := make(chan int, 100)

	for i := 0; i < 100; i++ {
		intChan <- i * 2

	}

	//遍历,不能使用普通的for循环,取出来的不是值
	// for i := 0; i < len(intChan); i++ {
	//  fmt.Println("i=", i)
	// }

	//使用for-range循环,取出来的是值
	close(intChan)
	for v := range intChan {
		fmt.Println("v=", v)
	}

}

```

### 协程和管道


![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/协程与管道.png)

看一个例子：
```go
package main

import (
	"fmt"
)

//write data
func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
		fmt.Printf("writeData写数据=%v\n", i)
		// time.Sleep(time.Second)
	}
	close(intChan)
}

//read data
func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		// time.Sleep(time.Second)
		fmt.Printf("readData 读到数据=%v\n", v)
	}

	//任务完成
	exitChan <- true
	close(exitChan)
}

func main() {

	//创建两个管道
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go readData(intChan, exitChan)
	go writeData(intChan)

	// time.Sleep(time.Second * 10)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}

}

```


再看一个例子：

![Alt text](https://raw.githubusercontent.com/Syncma/Figurebed/master/img/思路分析.png)

```go
package main

import (
    "fmt"
)

func putNum(intChan chan int) {
    for i := 1; i <= 80; i++ {
        intChan <- i
    }

    //关闭intChan
    close(intChan)

}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

    var flag bool
    for {
        num, ok := <-intChan
        //intChan取不到
        if !ok {
            break
        }
        flag = true
        //判断是不是素数
        for i := 2; i < num; i++ {
            //说明num不是素数
            if num%i == 0 {
                flag = false
                break

            }
        }
        if flag {
            //放入primeChan
            primeChan <- num
        }
    }

    fmt.Println("有一个primeNum 协程因为取不到数据退出")
    //还不能关闭primeChan
    //向exitChan写入true
    exitChan <- true
}

func main() {

    intChan := make(chan int, 1000)
    primeChan := make(chan int, 2000) //放入结果
    exitChan := make(chan bool, 4)    //退出管道

    //开启一个协程,向intChan写入1-8000
    go putNum(intChan)

    //开启4个协程,从intChan取出数据,并判断是否为素数
    //如果是,就放入到primeChan
    for i := 0; i < 4; i++ {
        go primeNum(intChan, primeChan, exitChan)
    }

    //主线程处理
    go func() {
        for i := 0; i < 4; i++ {
            <-exitChan
        }

        //关闭primeChan
        close(primeChan)

    }()

    //遍历primeChan
    for {
        res, ok := <-primeChan
        if !ok {
            break
        }

        //结果输出
        fmt.Printf("素数=%d\n", res)

    }

    fmt.Println("main主线程退出")

}

```

运行结果：
```
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
素数=1
素数=2
素数=3
素数=5
素数=7
素数=11
素数=13
素数=17
素数=19
素数=23
素数=29
素数=31
素数=37
素数=41
素数=43
素数=47
素数=53
素数=59
素数=61
素数=67
素数=71
素数=73
素数=79
main主线程退出
```



这里有个问题，就是结果显示不对：

代码里面增加休眠时间

修改后：
```go
package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)

}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		time.Sleep(time.Millisecond)
		num, ok := <-intChan
		//intChan取不到
		if !ok {
			break
		}
		flag = true
		//判断是不是素数
		for i := 2; i < num; i++ {
			//说明num不是素数
			if num%i == 0 {
				flag = false
				break

			}
		}
		if flag {
			//放入primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据退出")
	//还不能关闭primeChan
	//向exitChan写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) //放入结果
	exitChan := make(chan bool, 4)    //退出管道

	//开启一个协程,向intChan写入1-8000
	go putNum(intChan)

	//开启4个协程,从intChan取出数据,并判断是否为素数
	//如果是,就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		//关闭primeChan
		close(primeChan)

	}()

	//遍历primeChan
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}

		//结果输出
		fmt.Printf("素数=%d\n", res)

	}

	fmt.Println("main主线程退出")

}

```


运行结果：
```
素数=1
素数=2
素数=3
素数=5
素数=7
素数=11
素数=13
素数=17
素数=19
素数=23
素数=29
素数=31
素数=37
素数=41
素数=43
素数=47
素数=53
素数=59
素数=61
素数=67
素数=71
素数=73
素数=79
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
有一个primeNum 协程因为取不到数据退出
main主线程退出
```

### 代码效率

1.普通方法

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now().Unix()
	for num := 1; num <= 80000; num++ {

		flag := true
		//判断是不是素数
		for i := 2; i < num; i++ {
			//说明num不是素数
			if num%i == 0 {
				flag = false
				break

			}
		}
		if flag {

		}
	}
	end := time.Now().Unix()
	fmt.Println("普通方法耗时=", end-start) //普通方法耗时= 3

}
```


2.使用了协程+管道
```go
package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 80000; i++ {
		intChan <- i
	}

	//关闭intChan
	close(intChan)

}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	var flag bool
	for {
		// time.Sleep(time.Millisecond)
		num, ok := <-intChan
		//intChan取不到
		if !ok {
			break
		}
		flag = true
		//判断是不是素数
		for i := 2; i < num; i++ {
			//说明num不是素数
			if num%i == 0 {
				flag = false
				break

			}
		}
		if flag {
			//放入primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据退出")
	//还不能关闭primeChan
	//向exitChan写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 20000) //放入结果
	exitChan := make(chan bool, 4)     //退出管道

	start := time.Now().Unix()
	//开启一个协程,向intChan写入1-8000
	go putNum(intChan)

	//开启4个协程,从intChan取出数据,并判断是否为素数
	//如果是,就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//主线程处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}

		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start) //使用协程耗时= 1

		//关闭primeChan
		close(primeChan)

	}()

	//遍历primeChan
	for {
		_, ok := <-primeChan
		// res, ok := <-primeChan
		if !ok {
			break
		}

		//结果输出
		// fmt.Printf("素数=%d\n", res)

	}

	fmt.Println("main主线程退出")

}

```


3.优化版

在运行某个程序时，如何指定是否存在资源竞争问题？

**方法很简单，`在编译程序时，增加一个参数 -race`**


## golang管道细节总结


### 细节1
```go
package main

import (
    "fmt"
)

func main() {
    //管道可以声明只读或者只写

    //1.在默认情况下,管道是双向
    //var chan1 chan int //可读可写

    //2 声明为只写
    var chan2 chan<- int
    chan2 = make(chan int, 3)
    chan2 <- 20
    // num := <-chan2 //error

    //3 声明为只读
    var chan3 <-chan int
    chan3 = make(chan int, 3)
    // chan3 <- 20//error
    num := <-chan3

    fmt.Println("chan2=", chan2)

}
```


### 细节2
```go
package main

import (
    "fmt"
)

func main() {
    //使用select 可以解决从管道取数据的阻塞问题

    //1.定义一个管道 10个数据int
    intChan := make(chan int, 10)
    for i := 0; i < 10; i++ {
        intChan <- i
    }

    //2.定义一个管道 5个数据string
    StringChan := make(chan string, 5)
    for i := 0; i < 5; i++ {
        StringChan <- "hello" + fmt.Sprintf("%d", i)
    }

    //传统方法在遍历管道时候 如果不关闭会阻塞会导致deadlock
    //问题在实际开发中可能我们不好确定什么时候关闭管道
    //可以使用select 方法解决
    // label:
    for {
        select {
        //注意:这里如果intChan一直没有关闭不会一直阻塞而deadlock
        //会自动到下一个case匹配
        case v := <-intChan:
            fmt.Printf("从intChan读取数据%d\n", v)
        case v := <-StringChan:
            fmt.Printf("从StringChan读取数据%s\n", v)
        default:
            fmt.Printf("都取不到\n")
            // break label //跟label配合使用
            return
        }

    }

}
```


### 细节3
```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 10; i++ {
        time.Sleep(time.Second)
        fmt.Println("hello world")
    }
}

func test() {

    //使用defer+recover
    defer func() {
        //捕获抛出的panic
        if err := recover(); err != nil {
            fmt.Println("test()发生错误", err)
        }

    }()
    var myMap map[int]string
    myMap[0] = "golang"
}

func main() {
    go sayHello()
    go test()

    for i := 0; i < 10; i++ {
        fmt.Println("main() ok=", i)
    }

}

```