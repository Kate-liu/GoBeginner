package main

import (
	"errors"
	"fmt"
)

var a int = 2020

func checkYear() error {
	err := errors.New("wrong year")

	switch a, err := getYear(); a {
	case 2020:
		fmt.Println("it is", a, err)
	case 2021:
		fmt.Println("it is", a)
		err = nil
		// 修正，添加下面这句代码
		// return err
	}

	fmt.Println("after check, it is", a)
	return err
}

type new int

func getYear() (new, error) {
	var b int16 = 2021
	return new(b), nil
}

func main() {
	err := checkYear()
	if err != nil {
		fmt.Println("call checkYear error:", err)
		return
	}

	fmt.Println("call checkYear ok")

	// 遮蔽预定义标识符 new 的检验
	// p := new(int)
	// *p = 11
}
