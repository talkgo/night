## 2018-07-09

来源：《Go 夜读》微信群

时间：2018-07-09

----

## 问题1. golang 怎么建立一个一一映射，现在的映射 map 是可以多对一的是吧？

建立两个 map 互查。

segment tree？

## 问题2. golang 中 new 和 make 关键字的区别和使用场景？

首先，我们看一下 Go 标准包的 builtin 源码吧（*go/src/builtin/builtin.go*）

```go
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length. For example, make([]int, 0, 10) allocates an underlying array
//	of size 10 and returns a slice of length 0 and capacity 10 that is
//	backed by this underlying array.
//	Map: An empty map is allocated with enough space to hold the
//	specified number of elements. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
func make(t Type, size ...IntegerType) Type

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
```

`new` 它接受一个参数，这个参数是一个类型，不是一个值，分配好内存后，返回一个指向该类型内存地址的指针，同时请注意它同时把分配的内存置为零，也就是类型的零值。

new 不常用：new 在一些需要实例化接口的地方用的比较多，但是可以用 &A{} 替代。

但是 new 和 &A{} 也是有差别的，主要差别在于 &A{} 显示执行堆分配。

make 也是用于内存分配的，但是和new不同，它只用于channel,map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

>注意：这三种类型是引用类型，所以必须得初始化，但是不是置为零值，这个跟new是不一样的。

举例说明：

```go
package main

import "fmt"

func main() {
    p := new([]int) //p == nil; with len and cap 0
    fmt.Println(p)

    v := make([]int, 10, 50) // v is initialed with len 10, cap 50
    fmt.Println(v)

    /*********Output****************
        &[]
        [0 0 0 0 0 0 0 0 0 0]
    *********************************/

    (*p)[0] = 18        // panic: runtime error: index out of range
                        // because p is a nil pointer, with len and cap 0
    v[1] = 18           // ok
    
}

// 作者：iCaptain
// 链接：https://www.jianshu.com/p/c173dab0e71c
// 來源：简书
// 简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
```

go 的逃逸分析决定了是分配到堆上还是栈上。

----

## 参考

* [Effective Go#allocation_new](https://golang.org/doc/effective_go.html#allocation_new)
* [Effective Go#allocation_make](https://golang.org/doc/effective_go.html#allocation_make)
* [Go 语言机制之逃逸分析（Language Mechanics On Escape Analysis）](https://studygolang.com/articles/12444)
* [Go 的变量到底在堆还是栈中分配](http://www.zenlife.tk/go-allocated-on-heap-or-stack.md)
* [Golang 变量逃逸分析小探](http://reusee.github.io/post/escape_analysis/)
