package main

import (
	"fmt"
	"github.com/Kate-liu/GoBeginner/golanguagetype/compounddatatype/structcompoundliteral/typet"
	"unsafe"
)

type Book struct {
	Title   string         // 书名
	Pages   int            // 书的页数
	Indexes map[string]int // 书的索引
}

func main() {
	// 按字段顺序初始化
	var book = Book{"The Go Programming Language", 700, make(map[string]int)}
	fmt.Println(book)

	// // 按字段顺序初始化 之 包含非导出字段
	// var t1 = typet.T{11, "hello", 13}
	// // 错误：implicit assignment of unexported field 'f3' in typet.T literal
	// // too few values in T{...}
	// // 或
	// var t2 = typet.T{11, "hello", 13, 14, 15}
	// // 错误：implicit assignment of unexported field 'f3' in typet.T literal
	//
	// fmt.Println(t1)
	// fmt.Println(t2.f3) // t2.f3 undefined (cannot refer to unexported field or method f3)

	// field：value 初始化
	var t = typet.T{
		F2: "hello",
		F1: 11,
		F4: 67,
	}
	fmt.Println(t) // {11 hello 0 67 0}

	t1 := typet.T{}
	fmt.Println(t1) // {0  0 0 0}

	tp := new(typet.T)
	fmt.Println(tp) // &{0  0 0 0}

	// 结构体内存布局 值 占用内存大小 与 字段在内存中的地址偏移量
	var t2 typet.T
	fmt.Println(unsafe.Sizeof(t2))      // 48
	fmt.Println(unsafe.Offsetof(t2.F2)) // 8
	fmt.Println(unsafe.Offsetof(t2.F4)) // 32
}
