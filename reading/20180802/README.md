# golang中goroutine的调度
郑宝杨(boya) 2018-08-01 listomebao@gmail.com

## 阅读源码前可以阅读的资料
* [Goroutine背后的系统知识](http://blog.jobbole.com/35304/)
* [golang源码剖析-雨痕老师](https://github.com/qyuhen/book)
* [go-intervals](https://github.com/teh-cmc/go-internals)
* [也谈goroutine调度器](https://tonybai.com/2017/06/23/an-intro-about-goroutine-scheduler/)


## golang的调度模型概览
调度的机制用一句话描述：  
runtime准备好G,P,M，然后M绑定P，M从各种队列中获取G，切换到G的执行栈上并执行G上的任务函数，调用goexit做清理工作并回到M，如此反复。

### 基本概念
#### M（machine）
* M代表着真正的执行计算资源，可以认为它就是os thread（系统线程）。
* M是真正调度系统的执行者，每个M就像一个勤劳的工作者，总是从各种队列中找到可运行的G，而且这样M的可以同时存在多个。
* M在绑定有效的P后，进入调度循环，而且M并不保留G状态，这是G可以跨M调度的基础。

#### P（processor）
* P表示逻辑processor，是线程M的执行的上下文。
* P的最大作用是其拥有的各种G对象队列、链表、cache和状态。

#### G（goroutine）
* 调度系统的最基本单位goroutine，存储了goroutine的执行stack信息、goroutine状态以及goroutine的任务函数等。
* 在G的眼中只有P，P就是运行G的“CPU”。
* 相当于两级线程

#### 线程实现模型
来自`Go并发编程实战`
```
                    +-------+       +-------+      
                    |  KSE  |       |  KSE  |          
                    +-------+       +-------+      
                        |               |                       内核空间
- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -        
                        |               |                       用户空间
                    +-------+       +-------+
                    |   M   |       |   M   |
                    +-------+       +-------+
                  |          |         |          |
              +------+   +------+   +------+   +------+            
              |   P  |   |   P  |   |   P  |   |   P  |
              +------+   +------+   +------+   +------+   
           |     |     |     |     |     |     |     |     | 
         +---+ +---+ +---+ +---+ +---+ +---+ +---+ +---+ +---+ 
         | G | | G | | G | | G | | G | | G | | G | | G | | G | 
         +---+ +---+ +---+ +---+ +---+ +---+ +---+ +---+ +---+ 
```
* KSE（Kernel Scheduling Entity）是内核调度实体           
* M与P，P与G之前的关联都是动态的，可以变的

### 关系示意图
来自`golang源码剖析`
```
                            +-------------------- sysmon ---------------//------+ 
                            |                                                   |
                            |                                                   |
               +---+      +---+-------+                   +--------+          +---+---+
go func() ---> | G | ---> | P | local | <=== balance ===> | global | <--//--- | P | M |
               +---+      +---+-------+                   +--------+          +---+---+
                            |                                 |                 | 
                            |      +---+                      |                 |
                            +----> | M | <--- findrunnable ---+--- steal <--//--+
                                   +---+ 
                                     |
                                   mstart
                                     |
              +--- execute <----- schedule 
              |                      |   
              |                      |
              +--> G.fn --> goexit --+ 


              1. go func() 语气创建G。
              2. 将G放入P的本地队列（或者平衡到全局全局队列）。
              3. 唤醒或新建M来执行任务。
              4. 进入调度循环
              5. 尽力获取可执行的G，并执行
              6. 清理现场并且重新进入调度循环

```


## GPM的来由
### 特殊的g0和m0
g0和m0是在`proc.go`文件中的两个全局变量，m0就是进程启动后的初始线程，g0也是代表着初始线程的stack  
`asm_amd64.go` --> runtime·rt0_go(SB)
```go 
	// 程序刚启动的时候必定有一个线程启动（主线程）
	// 将当前的栈和资源保存在g0
	// 将该线程保存在m0
	// tls: Thread Local Storage
	// set the per-goroutine and per-mach "registers"
	get_tls(BX)
	LEAQ	runtime·g0(SB), CX
	MOVQ	CX, g(BX)
	LEAQ	runtime·m0(SB), AX

	// save m->g0 = g0
	MOVQ	CX, m_g0(AX)
	// save m0 to g0->m
	MOVQ	AX, g_m(CX)
```



### M的一生
#### M的创建
`proc.go`
```go
// Create a new m. It will start off with a call to fn, or else the scheduler.
// fn needs to be static and not a heap allocated closure.
// May run with m.p==nil, so write barriers are not allowed.
//go:nowritebarrierrec
// 创建一个新的m，它将从fn或者调度程序开始
func newm(fn func(), _p_ *p) {
	// 根据fn和p和绑定一个m对象
	mp := allocm(_p_, fn)
	// 设置当前m的下一个p为_p_
	mp.nextp.set(_p_)
	mp.sigmask = initSigmask
	...
	// 真正的分配os thread
	newm1(mp)
}
```

```go
func newm1(mp *m) {
	// 对cgo的处理
	...
	execLock.rlock() // Prevent process clone.
	// 创建一个系统线程
	newosproc(mp, unsafe.Pointer(mp.g0.stack.hi))
	execLock.runlock()
}
```
#### 状态
``` 
       mstart
          |
          v        找不到可执行任务，gc STW，
      +------+     任务执行时间过长，系统阻塞等   +------+
      | spin | ----------------------------> |unspin| 
      +------+          mstop                +------+
          ^                                      |
          |                                      v
      notewakeup <-------------------------  notesleep
```

#### M的一些问题
https://github.com/golang/go/issues/14592

### P的一生
#### P的创建
`proc.go`
```go
// Change number of processors. The world is stopped, sched is locked.
// gcworkbufs are not being modified by either the GC or
// the write barrier code.
// Returns list of Ps with local work, they need to be scheduled by the caller.
// 所有的P都在这个函数分配，不管是最开始的初始化分配，还是后期调整
func procresize(nprocs int32) *p {
	old := gomaxprocs
	// 如果 gomaxprocs <=0 抛出异常
	if old < 0 || nprocs <= 0 {
		throw("procresize: invalid arg")
	}
  ...
	// Grow allp if necessary.
	if nprocs > int32(len(allp)) {
		// Synchronize with retake, which could be running
		// concurrently since it doesn't run on a P.
		lock(&allpLock)
		if nprocs <= int32(cap(allp)) {
			allp = allp[:nprocs]
		} else {
			// 分配nprocs个*p
			nallp := make([]*p, nprocs)
			// Copy everything up to allp's cap so we
			// never lose old allocated Ps.
			copy(nallp, allp[:cap(allp)])
			allp = nallp
		}
		unlock(&allpLock)
	}

	// initialize new P's
	for i := int32(0); i < nprocs; i++ {
		pp := allp[i]
		if pp == nil {
			pp = new(p)
			pp.id = i
			pp.status = _Pgcstop            // 更改状态
			pp.sudogcache = pp.sudogbuf[:0] //将sudogcache指向sudogbuf的起始地址
			for i := range pp.deferpool {
				pp.deferpool[i] = pp.deferpoolbuf[i][:0]
			}
			pp.wbBuf.reset()
			// 将pp保存到allp数组里, allp[i] = pp
			atomicstorep(unsafe.Pointer(&allp[i]), unsafe.Pointer(pp))
		}
		...
	}
  ...

	_g_ := getg()
	// 如果当前的M已经绑定P，继续使用，否则将当前的M绑定一个P
	if _g_.m.p != 0 && _g_.m.p.ptr().id < nprocs {
		// continue to use the current P
		_g_.m.p.ptr().status = _Prunning
	} else {
		// release the current P and acquire allp[0]
		// 获取allp[0]
		if _g_.m.p != 0 {
			_g_.m.p.ptr().m = 0
		}
		_g_.m.p = 0
		_g_.m.mcache = nil
		p := allp[0]
		p.m = 0
		p.status = _Pidle
		// 将当前的m和p绑定
		acquirep(p)
		if trace.enabled {
			traceGoStart()
		}
	}
	var runnablePs *p
	for i := nprocs - 1; i >= 0; i-- {
		p := allp[i]
		if _g_.m.p.ptr() == p {
			continue
		}
		p.status = _Pidle
		if runqempty(p) { // 将空闲p放入空闲链表
			pidleput(p)
		} else {
			p.m.set(mget())
			p.link.set(runnablePs)
			runnablePs = p
		}
	}
	stealOrder.reset(uint32(nprocs))
	var int32p *int32 = &gomaxprocs // make compiler check that gomaxprocs is an int32
	atomic.Store((*uint32)(unsafe.Pointer(int32p)), uint32(nprocs))
	return runnablePs
}
```
所有的P在程序启动的时候就设置好了，并用一个allp slice维护，可以调用runtime.GOMAXPROCS调整P的个数，虽然代价很大
#### 状态转换
```
                                            acquirep(p)        
                          不需要使用的P       P和M绑定的时候       进入系统调用       procresize()
new(p)  -----+        +---------------+     +-----------+     +------------+    +----------+
            |         |               |     |           |     |            |    |          |
            |   +------------+    +---v--------+    +---v--------+    +----v-------+    +--v---------+
            +-->|  _Pgcstop  |    |    _Pidle  |    |  _Prunning |    |  _Psyscall |    |   _Pdead   |
                +------^-----+    +--------^---+    +--------^---+    +------------+    +------------+
                       |            |     |            |     |            |
                       +------------+     +------------+     +------------+
                           GC结束            releasep()        退出系统调用
                                            P和M解绑                      
``` 
P的数量默认等于cpu的个数，很多人认为runtime.GOMAXPROCS可以限制系统线程的数量，但这是错误的，M是按需创建的，和runtime.GOMAXPROCS没有关系。
### G的一生

#### G的创建
`proc.go`
```go
// Create a new g running fn with siz bytes of arguments.
// Put it on the queue of g's waiting to run.
// The compiler turns a go statement into a call to this.
// Cannot split the stack because it assumes that the arguments
// are available sequentially after &fn; they would not be
// copied if a stack split occurred.
//go:nosplit
// 新建一个goroutine，
// 􏳄 用fn + PtrSize 获取第一个参数的地址，也就是argp
// 用siz - 8 获取pc地址
func newproc(siz int32, fn *funcval) {
	argp := add(unsafe.Pointer(&fn), sys.PtrSize)
	pc := getcallerpc()
	// 用g0的栈创建G对象
	systemstack(func() {
		newproc1(fn, (*uint8)(argp), siz, pc)
	})
}
```

```go
// Create a new g running fn with narg bytes of arguments starting
// at argp. callerpc is the address of the go statement that created
// this. The new g is put on the queue of g's waiting to run.
// 根据函数参数和函数地址，创建一个新的G，然后将这个G加入队列等待运行
func newproc1(fn *funcval, argp *uint8, narg int32, callerpc uintptr) {
	_g_ := getg()

	if fn == nil {
		_g_.m.throwing = -1 // do not dump full stacks
		throw("go of nil func value")
	}
	_g_.m.locks++ // disable preemption because it can be holding p in a local var
	siz := narg
	siz = (siz + 7) &^ 7

	// We could allocate a larger initial stack if necessary.
	// Not worth it: this is almost always an error.
	// 4*sizeof(uintreg): extra space added below
	// sizeof(uintreg): caller's LR (arm) or return address (x86, in gostartcall).
	// 如果函数的参数大小比2048大的话，直接panic
	if siz >= _StackMin-4*sys.RegSize-sys.RegSize {
		throw("newproc: function arguments too large for new goroutine")
	}

	// 从m中获取p
	_p_ := _g_.m.p.ptr()
	// 从gfree list获取g
	newg := gfget(_p_)
	// 如果没获取到g，则新建一个
	if newg == nil {
		newg = malg(_StackMin)
		casgstatus(newg, _Gidle, _Gdead) //将g的状态改为_Gdead
		// 添加到allg数组，防止gc扫描清除掉
		allgadd(newg) // publishes with a g->status of Gdead so GC scanner doesn't look at uninitialized stack.
	}
	if newg.stack.hi == 0 {
		throw("newproc1: newg missing stack")
	}

	if readgstatus(newg) != _Gdead {
		throw("newproc1: new g is not Gdead")
	}

	totalSize := 4*sys.RegSize + uintptr(siz) + sys.MinFrameSize // extra space in case of reads slightly beyond frame
	totalSize += -totalSize & (sys.SpAlign - 1)                  // align to spAlign
	sp := newg.stack.hi - totalSize
	spArg := sp
	if usesLR {
		// caller's LR
		*(*uintptr)(unsafe.Pointer(sp)) = 0
		prepGoExitFrame(sp)
		spArg += sys.MinFrameSize
	}
	if narg > 0 {
		// copy参数
		memmove(unsafe.Pointer(spArg), unsafe.Pointer(argp), uintptr(narg))
		// This is a stack-to-stack copy. If write barriers
		// are enabled and the source stack is grey (the
		// destination is always black), then perform a
		// barrier copy. We do this *after* the memmove
		// because the destination stack may have garbage on
		// it.
		if writeBarrier.needed && !_g_.m.curg.gcscandone {
			f := findfunc(fn.fn)
			stkmap := (*stackmap)(funcdata(f, _FUNCDATA_ArgsPointerMaps))
			// We're in the prologue, so it's always stack map index 0.
			bv := stackmapdata(stkmap, 0)
			bulkBarrierBitmap(spArg, spArg, uintptr(narg), 0, bv.bytedata)
		}
	}

	memclrNoHeapPointers(unsafe.Pointer(&newg.sched), unsafe.Sizeof(newg.sched))
	newg.sched.sp = sp
	newg.stktopsp = sp
	// 保存goexit的地址到sched.pc
	newg.sched.pc = funcPC(goexit) + sys.PCQuantum // +PCQuantum so that previous instruction is in same function
	newg.sched.g = guintptr(unsafe.Pointer(newg))
	gostartcallfn(&newg.sched, fn)
	newg.gopc = callerpc
	newg.startpc = fn.fn
	if _g_.m.curg != nil {
		newg.labels = _g_.m.curg.labels
	}
	if isSystemGoroutine(newg) {
		atomic.Xadd(&sched.ngsys, +1)
	}
	newg.gcscanvalid = false
	// 更改当前g的状态为_Grunnable
	casgstatus(newg, _Gdead, _Grunnable)

	if _p_.goidcache == _p_.goidcacheend {
		// Sched.goidgen is the last allocated id,
		// this batch must be [sched.goidgen+1, sched.goidgen+GoidCacheBatch].
		// At startup sched.goidgen=0, so main goroutine receives goid=1.
		_p_.goidcache = atomic.Xadd64(&sched.goidgen, _GoidCacheBatch)
		_p_.goidcache -= _GoidCacheBatch - 1
		_p_.goidcacheend = _p_.goidcache + _GoidCacheBatch
	}
	// 生成唯一的goid
	newg.goid = int64(_p_.goidcache)
	_p_.goidcache++
	if raceenabled {
		newg.racectx = racegostart(callerpc)
	}
	if trace.enabled {
		traceGoCreate(newg, newg.startpc)
	}
	// 将当前新生成的g，放入队列
	runqput(_p_, newg, true)

	// 如果有空闲的p 且 m没有处于自旋状态 且 main goroutine已经启动，那么唤醒某个m来执行任务
	if atomic.Load(&sched.npidle) != 0 && atomic.Load(&sched.nmspinning) == 0 && mainStarted {
		wakep()
	}
	_g_.m.locks--
	if _g_.m.locks == 0 && _g_.preempt { // restore the preemption request in case we've cleared it in newstack
		_g_.stackguard0 = stackPreempt
	}
}
```

#### G的状态图
```
                                                      +------------+
                                      ready           |            |
                                  +------------------ |  _Gwaiting |
                                  |                   |            |
                                  |                   +------------+
                                  |                         ^ park_m
                                  V                         | 
  +------------+            +------------+  execute   +------------+            +------------+    
  |            |  newproc   |            | ---------> |            |   goexit   |            |
  |  _Gidle    | ---------> | _Grunnable |  yield     | _Grunning  | ---------> |   _Gdead   |      
  |            |            |            | <--------- |            |            |            |
  +------------+            +-----^------+            +------------+            +------------+
                                  |         entersyscall |      ^ 
                                  |                      V      | existsyscall
                                  |                   +------------+
                                  |   existsyscall    |            |
                                  +------------------ |  _Gsyscall |
                                                      |            |
                                                      +------------+

```
新建的G都是_Grunnable的，新建G的时候优先从gfree list从获取G，这样可以复用G，所以上图的状态不是完整的，_Gdead通过newproc会变为_Grunnable，
通过go func()的语法新建的G，并不是直接运行，而是放入可运行的队列中，什么时候运行用于并不能决定，而是搞调度系统去自发的运行。


