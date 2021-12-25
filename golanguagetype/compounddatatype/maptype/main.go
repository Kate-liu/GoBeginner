package main

import "fmt"

func main() {
	// // == 的比较操作
	// s1 := make([]int, 1)
	// s2 := make([]int, 2)
	// f1 := func() {}
	// f2 := func() {}
	// m1 := make(map[int]string)
	// m2 := make(map[int]string)
	//
	// println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
	// println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
	// println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)

	// // map 的声明
	// var m map[string]int // 一个map[string]int 类型的变量
	// m["key"] = 1         // panic: assignment to entry in nil map
	// fmt.Println(m)       // map[]

	// 复合字面值初始化 map 类型变量
	n := map[int]string{}
	n[1] = "liu"
	fmt.Println(n) // map[1:liu]

	// 复杂字面值初始化
	m1 := map[int][]string{
		1: []string{"val1_1", "val1_2"},
		3: []string{"val3_1", "val3_2", "val3_3"},
		7: []string{"val7_1"},
	}

	type Position struct {
		x float64
		y float64
	}

	m2 := map[Position]string{
		Position{29.935523, 52.568915}:  "school",
		Position{25.352594, 113.304361}: "shopping-mall",
		Position{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m1) // map[1:[val1_1 val1_2] 3:[val3_1 val3_2 val3_3] 7:[val7_1]]
	fmt.Println(m2) // map[{25.352594 113.304361}:shopping-mall {29.935523 52.568915}:school {73.224455 111.804306}:hospital]

	// 省略字面值中的元素类型
	m3 := map[Position]string{
		{29.935523, 52.568915}:  "school",
		{25.352594, 113.304361}: "shopping-mall",
		{73.224455, 111.804306}: "hospital",
	}
	fmt.Println(m3) // map[{25.352594 113.304361}:shopping-mall {29.935523 52.568915}:school {73.224455 111.804306}:hospital]

	// make 初始化
	m4 := make(map[int]string)    // 未指定初始容量
	m5 := make(map[int]string, 8) // 指定初始容量为8
	fmt.Println(m4)
	fmt.Println(m5)

}
