package main

import (
	"fmt"
	"time"
)

func deadloop() {
	for {
	}
}

func main() {
	go deadloop()
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("I got scheduled!")
	}
}

// 使用 Go 1.13.x 版本运行
// 1.在一个拥有多核处理器的主机上，使用 Go 1.13.x 版本运行这个示例代码，在命令行终端上是否能看到“I got scheduled!”输出呢？
// 也就是 main goroutine 在创建 deadloop goroutine 之后是否能继续得到调度呢？
// 2.通过什么方法可以让上面示例中的 main goroutine，在创建 deadloop goroutine 之后无法继续得到调度？
