package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	error
}

var ErrBad = MyError{
	error: errors.New("bad things happened"),
}

func bad() bool {
	return false
}

func returnsError() error {
	// 有坑
	var p *MyError = nil
	if bad() {
		p = &ErrBad
	}
	return p
	// 方法一
	// var p error = nil
	// if bad() {
	// 	p = &ErrBad
	// }
	// return p
	// 方法二
	// if bad() {
	// 	return &ErrBad
	// }
	// return nil
}

func main() {
	err := returnsError()
	if err != nil {
		fmt.Printf("error occur: %+v\n", err)
		return
	}
	fmt.Println("ok")
}
