package main

// // 函数的参数调用 之 整型和数组
// // 不改变传入参数的值
// func myFunction1(i int, arr [2]int) {
// 	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
// }
//
// // 函数的参数调用
// // 改变传入参数的值
// func myFunction2(i int, arr [2]int) {
// 	i = 29
// 	arr[1] = 88
// 	fmt.Printf("in my_funciton - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
// }
//
// func main() {
// 	i := 30
// 	arr := [2]int{66, 77}
// 	fmt.Printf("before calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
// 	myFunction1(i, arr)
// 	// myFunction2(i, arr)
// 	fmt.Printf("after  calling - i=(%d, %p) arr=(%v, %p)\n", i, &i, arr, &arr)
// }

// // 函数的参数调用 之 结构体和指针
// //  传递结构体时：会拷贝结构体中的全部内容；
// //  传递结构体指针时：会拷贝结构体指针；
// type MyStruct struct {
// 	i int
// }
//
// func myFunction(a MyStruct, b *MyStruct) {
// 	a.i = 31
// 	b.i = 41
// 	fmt.Printf("in my_function - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
// }
//
// func main() {
// 	a := MyStruct{i: 30}
// 	b := &MyStruct{i: 40}
// 	fmt.Printf("before calling - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
// 	myFunction(a, b)
// 	fmt.Printf("after calling  - a=(%d, %p) b=(%v, %p)\n", a, &a, b, &b)
// }

// // go 语言结构体内存布局
// type MyStruct struct {
// 	i int
// 	j int
// }
//
// func myFunction(ms *MyStruct) {
// 	ptr := unsafe.Pointer(ms)
// 	for i := 0; i < 2; i++ {
// 		c := (*int)(unsafe.Pointer((uintptr(ptr) + uintptr(8*i))))
// 		*c += i + 1
// 		fmt.Printf("[%p] %d\n", c, *c)
// 	}
// }
//
// func main() {
// 	a := &MyStruct{i: 40, j: 50}
// 	myFunction(a)
// 	fmt.Printf("[%p] %v\n", a, a)
// }

// go 语言结构体内存布局 之 简化代码
type MyStruct struct {
	i int
	j int
}

func myFunction(ms *MyStruct) *MyStruct {
	return ms
}
