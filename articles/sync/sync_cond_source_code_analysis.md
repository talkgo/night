# sync.Cond源码分析

Cond的主要作用就是获取锁之后，wait()方法会等待一个通知，来进行下一步锁释放等操作，以此控制锁合适释放，释放频率,适用于在并发环境下goroutine的等待和通知。

针对Golang 1.9的sync.Cond，与Golang 1.10一样。 源代码位置：sync\cond.go。

## 结构体
```go
type Cond struct {
	noCopy noCopy  // noCopy可以嵌入到结构中，在第一次使用后不可复制,使用go vet作为检测使用

	// 根据需求初始化不同的锁，如*Mutex 和 *RWMutex
	L Locker

	notify  notifyList  // 通知列表,调用Wait()方法的goroutine会被放入list中,每次唤醒,从这里取出
	checker copyChecker // 复制检查,检查cond实例是否被复制
}
```
再来看看等待队列`notifyList`结构体： 
```go  
type notifyList struct {
	wait   uint32
	notify uint32
	lock   uintptr
	head   unsafe.Pointer
	tail   unsafe.Pointer
}

```

## 函数
### NewCond
相当于`Cond`的构造函数，用于初始化`Cond`。

参数为Locker实例初始化,传参数的时候必须是引用或指针,比如&sync.Mutex{}或new(sync.Mutex)，不然会报异常:`cannot use lock (type sync.Mutex) as type sync.Locker in argument to sync.NewCond`。  

大家可以想想为什么一定要是指针呢？ 因为如果传入 Locker 实例，在调用 `c.L.Lock()` 和 `c.L.Unlock()` 的时候，会频繁发生锁的复制，会导致锁的失效，甚至导致死锁。

```go  
func NewCond(l Locker) *Cond {
	return &Cond{L: l}
}
```

### Wait
等待自动解锁c.L和暂停执行调用goroutine。恢复执行后,等待锁c.L返回之前。与其他系统不同，等待不能返回，除非通过广播或信号唤醒。


因为c。当等待第一次恢复时，L并没有被锁定，调用者通常不能假定等待返回时的条件是正确的。相反，调用者应该在循环中等待:

```go  
func (c *Cond) Wait() {
    // 检查c是否是被复制的，如果是就panic
	c.checker.check()
	// 将当前goroutine加入等待队列
	t := runtime_notifyListAdd(&c.notify)
	// 解锁
	c.L.Unlock()
	// 等待队列中的所有的goroutine执行等待唤醒操作
	runtime_notifyListWait(&c.notify, t)
	c.L.Lock()
}
```
判断cond是否被复制。
```go  
type copyChecker uintptr

func (c *copyChecker) check() {
	if uintptr(*c) != uintptr(unsafe.Pointer(c)) &&
		!atomic.CompareAndSwapUintptr((*uintptr)(c), 0, uintptr(unsafe.Pointer(c))) &&
		uintptr(*c) != uintptr(unsafe.Pointer(c)) {
		panic("sync.Cond is copied")
	}
}

```

### Signal
唤醒等待队列中的一个goroutine，一般都是任意唤醒队列中的一个goroutine，为什么没有选择FIFO的模式呢？这是因为FiFO模式效率不高，虽然支持，但是很少使用到。
```go  
func (c *Cond) Signal() {
    // 检查c是否是被复制的，如果是就panic
	c.checker.check()
	// 通知等待列表中的一个 
	runtime_notifyListNotifyOne(&c.notify)
}
```

### Broadcast
唤醒等待队列中的所有goroutine。
```go  
func (c *Cond) Broadcast() {
    // 检查c是否是被复制的，如果是就panic
	c.checker.check()
	// 检查c是否是被复制的，如果是就panic
	runtime_notifyListNotifyAll(&c.notify)
}
```

### 实例
```go  
package main

import (
	"fmt"
	"sync"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)

func main() {
	for i := 0; i < 40; i++ {
		go func(x int) {
			cond.L.Lock()         //获取锁
			defer cond.L.Unlock() //释放锁
			cond.Wait()           //等待通知,阻塞当前goroutine
			fmt.Println(x)
			time.Sleep(time.Second * 1)

		}(i)
	}
	time.Sleep(time.Second * 1)
	fmt.Println("Signal...")
	cond.Signal() // 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	cond.Signal() // 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 3)
	cond.Broadcast() //3秒之后 下发广播给所有等待的goroutine
	fmt.Println("Broadcast...")
	time.Sleep(time.Second * 60)
}


```

