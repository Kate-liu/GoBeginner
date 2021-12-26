package main

import (
	"github.com/Kate-liu/GoBeginner/golanguagetype/compounddatatype/structdefinition/book"
	"unsafe"
)

func main() {
	// 抽象现实中的书
	var b book.Book

	b.Title = "The Go Programming Language"
	b.Pages = 678

	// 定义空结构体
	type Empty struct{}

	var s Empty
	println(unsafe.Sizeof(s)) // 0

	// // goroutine之间通信的事件信息
	// var c = make(chan Empty) // 声明一个元素类型为Empty的channel
	// c <- Empty{}             // 向channel写入一个“事件”

	// 访问 book 中的 author 字段中的 Phone 字段
	var newbook book.Book
	println(newbook.Author.Phone)

	// // 访问 book 中的 author 字段中的 Phone 字段 之 简便的定义方法
	// println(newbook.Person.Phone) // 将类型名当作嵌入字段的名字
	// println(newbook.Phone)        // 支持直接访问嵌入字段所属类型中字段

	// // 声明结构体
	// type Book struct {
	// 	// ...
	// }
	// var book Book     // 变量声明方式
	// var book = Book{} // 标准变量声明语句
	// book := Book{}    // 短变量声明语句

}
