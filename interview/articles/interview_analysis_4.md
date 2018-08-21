# Golang面试题解析（四）

## 31. 算法
在utf8字符串判断是否包含指定字符串，并返回下标。
"北京天安门最美丽" , "天安门"
结果：2

解答：
```go  
import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println(Utf8Index("北京天安门最美丽", "天安门"))
	fmt.Println(strings.Index("北京天安门最美丽", "男"))
	fmt.Println(strings.Index("", "男"))
	fmt.Println(Utf8Index("12ws北京天安门最美丽", "天安门"))
}

func Utf8Index(str, substr string) int {
	asciiPos := strings.Index(str, substr)
	if asciiPos == -1 || asciiPos == 0 {
		return asciiPos
	}
	pos := 0
	totalSize := 0
	reader := strings.NewReader(str)
	for _, size, err := reader.ReadRune(); err == nil; _, size, err = reader.ReadRune() {
		totalSize += size
		pos++
		// 匹配到
		if totalSize == asciiPos {
			return pos
		}
	}
	return pos
}
```

## 32，编程
实现一个单例

解答：
```go  
package main

import "sync"

// 实现一个单例

type singleton struct{}

var ins *singleton
var mu sync.Mutex

//懒汉加锁:虽然解决并发的问题，但每次加锁是要付出代价的
func GetIns() *singleton {
	mu.Lock()
	defer mu.Unlock()

	if ins == nil {
		ins = &singleton{}
	}
	return ins
}

//双重锁:避免了每次加锁，提高代码效率
func GetIns1() *singleton {
	if ins == nil {
		mu.Lock()
		defer mu.Unlock()
		if ins == nil {
			ins = &singleton{}
		}
	}
	return ins
}

//sync.Once实现
var once sync.Once

func GetIns2() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}

```

## 33,执行下面的代码发生什么？

```go  
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

```
### 考点:**channel**
往已经关闭的channel写入数据会panic的。
结果：
```
panic: send on closed channel
```
## 34,执行下面的代码发生什么？
```
import "fmt"

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	c := &ConfigOne{}
	c.String()
}
```
### 考点:**fmt.Sprintf**
如果类型实现String()，％v和％v格式将使用String()的值。因此，对该类型的String()函数内的类型使用％v会导致无限递归。
编译报错：
```
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow
```
## 35，编程题
反转整数
反转一个整数，例如：

例子1: x = 123, return 321  
例子2: x = -123, return -321  

输入的整数要求是一个 32bit 有符号数，如果反转后溢出，则输出 0  

```
func reverse(x int) (num int) {
	for x != 0 {
		num = num*10 + x%10
		x = x / 10
	}
	// 使用 math 包中定义好的最大最小值
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return
}

```

## 36，编程题
合并重叠区间
给定一组 区间，合并所有重叠的 区间。

例如：
给定：[1,3],[2,6],[8,10],[15,18]
返回：[1,6],[8,10],[15,18]
```
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := make([]Interval, 0)
	swap := Interval{}
	for k, v := range intervals {
		if k == 0 {
			swap = v
			continue
		}
		if v.Start <= swap.End {
			swap.End = v.End
		} else {
			res = append(res, swap)
			swap = v
		}
	}
	res = append(res, swap)
	return res
}
```

## 37.输出什么？
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(len("你好bj!"))
}
```
### 考点:**编码长度**
输出9

## 38.编译并运行如下代码会发生什么？
```
package main

import "fmt"

type Test struct {
	Name string
}

var list map[string]Test

func main() {

	list = make(map[string]Test)
	name := Test{"xiaoming"}
	list["name"] = name
	list["name"].Name = "Hello"
	fmt.Println(list["name"])
}
```
### 考点:**map**
编程报错`cannot assign to struct field list["name"].Name in map`。
因为list["name"]不是一个普通的指针值，map的value本身是不可寻址的，因为map中的值会在内存中移动，并且旧的指针地址在map改变时会变得无效。
定义的是var list map[string]Test，注意哦Test不是指针，而且map我们都知道是可以自动扩容的，那么原来的存储name的Test可能在地址A，但是如果map扩容了地址A就不是原来的Test了，所以go就不允许我们写数据。你改为var list map[string]*Test试试看。

## 39.ABCD中哪一行存在错误？
```go  
type S struct {
}

func f(x interface{}) {
}

func g(x *interface{}) {
}

func main() {
	s := S{}
	p := &s
	f(s) //A
	g(s) //B
	f(p) //C
	g(p) //D

}
```
### 考点:**interface**
看到这道题需要第一时间想到的是Golang是强类型语言，interface是所有golang类型的父类，类似Java的Object。
函数中`func f(x interface{})`的`interface{}`可以支持传入golang的任何类型，包括指针，但是函数`func g(x *interface{})`只能接受`*interface{}`.


## 40.编译并运行如下代码会发生什么？
```go  
package main

import (
	"sync"
	//"time"
)

const N = 10

var wg = &sync.WaitGroup{}

func main() {

	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

}
```
### 考点:**WaitGroup**
这是使用WaitGroup经常犯下的错误！请各位同学多次运行就会发现输出都会不同甚至又出现报错的问题。
这是因为`go`执行太快了，导致`wg.Add(1)`还没有执行main函数就执行完毕了。
改为如下试试
```
for i := 0; i < N; i++ {
        wg.Add(1)
		go func(i int) {
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
```



## 附录
https://zhuanlan.zhihu.com/p/35058068?hmsr=toutiao.io&utm_medium=toutiao.io&utm_source=toutiao.io

https://stackoverflow.com/questions/42600920/runtime-goroutine-stack-exceeds-1000000000-byte-limit-fatal-error-stack-overf

https://studygolang.com/topics/3853