package main

import "fmt"

func main() {
	var s int8 = 127
	s += 1
	// 预期 128，实际结果 -128
	fmt.Println(s)

	var u uint8 = 1
	u -= 2
	// 预期 -1，实际结果 255
	fmt.Println(u)
}
