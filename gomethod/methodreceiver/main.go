package main

import "fmt"

// // 参数名与方法参数的形参名重复
// type T struct{}
//
// func (t T) M(t string) { // 编译器报错：duplicate argument t (重复声明参数t)
// 	// ... ...
// }

// // 省略 receiver 的参数 名
// type T struct{}
//
// func (T) M(t string) {
// 	// ... ...
// }

// // receiver 参数的基类型为指针类型
// type MyInt *int
//
// func (r MyInt) String() string { // r的基类型为MyInt，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
// 	return fmt.Sprintf("%d", *(*int)(r))
// }
//
// // receiver 参数的基类型为接口类型
// type MyReader io.Reader
//
// func (r MyReader) Read(p []byte) (int, error) { // r的基类型为MyReader，编译器报错：invalid receiver type MyReader (MyReader is an interface type)
// 	return r.Read(p)
// }

// // 不能为原生类型添加方法
// func (i int) Foo() string { // 编译器报错：cannot define new methods on non-local type int
// 	return fmt.Sprintf("%d", i)
// }

// // 不能跨越Go包为其他包的类型声明新方法
// func (s http.Server) Foo() { // 编译器报错：cannot define new methods on non-local type http.Server
//
// }

// 使用 Go 方法
type T struct{}

func (t T) M(n int) {
	fmt.Println(n)
}

func main() {
	var t T
	t.M(1) // 通过类型T的变量实例调用方法M

	p := &T{}
	p.M(2) // 通过类型*T的变量实例调用方法M
}
