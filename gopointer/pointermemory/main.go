package main

type foo struct {
	id   string
	age  int8
	addr string
}

func main() {
	// var p1 *int
	// var p2 *bool
	// var p3 *byte
	// var p4 *[20]int
	// var p5 *foo
	// var p6 unsafe.Pointer
	//
	// println(unsafe.Sizeof(p1)) // 8
	// println(unsafe.Sizeof(p2)) // 8
	// println(unsafe.Sizeof(p3)) // 8
	// println(unsafe.Sizeof(p4)) // 8
	// println(unsafe.Sizeof(p5)) // 8
	// println(unsafe.Sizeof(p6)) // 8

	// var a int = 17
	// var p *int = &a
	// println(*p) // 17
	//
	// (*p) += 3
	// println(a) // 20
	//
	// fmt.Printf("%p\n", p) // 0xc0000160d8

	// var a int = 5
	// var b int = 6
	//
	// var p *int = &a // 指向变量a所在内存单元
	// println(*p)     // 输出变量a的值
	// p = &b          // 指向变量b所在内存单元
	// println(*p)     // 输出变量b的值

	var a int = 5
	var p1 *int = &a // p1指向变量a所在内存单元
	var p2 *int = &a // p2指向变量a所在内存单元
	(*p1) += 5       // 通过p1修改变量a的值
	println(*p2)     // 10 对变量a的修改可以通过另外一个指针变量p2的解引用反映出来
}
