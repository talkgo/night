package main

import "fmt"

func main() {
	a := make([]int64, 0, 32)
	fmt.Println(cap(a), len(a))

	for i := 0; i < 30; i++ {
		a = append(a, 1)
		fmt.Println(cap(a), len(a))
	}

}
