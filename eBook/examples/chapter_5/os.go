package main

import (
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("windows")
	} else {
		fmt.Println("Not windows")
	}
}
