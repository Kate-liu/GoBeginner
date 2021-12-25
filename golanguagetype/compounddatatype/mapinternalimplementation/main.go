package main

func main() {
	// // 创建map类型变量实例
	// m := make(map[keyType]valType, capacityhint) → m := runtime.makemap(maptype, capacityhint, m)
	//
	// // 插入新键值对或给键重新赋值
	// m["key"] = "value" → v := runtime.mapassign(maptype, m, "key") // v是用于后续存储value的空间的地址
	//
	// // 获取某键的值
	// v := m["key"] → v := runtime.mapaccess1(maptype, m, "key")
	// v, ok := m["key"] → v, ok := runtime.mapaccess2(maptype, m, "key")
	//
	// // 删除某键
	// delete(m, "key") → runtime.mapdelete(maptype, m, "key")

}
