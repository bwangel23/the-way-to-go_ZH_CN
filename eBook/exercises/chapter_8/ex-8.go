package main

import (
	"fmt"
	"sort"
)

func main() {
	showFruit()
}

func showFruit() {
	fruitMap := map[string]string{
		"banana": "香蕉",
		"apple":  "苹果",
		"pear":   "梨",
		"mongo":  "芒果",
	}

	fmt.Println(fruitMap)

	keys := make([]string, len(fruitMap))

	i := 0
	for k, _ := range fruitMap {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, fruitMap[k])
	}
}
