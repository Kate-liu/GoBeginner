package main

func (t T) M1(x int) (err error) {
	// 代码块1
	m := 13
	// 代码块1 是包含 m、t、x 和 err 四个标识符的最内部代码块
	{
		// 代码块2
		// 代码块2 是包含类型 bar 标识符的最内部的那个包含代码块
		type bar struct{} // 类型标识符 bar 的作用域始于此
		{
			// 代码块3
			// 代码块3 是包含变量 a 标识符的最内部的那个包含代码块
			a := 5 // a 作用域开始于此
			{
				// 代码块4
				// ...
			}
			// a 作用域终止于此
		}
		// 类型标识符 bar 的作用域终止于此
	}
	// m、t、x 和 err 的作用域终止于此
}

func bar() {
	if a := 1; false {

	} else if b := 2; false {

	} else if c := 3; false {

	} else {
		println(a, b, c)
	}
}

func bar() {
	{ // 等价于第一个 if 的隐式代码块
		a := 1 // 变量 a 作用域始于此
		if false {

		} else {
			{ // 等价于第一个 else if 的隐式代码块
				b := 2 // 变量 b 作用域始于此
				if false {

				} else {
					{ // 等价于第二个 else if 的隐式代码块
						c := 3 // 变量 c 作用域始于此
						if false {

						} else {
							println(a, b, c)
						}
						// 变量 c 作用域终止于此
					}
				}
				// 变量 b 作用域终止于此
			}
		}
		// 变量 a 作用域终止于此
	}
}

func bar() {
	{ // 等价于第一个 if 的隐式代码块
		a := 1 // 变量 a 作用域始于此
		if false {

		} else {
			{ // 等价于第一个 else if 的隐式代码块
				b := 2 // 变量 b 的作用域始于此
				if false {

				} else {
					{ // 等价于第二个 else if 的隐式代码块
						c := 3 // 变量 c 作用域始于此
						if false {

						} else {
							println(a, b, c)
						}
						// 变量 c 的作用域终止于此
					}
				}
				// 变量 b 的作用域终止于此
			}
		}
		// 变量 a 作用域终止于此
	}
}
