package main

func main() {
	// for range 操作 channel 类型变量
	var c = make(chan int)
	for v := range c {
		// ...
	}
}
