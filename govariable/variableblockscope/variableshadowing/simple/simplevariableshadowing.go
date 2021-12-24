package main

import "fmt"

var a = 11

func foo(n int) {
	a := 1
	a += n
}

func main() {
	fmt.Println("a = ", a) // 11
	foo(5)
	fmt.Println("after calling foo, a = ", a) // 11
}
