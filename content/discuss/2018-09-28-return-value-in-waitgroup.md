---
title: 2018-09-28 goroutine 中怎么得到返回值
---

来源: Wechat discuss

时间：2018-09-28

## goroutine 中怎么得到返回值？


```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("开始。。。")
	wg.Add(2)

	time1 := time.Now()

	go f1(1, 4)
	go f2(2, 5)

	wg.Wait()

	time2 := time.Since(time1)
	fmt.Printf("cost:%v\n", time2)
}

func f1(n1, n2 int) int {
	time.Sleep(time.Second)
	wg.Done()
	fmt.Println("f1")
	return n1 + n2
}

func f2(n1, n2 int) int {
	time.Sleep(time.Second)
	wg.Done()
	fmt.Println("f2")
	return n1 + n2
}
```

如果我想要获取 `f1` ，`f2` 的返回值，有什么办法？

## 解决方案1

```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("开始。。。")
	wg.Add(2)

	time1 := time.Now()
	var x int
	go func() {
		x = f1(1, 4)
		wg.Done()
	}()

	var y int
	go func() {
		y = f2(2, 5)
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("f1:%d\n", x)
	fmt.Printf("f2:%d\n", y)

	time2 := time.Since(time1)
	fmt.Printf("cost:%v\n", time2)
}

func f1(n1, n2 int) int {
	time.Sleep(time.Second)

	fmt.Println("f1")
	return n1 + n2
}

func f2(n1, n2 int) int {
	time.Sleep(time.Second)
	fmt.Println("f2")
	return n1 + n2
}
```

这里需要注意的地方就是，我们要读取值，必须等待 goroutine 执行完之后才可以。

## 解决方案2

WaitGroup+Channel 结合的方式：

```golang
package main

import (
	"fmt"
	"sync"
	"time"
)

// R return value
type R struct {
	name string
	ret  int
}

func main() {
	fmt.Println("开始。。。")

	time1 := time.Now()
	var wg sync.WaitGroup
	retCh := make(chan *R, 2)
	wg.Add(1)
	go func() {
		retCh <- f1(1, 4)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		retCh <- f2(2, 5)
		wg.Done()
	}()
	wg.Wait()
	close(retCh)
	fmt.Println("after wait ...")
	for r := range retCh {
		fmt.Printf("func name:%s, return:%d\n", r.name, r.ret)
	}

	time2 := time.Since(time1)
	fmt.Printf("cost:%v\n", time2)
}

func f1(n1, n2 int) *R {
	time.Sleep(time.Second)

	fmt.Println("f1")
	return &R{
		name: "f1",
		ret:  n1 + n2,
	}
}

func f2(n1, n2 int) *R {
	time.Sleep(time.Second)

	fmt.Println("f2")
	return &R{
		name: "f2",
		ret:  n1 + n2,
	}
}
```

## 参考资料

1. [sync.WaitGroup](https://golang.org/pkg/sync/#WaitGroup)
>A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.

>A WaitGroup must not be copied after first use.
