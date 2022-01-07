package main

import (
	"fmt"
	"unsafe"
)

func foo(arr [5]int) {

}

func main() {
	// var arr1 [5]int
	// var arr2 [6]int
	// var arr3 [5]string
	// foo(arr1) // ok
	// foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型 // cannot use arr2 (type [6]int) as type [5]int in argument to foo
	// foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型 // cannot use arr3 (type [5]string) as type [5]int in argument to foo

	var arr4 = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println("数组长度: ", len(arr4))           // 6
	fmt.Println("数组大小: ", unsafe.Sizeof(arr4)) // 48

	// 数组默认初始化为零值
	var arr5 [6]int
	fmt.Println(arr5) // [0 0 0 0 0 0]

	// 大括号显示数组赋值
	var arr6 = [6]int{
		11, 12, 13, 14, 15, 16,
	}
	fmt.Println(arr6) // [11 12 13 14 15 16]

	// 自动计算数组长度
	var arr7 = [...]int{
		21, 22, 23,
	}
	fmt.Println(arr7)      // [21 22 23]
	fmt.Println(len(arr7)) // 3

	// 下标赋值的方式初始化
	var arr8 = [...]int{
		99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
	}
	fmt.Println(arr8) // [0 0 ... 99]

	// 数组的访问
	var arr9 = [5]int{11, 12, 13, 14, 15}
	fmt.Println(arr9[0], arr9[4]) // 11 15
	fmt.Println(arr9[-1])         // invalid array index -1 (index must be non-negative) 错误：下标值不能为负数
	fmt.Println(arr9[99])         // invalid array index 99 (out of bounds for 5-element array) 错误：下标值超出了arr的长度范围

	// 数组访问越界报错区分
	// var arr10 = [5]int{11, 12, 13, 14, 15}
	// var ind = 6
	// fmt.Println(arr10)
	// fmt.Println(arr10[6])   // 编译器报错：invalid array index 6 (out of bounds for 5-element array)
	// fmt.Println(arr10[ind]) // 运行时报错：panic: runtime error: index out of range [6] with length 5

}
