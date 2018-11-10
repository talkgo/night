## 2018-11-09

来源：《Go 夜读》微信群

时间：2018-11-09

---

## 强制使用字段命名方式初始化结构体

在定义结构体时,加入一个非导出且大小为0的字段, 编译器会强制开发者使用字段命名的方式初始化结构体, 而不能按顺序来赋值.

`这个小技巧利用了包的非导出字段可见性, 所以只能在不同包下初始化结构体才有用.`

``` go
// foo.go
package foo

type Config struct {
	_    [0]int
	Name string
	Size int
}
```

```go
// main.go
package main

import "foo"

func main() {
    //_ = foo.Config{[0]int{}, "bar", 123}
    // doesn't compile
    // 报错信息:implicit assignment of unexported field 'flag' in foo.Config literal

	_ = foo.Config{Name: "bar", Size: 123} // compile okay
}
```

>不要把非导出且大小为0的字段放在结构体的最后面, 可能会导致多分配内存.

## 参考资料

1. [How to force package users to use struct composite literals with field names?](https://go101.org/article/tips.html#force-to-use-keyed-struct-literals)
2. [Why does the final field of a zero-sized type in a struct contribute to the size of the struct sometimes?](https://go101.org/article/unofficial-faq.html#final-zero-size-field)
3. [golang 内存分析之字节对齐规则](https://my.oschina.net/u/2950272/blog/1829197)