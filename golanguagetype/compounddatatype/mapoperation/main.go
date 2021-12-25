package main

import (
	"fmt"
)

func main() {
	// 插入操作
	m := make(map[int]string)
	m[1] = "value1"
	m[2] = "value2"
	m[3] = "value3"
	fmt.Println(m) // map[1:value1 2:value2 3:value3]

	// 插入操作 之 新值覆盖旧值
	m1 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m1) // map[key1:1 key2:2]
	m1["key1"] = 11 // 11会覆盖掉"key1"对应的旧值1
	m1["key3"] = 3  // 此时m1为map[key1:11 key2:2 key3:3]
	fmt.Println(m1) // map[key1:11 key2:2 key3:3]

	// 获取键值对数量
	m2 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(len(m2)) // 2
	m2["key3"] = 3
	fmt.Println(len(m2)) // 3

	// 查找
	m3 := make(map[string]int)
	v := m3["key1"]
	fmt.Println(v) // 0

	m3["key1"] = 666
	v2 := m3["key1"]
	fmt.Println(v2) // 666

	// 查找 之 comma ok 手法
	m4 := make(map[string]int)
	m4["key1"] = 999
	v4, ok := m4["key1"]
	if !ok {
		// "key1" 不在 map 中
		fmt.Println("不存在的 key")
	}
	// "key1"在map中，v3将被赋予"key1"键对应的value
	fmt.Println("key1 在map中，值为:", v4)

	// 查找 之 comma ok 手法 之 空标识符
	m5 := make(map[string]int)
	_, ok1 := m5["key1"]
	// ... ...
	fmt.Println(ok1) // false

	// 删除操作
	m6 := map[string]int{
		"key1": 1,
		"key2": 2,
	}
	fmt.Println(m6)    // map[key1:1 key2:2]
	delete(m6, "key2") // 删除"key2"
	fmt.Println(m6)    // map[key1:1]

	// 遍历操作
	m7 := map[int]int{
		1: 11,
		2: 12,
		3: 13,
	}
	fmt.Printf("{ ")
	for k, v := range m7 {
		fmt.Printf("[%d, %d]", k, v)
	}
	fmt.Printf("}\n") // 输出 { [1, 11] [2, 12] [3, 13] }

	// 只关心键
	for k, _ := range m7 {
		// 只使用 k
		fmt.Printf("key: %d\n", k)
	}

	// 只关心键 更地道方式
	for k := range m7 {
		// 只使用 k
		fmt.Printf("key: %d\n", k)
	}

	// 只关心值
	for _, v := range m7 {
		// 只使用 k
		fmt.Printf("value: %d\n", v)
	}

	// // 获取 map 中的 value 地址，编译不通过
	// m8 := make(map[string]int)
	// m8["key1"] = 678
	// p := &m8["key1"] // cannot take the address of m[key]
	// fmt.Println(p)
}
