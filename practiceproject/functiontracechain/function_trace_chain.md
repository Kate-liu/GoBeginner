# Function Trace Chain

> 实战项目之函数调用链。

安排了一个小实战项目。

## 引子 

一道思考题：“除了捕捉 panic、延迟释放资源外，日常编码中还有哪些使用 defer 的小技巧呢？” 

可以得到：使用 defer 可以跟踪函数的执行过程。没错！这的确是 defer 的一个常见的使用技巧，很多 Go 教程在讲解 defer 时也会经常使用这个用途举例。

那么，具体是怎么用 defer 来实现函数执行过程的跟踪呢？这里，给出了一个最简单的实现：

```go
package main

func Trace(name string) func() {
   println("enter:", name)
   return func() {
      println("exit:", name)
   }
}

func foo() {
   defer Trace("foo")()
   bar()
}

func bar() {
   defer Trace("bar")()
}

func main() {
   defer Trace("main")()
   foo()
}
```

在讲解这段代码的原理之前，先看一下这段代码的执行结果，直观感受一下什么是函数调用跟踪：

```sh
enter: main
enter: foo
enter: bar
exit: bar
exit: foo
exit: main
```

这个 Go 程序的函数调用的全过程一目了然地展现在了面前：程序按main -> foo -> bar的函数调用次序执行，代码在函数的入口与出口处分别输出了跟踪日志。 

那这段代码是怎么做到的呢？简要分析一下。 

在这段实现中，在每个函数的入口处都使用 defer 设置了一个 deferred 函数。根据 defer 的运作机制，Go 会在 defer 设置 deferred 函数时对 defer 后 面的表达式进行求值。 

以 foo 函数中的defer Trace("foo")() 这行代码为例，Go 会对 defer 后面的表达式Trace("foo")() 进行求值。由于这个表达式包含一个函数调用Trace("foo")，所以这个函数会被执行。 

上面的 Trace 函数只接受一个参数，这个参数代表函数名，Trace 会首先打印进入某函数的日志，比如：“enter: foo”。然后返回一个闭包函数，这个闭包函数一旦被执行，就会输出离开某函数的日志。

在 foo 函数中，这个由 Trace 函数返回的闭包函数就被设置为了 deferred 函数，于是当 foo 函数返回后，这个闭包函数就会被执行，输出“exit: foo”的日志。 

搞清楚上面跟踪函数调用链的实现原理后，再来看看这个实现。会发现这里还是有一些“瑕疵”，也就是离期望的“跟踪函数调用链”的实现还有一些**不足之处**。这里列举了几点：

- 调用 Trace 时需手动显式传入要跟踪的函数名； 
- 如果是并发应用，不同 Goroutine 中函数链跟踪混在一起无法分辨； 
- 输出的跟踪结果缺少层次感，调用关系不易识别； 
- 对要跟踪的函数，需手动调用 Trace 函数。

逐一分析并解决上面提出的这几点问题，并经过逐步地代码演进，最终实现一个**自动注入跟踪代码**，并输出有层次感的函数调用链跟踪命令行工具。

## 自动获取所跟踪函数的函数名 

要解决“调用 Trace 时需要手动显式传入要跟踪的函数名”的问题，也就是要让 Trace 函数能够自动获取到它跟踪函数的函数名信息。

以跟踪 foo 为例，看看这样做能带来什么好处。 在手动显式传入的情况下，需要用下面这个代码对 foo 进行跟踪：

```go
defer Trace("foo")()
```

一旦实现了自动获取函数名，所有支持函数调用链跟踪的函数都只需使用下面调用形式的 Trace 函数就可以了：

```go
defer Trace()()
```

这种一致的 Trace 函数调用方式也为后续的自动向代码中注入 Trace 函数奠定了基础。那么如何实现 Trace 函数对它跟踪函数名的自动获取呢？

需要借助 **Go 标准库 runtime 包**的帮助。 这里，给出了新版 Trace 函数的实现以及它的使用方法，先看一下：

```go
package main

import "runtime"

func Trace() func() {
   pc, _, _, ok := runtime.Caller(1)
   if !ok {
      panic("not found caller")
   }

   fn := runtime.FuncForPC(pc)
   name := fn.Name()

   println("enter:", name)
   return func() {
      println("exit:", name)
   }
}

func foo() {
   defer Trace()()
   bar()
}

func bar() {
   defer Trace()()
}

func main() {
   defer Trace()()
   foo()
}
```

在这一版 Trace 函数中，通过 runtime.Caller 函数获得当前 Goroutine 的函数调用栈上的信息，runtime.Caller 的参数标识的是要获取的是哪一个栈帧的信息。当参数为 0 时，返回的是 Caller 函数的调用者的函数信息，在这里就是 Trace 函数。但需要的是 Trace 函数的调用者的信息，于是传入 1。 

