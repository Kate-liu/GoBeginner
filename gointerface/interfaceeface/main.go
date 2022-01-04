package main

import "fmt"

type T struct {
	n int
	s string
}

// func main() {
// 	var t = T{
// 		n: 17,
// 		s: "hello, interface",
// 	}
//
// 	var ei interface{} = t // Go运行时使用eface结构表示ei
//
// 	fmt.Printf("error occur: %+v\n", ei)
// }

func main() {
	var n int = 61
	var ei interface{} = n
	n = 62                          // n的值已经改变
	fmt.Println("data in box:", ei) // 输出仍是61
}
