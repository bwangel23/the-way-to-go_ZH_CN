package main

import (
	"fmt"
)

const LEN int = 5

func main() {
	// 数组的最大长度位2GB

	// basic()
	// initArray()
	// pointer()
	// arrayContants()
	// pointer2()
	// multidimArray()
	arraySum()
}

func arraySum() {
	var arr = [3]float64{1.2, 4.3, 6.5}

	res := sum(&arr)


func sum(arr *[3]float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}
	return
}

func multidimArray() {
	const (
		WIDTH  = 1920
		HEIGHT = 1080
	)

	var screen [WIDTH][HEIGHT]int

	for i := 0; i < WIDTH; i++ {
		for j := 0; j < HEIGHT; j++ {
			screen[i][j] = 0x111
		}
	}
}

func pointer2() {
	for i := 0; i < 5; i++ {
		showPointerArray(&[3]int{i, i * i, i * i * i})
	}
}

func showPointerArray(arr *[3]int) {
	fmt.Println(arr)
}

func arrayContants() {
	var arrKeys = [5]string{2: "two", 3: "three"}
	var numbers = [...]string{"one", "two", "three", "four"}

	for i, v := range arrKeys {
		fmt.Println(i, v)
	}
	fmt.Println(numbers)
}

func pointer() {
	var arr1 [5]int
	var arr2 = new([5]int)

	fmt.Printf("%T, %T\n", arr1, arr2)
}

func initArray() {
	// 注意数组的这种初始化的方式
	a := [...]string{"a", "b", "c"}

	// range 返回index, value 两个值，也可以只用一个变量来接收
	for i := range a {
		fmt.Println(i, a[i])
	}
}

func basic() {
	var arr [LEN]int

	for i := 0; i < LEN; i++ {
		arr[i] = i + 1
	}

	for i := 0; i < LEN; i++ {
		fmt.Println(i, arr[i])
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}

	for i := range arr {
		fmt.Println(i, arr[i])
	}
}
