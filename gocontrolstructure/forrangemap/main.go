package main

func main() {
	// for range 操作 map 类型变量
	var m = map[string]int{
		"Rob":  67,
		"Russ": 39,
		"John": 29,
	}

	for k, v := range m {
		println(k, v)
	}
}
