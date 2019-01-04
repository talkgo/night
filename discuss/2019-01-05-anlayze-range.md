## 1. 以下代码最终的结果是什么?

```go
func main(){
	sli := []int{6,7,8}
	for i := range sli {
		sli = append(sli,666)
		fmt.Println(i)
	}
}

// Output:

```

这段代码会形成死循环吗?

![](../images/2018-05-19-wechat-discuss-bitset.jpeg)

![import_cycle_not_allowed](../images/import_cycle_not_allowed.jpeg)





