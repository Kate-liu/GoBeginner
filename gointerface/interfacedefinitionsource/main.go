package main

// 接口的声明
// type error interface {
// 	Error() string
// }
//
// type RPCError struct {
// 	Code    int64
// 	Message string
// }
//
// func (e *RPCError) Error() string {
// 	return fmt.Sprintf("%s, code=%d", e.Message, e.Code)
// }
//
// func main() {
// 	var rpcErr error = NewRPCError(400, "unknown err") // typecheck1
// 	err := AsErr(rpcErr)                               // typecheck3
// 	println(err)
// }
//
// func NewRPCError(code int64, msg string) error {
// 	return &RPCError{ // typecheck3
// 		Code:    code,
// 		Message: msg,
// 	}
// }
//
// func AsErr(err error) error {
// 	return err
// }

// 接口的 interface{} 类型
// func main() {
// 	type Test struct{}
// 	v := Test{}
// 	Print(v)
// }
//
// func Print(v interface{}) {
// 	println(v)
// }

// 方法的接受者是结构体，而初始化的变量是结构体指针/结构体(报错)
// type Duck interface {
// 	Quack()
// }
//
// type Cat struct{}
//
// func (c *Cat) Quack() {
// 	fmt.Println("meow")
// }
//
// func main() {
// 	var c Duck = &Cat{}
// 	// var c Duck = Cat{}
// 	c.Quack()
// }

// nil 和 non-nil
// type TestStruct struct{}
//
// func NilOrNot(v interface{}) bool {
// 	return v == nil
// }
//
// func main() {
// 	var s *TestStruct
// 	fmt.Println(s == nil)    // #=> true
// 	fmt.Println(NilOrNot(s)) // #=> false
// }

type Duck interface {
	Quack()
}

type Cat struct {
	Name string
}

//go:noinline
func (c Cat) Quack() {
	// println(c.Name + " meow")
}

func main() {
	// 类型转换
	// var c Duck = &Cat{Name: "draven"}
	// c.Quack()

	// 类型断言 之 非空接口
	// var c Duck = &Cat{Name: "draven"}
	// switch c.(type) {
	// case *Cat:
	// 	cat := c.(*Cat)
	// 	cat.Quack()
	// }

	// 类型断言 之 空接口
	// var c interface{} = &Cat{Name: "draven"}
	// switch c.(type) {
	// case *Cat:
	// 	cat := c.(*Cat)
	// 	cat.Quack()
	// }

	// 动态派发
	var c Duck = &Cat{Name: "draven"}
	c.Quack()
	c.(*Cat).Quack()
}
