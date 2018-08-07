# 设计和算法相关

_1.要设计一个秒杀系统要注意什么？_

![](./images/miaosha.jpg)

> **前端秒杀页面**
>
>> `页面静态化`：将活动页面上的所有可以静态的元素全部静态化，并尽量减少动态
元素。通过CDN来抗峰值。              
>> `禁止重复提交`：用户提交之后按钮置灰，禁止重复提交。         
>> `用户限流`：在某一时间段内只允许用户提交一次请求，比如可以采取IP限流。              
>
> **服务端控制器(网关)**
>
>> `限制uid访问频率`：我们上面拦截了浏览器访问的请求，但针对某些恶意攻击或其它插件，在服务端控制层需要针对同一个访问uid，限制访问频率。                   
>
> **服务层**
>
>> `采用消息队列缓存请求`：既然服务层知道库存只有100台手机，那完全没有必要把100W个请求都传递到数据库啊，那么可以先把这些请求都写到消息队列缓存一下，数据库层订阅消息减库存，减库存成功的请求返回秒杀成功，失败的返回秒杀结束。      
>> `利用缓存应对读请求`：对类似于12306等购票业务，是典型的读多写少业务，大部分请求是查询请求，所以可以利用缓存分担数据库压力。      
>> `利用缓存应对写请求`：缓存也是可以应对写请求的，比如我们就可以把数据库中的库存数据转移到Redis缓存中，所有减库存操作都在Redis中进行，然后再通过后台进程把Redis中的用户秒杀请求同步到数据库中。           
>
> **数据库层**
>
>> 数据库层是最脆弱的一层，一般在应用设计时在上游就需要把请求拦截掉，数据库层只承担“能力范围内”的访问请求。所以，上面通过在服务层引入队列和缓存，让最底层的数据库高枕无忧。       

_2.要设计一个类似微信红包架构系统要注意什么？_

> `南北分区`        
> `快慢分离`       
> `Hash负载均衡`         
> `Cache屏蔽DB`          
> `双维度分库表`              
>
> [查看资料](https://blog.csdn.net/starsliu/article/details/51134473)

_3.如何在一个给定有序数组中找两个和为某个定值的数，要求时间复杂度为O(n),
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

_4.给定一个数组代表股票每天的价格，请问只能买卖一次的情况下，最大化利润是多少？日期不重叠的情况下，可以买卖多次呢？输入：{100,80,120,130,70,60,100,125}，只能买一次：65(60买进，125卖出)；可以买卖多次：115(80买进，130卖出；60买进，125卖出)？_

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

_5.给40亿个不重复的unsigned int的整数，没排过序的，然后再给一个数，如何快速判断这个数是否在那40亿个数当中？_

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