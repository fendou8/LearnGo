package main

import (
	"fmt"
)

func main() {
	//bool类型
	//变量声明,没有初始化，初始值为false
	var a bool
	fmt.Println("a = ", a)
	a = true
	fmt.Println("a = ", a)
	//自动推导
	var b = false
	fmt.Println("b = ", b)
	c := true
	fmt.Println("c = ", c)
}
