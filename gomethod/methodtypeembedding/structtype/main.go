package main

// 一般结构体类型定义

type S struct {
	A int
	b string
	c T
	p *P
	_ [10]int8
	F func()
}

// 带有 嵌入字段（Embedded Field）的结构体定义

type T1 int

type t2 struct {
	n int
	m int
}

type I interface {
	M1()
}

type S1 struct {
	T1
	*t2
	I
	a int
	b string
}
