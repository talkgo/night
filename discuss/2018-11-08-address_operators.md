## 2018-11-08

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

## 参考资料

1. [Go 夜读第一期 - cannot take address of temporary variables](https://github.com/developer-learning/night-reading-go/tree/master/reading/20180321#cannot-take-address-of-temporary-variables)
2. https://golang.org/ref/spec#Address_operators