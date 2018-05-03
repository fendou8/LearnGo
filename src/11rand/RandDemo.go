package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	//设置随机种子
	rand.Seed(time.Now().Unix())
	rand.
}
func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Println("第", i, "个随机整数：", a)
	}
	fmt.Println("-----------------------------------")
	for i := 0; i < 10; i++ {
		a := rand.Intn(100)
		fmt.Println("第", i, "个随机整数：", a)
	}
	fmt.Println("-----------------------------------")
	for i := 0; i < 10; i++ {
		a := rand.Float32()
		fmt.Println("第", i, "个随机浮点数：", a)
	}
}
