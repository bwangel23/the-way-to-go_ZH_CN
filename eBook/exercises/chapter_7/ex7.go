package main

import (
	"fmt"
)

func main() {
	fibonacci()
}

func fibonacci() {
	const LEN = 50
	var result [LEN]uint64

	for i := 0; i < LEN; i++ {
		if i == 0 || i == 1 {
			result[i] = uint64(1)
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}

	fmt.Println(result)
}
