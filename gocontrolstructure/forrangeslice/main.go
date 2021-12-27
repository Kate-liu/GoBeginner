package main

import "fmt"

func main() {
	// for 经典形式 遍历切片中的元素
	var sl = []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("sl[%d] = %d\n", i, sl[i])
	}

	// for range 形式 遍历切片中的元素
	for i, v := range sl {
		fmt.Printf("sl[%d] = %d\n", i, v)
	}

	// 省略元素值变量
	for i := range sl {
		// ...
	}

	// 空标识符代替下标变量
	for _, v := range sl {
		// ...
	}

	// 省略下标与元素值变量
	for _, _ = range sl {
		// ...
	}
	// 省略下标与元素值变量 优雅方式
	for range sl {
		// ...
	}
}
