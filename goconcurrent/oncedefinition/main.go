package main

import (
	"fmt"
	"sync"
)

func main() {
	// Once 示例
	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}
}
