package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	c := 30
	//一段一段处理，自动换行
	fmt.Println("a=", a)
	//格式化输出，将a的内容放在%d的位置
	fmt.Printf("a = %d\n", a)
	fmt.Printf("a = %d, b = %d, c = %d", a, b, c)
}
