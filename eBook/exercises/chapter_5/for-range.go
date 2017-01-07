package main

import (
	"fmt"
)

func main() {
	for ix, val := range "中国汉字" {
		var v int
		fmt.Printf("[%d]Character on position %d is: %c\n", v, ix, val)
		v = 5
	}

	s := ""
	for s != "aaaaa" {
		fmt.Println("Value of s:", s)
		s = s + "a"
	}

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
