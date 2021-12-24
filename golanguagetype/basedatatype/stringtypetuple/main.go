package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main() {
	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将 string 类型变量地址显式
	fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0

	p := (*[5]byte)(unsafe.Pointer(hdr.Data)) // 获取 Data 字段所指向的数组的指针
	dumpBytesArray((*p)[:])                   // [h e l l o ] // 输出底层数组的内容

	// 取出字符串的长度
	ll := len(s)
	fmt.Println(ll)
}
