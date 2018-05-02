package main

import (
	"fmt"
)

func main() {
	//float类型
	//变量声明,没有初始化，初始值为0.0
	var a float32
	fmt.Println("a = ", a)
	a = 3.14
	fmt.Println("a = ", a)
	/*
		自动推导
		1.默认为float64,
		2.float64存储小数比float32更准确
	*/
	var b = 3.14
	fmt.Println("b = ", b)
	c := 3.4
	fmt.Printf("c type is %T c = %f\n", c, c)
}
