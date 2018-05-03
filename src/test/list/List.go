package list

import (
	"fmt"
)

/*
1.大写意味着这个函数/变量是可以导出的，包外部可以访问
2.小写意味着这个函数/变量是私有的，包外部不能够访问
*/
var A int = 6
var b int = 7

func init() {
	fmt.Println("List init")
}

func Add(n int) {
	for i := 0; i <= n; i++ {
		fmt.Printf("aa %d+%d=%d\n", i, n-i, n)
	}
}
