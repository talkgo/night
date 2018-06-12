针对Golang 1.9的sync.RWMutex进行分析，与Golang 1.10基本一样除了将`panic`改为了`throw`之外其他的都一样。  
RWMutex是读写互斥锁。锁可以由任意数量的读取器或单个写入器来保持。
RWMutex的零值是一个解锁的互斥锁。  
**以下代码均去除race竞态检测代码**

源代码位置：`sync\rwmutex.go`
## 结构体
```
type RWMutex struct {
    w           Mutex  // 互斥锁
    writerSem   uint32 // 写锁信号量
    readerSem   uint32 // 读锁信号量
    readerCount int32  // 读锁计数器
    readerWait  int32  // 获取写锁时需要等待的读锁释放数量
}
```
常量
```go  
const rwmutexMaxReaders = 1 << 30   // 支持最多2^30个读锁
```


## 方法
### Lock
提供写锁操作.
```go  
func (rw *RWMutex) Lock() {
    // 竞态检测
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	// 使用Mutex锁
	rw.w.Lock()
	// Announce to readers there is a pending writer.
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	// Wait for active readers.
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_Semacquire(&rw.writerSem)
	}
	// 竞态检测
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
		race.Acquire(unsafe.Pointer(&rw.writerSem))
	}
}
```


### RLock
提供读锁操作,
```go 
func (rw *RWMutex) RLock() {
    // 竞态检测
	if race.Enabled {
		_ = rw.w.state
		race.Disable()
	}
	// 每次goroutine获取读锁时，readerCount+1
    // 如果写锁已经被获取，那么readerCount在-rwmutexMaxReaders与0之间，这时挂起获取读锁的goroutine，
    // 如果写锁没有被获取，那么readerCount>0，获取读锁,不阻塞
    // 通过readerCount判断读锁与写锁互斥,如果有写锁存在就挂起goroutine,多个读锁可以并行
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// 将goroutine排到G队列的后面,挂起goroutine
		runtime_Semacquire(&rw.readerSem)
	}
	// 竞态检测
	if race.Enabled {
		race.Enable()
		race.Acquire(unsafe.Pointer(&rw.readerSem))
	}
}
```


### RLocker
可以看到`RWMutex`实现接口`Locker`.
```go  
type Locker interface {
	Lock()
	Unlock()
}
```
而方法`RLocker`就是将`RWMutex`转换为`Locker`.
```
func (rw *RWMutex) RLocker() Locker {
	return (*rlocker)(rw)
}
```





## 总结
总结：

读写互斥锁的实现比较有技巧性一些，需要几点

1. 读锁不能阻塞读锁，引入readerCount实现

2. 读锁需要阻塞写锁，直到所以读锁都释放，引入readerSem实现

3. 写锁需要阻塞读锁，直到所以写锁都释放，引入wirterSem实现

4. 写锁需要阻塞写锁，引入Metux实现





