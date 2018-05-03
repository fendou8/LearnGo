package main

import (
	"fmt"
)

func main() {
	fmt.Println("4>3 :", 4 > 3)
	fmt.Println("4!=3 :", 4 != 3)

	fmt.Println("取反 4!=3 :", !(4 != 3))

	//&& 与，左右操作数都为真才为真
	fmt.Println("true&&true :", true && true)
	//|| 或者，左右操作数都为假才为假
	fmt.Println("false&&false :", false && false)

	//非0就是真，0就是假，但是go语音的bool类型和int不兼容

	/*
		&取地址运算符  &a 取变量a的地址
		*取值运算符      *a 指针变量a所指向内存的值
	*/
}
