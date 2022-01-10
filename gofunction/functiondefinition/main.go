package main

// func myAppend(sl []int, elems ...int) []int {
// 	fmt.Printf("%T\n", elems) // []int
// 	if len(elems) == 0 {
// 		println("no elems to append")
// 		return sl
// 	}
//
// 	sl = append(sl, elems...)
// 	return sl
// }
//
// func main() {
// 	sl := []int{1, 2, 3}
// 	sl = myAppend(sl) // no elems to append
// 	fmt.Println(sl)   // [1 2 3]
// 	sl = myAppend(sl, 4, 5, 6)
// 	fmt.Println(sl) // [1 2 3 4 5 6]
// }

// go 语言的调用惯例
// go tool compile -S -N -l main.go
func myFunction(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	myFunction(66, 77)
}
