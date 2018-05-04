package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("请输入n:")
	fmt.Scan(&n)
	sum := sumFactorial(n)
	fmt.Printf("%v阶乘之和为%v", n, sum)
}

func sumFactorial(n int) uint64 {
	var s, sum uint64 = 1, 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		fmt.Printf("%d!=%v\n", i, s)
		sum += s
	}
	return sum
}
