package main

func main() {
	// continue 语句的使用方法 + label
	var sum int
	var sl = []int{1, 2, 3, 4, 5, 6}
loop:
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 {
			// 忽略切片中值为偶数的元素
			continue loop
		}
		sum += sl[i]
	}
	println(sum) // 9
}
