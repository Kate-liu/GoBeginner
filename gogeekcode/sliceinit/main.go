package main

import "fmt"

func main() {
	// []int 初始化
	arr := []int{}
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(arr)

	// make 初始化
	arr2 := make([]int, 0, 5)
	arr2 = append(arr2, 1, 2, 3, 4, 5)
	fmt.Println(arr2)

	// 打印汇编代码命令：GOOS=linux GOARCH=amd64 go tool compile -S main.go
}
