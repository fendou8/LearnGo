package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE:calc command [arguments] ...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquareroot of a non-negative value.")
}

func main() {
	args := os.Args
	fmt.Println("args : ", args, "  len : ", len(args))
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	for i, value := range args {
		fmt.Printf("arg[%d] = %v\r\n", i, value)
	}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("err1  ", err1)
			fmt.Println("err2  ", err2)
			fmt.Println("USAGE: calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("USAGE: calc sqrt <integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
}
