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
