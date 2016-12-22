package main

import (
	"fmt"
)

func main() {
	// FizzBuzz()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d: %b and %d, %b\n", i, i, ^i, ^i)
	}

	var i uint8 = 10

	// ^为单目运算符的时候
	// ^i 运算表示:
	// -1^i (i 为有符号数)
	// MAXINT^i (i 为无符号数)
	fmt.Printf("%b: %b %t\n", 255^i, ^i, 255^i == ^i)
}

func FizzBuzz() {
	for i := 0; i < 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println(i, "FizzBuzz")
		case i%3 == 0:
			fmt.Println(i, "Fizz")
		case i%5 == 0:
			fmt.Println(i, "Buzz")
		}
	}
}
