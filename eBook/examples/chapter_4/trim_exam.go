package main

import (
	"fmt"
	"strings"
)

func main() {
	var orig string = "中文哈哈哈中文"

	afterTrim := strings.Trim(orig, "中文")

	fmt.Println(afterTrim)
}
