---
title: "2019-03-08 微信讨论"
date: 2019-03-08T21:00:00+08:00
---

## 读写锁引入

有下面一段程序，面试官问这段程序有什么问题？

```
type Store struct {
    a string
    b string

    sync.RWMutex
}

func (s *Store) GetA() string {
    fmt.Println("get a")
    s.RLock()
    fmt.Println("get a2")
    defer s.RUnlock()

    return s.a
}

func (s *Store) GetAB() (string, string) {
    fmt.Println("get ab")
    s.RLock()
    fmt.Println("get ab2")
    defer s.RUnlock()

    return s.GetA(), s.b
}

func (s *Store) Write(a, b string) {
    fmt.Println("write")
    s.Lock()
    defer s.Unlock()
    fmt.Println("write2")

    s.a = a
    s.b = b
}
```

1. 看到这段程序程序，首先想到的是读写锁的问题；
2. 其次，看 Store 这个结构体，各个函数都定义的是指针函数。那就说明：不存在读写锁的 copy 过程；
3. GetAB 方法中通过调用 GetA 方法，在 `s.RUnlock` 前通过调用 `s.GetA`，又做了一次读写锁上锁 `s.RLock`，但是读锁可以多次上锁，所以单看这里没什么问题；
4. 然后，想到会不会 `Write` 和 `GetAB` 并发调用的时候会存在问题呢？思考了一会，觉得没问题，就放弃了。

以上，是面试时整个思路。

