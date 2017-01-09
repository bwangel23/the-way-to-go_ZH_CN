package main

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

const LEN int = 5

func main() {
	// 数组的最大长度为2GB

	// basic()
	// initArray()
	// pointer()
	// arrayContants()
	// pointer2()
	// multidimArray()
	// arraySum()
	// basicSlice()
	// sliceSum()
	// makeSlice()
	// useBuffer()
	// strToSlice()
	// subString()
	// changeString()
	// StringCompareFunc()
	sortString()
}

func sortString() {
	s := []string{"xff", "abc", "python"}
	sort.Strings(s)

	fmt.Println("Strings: ", s)
}

func StringCompareFunc() {
	result := StringCompare([]byte("xff1"), []byte("xff"))

	fmt.Println(result)
}

func StringCompare(a []byte, b []byte) (result int) {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}

	return 0
}

func changeString() {
	// 字符串是不可变的，因为要首先将其转换成字节数组，再来转变
	s := "hello"
	s_bytes := []byte(s)
	s_bytes[0] = 'c'
	s = string(s_bytes)
	fmt.Println(s)
}

func subString() {
	// 字符串是一个双字结构，包含只想字符数组首部的指针和一个记录字符串长度的整数
	s := "Hello, 小芳芳"
	substr := s[5:]

	fmt.Println(substr, reflect.TypeOf(substr))
}

func strToSlice() {
	s := "\xF0\x9F\x98\x8F\u754c中文繁體"

	for i, v := range []rune(s) {
		fmt.Printf("%d: %c\n", i, v)
	}
}

func arraySum() {
	var arr = [3]float64{1.2, 4.3, 6.5}

	res := sum(&arr)

	fmt.Println(res)
}

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

func useBuffer() {
	var buffer bytes.Buffer
	// for {
	// 	if s, ok := getNextString(); ok { //method getNextString() not shown here
	// 		buffer.WriteString(s)
	// 	} else {
	// 		break
	// 	}
	// }
	buffer.WriteString("hello")
	fmt.Print(buffer.String(), "\n")
}

func makeSlice() {
	var slice1 []int = make([]int, 5, 10)

	for i := 0; i < len(slice1); i++ {
		slice1[i] = i * 5
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Println(i, slice1[i])
	}

	fmt.Println("The length of the slice1 is ", len(slice1))
	fmt.Println("The capiable of the slice1 is ", cap(slice1))
}

func sliceSum() {
	var arr = []uint64{0, 1, 2, 3, 4, 5, 6, 7}

	result := sumUInt(arr[:])
	fmt.Println(result)
}

func sumUInt(arr []uint64) (sum uint64) {
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
