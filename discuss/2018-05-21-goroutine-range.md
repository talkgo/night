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

答复：

- goroutine 的内部实现决定的，可以看一下《 Go 语言读书笔记》
- 这样写，i 是最后一个 a 的值了。

延伸出来的知识点：惰性求值和闭包。

## 2. []int 转 []int32 有没有什么好办法？

?

## 3. 大家的公司一般都用什么用于监控？

- zabbix
- grafana+influxdb
- openfalcon

## 参考资料

1. [https://github.com/golang/go/wiki/CommonMistakes](https://github.com/golang/go/wiki/CommonMistakes)
2. [https://golang.org/doc/faq#closures_and_goroutines](https://golang.org/doc/faq#closures_and_goroutines)
