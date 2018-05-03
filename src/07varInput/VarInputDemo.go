package main

import (
	"fmt"
)

func main() {
	var a int //声明变量
	fmt.Println("请输入：")
	//阻塞等待用户输入
	fmt.Scanf("%d", &a) //别忘了&
	fmt.Printf("a=%d\n", a)

	fmt.Scan(&a) //别忘了&
	fmt.Println("a=", a)
}
