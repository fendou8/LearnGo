package main

import (
	"fmt"
)

func main() {
	fmt.Println("类型转换")
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	//bool类型不能转换成int
	fmt.Printf("flag = %d\n", flag)
	//整型也不能转换成bool
	//0是假，非零为真
	//	flag = 0
	//不能转换的类型，叫做不兼容类型
	fmt.Printf("flag = %d\n", flag)

	var ch byte
	ch = 'a'
	var t int
	//类型转换，把ch的值取出来后，转成int后再给t赋值
	t = int(ch)
	fmt.Println("t = ", t)

}
