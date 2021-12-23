package main

func main() {
	// 复数字面值直接初始化一个复数类型变量
	var c = 5 + 6i
	var d = 0o123 + .12345e+5i // 83+12345i

	// 使用 complex 函数
	var c = complex(5, 6)             // 5 + 6i
	var d = complex(0o123, .12345e+5) // 83+12345i

	// 使用预定义的函数 real 和 imag
	var c = complex(5, 6) // 5 + 6i
	r := real(c)          // 5.000000
	i := imag(c)          // 6.000000
}
