package main

import "runtime"

func main() {
	// // if 分支表达式
	// if boolean_expression {
	// 	// 新分支
	// }
	// // 原分支

	// if 判断示例
	if runtime.GOOS == "darwin" {
		println("we are on darwin os")
	}

	// if 判断示例 逻辑操作符 + ()
	if (runtime.GOOS == "darwin") && (runtime.GOARCH == "amd64") && (runtime.Compiler != "gccgo") {
		println("we are using standard go compiler on darwin os for amd64")
	}

	// if 判断示例 逻辑操作符优先级 不加()
	a, b := false, true
	if a && b != true {
		println("(a && b) != true")
		return
	}
	println("a && (b != true) == false") // 输出：a && (b != true) == false

	// // 二分支结构
	// if boolean_expression {
	// 	// 分支1
	// } else {
	// 	// 分支2
	// }

	// 多分支结构
	if boolean_expression1 {
		// 分支1
	} else if boolean_expression2 {
		// 分支2
		// ... ...
	} else if boolean_expressionN {
		// 分支N
	} else {
		// 分支N+1
	}

	// 四分支结构
	if boolean_expression1 {
		// 分支1
	} else if boolean_expression2 {
		// 分支2
	} else if boolean_expression3 {
		// 分支3
	} else {
		// 分支4
	}

	// 四分支结构 等价变换
	if boolean_expression1 {
		// 分支1
	} else {
		if boolean_expression2 {
			// 分支2
		} else {
			if boolean_expression3 {
				// 分支3
			} else {
				// 分支4
			}
		}
	}

}
