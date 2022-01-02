package main

import "fmt"

func main() {
	// // array to slice
	// a := [3]int{11, 12, 13}
	// b := a[:] // 通过切片化将数组a转换为切片b
	// b[1] += 10
	// fmt.Printf("%v\n", a) // [11 22 13]
	//
	// // slice to array unsafe method
	// c := []int{11, 12, 13}
	// var p = (*[3]int)(unsafe.Pointer(&c[0]))
	// p[1] += 10
	// fmt.Printf("%v\n", c) // [11 22 13]
	//
	// // slice to array new feature method
	d := []int{11, 12, 13}
	q := (*[3]int)(d) // 将切片转换为数组类型指针
	q[1] += 10
	fmt.Printf("%v\n", d) // [11 22 13]

	// // example
	// var b = []int{11, 12, 13}
	// var p = (*[4]int)(b)     // cannot convert slice with length 3 to pointer to array with length 4
	// var p = (*[0]int)(b)     // ok，*p = []
	// var p = (*[1]int)(b)     // ok，*p = [11]
	// var p = (*[2]int)(b)     // ok，*p = [11, 12]
	// var p = (*[3]int)(b)     // ok，*p = [11, 12, 13]
	// var p = (*[3]int)(b[:1]) // cannot convert slice with length 1 to pointer to array with length 3
	//
	// // nil slice and empty slice
	// var b1 []int // nil切片
	// p1 := (*[0]int)(b1)
	// var b2 = []int{} // empty切片
	// p2 := (*[0]int)(b2)
}
