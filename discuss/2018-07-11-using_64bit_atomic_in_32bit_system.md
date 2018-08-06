## 2018-07-11

来源：《Go 夜读》微信群

时间：2018-07-11

----

## 在32位系统中使用64位原子操作的坑

### 先来个例子
```go
type Y struct {
	a bool
	v uint64
}

func TestAtomicY(t *testing.T) {
	var y Y
	atomic.AddUint64(&y.v, 1) // panic in 32bit system
}
```

在上面的例子中，如果在64位系统中运行是没问题的的，但是在32位系统中会panic。

### 为何会 Panic ？
这个简单回答可以查看，go官方的文档[atomic-pkt-note](https://golang.google.cn/pkg/sync/atomic/#pkg-note-BUG)的内容：  
```html
Bugs

On x86-32, the 64-bit functions use instructions unavailable before the Pentium MMX.

On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.

On both ARM and x86-32, it is the caller's responsibility to arrange for 64-bit alignment of 64-bit words accessed atomically. The first word in a variable or in an allocated struct, array, or slice can be relied upon to be 64-bit aligned.
```
意思就是你如果要在32位系统中用64位的原子操作，必须要自己保证64位对齐，也就是8字节对齐。  
如果要看汇编怎么判断的可以查看[atomic·Xadd64](https://github.com/golang/go/blob/master/src/runtime/internal/atomic/asm_386.s#L97)


### 如果需要在一个struct中维护多个64位字段，且都需要原子操作，怎么办？
最简单的办法就是将所有64位字段放在struct的头部，这样就可以保证8字节对齐，如果你把这个struct嵌入在别的结构体，也要记得嵌入到头部。


### 一些例子
```go
package aotmic_test

import (
	"log"
	"sync/atomic"
	"testing"
	"unsafe"
)

type X struct {
	v uint64
	x uint64
	a bool
	z uint64
	y uint32
}

func TestAtomic(t *testing.T) {
	var x X
	log.Printf("x.a=%p, offset=%d, alig=%d", &x.a, unsafe.Offsetof(x.a), unsafe.Alignof(x.a))
	log.Printf("x.v=%p, offset=%d, alig=%d", &x.v, unsafe.Offsetof(x.v), unsafe.Alignof(x.v))
	log.Printf("x.x=%p, offset=%d, alig=%d", &x.x, unsafe.Offsetof(x.x), unsafe.Alignof(x.x))
	log.Printf("x.y=%p, offset=%d, alig=%d", &x.y, unsafe.Offsetof(x.y), unsafe.Alignof(x.y))
	log.Printf("x.z=%p, offset=%d, alig=%d", &x.z, unsafe.Offsetof(x.z), unsafe.Alignof(x.z))
	log.Printf("x.v=%p", &x.v)
	atomic.AddUint64(&x.z, 1) // panic
}

type Y struct {
	a bool
	X
}

func TestAtomicY(t *testing.T) {
	var y Y
	x := y.X
	atomic.AddUint64(&x.v, 1)
	atomic.AddUint64(&y.X.v, 1) // panic
}

type Y2 struct {
	X
	a bool
}

func TestAtomicY2(t *testing.T) {
	y := &Y2{}
	atomic.AddUint64(&y.X.v, 1)
}

type Temp struct {
	A byte
	B [2]byte
	C int64
}

func TestAtomicTemp(t *testing.T) {
	var x Temp
	log.Printf("sizof=%d", unsafe.Sizeof(x))
	log.Printf("x.A=%p, offset=%d, alig=%d", &x.A, unsafe.Offsetof(x.A), unsafe.Alignof(x.A))
	log.Printf("x.B=%p, offset=%d, alig=%d", &x.B, unsafe.Offsetof(x.B), unsafe.Alignof(x.B))
	log.Printf("x.C=%p, offset=%d, alig=%d", &x.C, unsafe.Offsetof(x.C), unsafe.Alignof(x.C))
}

```

### 参考链接
 https://go101.org/article/memory-layout.html

https://github.com/golang/go/issues/5278




