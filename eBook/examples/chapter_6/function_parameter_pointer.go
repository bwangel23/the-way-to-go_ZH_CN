package main

import (
	"fmt"
)

func main() {
	var a int = 3

	Pointer(&a)
	Value(a)
}

func Pointer(a *int) {
	fmt.Println(a)
}

func Value(a int) {
	fmt.Println(&a)
}
