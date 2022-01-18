package main

import "fmt"

func main() {
	// panic 只会触发当前 goroutine 的延迟函数调用
	// 跨协程失效
	// defer println("in main")
	// go func() {
	// 	defer println("in goroutine")
	// 	panic("")
	// }()
	//
	// time.Sleep(1 * time.Second)

	// recover 只有在 defer 中调用才会生效
	// 失效的崩溃恢复
	// defer println("in main")
	// if err := recover(); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// panic("unknown err")

	// panic 允许在 defer 中嵌套多次调用
	// 嵌套崩溃
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
}
