---
title:  2018-11-08
---
来源: Wechat discuss

## Address_operators

```golang
package main

import (
   "fmt"
   "math"
)

type abser interface {
   Abs() float64
}

type onef float64

func (one onef) Abs() float64{
   return math.Pi
}

func main(){
   var a,b  abser
   a = &onef(3.14)
   aa := onef(3.14)
   b = &aa
   fmt.Println(a.Abs())
   fmt.Println(b.Abs())
}
```

为什么 `a = &onef(3.14)` 会报错，而 `aa := onef(3.14)` 是可以的呢？

For an operand x of type T, the address operation &x generates a pointer of type \*T to x. The operand must be addressable, that is, either a variable, pointer indirection, or slice indexing operation; or a field selector of an addressable struct operand; or an array indexing operation of an addressable array. As an exception to the addressability requirement, x may also be a (possibly parenthesized) composite literal. If the evaluation of x would cause a run-time panic, then the evaluation of &x does too.

>对于类型为 T 的操作数 x，地址操作 ＆x 生成类型为 \*T 到 x 的指针。操作数必须是可寻址的，即，变量，指针间接或切片索引操作;
或可寻址结构操作数的字段选择器;
或者可寻址数组的数组索引操作。
作为可寻址性要求的例外，x也可以是（可能带括号的）复合文字。
如果x的评估会导致运行时恐慌，那么＆x的评估也会发生。

For an operand x of pointer type \*T, the pointer indirection \*x denotes the variable of type T pointed to by x. If x is nil, an attempt to evaluate \*x will cause a run-time panic.

>对于指针类型* T的操作数x，指针间接* x表示由x指向的类型T的变量。如果x为nil，则尝试评估* x将导致运行时出现紧急情况。

## 参考资料

1. [Go 夜读第一期 - cannot take address of temporary variables](https://github.com/developer-learning/night-reading-go/tree/master/reading/20180321#cannot-take-address-of-temporary-variables)
2. https://golang.org/ref/spec#Address_operators