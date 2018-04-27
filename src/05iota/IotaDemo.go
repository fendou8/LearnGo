package main

import (
	"fmt"
)

func main() {
	/*
		iota枚举
		常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量
			1.iota常量自动生成器，每隔一行自动累加1
			2.iota用来给常量赋值
			3.iota遇到const，重置为0
			4.可以只写一个iota
			5.如果是同一行，值都一样
	*/
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Printf("a = %d b = %d c = %d\n", a, b, c)
	//iota遇到const，重置为0
	const d = iota
	fmt.Printf("d=%d", d)
	//可以只写一个iota
	const (
		a1 = iota
		b1
		c1
	)
	fmt.Printf("a1 = %d b1 = %d c1 = %d\n", a1, b1, c1)
	//如果是同一行，值都一样
	const (
		a2         = iota
		b2, b3, b4 = iota, iota, iota
		c2         = iota
	)
	fmt.Printf("a2 = %d b2 = %d b3 = %d b4 = %d c2 = %d\n", a2, b2, b3, b4, c2)
}
