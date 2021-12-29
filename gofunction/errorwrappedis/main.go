package main

import (
	"errors"
	"fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
	err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
	err2 := fmt.Errorf("wrap err1: %w", err1)
	println(err2 == ErrSentinel) // false

	if errors.Is(err2, ErrSentinel) {
		println("err2 is ErrSentinel")
		return
	}

	println("err2 is not ErrSentinel")

}
