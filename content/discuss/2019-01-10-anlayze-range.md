---
title: 2019-01-10 关于range的一些不易察觉的"坑"
date: 2019-01-10T00:00:00+08:00
---

来源: Wechat discuss


### 1. 以下代码最终的结果是什么?

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
输出结果为:
```
	0
	1
	2
```

看上去 for 循环的次数,在进入循环体前已经确定了,且次数为 range 后 len(sli)

实际上,range 为 golang 的语法糖,其实际执行相当于

```go
func main()  {
	sli := []int{6,7,8}
	len_sli := len(sli)
	for index := 0 ; index < len_sli; index++ {
		sli = append(sli,666)
		fmt.Println(index)
	}
}
```

**即在进入循环之前,控制循环次数的这个 len_sli 参数已经确定;**

**在循环体内对原切片进行 append 操作,并不会影响 len_sli 的值**



### 2. 写出以下代码的输出:

```go
package main

const N  = 3

func main(){
	m := make(map[int]*int)
	
	for i := 0; i < N; i++ {
		m[i] = &i
	}
	
	for _, v := range m {
		print(*v)
	}
}

```

初步分析:
- 在第一个循环中为m赋值,键名0,1,2分别对应着键值0,1,2的内存地址
- 在第二个循环中迭代m,每次循环用*取出指针键值对应内存地址里存的值
- 初步分析,结果应该是0,1,2

运行结果为:

```
	3
	3
	3
```

结果却是3,3,3

- 我们加一下注释代码再来看:



```
func main()  {
   	m := make(map[int]*int)
   
   	for i := 0; i < 3; i++{
   		m[i] = &i //A
   		fmt.Println("&i的值是:",&i)
   		fmt.Println("i的值是:",i)
   	}
   	for c,v := range m {
   		fmt.Println(c)
   		time.Sleep(1e9)
   		fmt.Println(*v)
   		time.Sleep(1e9)
   	}
}
```

结果如下:
```
&i的值是: 0xc420016468
i的值是: 0
&i的值是: 0xc420016468
i的值是: 0
&i的值是: 0xc420016468
i的值是: 0
0
3
1
3
2
3
```




即在迭代中m的三个元素的指针相同,都指向了最后一个迭代对象的地址,在此即3的值

- 如果在迭代体中需要访问数组/map元素的指针，那么务必小心.这类 bug 无形极难轻易寻获


- 改进办法:引入中间变量,如下:

```go
func main(){
	m := make(map[int]*int)
	
	for i := 0; i < 3; i++ {
		x := i
		fmt.Println(x)
		fmt.Println(&x)
		m[i] = &x
		fmt.Println("&i的值是:",&i)
		fmt.Println("i的值是:",i)
	}
	
	for c,v := range m {
		
		fmt.Println(c)
		time.Sleep(1e9)
		fmt.Println(*v)
		time.Sleep(1e9)
	}
}

```

输出为:


```
0
0xc420016470
&i的值是: 0xc420016468
i的值是: 0
1
0xc420016490
&i的值是: 0xc420016468
i的值是: 1
2
0xc4200164a8
&i的值是: 0xc420016468
i的值是: 2
0
0
1
1
2
2
```

- 即此处v其实是一个全局变量,只分配了一次内存地址


## 总结:
关于range,有两点需要注意:
-   一个是长度在循环之前就已经确定
-   另一个是迭代出的值是全局变量~

