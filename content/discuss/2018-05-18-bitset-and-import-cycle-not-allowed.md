---
title:  2018-05-18 包循环依赖如何解决
source: 《Go 夜读》微信群
data: 2018-05-18
---
## 1. 包循环依赖如何解决？

![import_cycle_not_allowed](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/import_cycle_not_allowed.jpeg)

方法：

用 `govendor list -v` 可以查看一个包被哪些包依赖。

详解：

使用 `govendor list -v` 可以查看一个包被哪些包依赖：

那么反过来，你可能想知道一个包依赖了哪些包？这个是 go 工具链里面提供的方法，直接使用 `go list`，比如：`go list -f '{{ .Imports }}' github.com/developer-learning/night-go`

## 2. bitset

strings 包之 bitset，一个实现：[https://github.com/henrylee2cn/goutil/blob/master/bitset/bitset.go](https://github.com/henrylee2cn/goutil/blob/master/bitset/bitset.go)

![](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/2018-05-19-wechat-discuss-bitset.jpeg)

对于以上代码的一些讨论：

```
Q：1. 没有看到 unset，或者 clear 清除标志，2. 其次位的设置可以用 atomic 的 or 或操作，无锁编程效率更高。
A：可以加 clear，如果用原子锁就不能用 []byte 类型实现了。

Q：count 实现的太粗暴了，直接就一个一个的遍历，无法利用到 simd 和底层 cpu 指令。
A：count 字段可以用缓存字段，也就是可以用 map，也可以用 `math/bits` 这个库来实现。

Q：[]byte 为什么没办法使用原子锁，地址是有办法获得的。
A：那你知道锁的底层也是原子锁实现的吧？既然还是整个切片上锁，换成原子锁也一样的。
```

其实用库是最后的，这种实现不是最好的，往往 Go 的库能够实现 simd，在 go 1.x，编译期开始了对 simd 的支持。刚好算 byte 中置 1 的数目这个是可以很好的被 simd 的，这个叫 POPCNT ，机器指令，真正的 o(1)。

### simd 是什么？

![simd](https://raw.githubusercontent.com/developer-learning/night-reading-go/master/images/450px-SIMD.svg.png)

>单指令流多数据流（英语：Single Instruction Multiple Data，缩写：SIMD）是一种采用一个控制器来控制多个处理器，同时对一组数据（又称“数据向量”）中的每一个分别执行相同的操作从而实现空间上的并行性的技术。

>在微处理器中，单指令流多数据流技术则是一个控制器控制多个平行的处理微元，例如Intel的MMX或SSE，以及AMD的3D Now!指令集。

>图形处理器（GPU）拥有强大的并行处理能力和可程式流水线，面对单指令流多数据流时，运算能力远超传统CPU。OpenCL和CUDA分別是目前最广泛使用的开源和专利通用图形处理器（GPGPU）预算语言。--摘自[simd-维基百科](https://zh.wikipedia.org/wiki/单指令流多数据流)。

Go 官方库的实现：

```go
// math/bits/bits.go
// 位计数中的分治法
// OnesCount64 returns the number of one bits ("population count") in x.
func OnesCount64(x uint64) int {
	// Implementation: Parallel summing of adjacent bits.
	// See "Hacker's Delight", Chap. 5: Counting Bits.
	// The following pattern shows the general approach:
	//
	// 实现：相邻位的并行求和。
	// 见“黑客的喜悦”（见：算法心得：高效算法的奥秘（中文第2版）），章节.5：计数位。
	// 以下模式显示了一般方法：
	//   x = x>>1&(m0&m) + x&(m0&m)
	//   x = x>>2&(m1&m) + x&(m1&m)
	//   x = x>>4&(m2&m) + x&(m2&m)
	//   x = x>>8&(m3&m) + x&(m3&m)
	//   x = x>>16&(m4&m) + x&(m4&m)
	//   x = x>>32&(m5&m) + x&(m5&m)
	//   return int(x)
	//
	// Masking (& operations) can be left away when there's no
	// danger that a field's sum will carry over into the next
	// field: Since the result cannot be > 64, 8 bits is enough
	// and we can ignore the masks for the shifts by 8 and up.
	// Per "Hacker's Delight", the first line can be simplified
	// more, but it saves at best one instruction, so we leave
	// it alone for clarity.
	// 如果不存在字段总和将传送到下一个字段的危险，则可以将掩码（＆操作）留空：由于结果不能大于64，所以8位就足够了，我们可以忽略8位以上的移位掩码。 根据“Hacker's Delight”，第一行可以简化得更多，但它最多可以节省一条指令，所以为了清晰起见，我们只保留一条。
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}
```

另外一个文档解释的非常清楚：[Efficient_implementation](https://www.wikiwand.com/en/Hamming_weight#/Efficient_implementation)

强烈推荐阅读(只能用于学习查阅，请勿分享传播，如有侵权，请联系我)：

- **[《Hacker's Delit》](../docs/Hacker's-Delight-2nd-Edition.pdf)**
- [算法心得：高效算法的奥秘（中文第2版）](../docs/算法心得：高效算法的奥秘（中文第2版）.pdf)

## 其他

>Wikiwand是一款能够改变维基百科条目界面的软件，2013年由利奥尔·格罗斯曼和依兰·列文创建，2014年8月正式上线。软件界面包含工具栏菜单、导航栏、其他语言版本的个性化链接、新版面和链接条目的预览。内容列表将在左侧不断显示。--摘自[wikiwand 维基百科](https://zh.wikipedia.org/wiki/Wikiwand)

## 参考资料

1. [https://golang.org/pkg/math/bits/](https://golang.org/pkg/math/bits/)
2. [https://golang.org/src/math/bits/bits.go](https://golang.org/src/math/bits/bits.go)
3. [https://github.com/henrylee2cn/goutil#bitset](https://github.com/henrylee2cn/goutil#bitset)
4. [simd - 单指令流多数据流](https://zh.wikipedia.org/wiki/单指令流多数据流)
5. [https://www.wikiwand.com/en/Hamming_weight#/Efficient_implementation](https://www.wikiwand.com/en/Hamming_weight#/Efficient_implementation)
6. [https://zh.wikipedia.org/wiki/Wikiwand](https://zh.wikipedia.org/wiki/Wikiwand)