package main

import "fmt"

// // 定义 新接口
//
// type MyInterface interface {
// 	M1(int) error
// 	M2(io.Writer, ...string)
// }

// // 等价的 新接口
//
// type MyInterface interface {
// 	M1(a int) error
// 	M2(w io.Writer, strs ...string)
// }
//
// type MyInterface interface {
// 	M1(n int) error
// 	M2(w io.Writer, args ...string)
// }

// // 类型嵌入中的方法存在交集
//
// type Interface1 interface {
// 	M1()
// }
//
// type Interface2 interface {
// 	M1(string)
// 	M2()
// }
//
// type Interface3 interface {
// 	Interface1
// 	Interface2 // 编译器报错：duplicate method M1
// 	M3()
// }

// // 空接口类型
//
// type EmptyInterface interface {
// }

// // 定义接口类型变量
// var err error   // err是一个error接口类型的实例变量
// var r io.Reader // r是一个io.Reader接口类型的实例变量

func main() {
	// // 赋值给空接口类型的变量
	// var i interface{} = 15 // ok
	// i = "hello, golang"    // ok
	// type T struct{}
	//
	// var t T
	// i = t  // ok
	// i = &t // ok

	// // 类型断言
	// v := i.(T)

	// 类型断言例子
	var a int64 = 13
	var i interface{} = a
	v1, ok := i.(int64)
	fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
	v2, ok := i.(string)
	fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
	v3 := i.(int64)
	fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
	v4 := i.([]int)                                     // panic: interface conversion: interface {} is int64, not []int
	fmt.Printf("the type of v4 is %T\n", v4)
}
