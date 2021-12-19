package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	// 1.this md5 data
	data := []byte("liu") // 明文数据
	fmt.Printf("%x", md5.Sum(data))
	fmt.Println()

	// 2.this md5 md5hole data
	newdata := []byte("liu\n") // 明文数据
	fmt.Printf("%x", md5.Sum(newdata))

	// 3.shell md5 data
	// $echo "liu" | md5sum

	// 4.shell -n md5 data
	// $echo -n "liu" | md5sum

	// 1.this md5 data: 9d4d6204ee943564637f06093236b181
	// 2.this md5 md5hole data: a5157352b835b3061f66f8448387bec2
	// 3.shell md5 data: a5157352b835b3061f66f8448387bec2
	// 4.shell -n md5 data: 9d4d6204ee943564637f06093236b181

	// conclusion
	// true: md5.Sum("liu") is 9d4d6204ee943564637f06093236b181
}
