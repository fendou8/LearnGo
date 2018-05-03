package main

import (
	"fmt"
)

func main() {
	var m, n int
	fmt.Println("请输入m ,n：")
	fmt.Scanf("%d%d", &m, &n)
	if m > n {
		m, n = n, m
	}
	fmt.Printf("m = %d  n = %d\n", m, n)
	for i := m; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
