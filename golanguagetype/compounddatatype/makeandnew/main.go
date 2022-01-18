package main

func main() {
	// make
	slice := make([]int, 0, 100)
	hash := make(map[int]bool, 10)
	ch := make(chan int, 5)

	// new
	i := new(int)

	var v int // 等价于 new 初始化
	i := &v
}
