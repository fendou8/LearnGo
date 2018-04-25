/*
1.Go语言以包作为管理单位
2.每个文件必须先声明包
3.程序必须有一个main包(重要)
*/

package main

import (
	"fmt" //导入包
)

//入口函数，一个文件夹中有且只有一个main函数
func main() { //左括号必须和函数名同行
	//将"hello,world"打印到屏幕，Println()会自动换行
	//调用函数，大部分需要导入包
	fmt.Println("hello,world") //go语言语句结尾没有分号

	//单行注释
	/*
		块注释
	*/

	/*
		常用命令：
		1.go build xxx.go 编译go代码，生成一个可执行程序，然后运行可执行程序 xxx
		2.go run xxx.go 不生成可执行程序，直接运行

	*/
}
