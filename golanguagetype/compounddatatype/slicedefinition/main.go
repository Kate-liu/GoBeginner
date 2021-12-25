package main

import "fmt"

func main() {
	// 切片变量的声明
	var nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(nums) // [1 2 3 4 5 6]

	// 获取切片长度
	fmt.Println(len(nums)) // 6

	// 添加切片元素
	nums = append(nums, 7)
	fmt.Println(nums)      // [1 2 3 4 5 6 7]
	fmt.Println(len(nums)) // 7

	//  make 函数创建切片
	sl1 := make([]byte, 6, 19) // 其中10为cap值，即底层数组长度，6为切片的初始长度
	fmt.Println(sl1)           // [0 0 0 0 0 0]
	fmt.Println(len(sl1))      // 6
	fmt.Println(cap(sl1))      // 19
	//  make 函数创建切片，默认 cap = len = 6
	sl2 := make([]byte, 6) // 其中默认6为cap值，即底层数组长度，6为切片的初始长度 // cap = len = 6
	fmt.Println(sl2)       // [0 0 0 0 0 0]
	fmt.Println(len(sl2))  // 6
	fmt.Println(cap(sl2))  // 6

	// 数组的切片化
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sl3 := arr[3:7:9]
	fmt.Println(arr)      // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(sl3)      // [4 5 6 7]
	fmt.Println(len(sl3)) // 4
	fmt.Println(cap(sl3)) // 6

	// 更改切片元素的值，会改变原数组的值
	sl3[0] += 10
	fmt.Println(arr)                // [1 2 3 14 5 6 7 8 9 10]
	fmt.Println(sl3)                // [14 5 6 7]
	fmt.Println("arr[3] =", arr[3]) // arr[3] = 14
	fmt.Println(sl3[5])             // 测试：在切片在访问 大于长度 小于 cap 的元素，会报错：panic: runtime error: index out of range [5] with length 4
}