Caller 函数有四个返回值：第一个返回值代表的是程序计数（pc）；第二个和第三个参数代表对应函数所在的源文件名以及所在行数，这里暂时不需要；最后一个参数代表是否能成功获取这些信息，如果获取失败，抛出 panic。 

接下来，通过 runtime.FuncForPC 函数和程序计数器（PC）得到被跟踪函数的函数名称。运行一下改造后代码：

```sh
enter: main.main
enter: main.foo
enter: main.bar
exit: main.bar
exit: main.foo
exit: main.main
```

runtime.FuncForPC 返回的名称中不仅仅包含函数名，还包含了被跟踪函数所在的包名。也就是说，第一个问题已经圆满解决了。 

接下来，来解决第二个问题，也就是当程序中有多 Goroutine 时，Trace 输出的跟踪信息混杂在一起难以分辨的问题。

## 增加 Goroutine 标识 

上面的 Trace 函数在面对只有一个 Goroutine 的时候，还是可以支撑的，但当程序中并发运行多个 Goroutine 的时候，多个函数调用链的出入口信息输出就会混杂在一起，无法分辨。 

那么，接下来还继续对 Trace 函数进行改造，让它支持多 Goroutine 函数调用链的跟踪。方案就是**在输出的函数出入口信息时，带上一个在程序每次执行时能唯一区分 Goroutine 的 Goroutine ID**。 

可能会说，Goroutine 也没有 ID 信息啊！的确如此，Go 核心团队为了避免 Goroutine ID 的滥用，故意没有将 Goroutine ID 暴露给开发者。

但在 **Go 标准库的 h2_bundle.go** 中，却发现了一个获取 Goroutine ID 的标准方法，看下面代码：

```go
// net/http/h2_bundle.go
var http2goroutineSpace = []byte("goroutine ")

func http2curGoroutineID() uint64 {
   bp := http2littleBuf.Get().(*[]byte)
   defer http2littleBuf.Put(bp)
   b := *bp
   b = b[:runtime.Stack(b, false)]
   // Parse the 4707 out of "goroutine 4707 ["
   b = bytes.TrimPrefix(b, http2goroutineSpace)
   i := bytes.IndexByte(b, ' ')
   if i < 0 {
      panic(fmt.Sprintf("No space found in %q", b))
   }
   b = b[:i]
   n, err := http2parseUintBytes(b, 10, 64)
   if err != nil {
      panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
   }
   return n
}

var http2littleBuf = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 64)
		return &buf
	},
}

// parseUintBytes is like strconv.ParseUint, but using a []byte.
func http2parseUintBytes(s []byte, base int, bitSize int) (n uint64, err error) {
  // ...
}
```

不过，由于 http2curGoroutineID 不是一个导出函数，无法直接使用。可以把它 复制出来改造一下：

```go
package main

import (
   "bytes"
   "fmt"
   "runtime"
   "strconv"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
   b := make([]byte, 64)
   b = b[:runtime.Stack(b, false)]
   // Parse the 4707 out of "goroutine 4707 ["
   b = bytes.TrimPrefix(b, goroutineSpace)
   i := bytes.IndexByte(b, ' ')
   if i < 0 {
      panic(fmt.Sprintf("No space found in %q", b))
   }
   b = b[:i]
   n, err := strconv.ParseUint(string(b), 10, 64)
   if err != nil {
      panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
   }
   return n
}
```

这里，改造了两个地方。

- 一个地方是通过直接创建一个 byte 切片赋值给 b，替代原 http2curGoroutineID 函数中从一个 pool 池获取 byte 切片的方式，
- 另外一个是使用 strconv.ParseUint 替代了原先的 http2parseUintBytes。

改造后，就可以直接使用 curGoroutineID 函数来获取 Goroutine 的 ID 信息了。 

接下来，**在 Trace 函数中添加 Goroutine ID 信息**的输出：

```go
func Trace() func() {
   pc, _, _, ok := runtime.Caller(1)
   if !ok {
      panic("not found caller")
   }

   fn := runtime.FuncForPC(pc)
   name := fn.Name()
   gid := curGoroutineID()

   fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
   return func() {
      fmt.Printf("g[%05d]: exit: [%s]\n", gid, name)
   }
}
```

从上面代码看到，在出入口输出的跟踪信息中加入了 Goroutine ID 信息，输出的 Goroutine ID 为 5 位数字，如果 ID 值不足 5 位，则左补零，这一切都是 Printf 函数的格式控制字符串“%05d”实现的。

