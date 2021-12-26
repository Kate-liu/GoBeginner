package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	// 零值可用 之 sync.Mutex
	var mu sync.Mutex
	mu.Lock()
	// ... ...
	mu.Unlock()

	// 零值可用 之 bytes.Buffer 结构体类型
	var b bytes.Buffer

	b.Write([]byte("Hello, Go"))
	fmt.Println(b.String()) // 输出：Hello, Go

}
