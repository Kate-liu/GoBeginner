package main

import "fmt"

func main() {
	/*	// error example
		var a int16 = 5

		var b int = 8
		var c int64

		c = a + b

		fmt.Printf("%d\n", c)*/

	// correct example
	var a int16 = 5

	var b int = 8
	var c int64

	c = int64(a) + int64(b)

	fmt.Printf("%d\n", c)
}
