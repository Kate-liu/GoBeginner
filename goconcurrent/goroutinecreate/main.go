package main

import (
	"fmt"
	"time"
)

func main() {
	go fmt.Println("I am a goroutine")

	var c = make(chan int)
	go func(a, b int) {
		c <- a + b
	}(3, 4)

	// $GOROOT/src/net/http/server.go
	c := srv.newConn(rw)
	go c.serve(connCtx)

	println(<-c)
	time.Sleep(2000 * time.Second)
}
