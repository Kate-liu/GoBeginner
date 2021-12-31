package main

// // 多个嵌入字段的方法集合中都包含这个方法，有交集
//
// type E1 interface {
// 	M1()
// 	M2()
// 	M3()
// }
//
// type E2 interface {
// 	M1()
// 	M2()
// 	M4()
// }
//
// type T struct {
// 	E1
// 	E2
// }
//
// func main() {
// 	t := T{}
// 	t.M1()
// 	t.M2()
// }

// 多个嵌入字段的方法集合中都包含这个方法，有交集 解决方案

type E1 interface {
	M1()
	M2()
	M3()
}

type E2 interface {
	M1()
	M2()
	M4()
}

type T struct {
	E1
	E2
}

func (T) M1() { println("T's M1") }
func (T) M2() { println("T's M2") }

func main() {
	t := T{}
	t.M1() // T's M1
	t.M2() // T's M2
}