回头，越想越觉得这里哪里有问题，就在`夜读群`里求教了一下，群里大神发了一篇[读写锁优先级的文章](https://blog.csdn.net/xyz347/article/details/83902123)，然后给了一段测试样例，瞬间豁然开朗。

main函数逻辑如下：

```
func main() {
    store := Store{}
    wg := sync.WaitGroup{}
    wg.Add(2)
    go func() {
        defer wg.Done()
        for i := 1; i < 10000; i += 1 {
            fmt.Println("main write ", i)
            store.Write("111", "1111")
        }
    }()
    go func() {
        defer wg.Done()
        for i := 1; i < 10000; i += 1 {
            fmt.Println("main get ab", i)
            store.GetAB()
        }
    }()
    wg.Wait()
}
```

执行结果

```
main get ab 12  //main函数读取ab
get ab          //进入 s.GetAB 函数
main write  1   //main 函数写数据
write           //进入 s.Write 函数
write2          //获取写锁
main write  2
....            //写锁一直抢占
....
main write  13  //main 函数写数据
write           //进入 s.Write 函数
write2          //获取写锁
main write  14  
write
write2
get ab2         //之前 get ab 12 才获得读锁
get a           //进入 GetA 
get a2          //获取读锁
main get ab 13  //main函数 get ab
get ab          //进入 s.GetAB 函数
get ab2         //获取读锁
main write  15  //注意⚠️ 这个时候写数据开始了
write           //进入 Write 函数，后面尝试获取写锁
get a           //这个时候 GetAB 进入了 GetA，尝试获取读锁
fatal error: all goroutines are asleep - deadlock!  //出现了死锁
```

分析：

```

    GetAB       |            GetA               |                Write
                |                               |
r0 占用读锁      |                               |
                |                               |    w0 尝试获取写锁 等待r0释放读锁
                |   r1 尝试获取读锁，排在w0后面     |

```

由于读写锁的优先级，读锁和写锁同时竞争时，读锁要排在写锁后面，导致了 r1 竞争 w0的锁，w0竞争r0，r0执行不下去，最后死锁。

## 读写锁底层

读写锁前置条件：

1. 读写互斥，但是读读不互斥；
2. 读、写锁都不会出现饥饿；
3. 保证读上锁数量与解锁数量一致；

可以思考下，如果让你设计一个这样的锁，你会怎么设计？

----

go中读写锁的结构，如下：

```
type RWMutex struct {
    w           Mutex  // 用来保证同一时间只有一个写锁能够抢到锁
    writerSem   uint32 // 写锁信号量，在读锁全部解锁时通知阻塞的写锁
    readerSem   uint32 // 读锁信号量，在写锁解锁时通知阻塞的读操作
    readerCount int32  // 等待、已上锁的读锁数量
    readerWait  int32  // 写锁获得锁前，已经上锁的读锁数量
}
```

### 读锁逻辑

首先，看一下读上锁逻辑：

```
func (rw *RWMutex) RLock() {
    ...
    if atomic.AddInt32(&rw.readerCount, 1) < 0 {
        // A writer is pending, wait for it.
        runtime_SemacquireMutex(&rw.readerSem, false)
    }
    ...
}
```

上面，上读锁逻辑获试图获取读锁数量原子性加一： `atomic.AddInt32(&rw.readerCount, 1)`。自增操作返回值如果小于0，则阻塞等待信号量 `readerSem` 唤醒。

疑问：

1. 什么情况下 `readerCount` 小于0；
2. `runtime_SemacquireMutex` 不会造成读读互斥么？
3. 如何保证读、写互斥？

再来看一下，读解锁逻辑：

```
func (rw *RWMutex) RUnlock() {
    ...
    if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 { 
        if r+1 == 0 || r+1 == -rwmutexMaxReaders {
            race.Enable()
            throw("sync: RUnlock of unlocked RWMutex")
        }

        // A writer is pending.
        if atomic.AddInt32(&rw.readerWait, -1) == 0 {
            // The last reader unblocks the writer.
            runtime_Semrelease(&rw.writerSem, false)
        }
    }
    ...
}
```

解锁逻辑：先对 `atomic.AddInt32(&rw.readerCount, -1)` 进行原子性减一操作。

* r大于 0 ：直接释放锁完成；
* r小于 0 ：进行读锁数量一致性判断，`atomic.AddInt32(&rw.readerWait, -1)` 针对 `readerWait` 原子性减一后判断是否为 0，为 0 则唤起写锁信号量；

与读加锁类似，同样有 `atomic.AddInt32(&rw.readerCount, -1)` 小于 0 判断。可以有结论 `rw.readerCount` 小于 0，为写锁上锁的充要条件，后面分析写锁时进行验证。

解决了的问题：

1. 释放读锁，读锁全部释放后唤起写锁；
2. 上锁与解锁数量一致性保证；

疑问：

1. `readerCount` 修改成一个负数？如何保证这个负数足够小呢？

### 写锁逻辑

先上代码：

```
func (rw *RWMutex) Lock() {
    ...
    // First, resolve competition with other writers.
    rw.w.Lock()
    // Announce to readers there is a pending writer.
    r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
    // Wait for active readers.
    if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
        runtime_SemacquireMutex(&rw.writerSem, false)
    }
    ...
}
```

写上锁逻辑：

1. 首先，互斥量上锁，保证只有一个写锁加锁成功。
2. 然后，令 `readerCount` 原子性减去 `rwmutexMaxReaders`（这是个常量，具体定义 `const rwmutexMaxReaders = 1 << 30`）。这里可以验证之前猜想，`rw.readerCount` 小于0，是持有锁的充要条件。
    * `atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders` 返回结果是在写锁获取前，已持有读锁的数量 r。
        - r=0，说明没有读锁；
        - r<0，只有在`读解锁数量>读加锁数量`，或写锁多次时发生；第一个情况，读解锁会 `check`；第二种情况，`mutex` 保证同时只有一个写锁；
        - r>0，存在读锁；
3. 再进行 r!=0 判断（即存在读锁）。原子性操作 `atomic.AddInt32(&rw.readerWait, r)`，记录需要等待的读锁数量，然后等待`writerSem`唤醒。

最终，保证：1. 写锁唯一性；2. 等待读锁完全释放；3. 阻塞后面读锁的获取；

再来看一下，写锁解锁逻辑：

```
func (rw *RWMutex) Unlock() {
    ...
    // Announce to readers there is no active writer.
    r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
    if r >= rwmutexMaxReaders {
        race.Enable()
        throw("sync: Unlock of unlocked RWMutex")
    }
    // Unblock blocked readers, if any.
    for i := 0; i < int(r); i++ {
        runtime_Semrelease(&rw.readerSem, false)
    }
    // Allow other writers to proceed.
    rw.w.Unlock()
    ...
}
```

解锁逻辑：

1. 原子性操作 `atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)`。这里，能够看到两个隐含的点：
    * 原子操作结束后，如果有其他读锁试图获取读锁，不需要阻塞；
    * 这个时候其他线程还是不能够获取写锁；
    * 即：`写锁释放锁时，读锁要比写锁优先级高`；
2. 原子操作返回值，是当前读锁数量。包括在写锁前读锁（写锁未完全获得情况下写锁解锁），和写锁后阻塞读锁；然后 `runtime_Semrelease` 唤起阻塞着的读锁。
    * `runtime_Semrelease > runtime_SemacquireMutex` 会不会存在问题？验证过不会。
3. 然后写锁释放；

## 总结

通过分析，可以得出结论：

1. 写锁释放过程中，读锁优先级要高于写锁；
2. 读锁加锁后，写锁可以进入加锁过程，但是要等待之前读锁释放；即，并不少写锁优先级高于写锁，而是在`读锁已经上锁，或没有持有读写锁的协程`条件下，读写锁都有机会获取锁；

所以，针对之前的面试题，读锁嵌套读锁，在有写锁的时候，依据结论2会发生死锁。

通过上面分析，存在待验证问题：

1. `一个协程个已获取读锁，另个协程试图获取写锁，还有一个协程在完全获取写锁前调用Unlock，再一个协程释放读锁，按顺序进行流程`。会发生死锁具体可以自己分析（写锁信号量永远阻塞）；

2. `一个协程已上写锁锁，一个协程试图获取读锁，然后另一个协程释放读锁，最后一个协程释放写锁`，同样会发生死锁（读信号量永远阻塞）；

在以后用锁的时候不管有没有优先级，都要时刻记住死锁的四个必要条件：

1. 互斥条件：一个资源每次只能被一个进程使用。 
2. 锁的不可抢占：进程已获得的资源，在末使用完之前，不能强行剥夺。
3. 占有且等待：一个进程因请求资源而阻塞时，对已获得的资源保持不放。 
4. 循环等待条件: 若干进程之间形成一种头尾相接的循环等待资源关系。


## 参考

* [sync.RWMutex]https://medium.com/golangspec/sync-rwmutex-ca6c6c3208a0

