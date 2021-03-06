package main

import "fmt"

func main() {
	// 动态扩容 例子
	var s []int
	s = append(s, 11)
	fmt.Println(len(s), cap(s)) // 1 1
	s = append(s, 12)
	fmt.Println(len(s), cap(s)) // 2 2
	s = append(s, 13)
	fmt.Println(len(s), cap(s)) // 3 4
	s = append(s, 14)
	fmt.Println(len(s), cap(s)) // 4 4
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 5 8
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 6 8
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 7 8
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 8 8
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 9 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 10 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 11 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 12 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 13 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 14 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 15 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 16 16
	s = append(s, 15)
	fmt.Println(len(s), cap(s)) // 17 32

	// 进行新容量的申请，需要进行向上取整
	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(len(arr), cap(arr)) // 5 6

	// 自动扩容问题：切片与数组解除绑定
	// 定义数组
	// u := [...]int{11, 12, 13, 14, 15}
	// fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	// // 开始切片
	// s := u[1:3]
	// fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	// s = append(s, 24)
	// fmt.Println("after append 24, array:", u)
	// fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	// s = append(s, 25)
	// fmt.Println("after append 25, array:", u)
	// fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	// // 切片和原数组解除绑定
	// s = append(s, 26)
	// fmt.Println("after append 26, array:", u)
	// fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	// // 测试是否真的解除绑定
	// s[0] = 22
	// fmt.Println("after reassign 1st elem of slice, array:", u)
	// fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

}
