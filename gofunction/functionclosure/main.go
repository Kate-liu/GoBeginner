package main

import "fmt"

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	// 简单方式
	times(2, 5) // 计算2 x 5
	times(3, 5) // 计算3 x 5
	times(4, 5) // 计算4 x 5

	// 高级方式
	timesTwo := partialTimes(2)   // 以高频乘数2为固定乘数的乘法函数
	timesThree := partialTimes(3) // 以高频乘数3为固定乘数的乘法函数
	timesFour := partialTimes(4)  // 以高频乘数4为固定乘数的乘法函数
	fmt.Println(timesTwo(5))      // 10，等价于times(2, 5)
	fmt.Println(timesTwo(6))      // 12，等价于times(2, 6)
	fmt.Println(timesThree(5))    // 15，等价于times(3, 5)
	fmt.Println(timesThree(6))    // 18，等价于times(3, 6)
	fmt.Println(timesFour(5))     // 20，等价于times(4, 5)
	fmt.Println(timesFour(6))     // 24，等价于times(4, 6)
}
