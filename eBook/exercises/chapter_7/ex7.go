package main

import (
	"bytes"
	"fmt"
)

func main() {
	var byteArray = [...]byte{'c', 'f', 'b'}
	s := string(byteArray[:])
	fmt.Printf("%s\n", s)
}

func splitBuffer() {
	var str string = "hello, xff"
	var buf *bytes.Buffer = bytes.NewBufferString(str)
	s1, s2 := splitBufferFunc(buf)
	fmt.Println(s1, s2)
}

func splitBufferFunc(buf *bytes.Buffer) (s1 []byte, s2 []byte) {
	// s1 = buf.ReadBytes(3)
	// s2 = buf.ReadByte()
	return s1, s2
}
