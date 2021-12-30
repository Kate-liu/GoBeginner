package main

import "fmt"

// T 自定义方法
type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

func main() {

	// var t T
	// t.Get()
	// t.Set(1)

	// var t T
	// T.Get(t)
	// (*T).Set(&t, 1)

	var t T
	f1 := (*T).Set                           // f1的类型，也是T类型Set方法的类型：func (t *T, int)int
	f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
	fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
	fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
	f1(&t, 3)
	fmt.Println(f2(t)) // 3
}
