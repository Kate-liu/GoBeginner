package main

func main() {
	// for 语句
	var sum int
	for i := 0; i < 10; i++ {
		sum += i
	}
	println(sum) // 45

	// 声明多循环变量
	for i, j, k := 0, 1, 2; (i < 20) && (j < 10) && (k < 30); i, j, k = i+1, j+1, k+5 {
		sum += (i + j + k)
		println(sum)
	}

	// 省略后置循环语句
	for i := 0; i < 10; {
		i++
	}

	// 省略循环前置语句
	i := 0
	for ; i < 10; i++ {
		println(i)
	}

	// // 省略后置与前置语句
	// i := 0
	// for ; i < 10; {
	// 	println(i)
	// 	i++
	// }

	// 省略经典 for 循环形式中的分号
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// 省略循环判断条件表达式
	for {
		// 循环体代码
		i += 6
	}
	// 等价形式（无限循环）
	for true {
		// 循环体代码
		i += 1
	}
	// // 等价形式（无限循环）
	// for ; ; {
	// 	// 循环体代码
	// 	i += 1
	// }

}
