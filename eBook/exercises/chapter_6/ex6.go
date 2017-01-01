package main

import (
	"fmt"
	"strings"
)

func main() {
	fib(10)
}

func fib(times int) {
	fibo := fibonacci()
	for i := 0; i < times; i++ {
		fmt.Printf("%d ", fibo())
	}
	fmt.Println()

	addJPEGSuffix := MakeAddSuffix(".jpeg")
	fmt.Println(addJPEGSuffix("life.jpeg"))
	fmt.Println(addJPEGSuffix("life"))
}

func fibonacci() func() int {
	// 这里一开始设置p1为０，是为了让前两次调用返回１
	var p1, p2 int = 0, 1

	return func() (result int) {
		result = p1 + p2
		p1, p2 = p2, result

		if p1 == 1 {
			return 1
		}

		return
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	// 这就是一个添加后缀的工厂函数的例子
	return func(name string) string {
		if strings.HasSuffix(name, suffix) {
			return name
		}
		return name + suffix
	}
}
