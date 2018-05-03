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

}
