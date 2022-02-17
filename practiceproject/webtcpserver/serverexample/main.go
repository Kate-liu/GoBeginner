package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// 一般处理逻辑
// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		// read from the connection
// 		// ... ...
// 		// write to the connection
// 		// ... ...
// 	}
// }

// // 使用 SetReadDeadline 的处理逻辑
// func handleConn(c net.Conn) {
// 	defer c.Close()
// 	for {
// 		// read from the connection
// 		var buf = make([]byte, 128)
// 		c.SetReadDeadline(time.Now().Add(time.Second))
// 		n, err := c.Read(buf)
// 		if err != nil {
// 			log.Printf("conn read %d bytes, error: %s", n, err)
// 			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
// 				// 进行其他业务逻辑的处理
// 				continue
// 			}
// 			return
// 		}
// 		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
// 	}
// }

// socket 写操作阻塞的处理逻辑
func handleConn(c net.Conn) {
	defer c.Close()
	time.Sleep(time.Second * 10)
	for {
		// read from the connection
		time.Sleep(5 * time.Second)
		var buf = make([]byte, 60000)
		log.Println("start to read from conn")
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes, error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				continue
			}
		}

		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
