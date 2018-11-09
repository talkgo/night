# sync.WaitGroup源码分析

针对Golang 1.9的sync.WaitGroup进行分析，与Golang 1.10基本一样除了将`panic`改为了`throw`之外其他的都一样。
源代码位置：`sync\waitgroup.go`。  
## 结构体
```
type WaitGroup struct {
	noCopy noCopy  // noCopy可以嵌入到结构中，在第一次使用后不可复制,使用go vet作为检测使用
	// 位值:高32位是计数器，低32位是goroutine等待计数。
	// 64位的原子操作需要64位的对齐，但是32位。编译器不能确保它,所以分配了12个byte对齐的8个byte作为状态。
	state1 [12]byte // byte=uint8范围：0~255，只取前8个元素。转为2进制：0000 0000，0000 0000... ...0000 0000
	sema   uint32   // 信号量，用于唤醒goroutine
}
```
不知道大家是否和我一样，不论是使用Java的CountDownLatch还是Golang的WaitGroup，都会疑问，可以装下多个线程|协程等待呢？看了源码后可以回答了，可以装下
```
1111 1111 1111 ... 1111
\________32___________/
```
2^32个辣么多！所以不需要担心单机情况下会被撑爆了。
## 函数
以下代码已经去掉了与核心代码无关的race代码。
### Add
添加或者减少等待goroutine的数量。

参数delta可能是负的，加到WaitGroup计数器,可能出现如下结果
- 如果计数器变为零，所有被阻塞的goroutines都会被释放。
- 如果计数器变成负数，就增加恐慌。

```
func (wg *WaitGroup) Add(delta int) {
    // 获取到wg.state1数组中元素组成的二进制对应的十进制的值
	statep := wg.state()
	// 高32位是计数器
	state := atomic.AddUint64(statep, uint64(delta)<<32)
	// 获取计数器
	v := int32(state >> 32)
	w := uint32(state)
	// 计数器为负数，报panic
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
	// 添加与等待并发调用，报panic
	if w != 0 && delta > 0 && v == int32(delta) {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// 计数器添加成功
	if v > 0 || w == 0 {
		return
	}

	// 当等待计数器> 0时，而goroutine设置为0。
	// 此时不可能有同时发生的状态突变:
	// - 增加不能与等待同时发生，
	// - 如果计数器counter == 0，不再增加等待计数器
	if *statep != state {
		panic("sync: WaitGroup misuse: Add called concurrently with Wait")
	}
	// Reset waiters count to 0.
	*statep = 0
	for ; w != 0; w-- {
		// 目的是作为一个简单的wakeup原语，以供同步使用。true为唤醒排在等待队列的第一个goroutine
		runtime_Semrelease(&wg.sema, false)
	}
}

```

```
// unsafe.Pointer其实就是类似C的void *，在golang中是用于各种指针相互转换的桥梁。
// uintptr是golang的内置类型，是能存储指针的整型，uintptr的底层类型是int，它和unsafe.Pointer可相互转换。
// uintptr和unsafe.Pointer的区别就是：unsafe.Pointer只是单纯的通用指针类型，用于转换不同类型指针，它不可以参与指针运算；
// 而uintptr是用于指针运算的，GC 不把 uintptr 当指针，也就是说 uintptr 无法持有对象，uintptr类型的目标会被回收。
// state()函数可以获取到wg.state1数组中元素组成的二进制对应的十进制的值
func (wg *WaitGroup) state() *uint64 {
	if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
		return (*uint64)(unsafe.Pointer(&wg.state1))
	} else {
		return (*uint64)(unsafe.Pointer(&wg.state1[4]))
	}
}
```
### Done
相当于Add(-1)。  
```
func (wg *WaitGroup) Done() {
    // 计数器减一
	wg.Add(-1)
}
```


### Wait
执行阻塞，直到所有的WaitGroup数量变成0。  
```
func (wg *WaitGroup) Wait() {
	// 获取到wg.state1数组中元素组成的二进制对应的十进制的值
	statep := wg.state()
	// cas算法
	for {
		state := atomic.LoadUint64(statep)
		// 高32位是计数器
		v := int32(state >> 32)
		w := uint32(state)
		// 计数器为0，结束等待
		if v == 0 {
			// Counter is 0, no need to wait.
			return
		}
		// 增加等待goroutine计数，对低32位加1，不需要移位
		if atomic.CompareAndSwapUint64(statep, state, state+1) {
			// 目的是作为一个简单的sleep原语，以供同步使用
			runtime_Semacquire(&wg.sema)
			if *statep != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			return
		}
	}
}
```

## 使用注意事项
1. WaitGroup不能保证多个 goroutine 执行次序
2. WaitGroup无法指定固定的goroutine数目




