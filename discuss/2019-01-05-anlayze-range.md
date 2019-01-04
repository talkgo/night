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

看上去for循环的次数,在进入循环体前已经确定了,且次数为range后len(sli)

实际上,range为golang的语法糖,其实际执行相当于

```
func main()  {
	sli := []int{6,7,8}
	len_sli := len(sli)
	for index := 0 ; index < len_sli; index++ {
		sli = append(sli,666)
		fmt.Println(index)
	}
}
```

**即在进入循环之前,控制循环次数的这个len_sli参数已经确定;**

**在循环体内对原切片进行append操作,并不会影响len_sli的值**



## 2. 写出以下代码的输出:

![](../images/2019-01-05-range3.png)








