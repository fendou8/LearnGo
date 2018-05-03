package main

import (
	"fmt"
)

func main() {
	s := "王"
	//if和{之间就是条件，条件通常是关系运算符
	if s == "王" {
		fmt.Println("1111")
	} else {
		fmt.Println("2222")
	}

	//if支持一个初始化语句，初始化语句和判断条件以分号来分隔
	if a := 12; a == 11 {
		fmt.Println("a==10")
	} else if a > 10 {
		fmt.Println("a>10")
	} else { //else后面没有条件
		fmt.Println("a!=10")
	}
}
