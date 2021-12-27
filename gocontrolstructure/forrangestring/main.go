package main

import "fmt"

func main() {
	// for range 循环 string 类型
	var s = "中国人"
	for i, v := range s {
		fmt.Printf("%d %s 0x%x\n", i, string(v), v)
	}

	// for 经典形式
	var t = "中国人"
	for i := 0; i < len(t); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}
	// 输出
	// index: 0, value: 0xe4
	// index: 1, value: 0xb8
	// index: 2, value: 0xad
	// index: 3, value: 0xe5
	// index: 4, value: 0x9b
	// index: 5, value: 0xbd
	// index: 6, value: 0xe4
	// index: 7, value: 0xba
	// index: 8, value: 0xba

}
