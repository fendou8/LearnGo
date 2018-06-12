package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {

	files := os.Args[1:]
	if len(files) == 0 {
		_, counts := countLines2(os.Stdin)
		printCounts(counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			isrepeated, counts := countLines2(f)
			if isrepeated {
				fmt.Println(arg, " 中出现重复")
				printCounts(counts)
			}
			f.Close()
		}
	}

}

func printCounts(counts map[string]int) {
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines2(f *os.File) (bool, map[string]int) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	isrepeated := false
	for input.Scan() {
		temp := counts[input.Text()]
		counts[input.Text()]++
		if counts[input.Text()] > temp && temp > 0 {
			isrepeated = true
		}
	}
	return isrepeated, counts
}
