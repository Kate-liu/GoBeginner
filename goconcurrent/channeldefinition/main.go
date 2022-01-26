package main

func main() {
	// // 定义 channel
	// var ch chan int
	//
	// ch1 := make(chan int)
	// ch2 := make(chan int, 5)
	//
	// // 发送 和 接受
	// ch1 <- 13  // 将整型字面值13发送到无缓冲channel类型变量ch1中
	// n := <-ch1 // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
	// ch2 <- 17  // 将整型字面值17发送到带缓冲channel类型变量ch2中
	// m := <-ch2 // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中

	// // 无缓冲 channel 类型
	// ch1 := make(chan int)
	// ch1 <- 13 // fatal error: all goroutines are asleep - deadlock!
	// n := <-ch1
	// println(n)

	// // 无缓冲 channel 类型 - 改进
	// ch1 := make(chan int)
	// go func() {
	// 	ch1 <- 13 // 将发送操作放入一个新goroutine中执行
	// }()
	// n := <-ch1
	// println(n)

	// // 带缓冲 channel
	// ch2 := make(chan int, 1)
	// n := <-ch2 // 由于此时ch2的缓冲区中无数据，因此对其进行接收操作将导致goroutine挂起
	// println(n)
	// ch3 := make(chan int, 1)
	// ch3 <- 17 // 向ch3发送一个整型数17
	// ch3 <- 27 // 由于此时ch3中缓冲区已满，再向ch3发送数据也将导致goroutine挂起

	// // 设置发送与接受类型
	// ch1 := make(chan<- int, 1) // 只发送channel类型
	// ch2 := make(<-chan int, 1) // 只接收channel类型
	// <-ch1                      // invalid operation: <-ch1 (receive from send-only type chan<- in
	// ch2 <- 13                  // invalid operation: ch2 <- 13 (send to receive-only type <-chan

	// // 关闭 channel
	// ch := make(chan int, 5)
	// close(ch)
	// ch <- 13 // panic: send on closed channel

	// // channel 与 select
	// select {
	// case x := <-ch1: // 从channel ch1接收数据
	// 	... ...
	// case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
	// 	... ...
	// case ch3 <- z: // 将z值发送到channel ch3中:
	// 	... ...
	// default: // 当上面case中的channel通信均无法实施时，执行该默认分支
	// }
}
