package main

import (
	"fmt"
)

func main() {
	for i := uint64(0); i < uint64(30); i++ {
		fmt.Printf("factorial(%d): %d\n", i, factorial(i))
	}
}

func factorial(n uint64) (fac uint64) {
	fac = 1
	if n >= 1 {
		fac = factorial(n-1) * n
		return
	}

	return
}
