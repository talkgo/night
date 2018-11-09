# sync.Once源码分析

sync.Once可以实现单例模式，确保sync.Once.Do(f func())只会被执行一次，可以初始化某个实例单例。

针对Golang 1.9的sync.Once，与Golang 1.10一样。 源代码位置：sync\once.go。

## 结构体
Once结构体定义如下：
```go  
type Once struct {
	m    Mutex
	done uint32   // 初始值为0表示还未执行过，1表示已经执行过
}
```

## Do
```go  
func (o *Once) Do(f func()) {
    // done==1表示已经执行过了，直接结束返回
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// 锁住对象，避免并发问题
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
	    // 执行f函数后将done设置为1
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

需要注意的是执行f函数是同步进行的，也就是说可能存在阻塞问题。

