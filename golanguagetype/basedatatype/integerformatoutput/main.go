package main

import "fmt"

func main() {

	// 格式化输出
	var a int8 = 59
	fmt.Printf("%b\n", a) // 输出二进制：111011
	fmt.Printf("%d\n", a) // 输出十进制：59
	fmt.Printf("%o\n", a) // 输出八进制：73
	fmt.Printf("%O\n", a) // 输出八进制(带0o前缀)：0o73
	fmt.Printf("%x\n", a) // 输出十六进制(小写)：3b
	fmt.Printf("%X\n", a) // 输出十六进制(大写)：3B
}
