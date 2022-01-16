package main

import (
	"time"
)

// 2）关心是否存在错误的任务
// func taskselect() error {
// 	errCh := make(chan error, len(tasks))
// 	wg := sync.WaitGroup{}
// 	wg.Add(len(tasks))
// 	for i := range tasks {
// 		go func() {
// 			defer wg.Done()
// 			if err := tasks[i].Run(); err != nil {
// 				errCh <- err
// 			}
// 		}()
// 	}
// 	wg.Wait()
//
// 	select {
// 	case err := <-errCh:
// 		return err
// 	default:
// 		return nil
// 	}
// }

func main() {
	// 1）直接执行 default
	// ch := make(chan int)
	//
	// select {
	// case i := <-ch:
	// 	println(i)
	// default:
	// 	println("default")
	// }

	// 3）多个 case 就绪时，随机选择一个执行
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		}
	}
}
