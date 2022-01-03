# Functions

> 提供常用的 Go 语言内置函数的用法与解释。

## println 

`println` 是 Go 语言运行时提供的内置方法，它不需要依赖任何包就可以向标准输出打印字符串：

```go
// builtin/builtin.go

// The println built-in function formats its arguments in an
// implementation-specific way and writes the result to standard error.
// Spaces are always added between arguments and a newline is appended.
// Println is useful for bootstrapping and debugging; it is not guaranteed
// to stay in the language.
func println(args ...Type)
```

## Println

```go
// fmt/print.go

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
   return Fprintln(os.Stdout, a...)
}
```