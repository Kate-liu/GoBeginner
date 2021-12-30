package main

import (
	"fmt"
	"reflect"
)

// 发现问题
// type Interface interface {
// 	M1()
// 	M2()
// }
//
// type T struct{}
//
// func (t T) M1()  {}
// func (t *T) M2() {}
//
// func main() {
// 	var t T
// 	var pt *T
// 	var i Interface
// 	i = pt
// 	i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
// }

// 输出方法集合，确定问题

type Interface interface {
	M1()
	M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

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

func main() {
	var t T
	var pt *T
	dumpMethodSet(t)
	dumpMethodSet(pt)
}
