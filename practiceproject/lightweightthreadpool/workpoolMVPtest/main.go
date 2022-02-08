package main

import (
	"github.com/Kate-liu/GoBeginner/practiceproject/lightweightthreadpool/workpoolMVP"
	"time"
)

func main() {
	p := workpoolMVP.New(5)

	for i := 0; i < 10; i++ {
		err := p.Schedule(func() {
			time.Sleep(time.Second * 3)
		})
		if err != nil {
			println("task: ", i, "err:", err)
		}
	}

	p.Free()
}
