# Golang面试题解析（二）

## 12.是否可以编译通过？如果通过，输出什么？
```go  
func main() {
	i := GetValue()

	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}

```
### 解析
考点：**type**

编译失败，因为type只能使用在interface


## 13.下面函数有什么问题？
```go  
func funcMui(x,y int)(sum int,error){
    return x+y,nil
}
```
### 解析
考点：**函数返回值命名**
在函数有多个返回值时，只要有一个返回值有指定命名，其他的也必须有命名。
如果返回值有有多个返回值必须加上括号；
如果只有一个返回值并且有命名也需要加上括号；
此处函数第一个返回值有sum名称，第二个为命名，所以错误。

## 14.是否可以编译通过？如果通过，输出什么？
```go  
package main

func main() {

	println(DeferFunc1(1))
	println(DeferFunc2(1))
	println(DeferFunc3(1))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
```
### 解析
考点:**defer和函数返回值**
需要明确一点是defer需要在函数结束前执行。
函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4;
DeferFunc2函数中t的作用域为函数，返回1;
DeferFunc3返回3

## 15.是否可以编译通过？如果通过，输出什么？
```go  
func main() {
	list := new([]int)
	list = append(list, 1)
	fmt.Println(list)
}
```
### 解析
考点：**new**
list:=make([]int,0)

## 16.是否可以编译通过？如果通过，输出什么？
```go  
package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}
```
### 解析
考点：**append**
append切片时候别漏了'...'

## 17.是否可以编译通过？如果通过，输出什么？
```go  
func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
}
```
### 解析
考点:**结构体比较**
进行结构体比较时候，只有相同类型的结构体才可以比较，结构体是否相同不但与属性类型个数有关，还与属性顺序相关。
```
sn3:= struct {
    name string
    age  int
}{age:11,name:"qq"}
```
sn3与sn1就不是相同的结构体了，不能比较。
还有一点需要注意的是结构体是相同的，但是结构体属性中有不可以比较的类型，如map,slice。
如果该结构属性都是可以比较的，那么就可以使用“==”进行比较操作。

可以使用reflect.DeepEqual进行比较
```
if reflect.DeepEqual(sn1, sm) {
    fmt.Println("sn1 ==sm")
}else {
    fmt.Println("sn1 !=sm")
}
```
所以编译不通过： invalid operation: sm1 == sm2



## 18.是否可以编译通过？如果通过，输出什么？
```go
func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
func main() {
	var x *int = nil
	Foo(x)
}
```
### 解析
考点：**interface内部结构**
```
non-empty interface
```

## 19.是否可以编译通过？如果通过，输出什么？
```go  
func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return nil, false
}
func main()  {
	intmap:=map[int]string{
		1:"a",
		2:"bb",
		3:"ccc",
	}

	v,err:=GetValue(intmap,3)
	fmt.Println(v,err)
}
```
### 解析
考点：**函数返回值类型**
nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。但是如果不特别指定的话，Go 语言不能识别类型，所以会报错。通常编译的时候不会报错，但是运行是时候会报:`cannot use nil as type string in return argument`.

## 20.是否可以编译通过？如果通过，输出什么？
```go 
const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)

func main()  {
	fmt.Println(x,y,z,k,p)
}
```
### 解析
考点：**iota**
结果:
```
0 1 zz zz 4
```



