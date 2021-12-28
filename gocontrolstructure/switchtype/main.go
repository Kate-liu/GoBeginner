package main

// func main() {
// 	// // type switch 用法 之 获得变量 x 的动态类型信息
// 	// var x interface{} = 13
// 	// switch x.(type) {
// 	// case nil:
// 	// 	println("x is nil")
// 	// case int:
// 	// 	println("the type of x is int")
// 	// case string:
// 	// 	println("the type of x is string")
// 	// case bool:
// 	// 	println("the type of x is string")
// 	// default:
// 	// 	println("don't support the type")
// 	// }
//
// 	// // type switch 用法 之 获得变量 x 的动态类型的值
// 	// var x interface{} = 13
// 	// switch v := x.(type) {
// 	// case nil:
// 	// 	println("v is nil")
// 	// case int:
// 	// 	println("the type of v is int, v =", v)
// 	// case string:
// 	// 	println("the type of v is string, v =", v)
// 	// case bool:
// 	// 	println("the type of v is bool, v =", v)
// 	// default:
// 	// 	println("don't support the type")
// 	// }
// }

type I interface {
	M()
}

type T struct {
}

func (T) M() {
}

func main() {
	// 使用特定的接口类型 I
	var t T
	var i I = t
	switch i.(type) {
	case T:
		println("it is type T")
	case int:
		println("it is type int")
	case string:
		println("it is type string")
	}
}
