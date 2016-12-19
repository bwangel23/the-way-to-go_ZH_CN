package main

import "fmt"

func main() {
	f()
}

func a() {
	i := 0
	defer fmt.Println(i)
	// 注意这里会输出0，传递的是形参
	i++
	return
}

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
