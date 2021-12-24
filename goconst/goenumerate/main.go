package main

import "fmt"

// // 隐式重复前一个非空表达式
// const (
// 	Apple, Banana = 11, 22
// 	Strawberry, Grape
// 	Pear, Watermelon
// )

// // 隐式重复前一个非空表达式
// const (
// 	Apple, Banana     = 11, 22
// 	Strawberry, Grape = 11, 22 // 使用上一行的初始化表达式
// 	Pear, Watermelon  = 11, 22 // 使用上一行的初始化表达式
// )

// 测试 iota 的值
const (
	Apple, Banana     = iota, iota + 10 // 0, 10 (iota = 0)
	Strawberry, Grape                   // 1, 11 (iota = 1)
	Pear, Watermelon                    // 2, 12 (iota = 2)
)

// 空白标识符
const (
	_ = iota // 0
	Pin1
	Pin2
	Pin3
	_
	Pin5 // 5
)

// // 首字母排序的颜色变量
// const (
// 	Black  = 1
// 	Red    = 2
// 	Yellow = 3
// )

// // 首字母排序的颜色变量 增加新颜色 Blue
// const (
// 	Blue   = 1
// 	Black  = 2
// 	Red    = 3
// 	Yellow = 4
// )

// 首字母排序的颜色变量 使用iota
const (
	_ = iota
	Blue
	Red
	Yellow
)

const (
	a = iota + 1 // 1, iota = 0
	b            // 2, iota = 1
	c            // 3, iota = 2
)
const (
	i = iota << 1 // 0, iota = 0
	j             // 2, iota = 1
	k             // 4, iota = 2
)

func main() {
	fmt.Println(Apple)
	fmt.Println(Banana)
	fmt.Println(Strawberry)
	fmt.Println(Grape)
	fmt.Println(Pear)
	fmt.Println(Watermelon)

	fmt.Println(Pin1)
	fmt.Println(Pin2)
	fmt.Println(Pin3)
	fmt.Println(Pin5)

	fmt.Println(Blue)
	fmt.Println(Red)
	fmt.Println(Yellow)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}
