package main

import "fmt"

func foo(m map[string]int) {
	m["key1"] = 11
	m["key2"] = 12
}

func main() {
	m := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	fmt.Println(m) // map[key1:1 key2:2]
	foo(m)
	fmt.Println(m) // map[key1:11 key2:12]
}
