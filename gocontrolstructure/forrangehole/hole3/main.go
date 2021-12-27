package main

import "fmt"

func main() {
	// // 遍历 map 中元素的随机性
	// var m = map[string]int{
	// 	"tony": 21,
	// 	"tom":  22,
	// 	"jim":  23,
	// }
	//
	// counter := 0
	// for k, v := range m {
	// 	if counter == 0 {
	// 		delete(m, "tony")
	// 	}
	// 	counter++
	// 	fmt.Println(k, v)
	// }
	// fmt.Println("counter is ", counter)

	// 新创建一个 map 元素项
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)
}
