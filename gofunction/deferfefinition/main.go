package main

import (
	"fmt"
	"time"
)

func main() {
	// defer 逆序执行
	// for i := 0; i < 5; i++ {
	// 	defer fmt.Println(i)
	// }

	// defer 执行时机
	// {
	// 	defer fmt.Println("defer runs")
	// 	fmt.Println("block ends")
	// }
	// fmt.Println("main ends")

	// 函数运行时间 - 错误示例
	// startedAt := time.Now()
	// defer fmt.Println(time.Since(startedAt))
	// time.Sleep(time.Second)
	// 函数运行时间 - 正确示例 - 使用匿名函数实现
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()
	time.Sleep(time.Second)

}
