package main

import (
	"fmt"
)

func main() {
	//complex类型
	//变量声明,没有初始化，初始值为0+0i
	var c complex64
	fmt.Println("c = ", c)
	c = 2.1 + 3.14i
	fmt.Println("c = ", c)
	/*
		自动推导,默认类型为 complex128
	*/
	c2 := 2.1 + 3.14i
	fmt.Printf("c2 type is %T\n", c2)

	//通过内建函数取实部和虚部
	fmt.Println("real(c2) = ", real(c2), "imag(c2) = ", imag(c2))

}
