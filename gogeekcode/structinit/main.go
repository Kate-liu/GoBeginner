package main

import (
	"github.com/Kate-liu/GoBeginner/gogeekcode/structinit/foo"
	"github.com/Kate-liu/GoBeginner/gogeekcode/structinit/foooption"
)

func main() {
	// foo
	f1 := foo.NewFoo("rmliu", 1, 25, nil)
	f2 := foo.NewFoo("jianfengye", 2, 0, nil)

	println(f1)
	println(f2)

	// foo option
	f3 := foooption.NewFoo(1, foooption.WithName("rmliu"), foooption.WithAge(26))
	println(f3)
}
