package main

var a string

func main() {
	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	// 注意这里会输出G
	// 因为f1中的变量a属于f1的作用域，并没有传递到这个作用域中
	print(a)
}