这样对齐 Goroutine ID 的位数，为的是输出信息格式的一致性更好。如果 Go 程序中 Goroutine 的数量超过了 5 位数可以表示的数值范围，也可以自行调整控制字符串。 

接下来，也要对示例进行一些调整，将这个程序**由单 Goroutine 改为多 Goroutine 并发**的，这样才能验证支持多 Goroutine 的新版 Trace 函数是否好用：

```go
package main

import (
   "bytes"
   "fmt"
   "runtime"
   "strconv"
   "sync"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
   b := make([]byte, 64)
   b = b[:runtime.Stack(b, false)]
   // Parse the 4707 out of "goroutine 4707 ["
   b = bytes.TrimPrefix(b, goroutineSpace)
   i := bytes.IndexByte(b, ' ')
   if i < 0 {
      panic(fmt.Sprintf("No space found in %q", b))
   }
   b = b[:i]
   n, err := strconv.ParseUint(string(b), 10, 64)
   if err != nil {
      panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
   }
   return n
}

func Trace() func() {
   pc, _, _, ok := runtime.Caller(1)
   if !ok {
      panic("not found caller")
   }

   fn := runtime.FuncForPC(pc)
   name := fn.Name()
   gid := curGoroutineID()

   fmt.Printf("g[%05d]: enter: [%s]\n", gid, name)
   return func() {
      fmt.Printf("g[%05d]: exit: [%s]\n", gid, name)
   }
}

func A1() {
   defer Trace()()
   B1()
}

func B1() {
   defer Trace()()
   C1()
}

func C1() {
   defer Trace()()
   D()
}

func D() {
   defer Trace()()
}

func A2() {
   defer Trace()()
   B2()
}

func B2() {
   defer Trace()()
   C2()
}

func C2() {
   defer Trace()()
   D()
}

func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
      A2()
      wg.Done()
   }()

   A1()
   wg.Wait()
}
```

新示例程序共有两个 Goroutine，main groutine 的调用链为A1 -> B1 -> C1 -> D， 而另外一个 Goroutine 的函数调用链为A2 -> B2 -> C2 -> D。

来看一下这个程序的执行结果是否和原代码中两个 Goroutine 的调用链一致：

```sh
g[00001]: enter: [main.A1]
g[00001]: enter: [main.B1]
g[00001]: enter: [main.C1]
g[00006]: enter: [main.A2]
g[00001]: enter: [main.D]
g[00001]: exit: [main.D]
g[00001]: exit: [main.C1]
g[00001]: exit: [main.B1]
g[00001]: exit: [main.A1]
g[00006]: enter: [main.B2]
g[00006]: enter: [main.C2]
g[00006]: enter: [main.D]
g[00006]: exit: [main.D]
g[00006]: exit: [main.C2]
g[00006]: exit: [main.B2]
g[00006]: exit: [main.A2]
```

新示例程序输出了带有 Goroutine ID 的出入口跟踪信息，通过 Goroutine ID 可以快速确认某一行输出是属于哪个 Goroutine 的。 

但由于 **Go 运行时对 Goroutine 调度顺序的不确定性**，各个 Goroutine 的输出还是会存在交织在一起的问题，这给查看某个 Goroutine 的函数调用链跟踪信息带来阻碍。

这里提供一个**小技巧**：可以将程序的输出重定向到一个本地文件中，然后通过 Goroutine ID 过滤出（可使用 grep 工具）想查看的 groutine 的全部函数跟踪信息。 

到这里，就实现了输出带有 Goroutine ID 的函数跟踪信息，不过，有没有觉得输出的函数调用链信息还是不够美观，缺少层次感，体验依旧不那么优秀呢？

下面就来美化一下信息的输出形式。



## 让输出的跟踪信息更具层次感 

对于程序员来说，缩进是最能体现出“层次感”的方法，如果将上面示例中 Goroutine 00001 的函数调用跟踪信息以下面的形式展示出来，函数的调用顺序是不是更加一目了然了呢？

```go
g[00001]:       ->main.A1
g[00001]:               ->main.B1
g[00001]:                       ->main.C1
g[00001]:                               ->main.D
g[00001]:                               <-main.D
g[00001]:                       <-main.C1
g[00001]:               <-main.B1
g[00001]:       <-main.A1
```

那么就以这个形式为目标，考虑如何实现输出这种带缩进的函数调用跟踪信息。还是直接上代码吧：

