# 算法题

_1.如何在一个给定有序数组中找两个和为某个定值的数，要求时间复杂度为O(n),
比如给｛1，2，4，5，8，11，15｝和15？_

```Golang
func Lookup(meta []int32, target int32) {
	left := 0
	right := len(meta) - 1
	for i := 0; i < len(meta); i++ {
		if meta[left]+meta[right] > target {
			right--
		} else if meta[left]+meta[right] < target {
			left++
		} else {
			fmt.Println(fmt.Sprintf("%d, %d", meta[left], meta[right]))
			return
		}
	}
	fmt.Println("未找到匹配数据")
}
```

_2.给定一个数组代表股票每天的价格，请问只能买卖一次的情况下，最大化利润是多少？日期不重叠的情况下，可以买卖多次呢？输入：{100,80,120,130,70,60,100,125}，只能买一次：65(60买进，125卖出)；可以买卖多次：115(80买进，130卖出；60买进，125卖出)？_

```Golang
func main() {
	a := []int{100, 80, 120, 130, 70, 60, 100, 125}
	// a := []int{68, 0, 1, 67}
	var buyPrice, salePrice = 1<<31 - 1, -1 << 31
	var buyDay, saleDay = -1, -1

	type Op struct {
		BuyDay, SaleDay     int
		BuyPrice, SalePrice int
		Earnings            int
	}
	var opList = []Op{}

	// 遇到的第一个波谷买入，下一个波峰卖出，可以获取最大收益
	for k, todayPrice := range a {
		if buyDay == -1 {
			// 寻找买入点
			if todayPrice < buyPrice {
				buyPrice = todayPrice
				continue
			}
			// 买入
			buyDay = k - 1
			continue
		}
		// 寻找卖出点
		if todayPrice > salePrice {
			salePrice = todayPrice
			if k < len(a)-1 {
				continue
			}
		}
		// 卖出
		if k < len(a)-1 {
			saleDay = k - 1
		} else {
			saleDay = k
		}
		opList = append(opList, Op{
			BuyDay:    buyDay,
			SaleDay:   saleDay,
			BuyPrice:  buyPrice,
			SalePrice: salePrice,
			Earnings:  salePrice - buyPrice,
		})
		// 重复下一轮操作
		buyPrice, salePrice = 1<<31-1, -1<<31
		buyDay, saleDay = -1, -1
	}
	fmt.Printf("%+v", opList)
}
```

_3.给40亿个不重复的unsigned int的整数，没排过序的，然后再给一个数，如何快速判断这个数是否在那40亿个数当中？_

> **在这里只分享思路**
>
```
	首先，40亿个unsigned int的整数，如果放到内存，那就是大约16G的空间，
那么直接放到内存空间进行排序然后二分查找的方式是行不通的。  
```
>
> **方案一**
>
```
	在这里可以考虑使用bitmap，需要4*10^9bit内存， 大约500MB就可一把40
亿的数全部进行hash，时间复杂度是O（n），然后可以在O(1)的时间内进行判断此
数是否在40亿中；此过程在内存中完成。
```
>
> **方案二**
>
```
	考虑在磁盘中操作。
	因为2^32为40亿多，所以给定一个数可能在，也可能不在其中；这里我们把40亿
个数中的每一个用32位的二进制来表示，假设这40亿个数开始放在一个文件中。
	然后将这40亿个数分成两类:
      1.最高位为0
      2.最高位为1
  并将这两类分别写入到两个文件中；
  再然后把这两个文件为又分成两类:
      1.次最高位为0
      2.次最高位为1
  ...
  以此类推就可以找到了,而且时间复杂度为O(logn)
```