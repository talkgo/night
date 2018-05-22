## 2018-05-21

来源：《Go 夜读》微信群

时间：2018-05-21

----

## 1. 如下代码何解？

```go
a := []int{1,2,3}
for _, i:= range a {
	go func() {
		fmt.Println(i)
	}()
}
time.Sleep(time.Second)
// Output:
// 3
// 3
// 3
```

大家可以看参考资料 1, 2，上面有完整的解答。

延伸出来的知识点：惰性求值和闭包。

解答：
根据问题的描述推断出是闭包的特性，通过Google检索关键字 **glang faq loop closure** 可找到官方的[faq](https://golang.org/doc/faq#closures_and_goroutines)，里面有详细的解释，并提醒了用户可以通过 *go vet*命令来检查程序中类似的问题。

最后官方的faq还给出了针对这个问题的两种解决方法。其中第一种方法很好理解，就是把变量v的值通过参数的形式传递给goroutine，因为go中func的参数传递都是值传递，所以就在goroutine启动时获得了当前v变量的值。
```go
for _, v := range values {
    go func(u string) {
        fmt.Println(u)
        done <- true
    }(v)
}
```
第二种方法是更巧妙地通过一个赋值语句来解决的。
```go
for _, v := range values {
    v := v // create a new 'v'.
    go func() {
        fmt.Println(v)
        done <- true
    }()
}
```
如果没有明白其中的原理，只要自己把赋值前后变量v的内存地址打印出来就明白了，在赋值后新的变量v实际上是在内存中开辟了一个新的空间并保存了当前变量v的值，两个变量并不是指向相同的内存地址。
```go
for _, v := range values {
    fmt.Printf("Before assignment:%p\n", &v)
    v := v
    fmt.Printf("After assignment:%p\n", &v)
    go func() {
        fmt.Println(v)
        done <- true
    }()
}
// Output:
// Before assignment:0xc04206c080
// After assignment:0xc04206c090
// Before assignment:0xc04206c080
// After assignment:0xc04206c0a0
// Before assignment:0xc04206c080
// After assignment:0xc04206c0b0
```



## 2. []int 转 []int32 有没有什么好办法？

?

## 3. 大家的公司一般都用什么用于监控？

- zabbix
- grafana+influxdb
- openfalcon

## 参考资料

1. [https://github.com/golang/go/wiki/CommonMistakes](https://github.com/golang/go/wiki/CommonMistakes)
2. [https://golang.org/doc/faq#closures_and_goroutines](https://golang.org/doc/faq#closures_and_goroutines)
