package main

// if 分支判断
func readByExt(ext string) {
	if ext == "json" {
		println("read json file")
	} else if ext == "jpg" || ext == "jpeg" || ext == "png" || ext == "gif" {
		println("read image file")
	} else if ext == "txt" || ext == "md" {
		println("read text file")
	} else if ext == "yml" || ext == "yaml" {
		println("read yaml file")
	} else if ext == "ini" {
		println("read ini file")
	} else {
		println("unsupported file extension:", ext)
	}
}

// switch 语句
func readByExtBySwitch(ext string) {
	switch ext {
	case "json":
		println("read json file")
	case "jpg", "jpeg", "png", "gif":
		println("read image file")
	case "txt", "md":
		println("read text file")
	case "yml", "yaml":
		println("read yaml file")
	case "ini":
		println("read ini file")
	default:
		println("unsupported file extension:", ext)
	}
}

func main() {
	// switch 语句一般形式
	switch initStmt; expr {
	case expr1:
		// 执行分支1
	case expr2:
		// 执行分支2
	case expr3_1, expr3_2, expr3_3:
		// 执行分支3
	case expr4:
		// 执行分支4
		// ... ...
	case exprN:
		// 执行分支N
	default:
		// 执行默认分支
	}

	// switch 表达式的类型为布尔类型, 省略 switch 后面的表达式
	// 带有initStmt语句的switch语句
	switch initStmt; {
	case bool_expr1:
	case bool_expr2:
		// ... ...
	}

	// 没有initStmt语句的switch语句
	switch {
	case bool_expr1:
	case bool_expr2:
		// ... ...
	}

}
