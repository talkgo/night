# Golang面试题解析（一）
最近在很多地方看到了[golang的面试题](https://zhuanlan.zhihu.com/p/26972862)，看到了很多人对Golang的面试题心存恐惧，也是为了复习基础，我把解题的过程总结下来。
## 面试题
### 1. 写出下面代码输出内容。
```go

package main

import (
    "fmt"
)

func main() {
    defer_call()
}

func defer_call() {
    defer func() { fmt.Println("打印前") }()
    defer func() { fmt.Println("打印中") }()
    defer func() { fmt.Println("打印后") }()

    panic("触发异常")
}
```
考点：**defer执行顺序**
解答：
defer 是**后进先出**。
panic 需要等defer 结束后才会向上传递。
出现panic恐慌时候，会先按照defer的后入先出的顺序执行，最后才会执行panic。
```go
打印后
打印中
打印前
panic: 触发异常

```
### 2. 以下代码有什么问题，说明原因。
```go
type student struct {
    Name string
    Age  int
}

func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    for _, stu := range stus {
        m[stu.Name] = &stu
    }

}

```
考点：**foreach**
解答：
这样的写法初学者经常会遇到的，很危险！
与Java的foreach一样，都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针，
最终该指针的值为遍历的最后一个struct的值拷贝。
就像想修改切片元素的属性：
```go
for _, stu := range stus {
    stu.Age = stu.Age+10
}
```
也是不可行的。
大家可以试试打印出来：
```
func pase_student() {
    m := make(map[string]*student)
    stus := []student{
        {Name: "zhou", Age: 24},
        {Name: "li", Age: 23},
        {Name: "wang", Age: 22},
    }
    // 错误写法
    for _, stu := range stus {
        m[stu.Name] = &stu
    }

    for k,v:=range m{
        println(k,"=>",v.Name)
    }

    // 正确
    for i:=0;i<len(stus);i++  {
        m[stus[i].Name] = &stus[i]
    }
    for k,v:=range m{
        println(k,"=>",v.Name)
    }
}

```

### 3. 下面的代码会输出什么，并说明原因
```go
func main() {
    runtime.GOMAXPROCS(1)
    wg := sync.WaitGroup{}
    wg.Add(20)
    for i := 0; i < 10; i++ {
        go func() {
            fmt.Println("A: ", i)
            wg.Done()
        }()
    }
    for i := 0; i < 10; i++ {
        go func(i int) {
            fmt.Println("B: ", i)
            wg.Done()
        }(i)
    }
    wg.Wait()
}

```
考点：**go执行的随机性和闭包**
解答：
谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。
但是`A: `均为输出10，`B: `从0~9输出(顺序不定)。
第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。
故go func执行时，i的值始终是10。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。
尾部(i)将发生值拷贝，go func内部指向值拷贝地址。

### 4. 下面代码会输出什么？
```go
type People struct{}

func (p *People) ShowA() {
    fmt.Println("showA")
    p.ShowB()
}
func (p *People) ShowB() {
    fmt.Println("showB")
}

type Teacher struct {
    People
}

func (t *Teacher) ShowB() {
    fmt.Println("teacher showB")
}

func main() {
    t := Teacher{}
    t.ShowA()
}

```
考点：**go的组合继承**
解答：
这是Golang的组合模式，可以实现OOP的继承。
被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法(ShowA())调用时接受者并没有发生变化。
此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。

```go
showA
showB
```

### 5. 下面代码会触发异常吗？请详细说明
```go
func main() {
    runtime.GOMAXPROCS(1)
    int_chan := make(chan int, 1)
    string_chan := make(chan string, 1)
    int_chan <- 1
    string_chan <- "hello"
    select {
    case value := <-int_chan:
        fmt.Println(value)
    case value := <-string_chan:
        panic(value)
    }
}

```
考点：**select随机性**
解答：
select会随机选择一个可用通用做收发操作。
所以代码是有肯触发异常，也有可能不会。
单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则：
* select 中只要有一个case能return，则立刻执行。
* 当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。
* 如果没有一个case能return则可以执行”default”块。

### 6. 下面代码输出什么？
```go
func calc(index string, a, b int) int {
    ret := a + b
    fmt.Println(index, a, b, ret)
    return ret
}

func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("10", a, b))
    a = 0
    defer calc("2", a, calc("20", a, b))
    b = 1
}

```

考点：**defer执行顺序**
解答：
这道题类似第1题
需要注意到defer执行顺序和值传递
index:1肯定是最后执行的，但是index:1的第三个参数是一个函数，所以最先被调用calc("10",1,2)==>10,1,2,3
执行index:2时,与之前一样，需要先调用calc("20",0,2)==>20,0,2,2
执行到b=1时候开始调用，index:2==>calc("2",0,2)==>2,0,2,2
最后执行index:1==>calc("1",1,3)==>1,1,3,4
```go
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4
```

### 7. 请写出以下输入内容
```go
func main() {
    s := make([]int, 0)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```
考点：**make默认值和append**
解答：
make初始化是由默认值的哦，此处默认值为0
```go
[0 0 0 0 0 1 2 3]
```
大家试试改为:
```
s := make([]int, 0)
s = append(s, 1, 2, 3)
fmt.Println(s)//[1 2 3]
```

### 8. 下面的代码有什么问题?
```go
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (ua *UserAges) Add(name string, age int) {
	ua.Lock()
	defer ua.Unlock()
	ua.ages[name] = age
}

func (ua *UserAges) Get(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age
	}
	return -1
}
```
考点：**map线程安全**
解答：
可能会出现`fatal error: concurrent map read and map write`.
修改一下看看效果
```go
func (ua *UserAges) Get(name string) int {
    ua.Lock()
    defer ua.Unlock()
    if age, ok := ua.ages[name]; ok {
        return age
    }
    return -1
}
```

### 9. 下面的迭代会有什么问题？
```go
func (set *threadSafeSet) Iter() <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		set.RLock()

		for elem := range set.s {
			ch <- elem
		}

		close(ch)
		set.RUnlock()

	}()
	return ch
}
```
考点：**chan缓存池**
解答：
看到这道题，我也在猜想出题者的意图在哪里。
chan?sync.RWMutex?go?chan缓存池?迭代?
所以只能再读一次题目，就从迭代入手看看。
既然是迭代就会要求set.s全部可以遍历一次。但是chan是为缓存的，那就代表这写入一次就会阻塞。
我们把代码恢复为可以运行的方式，看看效果
```go
package main

import (
    "sync"
    "fmt"
)

//下面的迭代会有什么问题？

type threadSafeSet struct {
    sync.RWMutex
    s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
    // ch := make(chan interface{}) // 解除注释看看！
    ch := make(chan interface{},len(set.s))
    go func() {
        set.RLock()

        for elem,value := range set.s {
            ch <- elem
            println("Iter:",elem,value)
        }

        close(ch)
        set.RUnlock()

    }()
    return ch
}

func main()  {

    th:=threadSafeSet{
        s:[]interface{}{"1","2"},
    }
    v:=<-th.Iter()
    fmt.Sprintf("%s%v","ch",v)
}

```

### 10. 以下代码能编译过去吗？为什么？
```go
package main

import (
	"fmt"
)

type People interface {
	Speak(string) string
}

type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
}
```
考点：**golang的方法集**
解答：
编译不通过！
做错了！？说明你对golang的方法集还有一些疑问。
一句话：golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。


### 11. 以下代码打印出来什么内容，说出为什么。
```go
package main

import (
	"fmt"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func main() {
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}
}

```
考点：**interface内部结构**
解答：
很经典的题！
这个考点是很多人忽略的interface内部结构。
go中的接口分为两种一种是空的接口类似这样：
```
var in interface{}
```
另一种如题目：
```
type People interface {
    Show()
}
```
他们的底层结构如下：
```
type eface struct {      //空接口
    _type *_type         //类型信息
    data  unsafe.Pointer //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type iface struct {      //带有方法的接口
    tab  *itab           //存储type信息还有结构实现方法的集合
    data unsafe.Pointer  //指向数据的指针(go语言中特殊的指针类型unsafe.Pointer类似于c语言中的void*)
}
type _type struct {
    size       uintptr  //类型大小
    ptrdata    uintptr  //前缀持有所有指针的内存大小
    hash       uint32   //数据hash值
    tflag      tflag
    align      uint8    //对齐
    fieldalign uint8    //嵌入结构体时的对齐
    kind       uint8    //kind 有些枚举值kind等于0是无效的
    alg        *typeAlg //函数指针数组，类型实现的所有方法
    gcdata    *byte
    str       nameOff
    ptrToThis typeOff
}
type itab struct {
    inter  *interfacetype  //接口类型
    _type  *_type          //结构类型
    link   *itab
    bad    int32
    inhash int32
    fun    [1]uintptr      //可变大小 方法集合
}
```
可以看出iface比eface 中间多了一层itab结构。
itab 存储_type信息和[]fun方法集，从上面的结构我们就可得出，因为data指向了nil 并不代表interface 是nil，
所以返回值并不为空，这里的fun(方法集)定义了接口的接收规则，在编译的过程中需要验证是否实现接口
结果：
```go
BBBBBBB
```