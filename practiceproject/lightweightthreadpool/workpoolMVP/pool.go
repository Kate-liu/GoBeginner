package workpoolMVP

import (
	"errors"
	"fmt"
	"sync"
)

type Task func()

type Pool struct {
	capacity int // workerpool 大小

	active chan struct{} // 对应上图中的active channel
	tasks  chan Task     // 对应上图中的task channel

	wg   sync.WaitGroup // 用于在pool销毁时等待所有worker退出
	quit chan struct{}  // 用于通知各个worker退出的信号channel
}

const (
	defaultCapacity = 100
	maxCapacity     = 10000
)

func New(capacity int) *Pool {
	if capacity <= 0 {
		capacity = defaultCapacity
	}
	if capacity > maxCapacity {
		capacity = maxCapacity
	}

	p := &Pool{
		capacity: capacity,
		tasks:    make(chan Task),
		quit:     make(chan struct{}),
		active:   make(chan struct{}, capacity),
	}
	fmt.Printf("workerpool start\n")

	go p.run()

	return p
}

func (p *Pool) run() {
	idx := 0

	for {
		select {
		case <-p.quit:
			return
		case p.active <- struct{}{}:
			// create a new worker
			idx++
			p.newWorker(idx)
		}
	}
}

func (p *Pool) newWorker(i int) {
	p.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("worker[%03d]: recover panic[%s] and exit\n", i, err)
				<-p.active
			}
			p.wg.Done()
		}()

		fmt.Printf("worker[%03d]: start\n", i)

		for {
			select {
			case <-p.quit:
				fmt.Printf("worker[%03d]: exit\n", i)
				<-p.active
				return
			case t := <-p.tasks:
				fmt.Printf("worker[%03d]: receive a task\n", i)
				t()
			}
		}
	}()
}

var ErrWorkerPoolFreed = errors.New("workerpool freed") // workerpool

func (p *Pool) Schedule(t Task) error {
	select {
	case <-p.quit:
		return ErrWorkerPoolFreed
	case p.tasks <- t:
		return nil
	}
}

func (p *Pool) Free() {
	close(p.quit) // make sure all worker and p.run exit and schedule return error
	p.wg.Wait()
	fmt.Printf("workerpool freed\n")
}
