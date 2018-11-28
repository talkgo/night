package main

import (
	"fmt"
	"testing"
)

// TestLogColor 测试日志颜色
func TestLogColor(t *testing.T) {
	fmt.Println("")

	// | 前景 | 背景 | 颜色 |
	// |----|----|----|
	// | 30 | 40 | 黑色 |
	// | 31 | 41 | 红色 |
	// | 33 | 42 | 绿色 |
	// | 33 | 43 | 黄色 |
	// | 34 | 44 | 蓝色 |
	// | 35 | 45 | 紫红色 |
	// | 36 | 46 | 青蓝色 |
	// | 37 | 47 | 白色 |
	//
	// | 代码 | 意义 |
	// |----|----|
	// | 0 | 终端默认设置 |
	// | 1 | 高亮显示 |
	// | 4 | 使用下划线 |
	// | 5 | 闪烁 |
	// | 7 | 反白显示 |
	// | 8 | 不可见 |

	for b := 40; b <= 47; b++ {
		for f := 30; f <= 37; f++ {
			for d := range []int{0, 1, 4, 5, 7, 8} {
				fmt.Printf(" %c[%d;%d;%dm%s(f=%d,b=%d,d=%d)%c[0m ", 0x1B, d, b, f, "", f, b, d, 0x1B)
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
	// 其中 0x1B 是标记，[开始定义颜色，1代表高亮，40代表黑色背景，32代表绿色前景色，0代表恢复默认颜色。
	fmt.Printf("%c[1;40;32m%s%c[0m", 0x1B, "testPrintColor", 0x1B)
	fmt.Println()
	fmt.Printf("%s\n", "testPrintColor")
}
