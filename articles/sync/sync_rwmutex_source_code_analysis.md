# sync.RWMutex源码分析

针对 Golang 1.9 的 sync.RWMutex 进行分析，与 Golang 1.10 基本一样除了将`panic`改为了`throw`之外其他的都一样

RWMutex 是读写互斥锁，锁可以由任意数量的读取器或单个写入器来保持

RWMutex 的零值是一个解锁的互斥锁

RWMutex 是抢占式的读写锁，写锁之后来的读锁是加不上的
  
**以下代码均去除race竞态检测代码**

源代码位置：`sync/rwmutex.go`

## 结构体

```go
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

以下是 `sync.RWMutex` 提供的4个方法

### Lock

提供写锁加锁操作

```go  
func (rw *RWMutex) Lock() {
	// 使用 Mutex 锁
	rw.w.Lock()
	// 将当前的 readerCount 置为负数，告诉 RUnLock 当前存在写锁等待
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
	// 等待读锁释放
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_Semacquire(&rw.writerSem)
	}
}
```

### Unlock

提供写锁释放操作

```go
func (rw *RWMutex) Unlock() {
	// 加上 Lock 的时候减去的 rwmutexMaxReaders
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	// 没执行Lock调用Unlock，抛出异常
	if r >= rwmutexMaxReaders {
		race.Enable()
		throw("sync: Unlock of unlocked RWMutex")
	}
	// 通知当前等待的读锁
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false)
	}
	// 释放 Mutex 锁
	rw.w.Unlock()
}
```

### RLock

提供读锁操作

```go 
func (rw *RWMutex) RLock() {
	// 每次 goroutine 获取读锁时，readerCount+1
    // 如果写锁已经被获取，那么 readerCount 在 -rwmutexMaxReaders 与 0 之间，这时挂起获取读锁的 goroutine
    // 如果写锁没有被获取，那么 readerCount > 0，获取读锁, 不阻塞
    // 通过 readerCount 判断读锁与写锁互斥, 如果有写锁存在就挂起goroutine, 多个读锁可以并行
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
		// 将 goroutine 排到G队列的后面,挂起 goroutine
		runtime_Semacquire(&rw.readerSem)
	}
}
```

### RUnLock

RUnLock 方法对读锁进行解锁

```go
func (rw *RWMutex) RUnlock() {
	// 写锁等待状态，检查当前是否可以进行获取
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
		// r + 1 == 0表示直接执行RUnlock()
		// r + 1 == -rwmutexMaxReaders表示执行Lock()再执行RUnlock()
		// 两总情况均抛出异常
		if r+1 == 0 || r+1 == -rwmutexMaxReaders {
			race.Enable()
			throw("sync: RUnlock of unlocked RWMutex")
		}
		// 当读锁释放完毕后，通知写锁
		if atomic.AddInt32(&rw.readerWait, -1) == 0 {
			// The last reader unblocks the writer.
			runtime_Semrelease(&rw.writerSem, false)
		}
	}
}
```

### RLocker

可以看到 `RWMutex` 实现接口 `Locker`

```go  
type Locker interface {
	Lock()
	Unlock()
}
```

而方法 `RLocker` 就是将 `RWMutex` 转换为 `Locker`

```go
func (rw *RWMutex) RLocker() Locker {
	return (*rlocker)(rw)
}
```

## 总结

读写互斥锁的实现比较有技巧性一些，需要几点

1. 读锁不能阻塞读锁，引入readerCount实现

2. 读锁需要阻塞写锁，直到所有读锁都释放，引入readerSem实现

3. 写锁需要阻塞读锁，直到所有写锁都释放，引入wirterSem实现

4. 写锁需要阻塞写锁，引入Metux实现





