package main

import (
	"fmt"
	"github.com/Kate-liu/GoBeginner/gogeekcode/pkgerrdemo/sub1"
)

func main() {
	err := sub1.Diff(1, 2)
	fmt.Printf("%+v", err)
}
