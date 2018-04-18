1. defer runtime.GOMAXPROCS(runtime.GOMAXPROCS(1))
2. runtime.ReadMemStats(&m1)
3. defer_lock.go

```go
package main

import (
	"sync"
)

func main() {
	var mu sync.Locker = new(I)
	defer LockUnlock(mu)()
	println("doing")
}

func LockUnlock(mu sync.Locker) (unlock func()) {
	mu.Lock()
	return mu.Unlock
}

type I struct{}

func (i *I) Lock() {
	println("lock")
}

func (i *I) Unlock() {
	println("unlock")
}
```

4. 	buf := make([]byte, len(b.buf), 2*cap(b.buf)+n)  为什么是2倍呢？
5. // NOTE(rsc): This function does NOT call the runtime cmpstring function,
	// because we do not want to provide any performance justification for
	// using strings.Compare. Basically no one should use strings.Compare.
	// As the comment above says, it is here only for symmetry with package bytes.
	// If performance is important, the compiler should be changed to recognize
	// the pattern so that all code doing three-way comparisons, not just code
	// using strings.Compare, can benefit.

6. b.buf = append(b.buf, s...) s是string，b.buf是[]byte
7. int int64的问题？ 在32位机器上进行int64原子操作时的panic
8. defer LockUnlock(mu),如果LockUnlock(mu)没有带()，则会丢失func函数的执行
9. 	if r < utf8.RuneSelf 
10. 	if cap(b.buf)-l < utf8.UTFMax {
11. Example
12. xxx_test.go 其实是xxx包，但是又不想放到xxx包里面，因为它只是提供给_test.go包使用的函数。
13. rune 码点的处理（reader.go 	prevRune int   // index of previous rune; or < 0）
14. off 与 offset 尴尬的问题
15. whence where when ？[wiki-whence](https://en.wiktionary.org/wiki/whence)
16. // It is similar to bytes.NewBufferString but more efficient and read-only.

## 参考资料

1. [wiki-whence](https://en.wiktionary.org/wiki/whence)
2. [Go 延迟函数 defer 详解](https://mp.weixin.qq.com/s/5xeAOYi3OoxCEPe-S2RE2Q)
