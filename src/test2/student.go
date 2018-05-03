package test2

import (
	"fmt"
)

var Name string
var Age int

//每个源文件都可以包含一个init函数，这个init函数自动被go运行框架调用，并且在main函数前调用
func init() {
	Name = "李四"
	Age = 20
	fmt.Printf("Name:%s Age:%d\n", Name, Age)
}

func SetInfo(name string, age int) {
	fmt.Printf("Name:%s Age:%d\n", Name, Age)
	Name = name
	Age = age
	fmt.Printf("Name:%s Age:%d\n", Name, Age)
}
