# Golang面试题解析（三）

# 21.编译执行下面代码会出现什么?
```go  
package main
var(
    size :=1024
    max_size = size*2
)
func main()  {
    println(size,max_size)
}
```
## 解析
考点:**变量简短模式**  
变量简短模式限制：
- 定义变量同时显式初始化
- 不能提供数据类型
- 只能在函数内部使用

结果：
```
syntax error: unexpected :=
```  

# 22.下面函数有什么问题？
```
package main
const cl  = 100

var bl    = 123

func main()  {
    println(&bl,bl)
    println(&cl,cl)
}

```   
## 解析  
考点:**常量**  
常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，作为指令数据使用，

```
cannot take the address of cl
```

# 23.编译执行下面代码会出现什么?
```
package main

func main()  {

    for i:=0;i<10 ;i++  {
    loop:
        println(i)
    }
    goto loop
}
```
## 解析   
考点：**goto**  
goto不能跳转到其他函数或者内层代码
```
goto loop jumps into block starting at
```  
# 24.编译执行下面代码会出现什么?
```
package main
import "fmt"

func main()  {
    type MyInt1 int
    type MyInt2 = int
    var i int =9
    var i1 MyInt1 = i
    var i2 MyInt2 = i
    fmt.Println(i1,i2)
}
```  
## 解析   
考点：**Go 1.9 新特性 Type Alias **  
基于一个类型创建一个新类型，称之为defintion；基于一个类型创建一个别名，称之为alias。
MyInt1为称之为defintion，虽然底层类型为int类型，但是不能直接赋值，需要强转；
MyInt2称之为alias，可以直接赋值。

结果:
```
cannot use i (type int) as type MyInt1 in assignment
```

# 25.编译执行下面代码会出现什么?
```
package main
import "fmt"

type User struct {
}
type MyUser1 User
type MyUser2 = User
func (i MyUser1) m1(){
    fmt.Println("MyUser1.m1")
}
func (i User) m2(){
    fmt.Println("User.m2")
}

func main() {
    var i1 MyUser1
    var i2 MyUser2
    i1.m1()
    i2.m2()
}
```
## 解析   
考点：**Go 1.9 新特性 Type Alias **  
因为MyUser2完全等价于User，所以具有其所有的方法，并且其中一个新增了方法，另外一个也会有。
但是
```
i1.m2()
```
是不能执行的，因为MyUser1没有定义该方法。
结果:
```
MyUser1.m1
User.m2
```

# 26.编译执行下面代码会出现什么?
```
package main

import "fmt"

type T1 struct {
}
func (t T1) m1(){
    fmt.Println("T1.m1")
}
type T2 = T1
type MyStruct struct {
    T1
    T2
}
func main() {
    my:=MyStruct{}
    my.m1()
}
```
## 解析  
考点：**Go 1.9 新特性 Type Alias **  
是不能正常编译的,异常：
```
ambiguous selector my.m1
```
结果不限于方法，字段也也一样；也不限于type alias，type defintion也是一样的，只要有重复的方法、字段，就会有这种提示，因为不知道该选择哪个。
改为:
```
my.T1.m1()
my.T2.m1()
```
type alias的定义，本质上是一样的类型，只是起了一个别名，源类型怎么用，别名类型也怎么用，保留源类型的所有方法、字段等。

# 27.编译执行下面代码会出现什么?
```
package main

import (
    "errors"
    "fmt"
)

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
    if reallyDoIt {
        result, err := tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}

func tryTheThing() (string,error)  {
    return "",ErrDidNotWork
}

func main() {
    fmt.Println(DoTheThing(true))
    fmt.Println(DoTheThing(false))
}
```
## 解析  
考点：**变量作用域**    
因为 if 语句块内的 err 变量会遮罩函数作用域内的 err 变量，结果：
```
<nil>
<nil>
```
改为：
```
func DoTheThing(reallyDoIt bool) (err error) {
    var result string
    if reallyDoIt {
        result, err = tryTheThing()
        if err != nil || result != "it worked" {
            err = ErrDidNotWork
        }
    }
    return err
}
```
# 28.编译执行下面代码会出现什么?
```
package main

func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        funs = append(funs, func() {
            println(&i,i)
        })
    }
    return funs
}

func main(){
    funs:=test()
    for _,f:=range funs{
        f()
    }
}
```
## 解析  
考点：**闭包延迟求值**  
for循环复用局部变量i，每一次放入匿名函数的应用都是想一个变量。
结果：
```
0xc042046000 2
0xc042046000 2
```
如果想不一样可以改为：
```
func test() []func()  {
    var funs []func()
    for i:=0;i<2 ;i++  {
        x:=i
        funs = append(funs, func() {
            println(&x,x)
        })
    }
    return funs
}
```  

# 29.编译执行下面代码会出现什么?
```
package main

func test(x int) (func(),func())  {
    return func() {
        println(x)
        x+=10
    }, func() {
        println(x)
    }
}

func main()  {
    a,b:=test(100)
    a()
    b()
}

```
## 解析  
考点：**闭包引用相同变量***  
结果：
```
100
110
```

# 30.编译执行下面代码会出现什么?
```
package main

import (
    "fmt"
    "reflect"
)

func main1()  {
    defer func() {
       if err:=recover();err!=nil{
           fmt.Println(err)
       }else {
           fmt.Println("fatal")
       }
    }()

    defer func() {
        panic("defer panic")
    }()
    panic("panic")
}

func main()  {
    defer func() {
        if err:=recover();err!=nil{
            fmt.Println("++++")
            f:=err.(func()string)
            fmt.Println(err,f(),reflect.TypeOf(err).Kind().String())
        }else {
            fmt.Println("fatal")
        }
    }()

    defer func() {
        panic(func() string {
            return  "defer panic"
        })
    }()
    panic("panic")
}
```
## 解析  
考点：**panic仅有最后一个可以被revover捕获**  
触发`panic("panic")`后顺序执行defer，但是defer中还有一个panic，所以覆盖了之前的`panic("panic")`
```
defer panic
```