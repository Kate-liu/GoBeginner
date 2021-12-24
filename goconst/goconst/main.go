package main

import "fmt"

// Pi 常量的声明
const Pi float64 = 3.14159265358979323846 // 单行常量声明
// 以const代码块形式声明常量
const (
	size    int64 = 4096
	i, j, s       = 13, 14, "bar" // 单行声明多个常量
)

// 无类型常量 类型安全要求
type myInt int

const n myInt = 13
const m int = n + 5 // 编译器报错：cannot use n + 5 (type myInt) as type int in const initializer

func main() {
	var a int = 5
	fmt.Println(a + n) // 编译器报错：invalid operation: a + n (mismatched types int and myInt)
}

// 无类型常量 显示转型
type myInt int

const n myInt = 13
const m int = int(n) + 66 // OK

func main() {
	var a int = 5
	fmt.Println(a + int(n)) // 输出：18
	fmt.Println(m)          // 输出：79
}

// 无类型常量 无类型常量（Untyped Constant）
type myInt int

const n = 13

func main() {
	var a myInt = 5
	fmt.Println(a + n) // 输出：18
}

// 隐式转型 无法转换为目标类型
func main() {
	const m = 1333333333

	var k int8 = 1
	j := k + m // 编译器报错：constant 1333333333 overflows int8
}
