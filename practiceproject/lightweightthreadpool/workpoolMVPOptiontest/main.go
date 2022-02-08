package main

import (
	"fmt"
	"github.com/Kate-liu/GoBeginner/practiceproject/lightweightthreadpool/workpoolMVPOption"
	"time"
)

func main() {
	p := workpoolMVPOption.New(5, workpoolMVPOption.WithPreAllocWorkers(false), workpoolMVPOption.WithBlock(false))

	time.Sleep(time.Second * 2)
	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			fmt.Printf("task[%d]: error: %s\n", i, err.Error())
		}
	}

	p.Free()
}
