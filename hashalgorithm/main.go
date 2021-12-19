package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"sync/atomic"
	"time"
)

// 测试常用到的 MD5 跟 SHA-256 这两个哈希算法的计算到底有多快
func main() {
	data := []byte("liu") // 明文数据

	// 两个计数
	md5Count := int64(0)
	sha256Count := int64(0)

	go func() {
		for {
			md5.Sum(data)                 // MD5哈希计算
			atomic.AddInt64(&md5Count, 1) // 原子计数+1
		}
	}()

	go func() {
		for {
			sha256.Sum256(data)              // SHA-256哈希计算
			atomic.AddInt64(&sha256Count, 1) // 原子计数+1
		}
	}()

	time.Sleep(time.Second) // 等待1秒

	// 输出结果
	fmt.Printf("md5 count: %d\n", md5Count)
	fmt.Printf("sha256 count: %d\n", sha256Count)
}

// 计算结果
// md5 count: 7345719
// sha256 count: 4432737