```go
package main

import (
   "bytes"
   "fmt"
   "runtime"
   "strconv"
   "sync"
)

var goroutineSpace = []byte("goroutine ")

func curGoroutineID() uint64 {
   b := make([]byte, 64)
   b = b[:runtime.Stack(b, false)]
   // Parse the 4707 out of "goroutine 4707 ["
   b = bytes.TrimPrefix(b, goroutineSpace)
   i := bytes.IndexByte(b, ' ')
   if i < 0 {
      panic(fmt.Sprintf("No space found in %q", b))
   }
   b = b[:i]
   n, err := strconv.ParseUint(string(b), 10, 64)
   if err != nil {
      panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
   }
   return n
}

var mu sync.Mutex
var m = make(map[uint64]int)

func Trace() func() {
   pc, _, _, ok := runtime.Caller(1)
   if !ok {
      panic("not found caller")
   }

   fn := runtime.FuncForPC(pc)
   name := fn.Name()
   gid := curGoroutineID()

   mu.Lock()
   indents := m[gid]    // 获取当前 gid 对应的缩进层次
   m[gid] = indents + 1 // 缩进层次 +1 后存入 map
   mu.Unlock()

   printTrace(gid, name, "->", indents+1)

   return func() {
      mu.Lock()
      indents := m[gid]    // 获取当前 gid 对应的缩进层次
      m[gid] = indents - 1 // 缩进层次 -1 后存入 map
      mu.Unlock()

      printTrace(gid, name, "<-", indents)
   }
}

func printTrace(id uint64, name string, arrow string, indent int) {
   indents := ""
   for i := 0; i < indent; i++ {
      indents += "   "
   }
   fmt.Printf("g[%05d]: %s%s%s\n", id, indents, arrow, name)
}

func A1() {
   defer Trace()()
   B1()
}

func B1() {
   defer Trace()()
   C1()
}

func C1() {
   defer Trace()()
   D()
}

func D() {
   defer Trace()()
}

func A2() {
   defer Trace()()
   B2()
}

func B2() {
   defer Trace()()
   C2()
}

func C2() {
   defer Trace()()
   D()
}

func main() {
   var wg sync.WaitGroup
   wg.Add(1)
   go func() {
      A2()
      wg.Done()
   }()

   A1()
   wg.Wait()
}
```

在上面这段代码中，使用了一个 map 类型变量 m 来保存每个 Goroutine 当前的缩进信息：m 的 key 为 Goroutine 的 ID，值为缩进的层次。

然后，考虑到 Trace 函数可能在并发环境中运行，根据“map 不支持并发写”的注意事项，增加了一个 sync.Mutex 实例 mu 用于同步对 m 的写操作。 

这样，对于一个 Goroutine 来说，

- 每次刚进入一个函数调用，就在输出入口跟踪信息之前，将缩进层次加一，并输出入口跟踪信息，加一后的缩进层次值也保存到 map 中。
- 然后，在函数退出前，取出当前缩进层次值并输出出口跟踪信息，之后再将缩进层次减一后保存到 map 中。 

除了增加缩进层次信息外，在这一版的 Trace 函数实现中，也把输出出入口跟踪信息 的操作提取到了一个独立的函数 printTrace 中，这个函数会根据传入的 Goroutine ID、 函数名、箭头类型与缩进层次值，按预定的格式拼接跟踪信息并输出。 

运行新版示例代码，会得到下面的结果：

```sh
g[00001]:       ->main.A1
g[00001]:               ->main.B1
g[00001]:                       ->main.C1
g[00001]:                               ->main.D
g[00001]:                               <-main.D
g[00001]:                       <-main.C1
g[00001]:               <-main.B1
g[00001]:       <-main.A1
g[00018]:       ->main.A2
g[00018]:               ->main.B2
g[00018]:                       ->main.C2
g[00018]:                               ->main.D
g[00018]:                               <-main.D
g[00018]:                       <-main.C2
g[00018]:               <-main.B2
g[00018]:       <-main.A2
```

显然，通过这种带有缩进层次的函数调用跟踪信息，可以更容易地识别某个 Goroutine 的函数调用关系。 

到这里，函数调用链跟踪已经支持了多 Goroutine，并且可以输出有层次感的跟踪 信息了，但对于 Trace 特性的使用者而言，依然需要手工在自己的函数中添加对 Trace 函数的调用。

那么是否可以将 Trace 特性自动注入特定项目下的各个源码文件中呢？接下来继续来改进 Trace 工具。



---

waiting for a moment.....





## 利用代码生成自动注入 Trace 函数 

要实现向目标代码中的函数 / 方法自动注入 Trace 函数，首先要做的就是将上面 Trace 函数相关的代码打包到一个 module 中以方便其他 module 导入。

下面就先来 看看将 Trace 函数放入一个独立的 module 中的步骤。 

### 将 Trace 函数放入一个独立的 module 中 

创建一个名为 instrument_trace 的目录，进入这个目录后，通过 go mod init 命令创 建一个名为 github.com/bigwhite/instrument_trace 的 module：

















