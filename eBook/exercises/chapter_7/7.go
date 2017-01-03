package main

import (
	"fmt"
)

func main() {
	// basicSlice()
	sliceSum()
}

func sliceSum() {
	var arr = []uint64{0, 1, 2, 3, 4, 5, 6, 7}

	result := sum(arr[:])
	fmt.Println(result)
}

func sum(arr []uint64) (sum uint64) {
	// 传递了一个切片参数过来
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return
}

func basicSlice() {
	var arr = []int{0, 1, 2, 3, 4, 5}
	var slice1 = arr[2:5]

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d ", slice1[i])
	}
	fmt.Println()

	fmt.Println("Length of the array", len(arr))
	fmt.Println("Length of the slice1", len(slice1))
	fmt.Println("Capiable of the slice1", len(slice1))

	// 从切片的当前位置开始计算，如果超出了范围，则继续取数组中的值，如果超出了数组的范围，抛出异常
	slice1 = slice1[0:4]

	// 切片不能向前移动，如果像下面这样操作将会抛出异常
	// slice_error = slice1[-1:]

	fmt.Println(slice1)
}
