package main

import (
	"fmt"
	"reflect"
)

// 类型声明语法声明的新类型
// type I interface {
// 	M1()
// 	M2()
// }
//
// type T int
//
// type NT T // 基于已存在的类型T创建新的defined类型NT
// type NI I // 基于已存在的接口类型I创建新defined接口类型NI

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

type T1 T

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
