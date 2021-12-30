package main

import (
	"fmt"
	"reflect"
)

func dumpMethodSet(i interface{}) {
	dynTyp := reflect.TypeOf(i)
	if dynTyp == nil {
		fmt.Printf("there is no dynamic type\n")
		return
	}
	n := dynTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", dynTyp)
		return
	}
	fmt.Printf("%s's method set:\n", dynTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", dynTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

//
// type T struct{}
//
// func (T) M1()  {}
// func (T) M2()  {}
//
// func (*T) M3() {}
// func (*T) M4() {}
//
// func main() {
// 	var n int
// 	dumpMethodSet(n)
// 	dumpMethodSet(&n)
//
// 	var t T
// 	dumpMethodSet(t)
// 	dumpMethodSet(&t)
// }

// 思考题验证

type T struct{}

func (T) M1() {}
func (T) M2() {}

type S T

// type S =  T

func main() {
	var t T
	dumpMethodSet(t)
	dumpMethodSet(&t)

	var s S
	dumpMethodSet(s)
	dumpMethodSet(&s)
}
