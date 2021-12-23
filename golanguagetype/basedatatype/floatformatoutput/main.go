package main

import "fmt"

func main() {
	// 浮点数格式化输出
	var f float64 = 123.45678
	fmt.Printf("%f\n", f) // 123.456780

	// 科学计数法形式输出
	fmt.Printf("%e\n", f) // 1.234568e+02
	fmt.Printf("%x\n", f) // 0x1.edd3be22e5de1p+06
}
