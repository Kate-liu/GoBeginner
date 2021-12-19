# Hash Algorithm

## MD5 VS SHA-256

测试常用到的 MD5 跟 SHA-256 这 两个哈希算法的计算到底有多快。

参见：[main.go](main.go)



## Mac 安装 md5sum和sha1sum

命令行执行：

```sh
$brew install md5sha1sum
```





## Shell md5.Sum 的坑

直接在命令行执行的时候，使用：

```sh
$echo "liu" | md5sum 
a5157352b835b3061f66f8448387bec2  -
```

但是写go程序得到的结果是：

```go
data := []byte("liu") // 明文数据
fmt.Printf("%x", md5.Sum(data))

// 9d4d6204ee943564637f06093236b181
```

这两个的结果不一样，这是因为：echo默认是带换行符做结尾的。echo  'liu' | md5sum 计算的其实是 liu 加上换行(\n)的 md5值 ，而不是liu 的md5值。

解决办法：echo -n 可以去掉换行符。

命令行输入以下命令即可：

```sh
$echo -n "liu" | md5sum
9d4d6204ee943564637f06093236b181  -
```

