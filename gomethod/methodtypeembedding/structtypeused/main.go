package main

import (
	"fmt"
	"io"
	"strings"
)

type MyInt int

func (n *MyInt) Add(m int) {
	*n = *n + MyInt(m)
}

type t struct {
	a int
	b int
}

type S struct {
	*MyInt
	t
	io.Reader
	s string
	n int
}

func main() {
	m := MyInt(17)
	r := strings.NewReader("hello, go")
	s := S{
		MyInt: &m,
		t: t{
			a: 1,
			b: 2,
		},
		Reader: r,
		s:      "demo",
	}

	// var sl = make([]byte, len("hello, go"))
	// s.Reader.Read(sl)
	// fmt.Println(string(sl)) // hello, go
	// s.MyInt.Add(5)
	// fmt.Println(*(s.MyInt)) // 22

	var sl = make([]byte, len("hello, go"))
	s.Read(sl)
	fmt.Println(string(sl))
	s.Add(5)
	fmt.Println(*(s.MyInt))
}
