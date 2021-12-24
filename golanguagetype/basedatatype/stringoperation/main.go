package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 字符串常见操作 之 下标操作
	var s = "中国人"
	fmt.Printf("0x%x\n", s[0]) // 0xe4：字符“中” utf-8 编码的第一个字节

	// 字符串常见操作 之 字符迭代 之 for 迭代
	var t = "中国人"
	for i := 0; i < len(t); i++ {
		fmt.Printf("index: %d, value: 0x%x\n", i, s[i])
	}

	// 字符串常见操作 之 字符迭代 之 for range 迭代
	var u = "中国人"
	for i, v := range u {
		fmt.Printf("index: %d, value: 0x%x\n", i, v)
	}

	count := utf8.RuneCountInString(u) // 获取字符串中字符个数，反查出使用的是什么编码方式
	fmt.Println(count)

	// 字符串常见操作 之 字符串连接
	v := "Rob Pike, "
	v = v + "Robert Griesemer, "
	v += " Ken Thompson"
	fmt.Println(v) // Rob Pike, Robert Griesemer, Ken Thompson

	// 字符串常见操作 之 字符串比较
	// ==
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2) // true
	// !=
	s1 = "Go"
	s2 = "C"
	fmt.Println(s1 != s2) // true
	// < and <=
	s1 = "12345"
	s2 = "23456"
	fmt.Println(s1 < s2)  // true
	fmt.Println(s1 <= s2) // true
	// > and >=
	s1 = "12345"
	s2 = "123"
	fmt.Println(s1 > s2)  // true
	fmt.Println(s1 >= s2) // true

	// 字符串常见操作 之 字符串转换
	var w string = "中国人"
	// string -> []rune
	rs := []rune(w)
	fmt.Printf("%x\n", rs) // [4e2d 56fd 4eba]
	// string -> []byte
	bs := []byte(w)
	fmt.Printf("%x\n", bs) // e4b8ade59bbde4baba
	// []rune -> string
	w1 := string(rs)
	fmt.Println(w1) // 中国人
	// []byte -> string
	w2 := string(bs)
	fmt.Println(w2) // 中国人
}
