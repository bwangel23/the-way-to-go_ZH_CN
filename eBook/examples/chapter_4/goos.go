package main

import (
	"fmt"
	"os"
)

func main() {
	var goos string = os.Getenv("OSTYPE")
	fmt.Printf("The operating system is: %s\n", goos)

	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	fmt.Printf("%v\n", 3+2i)
}
