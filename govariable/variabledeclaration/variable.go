package main

import "fmt"

func main() {
	// 通用变量声明
	var a int = 10
	fmt.Println(a)

	// 默认赋予 int 类型的零值
	var b int
	fmt.Println(b)

	// 变量声明块
	var (
		c int    = 128
		d int8   = 6
		e string = "hello"
		f rune   = 'A'
		g bool   = true
	)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	// 一行变量声明实现方式
	var h, i, j int = 6, 7, 8
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)

	// 变量声明块(一行变量声明实现方式)
	var (
		k, l, m int  = 7, 8, 9
		n, o, p rune = 'B', 'C', 'D'
	)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)

	// 省略类型信息的声明
	var q = 13
	fmt.Println(q)

	// 显示类型转型
	var r = int32(61)
	fmt.Println(r)

	// 没有初值的声明（不被允许）
	// var s

	// 多变量声明 + 省略类型信息的声明
	var t, u, v = 23, 'A', "world"
	fmt.Println(t)
	fmt.Println(u)
	fmt.Println(v)

	// 短变量声明
	w := 45
	x := 'S'
	y := "liu"
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)

	// 短变量声明 + 一次声明多个变量
	aa, bb, cc := 22, 'M', "ming"
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	// 变量声明块 + 显示类型转型 (推荐使用)
	var (
		dd = 13
		ee = int32(34)
		ff = float32(3.24)
	)
	fmt.Println(dd)
	fmt.Println(ee)
	fmt.Println(ff)

}
