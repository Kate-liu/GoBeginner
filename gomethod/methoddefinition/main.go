package main

// T 自定义方法
type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(a int) int {
	t.a = a
	return t.a
}

// Get 类型T的方法Get的等价函数
func Get(t T) int {
	return t.a
}

// Set 类型*T的方法Set的等价函数
func Set(t *T, a int) int {
	t.a = a
	return t.a
}
