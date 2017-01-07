package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d is Even: %t\n", 16, even(16))
}

func even(nr int) bool {
	if nr == 0 {
		return true
	}

	return odd(ReverSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}

	return even(ReverSign(nr) - 1)
}

func ReverSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}
