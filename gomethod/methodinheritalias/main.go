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

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type T1 = T

func main() {
	var t T
	var pt *T
	var t1 T1
	var pt1 *T1

	dumpMethodSet(t)
	dumpMethodSet(t1)

	dumpMethodSet(pt)
	dumpMethodSet(pt1)
}
