package main

import (
	"bufio"
	"errors"
	"fmt"
)

func main() {
	err := errors.New("your first demo error")
	errWithCtx = fmt.Errorf("index %d is out of bounds", i)

	err := doSomething()
	if err != nil {
		// 不关心err变量底层错误值所携带的具体上下文信息
		// 执行简单错误处理逻辑并返回
		// ... ...
		return err
	}

	data, err := b.Peek(1)
	if err != nil {
		switch err.Error() {
		case "bufio: negative count":
			// ... ...
			return
		case "bufio: buffer full":
			// ... ...
			return
		case "bufio: invalid use of UnreadByte":
			// ... ...
			return
		default:
			// ... ...
			return
		}
	}

	data, err := b.Peek(1)
	if err != nil {
		switch err {
		case bufio.ErrNegativeCount:
			// ... ...
			return
		case bufio.ErrBufferFull:
			// ... ...
			return
		case bufio.ErrInvalidUnreadByte:
			// ... ...
			return
		default:
			// ... ...
			return
		}
	}

}
