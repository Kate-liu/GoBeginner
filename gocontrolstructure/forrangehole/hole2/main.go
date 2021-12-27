package main

import "fmt"

func main() {
	// // 参与循环的是 range 表达式的副本
	// var a = [5]int{1, 2, 3, 4, 5}
	// var r [5]int
	//
	// fmt.Println("original a =", a)
	//
	// for i, v := range a {
	// 	if i == 0 {
	// 		a[1] = 12
	// 		a[2] = 13
	// 	}
	// 	r[i] = v
	// }
	//
	// fmt.Println("after for range loop, r =", r)
	// fmt.Println("after for range loop, a =", a)

	// // 参与循环的是 range 表达式的副本 的 修改(使用切片替代数组)
	// var a = [5]int{1, 2, 3, 4, 5}
	// var r [5]int
	//
	// fmt.Println("original a =", a)
	//
	// for i, v := range a[:] {
	// 	if i == 0 {
	// 		a[1] = 12
	// 		a[2] = 13
	// 	}
	// 	r[i] = v
	// }
	//
	// fmt.Println("after for range loop, r =", r)
	// fmt.Println("after for range loop, a =", a)

	// 参与循环的是 range 表达式的副本 的 修改(使用数组指针替代数组)
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("original a =", a)

	for i, v := range &a { // a 改为 &a
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}

	fmt.Println("after for range loop, r =", r)
	fmt.Println("after for range loop, a =", a)

}
