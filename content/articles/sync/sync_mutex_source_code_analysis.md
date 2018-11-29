---
title: sync.Mutex 源码分析
---
针对 Golang 1.10.3 的 sync.Mutex 进行分析，代码位置：`sync/mutex.go`

## 结构体

```go
type Mutex struct {
	state int32  // 指代mutex锁当前的状态
	sema  uint32 // 信号量，用于唤醒goroutine
}
```

Mutex 中的 state 用于指代锁当前的状态，如下所示

```go
1111 1111 ...... 1111 1111
\_________29__________/|||
 存储等待 goroutine 数量 ||表示当前 mutex 是否加锁
                       |表示当前 mutex 是否被唤醒
                       表示 mutex 当前是否处于饥饿状态
                   	
```

## 几个常量

```go
const (
	mutexLocked = 1 << iota
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
	starvationThresholdNs = 1e6
)
```

* mutexLocked 值为1，根据 `mutex.state & mutexLocked` 得到 mutex 的加锁状态，结果为1表示已加锁，0表示未加锁
* mutexWoken 值为2（二进制：10），根据 `mutex.state & mutexWoken` 得到 mutex 的唤醒状态，结果为1表示已唤醒，0表示未唤醒
* mutexStarving 值为4（二进制：100），根据 `mutex.state & mutexStarving` 得到 mutex 的饥饿状态，结果为1表示处于饥饿状态，0表示处于正常状态
* mutexWaiterShift 值为3，根据 `mutex.state >> mutexWaiterShift` 得到当前等待的 goroutine 数目
* starvationThresholdNs 值为1e6纳秒，也就是1毫秒，当等待队列中队首 goroutine 等待时间超过 starvationThresholdNs，mutex 进入饥饿模式

## 饥饿模式与正常模式

Mutex 有两种工作模式：正常模式和饥饿模式

在正常模式中，等待者按照 FIFO 的顺序排队获取锁，但是一个被唤醒的等待者有时候并不能获取 mutex，它还需要和新到来的 goroutine 们竞争 mutex 的使用权。新到来的 goroutine 存在一个优势，它们已经在 CPU 上运行且它们数量很多，因此一个被唤醒的等待者有很大的概率获取不到锁，在这种情况下它处在等待队列的前面。如果一个 goroutine 等待 mutex 释放的时间超过1ms，它就会将 mutex 切换到饥饿模式

在饥饿模式中，mutex 的所有权直接从解锁的 goroutine 递交到等待队列中排在最前方的 goroutine。新到达的 goroutine 们不要尝试去获取 mutex，即使它看起来是在解锁状态，也不要试图自旋，而是排到等待队列的尾部

如果一个等待者获得 mutex 的所有权，并且看到以下两种情况中的任一种：1) 它是等待队列中的最后一个，或者 2) 它等待的时间少于1ms，它便将 mutex 切换回正常操作模式

## 函数

以下代码已经去掉了与核心代码无关的 race 代码

### Lock

Lock 方法申请对 mutex 加锁，Lock 执行的时候，分三种情况

1. **无冲突** 通过 CAS 操作把当前状态设置为加锁状态
2. **有冲突 开始自旋**，并等待锁释放，如果其他 goroutine 在这段时间内释放了该锁，直接获得该锁；如果没有释放，进入3
3. **有冲突，且已经过了自旋阶段** 通过调用 semacquire 函数来让当前 goroutine 进入等待状态

```go
func (m *Mutex) Lock() {
	// 查看 state 是否为0，如果是则表示可以加锁，将其状态转换为1，当前 goroutine 加锁成功，函数返回
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}

	var waitStartTime int64 // 当前 goroutine 开始等待的时间
	starving := false // mutex 当前所处的模式
	awoke := false // 当前 goroutine 是否被唤醒
	iter := 0 // 自旋迭代的次数
	old := m.state // old 保存当前 mutex 的状态
	for {
		// 当 mutex 处于正常工作模式且能够自旋的时候，进行自旋操作（汇编实现，内部持续调用 PAUSE 指令，消耗 CPU 时间）
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			// 将 mutex.state 的倒数第二位设置为1，用来告 Unlock 操作，存在 goroutine
			// 即将得到锁，不需要唤醒其他 goroutine
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
		new := old
		// 当 mutex 不处于饥饿状态的时候，将 new 的第一位设置为1，即加锁
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		// 当 mutex 处于加锁状态或饥饿状态的时候，新到来的 goroutine 进入等待队列
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		// 当前 goroutine 将 mutex 切换为饥饿状态，但如果当前 mutex 未加锁，则不需要切换
		// Unlock 操作希望饥饿模式存在等待者
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			// 当前 goroutine 被唤醒，将 mutex.state 倒数第二位重置
			if new&mutexWoken == 0 {
				throw("sync: inconsistent mutex state")
			}
			new &^= mutexWoken
		}
		// 调用 CAS 更新 state 状态
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			// mutex 处于未加锁，正常模式下，当前 goroutine 获得锁
			if old&(mutexLocked|mutexStarving) == 0 {
				break
			}
			// queueLifo 为 true 代表当前 goroutine 是等待状态的 goroutine
			queueLifo := waitStartTime != 0
			if waitStartTime == 0 {
				// 记录开始等待时间
				waitStartTime = runtime_nanotime()
			}
			// 将被唤醒却没得到锁的 goroutine 插入当前等待队列的最前端
			runtime_SemacquireMutex(&m.sema, queueLifo)
			// 如果当前 goroutine 等待时间超过starvationThresholdNs，mutex 进入饥饿模式
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 {
				if old&(mutexLocked|mutexWoken) != 0 || old>>mutexWaiterShift == 0 {
					throw("sync: inconsistent mutex state")
				}
				// 等待状态的 goroutine - 1
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				// 如果不是饥饿模式了或者当前等待着只剩下一个，退出饥饿模式
				if !starving || old>>mutexWaiterShift == 1 {
					delta -= mutexStarving
				}
				// 更新状态
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true
			iter = 0
		} else {
			old = m.state
		}
	}
}
```

### Unlock

Unlock方法释放所申请的锁

```go
func (m *Mutex) Unlock() {
	// mutex 的 state 减去1，加锁状态 -> 未加锁
	new := atomic.AddInt32(&m.state, -mutexLocked)
	// 未 Lock 直接 Unlock，报 panic
	if (new+mutexLocked)&mutexLocked == 0 {
		throw("sync: unlock of unlocked mutex")
	}
	// mutex 正常模式
	if new&mutexStarving == 0 {
		old := new
		for {
			// 如果没有等待者，或者已经存在一个 goroutine 被唤醒或得到锁，或处于饥饿模式
			// 无需唤醒任何处于等待状态的 goroutine
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			// 等待者数量减1，并将唤醒位改成1
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				// 唤醒一个阻塞的 goroutine，但不是唤醒第一个等待者
				runtime_Semrelease(&m.sema, false)
				return
			}
			old = m.state
		}
	} else {
		// mutex 饥饿模式，直接将 mutex 拥有权移交给等待队列最前端的 goroutine
		runtime_Semrelease(&m.sema, true)
	}
}
```
	