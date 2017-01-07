package main

import (
	"fmt"
)

func main() {
	// var minNumber int

	// minNumber = min(1, 2, 3, 4, 5)
	// fmt.Println(minNumber)

	// /* 数组初始化并赋值 */
	// var arr = []int{32, 543, 5, 43, 54}

	// fmt.Println(min(arr...))

	// printEachArgument("hello", "world", "xff")
	// printTypeForEachArgument("hello", "world", "xff", 2, true)

	for i := 0; i <= 10; i++ {
		fmt.Printf("fibonacci(%d): %d\n", i, fibonacci(i))
	}
}

func min(numbers ...int) int {
	min := numbers[0]

	for _, number := range numbers {
		if number < min {
			min = number
		}
	}

	return min
}

func printEachArgument(arguments ...string) {
	for _, val := range arguments {
		fmt.Println(val)
	}
}

func printTypeForEachArgument(arguments ...interface{}) {
	for _, val := range arguments {
		switch val.(type) {
		case int:
			fmt.Println(val, ": int")
		case bool:
			fmt.Println(val, ": bool")
		case string:
			fmt.Println(val, ": string")
		}
	}
}

func fibonacci(n int) (res int) {
	if n < 2 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}
