package main

import (
	"fmt"
)

func main() {
	var month int = 4

	switch {
	case month > 0 && month <= 3:
		fmt.Println("Sprint")
	case month > 3 && month <= 6:
		fmt.Println("Summar")
	case month > 6 && month <= 9:
		fmt.Println("autumn")
	case month > 9 && month <= 12:
		fmt.Println("Winter")
	}
}
