package main

import (
	"errors"
	"fmt"
)

func main() {
	// 静态特性
	// var err error = 1 // cannot use 1 (type int) as type error in assignment: int does not implement error (missing Error method)

	// 动态特性
	var err error
	err = errors.New("error1")
	fmt.Printf("%T\n", err) // *errors.errorString
}
