## 2018-05-22

来源：《Go 夜读》微信群

时间：2018-05-22

----

## 背景
去面试的时候遇到一道和 string 相关的题目，记录一下用到的知识点。题目如下：
```go
s:="123"
ps:=&s
b:=[]byte(s)
pb:=&b

s+="4"
*ps+="5"
(*pb)[1] = 0
(*pb)[2] = 4

fmt.Printf("%+v\n",*ps)
fmt.Printf("%+v\n",*pb)
```
> 问以上代码的输出是什么。

## 分析
很容易可以看出 s 和 ps 代表同一个 string，b 和 pb 代表同一个 byte 的切片。关键在于
```go
b:=[]byte(s)
```
根据 [The Go Programming Language](http://www.gopl.io/) 的解释：
> A string contains an array of bytes that, once created, is immutable. By contrast, the elements of a byte slice can be freely modified.  
Strings can be converted to byte slices and back again: 
>
> s := “abc”  
b := []byte(s)  
s2 := string(b)  
>
> Conceptually, the []byte(s) conversion allocates a new byte array holding a copy of the bytes of s, and yields a slice that references the entirety of that array. An optimizing compiler may be able to avoid the allocation and copying in some cases, but in general copying is required to ensure that the bytes of s remain unchanged even if those of b are subsequently modified. The conversion from byte slice back to string with string(b) also makes a copy, to ensure immutability of the resulting string s2.

因为 string 是不可变的，所以不管是从 string 转到 []byte 还是从 []byte 转换到 string 都会发生一次复制。因此 p 和 b 可以看作两个内容相同的两个对象，对 p，b各自的修改不会影响对方。  
先看 ps，经过两次拼接后就是 "12345"。  
再看 pb，因为 s 中的内容都是 ASCII 字符，在 b 中只需要用一个 byte 就可以表示一个字符，所以 b 的实际内容是[49,50,51]，分别对应1，2，3的 ASCII 编码。经过就修改后就变成了[49,0,4]。  

## 答案
经过上面的分析，最后输出的答案是"12345"和[49,0,4]

## 延伸
题目中对字符串的拼接也是常用的场景。因为 string 是不可变的，所以在拼接字符串时实际上也是将源字符串复制了一次，所以在 string 比较大时会消耗不少的内存和时间。关于字符串拼接的各种方法这里不详细展开了，有兴趣的可以参考以下几个链接：  
https://gocn.io/question/265  
https://gocn.io/article/704

## Reference
https://blog.golang.org/strings  
https://sheepbao.github.io/post/golang_byte_slice_and_string/  
http://nanxiao.me/golang-string-byte-slice-conversion/

