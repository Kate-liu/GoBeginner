package main

import (
	"fmt"
)

func main() {
	// go 对字符串的支持
	const (
		GO_SLOGAN = "less is more"  // GO_SLOGAN是string类型常量
		s1        = "hello, gopher" // s1是string类型常量
	)
	var s2 = "I love go" // s2是string类型变量

	// string 类型的数据是不可变的
	var s string = "hello"
	s[0] = 'k'   // 错误： 字符串的内容是不可改变的(cannot assign to s[0] (strings are immutable))
	s = "gopher" // ok

	// string 类型的数据所见即所得
	var s string = `         
 			   ,_---~~~~~----._
		_,,_,*^____      _____*g*\"*,--,
	   / __/ /'     ^.  /      \ ^@q   f
	  [  @f | @))    |  | @))   l  0 _/
	   \/   \~____ / __ \_____/     \
		|           _l__l_           I
		}          [______]          I
		]            | | |           |
		]             ~ ~            |
		|                            |
		|                            |`

	fmt.Println(s)

}
