package main

import (
	"fmt"
	"reflect"
)

func main() {
	author := "draven"
	fmt.Println("TypeOf author:", reflect.TypeOf(author))   // 类型
	fmt.Println("ValueOf author:", reflect.ValueOf(author)) // 值

	v := reflect.ValueOf(1)
	fmt.Println(v)
	fmt.Println(v.Interface().(int)) // 显示转换

	// i1 := 1
	// v1 := reflect.ValueOf(i1)
	// v1.SetInt(10) // panic: reflect: reflect.Value.SetInt using unaddressable value
	// fmt.Println(i1)

	i2 := 1
	v2 := reflect.ValueOf(&i2) // 借用指针进行值的更改
	v2.Elem().SetInt(10)
	fmt.Println(i2)

	// 迂回的方式进行值的更改，等效意图
	i3 := 1
	v3 := &i3
	*v3 = 10
	fmt.Println(i3)
}
