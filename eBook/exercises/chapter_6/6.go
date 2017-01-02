package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

var where = func() {
	_, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s %d\n", file, line)
}

var where2 = log.Print

func main2() {
	log.SetFlags(log.Llongfile)
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
		log.Print("")
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

const LIM = 100

var result [LIM]uint64

func main() {
	// 这里测试的结果，闭包反而比缓存更快一点
	fibsCache()
	fibs(LIM)
}

func fibs(times int) {
	start := time.Now()

	fibo := fibonacci()
	for i := 0; i < times; i++ {
		fmt.Printf("%d ", fibo())
	}
	fmt.Println()

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The time of the fibs is %s\n", delta)
}

func fibonacci() func() int {
	// 这里一开始设置为０，是为了让前两次调用返回１
	var p1, p2 int = 0, 0

	return func() (result int) {
		if p1 == 0 {
			p1 = 1
			return 1
		} else if p2 == 0 {
			p2 = 1
			return 1
		}

		result = p1 + p2
		p1, p2 = p2, result

		return
	}
}

func fibsCache() {
	start := time.Now()

	for i := 0; i < LIM; i++ {
		fmt.Printf("%d ", fibonacciCache(i))
	}
	fmt.Println()

	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The time of the fibsCache is %s\n", delta)
}

func fibonacciCache(n int) (res uint64) {
	if result[n] != 0 {
		return result[n]
	}

	if n <= 1 {
		result[n] = 1
	} else {
		result[n] = fibonacciCache(n-1) + fibonacciCache(n-2)
	}
	res = result[n]
	return
}

func f() {
	where2()
	// 这里这几个函数用的是同样的内存地址
	start := time.Now()
	for i := 0; i < 5; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		g(i)
		fmt.Printf("- g is of %T and has value %v\n", g, g)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("The time amout of the function f is %s\n", delta)
}

func test() (ret int) {

	// 这里的 defer 定义的函数，会把返回值 ret 改变
	// 这里的返回值是一个固定的变量 ret，所以return 1会把1赋值给ret变量
	defer func() {
		where()
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
