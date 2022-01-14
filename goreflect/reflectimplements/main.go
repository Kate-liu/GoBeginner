package main

import (
	"fmt"
	"reflect"
)

type CustomError struct{}

func (*CustomError) Error() string {
	return ""
}

func main() {
	// 判断一个类型是否实现了某个接口
	typeOfError := reflect.TypeOf((*error)(nil)).Elem()
	customErrorPtr := reflect.TypeOf(&CustomError{})
	customError := reflect.TypeOf(CustomError{})

	fmt.Println(customErrorPtr.Implements(typeOfError)) // #=> true
	fmt.Println(customError.Implements(typeOfError))    // #=> false
}
