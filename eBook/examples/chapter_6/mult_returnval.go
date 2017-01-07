package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculate(3, 4))
	fmt.Println(calculate_anonymous(3, 4))
}

func calculate(a int, b int) (sum int, mul int, sub int) {
	sum = a + b
	mul = a * b
	sub = a - b

	return
}

func calculate_anonymous(a int, b int) (int, int, int) {
	sum := a + b
	mul := a * b
	sub := a - b

	return sum, mul, sub
}
