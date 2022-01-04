package main

import "fmt"

type T struct {
	n int
	s string
}

func (T) M1() {}
func (T) M2() {}

type NonEmptyInterface interface {
	M1()
	M2()
}

func main() {
	var t = T{
		n: 18,
		s: "hello, interface",
	}

	var i NonEmptyInterface = t

	fmt.Printf("error occur: %+v\n", i)

}
