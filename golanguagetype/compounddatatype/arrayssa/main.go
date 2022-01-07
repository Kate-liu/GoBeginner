package arrayssa

func outOfRange() int {
	// v1-越界
	arr := [3]int{1, 2, 3}
	i := 4
	elem := arr[i]
	return elem
	// v2-字面值整数
	// arr := [3]int{1, 2, 3}
	// elem := arr[2]
	// return elem
	// v3-复制操作
	// arr := [3]int{1, 2, 3}
	// arr[0] = 666
	// return arr[2]
}
