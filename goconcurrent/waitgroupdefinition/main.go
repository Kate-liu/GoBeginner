package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Request http.Request

func main() {
	// // WaitGroup 使用场景
	// // 批量发出 Request 请求
	// requests := []*Request{...}
	// wg := &sync.WaitGroup{}
	// wg.Add(len(requests))
	//
	// for _, request := range requests {
	// 	go func(r *Request) {
	// 		defer wg.Done()
	// 		// 发送请求并处理
	// 		// res, err := service.call(r)
	// 	}(request)
	// }
	// wg.Wait()

	wg := sync.WaitGroup{}
	yawg := wg
	fmt.Println(wg, yawg)
}
