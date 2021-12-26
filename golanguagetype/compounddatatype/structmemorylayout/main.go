package main

import "unsafe"

type T struct {
	b byte   // 1 + padding(7)
	i int64  // 8
	u uint16 // 2 + padding(6)
}

type S struct {
	b byte   // 1
	u uint16 // 2 + padding(5)
	i int64  // 8
}

func main() {
	var t T
	println(unsafe.Sizeof(t)) // 24

	var s S
	println(unsafe.Sizeof(s)) // 16
}
