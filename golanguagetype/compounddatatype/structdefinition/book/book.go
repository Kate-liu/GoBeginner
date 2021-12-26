package book

type Book struct {
	Title   string         // 书名
	Pages   int            // 书的页数
	Indexes map[string]int // 书的索引
	Author  Person         // 作者
	// ...
}

// // Book 简便的定义：
// type Book struct {
// 	Title string
// 	Person
// 	// ... ...
// }

type Person struct {
	Name  string
	Phone string
	Addr  string
}
