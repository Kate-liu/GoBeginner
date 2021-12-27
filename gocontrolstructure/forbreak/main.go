package main

func main() {
	// for break 语句的使用
	var sl = []int{5, 19, 6, 3, 8, 12}
	var firstEven int = -1
	// 找出整型切片sl中的第一个偶数
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			firstEven = sl[i]
			break
		}
	}

	println(firstEven) // 6
}
