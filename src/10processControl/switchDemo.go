package main

import (
	"fmt"
)

func main() {
	var num int = 1
	fmt.Println("请按下楼层")
	fmt.Scan(&num)
	switch num { //switch后面是变量本身
	case 1:
		fmt.Println("按下的是1楼")
		break //go语言保留了break关键字，跳出switch，不写默认包含
	case 2:
		fmt.Println("按下的是2楼")
		fallthrough //不跳出switch
	case 3:
		fmt.Println("按下的是3楼")
		fallthrough
	case 4, 5, 6, 7, 8: //可以写多个
		fmt.Println("按下的是", num, "楼")
		fallthrough
	default:
		fmt.Println("按下的是无效楼层")
	}

	//switch 后面支持一个初始化语句，初始化语句和变量本身以分号来分隔
	switch num2 := 2; num2 {
	case 1:
		fmt.Println("按下的是1楼")
	case 2:
		fmt.Println("按下的是2楼")
	default:
		fmt.Println("按下的是无效楼")
	}

	score := 95
	switch { //可以没有条件
	case score > 90: //case后面可以放条件
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	case score > 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
