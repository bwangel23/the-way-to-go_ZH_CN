package main

import (
	"fmt"
)

var a = "G"

func main() {
	n()
	m()
	n()
}

func m() { fmt.Println(a) }

func n() {
	a := "0"
	fmt.Println(a)
}
