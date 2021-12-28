package main

type person struct {
	name string
	age  int
}

func main() {
	// 自定义结构体类型作为 switch 表达式类型的例子
	p := person{"tom", 13}
	switch p {
	case person{"tony", 33}:
		println("match tony")
	case person{"tom", 13}:
		println("match tom")
	case person{"lucy", 23}:
		println("match lucy")
	default:
		println("no match")
	}
}
