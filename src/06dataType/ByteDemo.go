package main

import (
	"fmt"
)

func main() {
	//byte类型
	//变量声明,没有初始化，初始值为0
	var ch byte
	fmt.Println("ch = ", ch)
	ch = 97
	//格式化输出，%c以字符方式打印，%d以整型方式打印
	fmt.Printf("ch = %c, ch = %d\n", ch, ch)
	/*
		自动推导
		1.默认为int
	*/
	var ch2 = 65
	fmt.Printf("ch2 = %c\n", ch2)
	ch3 := 66
	fmt.Printf("ch3 type is %T ch3 = %c\n", ch3, ch3)

	var ch4 byte
	ch4 = 'c' //字符用单引号
	fmt.Printf("ch4 = %c ch4 = %d\n", ch4, ch4)

}
