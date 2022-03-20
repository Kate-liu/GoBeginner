package main

// func main() {
// 	var a int = 5
// 	var p1 *int = &a
// 	println(*p1) // 5
//
// 	var b int = 55
// 	var p2 *int = &b
// 	println(*p2) // 55
//
// 	var pp **int = &p1
// 	println(**pp)            // 5
//
// 	println((*pp) == p1)     // true
// 	println((**pp) == (*p1)) // true
// 	println((**pp) == a)     // true
//
// 	pp = &p2
// 	println(**pp) // 55
//
// }

func foo(pp **int) {
	var b int = 55
	var p1 *int = &b
	(*pp) = p1
}

func main() {
	var a int = 5
	var p *int = &a
	println(*p) // 5

	foo(&p)
	println(*p) // 55
}
