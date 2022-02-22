package main

import "fmt"

type Vector[T any] []T

func (v Vector[T]) Dump() {
	fmt.Printf("%#v\n", v)
}

func main() {
	var iv = Vector[int]{1, 2, 3, 4}
	var sv Vector[string]

	sv = []string{"a", "b", "c", "d"}

	iv.Dump()
	sv.Dump()
}
