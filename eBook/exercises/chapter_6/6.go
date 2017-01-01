package main

import (
	"fmt"
)

func main() {

	// 直接调用匿名函数
	func() {
		sum := 0.0
		for i := 1.0; i <= 1e6; i++ {
			sum += i
		}

		fmt.Println(sum)
	}()

	f()

	fv := func() {
		fmt.Println("Hello, World")
	}

	fv()
	fmt.Printf("The fv is %T\n", fv)

	fmt.Println(test())

	fmt.Println("result of Add2()(3)", Add2()(3))

	add3 := Adder(3)
	fmt.Println("result of Adder(3)(4)", add3(4))

	total := Total()
	fmt.Printf("%d - %d - %d\n", total(3), total(200), total(1))
}

func f() {
	// 这里这几个函数用的是同样的内存地址
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf("- g is of %T and has value %v\n", g, g)
	}
}

func test() (ret int) {
	// 这里的 defer 定义的函数，会把返回值 ret 改变
	// 这里的返回值是一个固定的变量 ret，所以return 1会把1赋值给ret变量
	defer func() {
		ret++
	}()
	return 2
}

// Add2 不传入参数，返回一个函数，函数返回声明：func(b int) int

func Add2() func(int) int {
	return func(a int) int {
		return a + 2
	}
}

func Adder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func Total() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
