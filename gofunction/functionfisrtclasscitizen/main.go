package main

// // Go 函数可以存储在变量中
// var (
// 	myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
// 		return fmt.Fprintf(w, format, a...)
// 	}
// )
//
// func main() {
// 	fmt.Printf("%T\n", myFprintf)             // func(io.Writer, string, ...interface {}) (int, error)
// 	myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出 Hello，Go
// }

func setup(task string) func() {
	println("do some setup stuff for", task)
	return func() {
		println("do some teardown stuff for", task)
	}
}

func main() {
	teardown := setup("demo")
	defer teardown()
	println("do some bussiness stuff")

}
