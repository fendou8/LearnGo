package main

import (
	"fmt"
)

func main() {
	//string类型
	//变量声明,没有初始化，初始值为空字符串，占两个字节
	var str string
	fmt.Println("str = ---", str, "---")
	str = "abc" //字符串用双引号
	fmt.Println("str = ", str)
	/*
		自动推导
	*/
	var str2 = "abc"
	fmt.Println("str2 = ", str2)
	str3 := "abcd"
	fmt.Printf("str3 type is %T str3 = %s\n", str3, str3)
	//输出字符串有几个字符
	fmt.Println("len(str3) = ", len(str3))

	/*
		字符串和字符区别
		字符：
		1.单引号
		2.字符元素往往只有一个字符，转义字符除外'\n'
		字符串：
		1.双引号
		2.字符串有一个或多个字符组成
		3.字符串都是隐藏了一个结束符‘\0’
	*/

	//输出字符串中指定字符
	fmt.Printf("str3[0] = %c , str3[1] = %c \n", str3[0], str3[1])
}
