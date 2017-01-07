package main

import (
	"fmt"
	"strconv"
)

func main() {
	var ret int
	ret = atoi("3")
	fmt.Printf("%d\n", ret)
}

func atoi(s string) (n int) {
	n, err := strconv.Atoi(s)
	fmt.Println(n, err)
	return
}
