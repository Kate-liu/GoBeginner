package main

func main() {
	// for
	// for i := 0; i < 10; i++ {
	// 	println(i)
	// }

	// range
	// arr := []int{1, 2, 3}
	// for i, _ := range arr {
	// 	println(i)
	// }

	// 循环永动机
	// arr := []int{1, 2, 3}
	// for _, v := range arr {
	// 	arr = append(arr, v)
	// }
	// fmt.Println(arr)

	// 神奇的指针
	// arr := []int{1, 2, 3}
	// newArr := []*int{}
	// for i, v := range arr {
	// 	newArr = append(newArr, &v)      // wrong
	// 	newArr = append(newArr, &arr[i]) // right
	// }
	// for _, v := range newArr {
	// 	fmt.Println(*v)
	// }

	// 遍历清空数组
	// arr := []int{1, 2, 3}
	// for i, _ := range arr {
	// 	arr[i] = 0
	// }
	// println(arr)

	// 随机遍历
	hash := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	for k, v := range hash {
		println(k, v)
	}
}
