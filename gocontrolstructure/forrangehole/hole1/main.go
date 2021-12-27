package main

import (
	"fmt"
	"time"
)

func main() {
	// // for range 循环变量的重用
	// var m = []int{1, 2, 3, 4, 5}
	//
	// for i, v := range m {
	// 	go func() {
	// 		time.Sleep(time.Second * 3)
	// 		fmt.Println(i, v)
	// 	}()
	// }
	// time.Sleep(time.Second * 10)

	// // for range 循环变量的重用 的 等价转换
	// var m = []int{1, 2, 3, 4, 5}
	//
	// {
	// 	i, v := 0, 0
	// 	for i, v = range m {
	// 		go func() {
	// 			time.Sleep(time.Second * 3)
	// 			fmt.Println(i, v)
	// 		}()
	// 	}
	// }
	// time.Sleep(time.Second * 10)

	// for range 循环变量的重用 的 修改(绑定参数 i，v)
	var m = []int{1, 2, 3, 4, 5}

	for i, v := range m {
		go func(i, v int) {
			time.Sleep(time.Second * 3)
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(time.Second * 10)
}
