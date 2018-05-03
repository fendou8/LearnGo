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

	/*
	   值类型：
	   	1.变量直接存储值，内存通常在栈中分配。
	   	2.基本数据类型int、float、bool、string以及数组和struct(结构体)
	   引用类型：
	   	1.变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配。通过GC回收
	   	2.指针、slice、map、chan等都是引用类型
	*/
	fmt.Println("----------------------------------------")
	a := 5 //值类型
	fmt.Println(a)
	modify(a)
	fmt.Println(a)
	modify1(&a) //传递的是地址
	fmt.Println(a)

	b := make(chan int, 1) //引用类型
	fmt.Println(b)
	fmt.Println("----------------------------------------")

	/*
		   变量的作用域
		   1.在函数内部声明的变量叫做局部变量，生命周期仅限于函数内部
		   2.在函数外部声明的变量叫做全部变量，生命周期作用于整个包，
			   如果是大写的，则作用于整个程序
	*/

}
func modify(a int) {
	a = 10
	return
}
func modify1(a *int) {
	*a = 100
	return
}
