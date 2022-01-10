package main

func main() {
	// 字符串的组成
	var s = "中国人"
	// fmt.Println("the character count in s is", utf8.RuneCountInString(s)) // 3
	// for _, c := range s {
	// 	fmt.Printf("0x%x ", c) // 0x4e2d 0x56fd 0x4eba
	// }
	// fmt.Printf("\n")

	// 字符串是只读的字节数组
	println([]byte(s)) // go.string."中国人" SRODATA, 只读的字节数组
	// $ GOOS=linux GOARCH=amd64 go1.17 tool compile -S main.go
	// "".main STEXT size=109 args=0x0 locals=0x58 funcid=0x0
	//        0x0000 00000 (main.go:3)        TEXT    "".main(SB), ABIInternal, $88-0
	// ...
	// 	""..inittask SNOPTRDATA size=24
	//        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
	//        0x0010 00 00 00 00 00 00 00 00                          ........
	// go.string."中国人" SRODATA dupok size=9
	//        0x0000 e4 b8 ad e5 9b bd e4 ba ba                       .........
	// gclocals·69c1753bd5f81501d95132d08af04464 SRODATA dupok size=8
	//        0x0000 02 00 00 00 00 00 00 00                          ........
	// gclocals·9fb7f0986f647f17cb53dda1484e0f7a SRODATA dupok size=10
	//        0x0000 02 00 00 00 01 00 00 00 00 01                    ..........
}
