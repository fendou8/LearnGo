package main

import (
	"fmt"
)

func main() {
	//给int64起一个别名叫bigint
	type bigint int64

	var a bigint
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)
	var b long = 11
	var ch char = 's'
	fmt.Printf("b type is %T\n", b)
	fmt.Printf("ch type is %T\n", ch)
}
