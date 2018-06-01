## 2018-05-17 线下活动

>参与人数: 12 人

*Go 标准包阅读*

Go 版本：go 1.10.1

### strings

- strings.go

### 问题清单

以下是我们在阅读过程中的一些问题，希望可以引起大家的关注，也欢迎大家提出自己的理解，最好可以给以文章总结。

0. // Remove if golang.org/issue/6714 is fixed
1. bp := copy(b, a[0])
2. return len(s) >= len(prefix) && s[0:len(prefix)] == prefix 各种开发语言都有的短路机制；字符串底层也是可以用作切片的；
3. 为什么要判断这个错误：		if c == utf8.RuneError
4. c -= 'a' - 'A' （小写转大写的算法）
5. // Since we cannot return an error on overflow,
	// we should panic if the repeat will generate
	// an overflow.
	// See Issue golang.org/issue/16237
6. truth
7. asciiSet （bitset 标记位，存在标记为1）

传一个字符串，把字符串包含的ascii，对应的256位，进行映射。

8. Unicode 包很多都看不懂。

9. func isSeparator(r rune) bool

## 延伸阅读

1. 