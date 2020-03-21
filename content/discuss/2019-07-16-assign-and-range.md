---
title: 2019-07-16 := 的由来和 range 遍历
date: 2019-07-16T00:00:00+08:00
---

来源：『Go 夜读』微信群

----

## := 的由来

0.23 BNF 和 EBNF，非终结符和终结符，开始符号及产生式

这都是辅助解释文法的概念。
BNF（Backus-Naur Form），即巴科斯范式，是一种程序设计语言描述工具，是由 John Backus 和 Peter Naur 引入的一种形式化符号，用来描述给定语言的语法，即描述语言的语言，故可以称之为元语言。

-> 定义为（也可用 := 和 ::= 表示）

The symbol is called "becomes" and was introduced with IAL (later called Algol 58) and Algol 60. It is the symbol for assigning a value to a variable. One reads x := y; as "x becomes y".

Using ":=" rather than "=" for assignment is mathematical fastidiousness; to such a viewpoint, "x = x + 1" is nonsensical. Other contemporary languages might have used a left arrow for assignment, but that was not common (as a single character) in many character sets.

Algol 68 further distinguished identification and assignment; INT the answer = 42; says that "the answer" is declared identically equal to 42 (i.e., is a constant value). In INT the answer := 42; "the answer" is declared as a variable and is initially assigned the value 42.

There are other assigning symbols, like +:=, pronounced plus-and-becomes; x +:= y adds y to the current value of x, storing the result in x.

(Spaces have no significance, so can be inserted "into" identifiers rather than having to mess with underscores)

## range 删除元素

```golang
func main() {
	// 删除切片中的某个元素

	// 删除切片中指定元素
	{
		items := []int{1, 2, 4, 2, 3, 0}
		deleteItem := 2
		for i := 0; i < len(items); i++ {
			if items[i] == deleteItem {
				items = append(items[:i], items[i+1:]...)
				i--
			}
		}
		fmt.Println(items) // output: [1, 4, 3, 0]
	}

	// 删除找到的第一个元素
	{
		// 切片比较大的话，还是用普通的 for 循环比较好
		items := []int{1, 2, 4, 2, 3, 0}
		deleteItem := 2
		for i, item := range items {
			// 找到要删除的第一个元素，删除并退出循环
			if item == deleteItem {
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
		fmt.Println(items) // output: [1, 4, 2, 3, 0]
	}
}
```

## 参考资料：

1. [What is := operator?](https://stackoverflow.com/questions/10405820/what-is-the-operator/55894870)
2. [Python 相关语法参考](https://www.python.org/dev/peps/pep-0572)