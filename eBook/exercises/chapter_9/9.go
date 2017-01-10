package main

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"unsafe"
)

func main() {
	// List()
	// sizeof()
	replaceString()
}

func replaceString() {
	pat := "[0-9]+.[0-9]+"
	str := "今天的温度为6.3摄氏度"

	if ok, _ := regexp.MatchString(pat, str); ok {
		fmt.Println("Matched")
	}

	f := func(src string) string {
		number, _ := strconv.ParseFloat(src, 64)
		return strconv.FormatFloat(number*2, 'f', 2, 32)
	}

	re, _ := regexp.Compile(pat)
	result := re.ReplaceAllStringFunc(str, f)

	fmt.Printf("%s", result)
}

func sizeof() {
	var i int32 = 32

	fmt.Println(unsafe.Sizeof(i))
}

func List() {
	var numbers = list.New()

	numbers.PushBack(101)
	numbers.PushBack(102)
	numbers.PushBack(103)

	for e := numbers.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d", e.Value)
		if e.Prev() != nil {
			fmt.Printf(" -- %d\n", e.Prev().Value)
		} else {
			fmt.Printf(" -- %d\n", 0)
		}
	}
}
