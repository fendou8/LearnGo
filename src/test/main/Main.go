package main

import (
	"fmt"
	_ "test/list"   //应用该包，只进行初始化，不调用时，使用“_”
	student "test2" //包别名，设置了别名就必须通过别名调用
)

//
//func list(n int) {
//	for i := 0; i <= n; i++ {
//		fmt.Printf("%d+%d=%d\n", i, n-i, n)
//	}
//}

func main() {
	fmt.Println("----------main begin----------")
	//	fmt.Println("请输入：")
	//	fmt.Scan(&n)
	//	list(n) //调用本文件或同一个包源文件中的list函数
	/*
		1.不同包中函数，通过包名/别名+点+函数名进行调用
	*/
	//	list.Add(n)

	student.SetInfo("张三", 18)

	sum := add(10, 20, 30, 40)
	fmt.Println(sum)
	str := addString("ab", "cd", "add")
	fmt.Println(str)
}

/*
注意：其中arg是一个slice，我们可以通过arg[index]依次访问所有参数，
通过len(arg)来判断传递参数的个数
*/
//0个或多个参数
//func add(arg ...int) {}

//1个或多个参数
func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

//1个或多个参数
func addString(a string, arg ...string) string {
	var result string = a
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return result
}

//2个或多个参数
//func add(a int, b int, arg ...int) {}
