package main

func case1() int {
	println("eval case1 expr")
	return 1
}

func case2() int {
	println("eval case2 expr")
	return 2
}

func switchexpr() int {
	println("eval switch expr")
	return 1
}

func main() {
	// 使用 fallthrough 的 switch 语句
	switch switchexpr() {
	case case1():
		println("exec case1")
		fallthrough
	case case2():
		println("exec case2")
		fallthrough
	default:
		println("exec default")
	}
}
