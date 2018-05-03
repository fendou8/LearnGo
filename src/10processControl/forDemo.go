package main

import (
	"fmt"
)

func main() {
	/*
	   for 初始化条件；判断条件；条件变化{
	   ......
	   }
	*/
	var sum int = 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	/*
		迭代器
	*/
	str := "abcde"
	for i := 0; i < len(str); i++ { //传统写法
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	fmt.Println("--------------------")
	//迭代打印每个元素，默认返回2个值：一个是元素的位置，一个是元素的本身
	for i, data := range str {
		fmt.Printf("str[%d] = %c\n", i, data)
	}
	fmt.Println("--------------------")
	for i := range str { //第二个返回值默认丢弃
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
	fmt.Println("--------------------")
	for i, _ := range str { //第二个返回值默认丢弃
		fmt.Printf("str[%d] = %c\n", i, str[i])
	}
}
