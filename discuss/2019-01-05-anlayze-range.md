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
面试时遇到的一个问题,
这段代码会形成死循环吗?



回来之后试了一下出乎意料:
![](../images/2019-01-05-range1.jpeg)






