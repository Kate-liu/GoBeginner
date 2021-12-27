package main

func main() {
	// if 语句的自用变量
	if a, c := f(), h(); a > 0 {
		println(a)
	} else if b := f(); b > 0 {
		println(a, b)
	} else {
		println(a, b, c)
	}

}
