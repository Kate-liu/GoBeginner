package main

import (
	"errors"
	"fmt"
	"time"
)

func spawn(f func() error) <-chan error {
	c := make(chan error)
	go func() {
		c <- f()
	}()

	return c
}

func main() {
	c := spawn(func() error {
		time.Sleep(2 * time.Second)
		return errors.New("timeout")
	})
	fmt.Println(<-c)
}
