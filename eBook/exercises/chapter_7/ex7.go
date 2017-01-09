package main

import (
	"bytes"
	"fmt"
)

func main() {
	// fibonacci()
	// splitBuffer()
	// reverseStringFunc()
	callMapFunc()
	// splitBuffer()
	// forRange()
	// arraySumFunc()
	// arrayAppend()
	// sliceCopy()
	InsertStringSliceFunc()
}

func callMapFunc() {
	numbers := []int{1, 2, 3, 4, 5}
	result := mapFunc(multiTen, numbers)
	fmt.Println(result)
}

func multiTen(n int) int {
	return n * 10
}

func mapFunc(f func(int) int, numbers []int) (result []int) {
	for _, v := range numbers {
		r := f(v)
		result = append(result, r)
	}

	return
}

func InsertStringSliceFunc() {
	dest := [...]byte{'a', 'b', 'c'}
	src := [...]byte{'e', 'f', 'd'}
	fmt.Println(string(InsertStringSlice(dest[:], src[:], 1)))
}

func InsertStringSlice(dest []byte, src []byte, index int) (result []byte) {

	// 复制的时候一定要确保目的切片足够长
	// result = make([]byte, 10, 20)
	// copy(result, dest[:index])

	result = append(result, dest[:index]...)
	result = append(result, src...)
	result = append(result, dest[index:]...)

	return
}

func reverseStringFunc() {
	fmt.Println("Google is reversed in: ", reverseString("Google"))
}

func reverseString(s string) (result string) {
	s_bytes := []byte(s)

	length := len(s_bytes)

	for i := 0; i < length/2; i++ {
		s_bytes[i], s_bytes[length-i-1] = s_bytes[length-i-1], s_bytes[i]
	}

	return string(s_bytes)
}

func fibonacci() {
	const LEN = 50
	var result [LEN]uint64

	for i := 0; i < LEN; i++ {
		if i == 0 || i == 1 {
			result[i] = uint64(1)
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}

	fmt.Println(result)
}

func sliceCopy() {
	var sliceFrom = []int{1, 2, 3}
	var sliceTo = make([]int, len(sliceFrom), 10)

	copy(sliceTo, sliceFrom[:1])

	fmt.Println(sliceTo)
}

func arrayAppend() {
	// 上面的定义是一个数组，下面的定义是一个切片
	// var sl1 = [...]int{0, 1, 3, 4}
	var sl1 = []int{1, 2, 3}

	sl1 = append(sl1, 5, 8)

	fmt.Println(sl1)
}

func forRange() {
	var item int
	var items = [...]int{10, 20, 30, 40, 50}

	for _, item = range items {
		item *= 2
	}

	fmt.Println(item)
}

func arraySumFunc() {
	// 这里使用了数组指针
	var arrayP = &[...]int{1, 3, 4, 5, 6, 7, 9}

	fmt.Println(arraySum(arrayP))
}

func arraySum(array *[7]int) (sum int) {
	for _, value := range *array {
		sum += value
	}

	return
}

func splitBuffer() {
	var buf *bytes.Buffer = new(bytes.Buffer)
	buf.WriteString("hello, xff")

	s1, s2 := splitBufferFunc(buf, 5)
	fmt.Printf("%s\n%s\n", string(s1), string(s2))
}

func splitBufferFunc(buf *bytes.Buffer, n int) (s1 []byte, s2 []byte) {
	s1 = buf.Next(n)
	s2 = buf.Next(buf.Len())
	return s1, s2
}
