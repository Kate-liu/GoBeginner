package main

func main() {
	// type 关键字自定义
	type MyInt int32

	// 错误
	var m int = 5
	var n int32 = 6
	var a MyInt = m // 错误：在赋值中不能将m（int类型）作为MyInt类型使用
	var a MyInt = n // 错误：在赋值中不能将n（int32类型）作为MyInt类型使用

	// 正确，显式转型
	var m int = 5
	var n int32 = 6
	var a MyInt = MyInt(m) // ok
	var a MyInt = MyInt(n) // ok

	// 类型别名语法自定义
	type MyInt = int32

	var n int32 = 6
	var a MyInt = n // ok

}
