package main

import (
	"fmt"
)

func main() {
	var i int = 0

start:
	fmt.Println(i)
	if i < 15 {
		i++
		goto start
	}
}
