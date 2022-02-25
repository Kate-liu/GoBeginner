package main

import (
	"constraints"
)

func Add[T constraints.Integer](a, b T) T {
	return a + b
}

func main() {
	var m, n int = 5, 6
	println(Add(m, n)) // Add[int](m, n)

	var i, j int64 = 15, 16
	println(Add(i, j)) // Add[int64](i, j)

	var c, d byte = 0x11, 0x12
	println(Add(c, d)) // Add[byte](c, d)
}
