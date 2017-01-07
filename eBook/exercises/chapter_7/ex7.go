package main

import (
	"bytes"
	"fmt"
)

func main() {
	fibonacci()
	splitBuffer()
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
