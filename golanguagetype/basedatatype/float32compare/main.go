package main

import (
	"fmt"
	"math"
)

func main() {
	var f1 float32 = 16777216.0
	var f2 float32 = 16777217.0

	fmt.Println(f1 == f2) // true

	// 验证为啥不是 false
	bitsf1 := math.Float32bits(f1)
	bitsf2 := math.Float32bits(f2)
	fmt.Printf("%b\n", bitsf1) // 10010111_00000000000000000000000
	fmt.Printf("%b\n", bitsf2) // 10010111_00000000000000000000000
}
