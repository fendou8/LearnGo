package main

import (
	"fmt"
)

func main() {
	//变量：程序运行期间可以改变的量，变量声明需要 var
	//常量：程序运行期间不可以改变的量，常量声明需要const
	const a int = 10
	//	a = 20 //不能够改变
	fmt.Println(a)
	//常量自动推导,不能使用:=
	const b = 20.4
	fmt.Printf("b type is %T b=%f\n", b, b)

	/*
		不同类型变量的声明(定义)
	*/
	//方式1
	var a1 int
	var b1 float64
	a1, b1 = 10, 3.14
	fmt.Printf("a1 = %d b1 = %f\n", a1, b1)
	//方式2
	var (
		a2 int
		b2 float64
	)
	a2, b2 = 10, 3.14
	fmt.Printf("a2 = %d b2 = %f\n", a2, b2)
	//可以定义时赋值
	var (
		a3 int     = 10
		b3 float64 = 3.14
	)
	fmt.Printf("a3 = %d b3 = %f\n", a3, b3)
	//可以自动推导类型
	var (
		a4 = 10
		b4 = 3.14
	)
	fmt.Printf("a4 = %d b4 = %f\n", a4, b4)

	//常量定义时赋值
	const (
		a5 int     = 10
		b5 float64 = 3.14
	)
	fmt.Printf("a5 = %d b5 = %f\n", a5, b5)
	//常量自动推导
	const (
		a6 = 10
		b6 = 3.14
	)
	fmt.Printf("a6 = %d b6 = %f\n", a6, b6)
}
