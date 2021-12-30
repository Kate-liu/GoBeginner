# Go Function

> Go 函数

## Function

Go 代码中的基本功能逻辑单元：函 数。 

函数是现代编程语言的基本语法元素，无论是在命令式语言、面向对象语言还 是动态脚本语言中，函数都位列 C 位。 

在 Go 语言中，函数是唯一一种基于特定输入，实现特定任务并可返回 任务执行结果的代码块（**Go 语言中的方法本质上也是函数**）。如果忽略 Go 包在 Go 代码组织层面的作用，可以说 **Go 程序就是一组函数的集合**，实际上，日常的 Go 代 码编写大多都集中在实现某个函数上。 

但“一龙生九子，九子各不同”！虽然各种编程语言都加入了函数这个语法元素，但各个语言中函数的形式与特点又有不同。那么 Go 语言中函数又有哪些独特之处呢？

### Go 函数与函数声明 

函数对应的英文单词是 Function，Function 这个单词原本是功能、职责的意思。编程语 言使用 Function 这个单词，表示将一个大问题分解后而形成的、若干具有特定功能或职责 的小任务，可以说十分贴切。

函数代表的小任务可以在一个程序中被多次使用，甚至可以 在不同程序中被使用，因此函数的出现也提升了整个程序界代码复用的水平。 

那 Go 语言中，函数相关的语法形式是怎样的呢？

#### Go 函数声明

先来看最常用的 Go 函数声明。 在 Go 中，定义一个函数的最常用方式就是使用函数声明。以 Go 标准库 fmt 包 提供的 Fprintf 函数为例，看一下一个普通 Go 函数的声明长啥样：

![image-20211228162634165](go_function.assets/image-20211228162634165.png)

一个 Go 函数的声明由五部分组成，一个个来拆解一下。 

- 第一部分是**关键字 func**，
  - Go 函数声明必须以关键字 func 开始。
- 第二部分是**函数名**。
  - 函数名是指代函数定义的标识符，函数声明后，会通过函数名这个标识符来使用这个函数。
  - 在同一个 Go 包中，函数名应该是唯一的，
  - 并且它也遵守 Go 标识符的导出规则，也就是之前说的，**首字母大写的函数名指代的函数是可以在包外使用的**，**小写的就只在包内可见**。 
- 第三部分是**参数列表**。
  - 参数列表中声明了将要在函数体中使用的各个参数。
  - 参数列表 紧接在函数名的后面，并用一个括号包裹。它使用逗号作为参数间的分隔符，而且每个参 数的参数名在前，参数类型在后，这和变量声明中变量名与类型的排列方式是一致的。 
  - 另外，Go 函数支持变长参数，也就是一个形式参数可以对应数量不定的实际参数。 
  - Fprintf 就是一个支持变长参数的函数，可以看到它第三个形式参数 a 就是一个变长参 数，而且变长参数与普通参数在声明时的不同点，就在于它会在类型前面增加了一 个“…”符号。
- 第四部分是**返回值列表**。
  - 返回值承载了函数执行后要返回给调用者的结果，返回值列表声 明了这些返回值的类型，返回值列表的位置紧接在参数列表后面，两者之间用一个空格隔 开。
  - 不过，上图中比较特殊，Fprintf 函数的返回值列表不仅声明了返回值的类型，还声明 了返回值的名称，这种返回值被称为**具名返回值**。多数情况下，不需要这么做，只需 声明返回值的类型即可。 
- 最后，放在一对大括号内的是**函数体**，函数的具体实现都放在这里。
  - 不过，函数声明中的 **函数体是可选的**。如果没有函数体，说明这个函数可能是在 Go 语言之外实现的，比如使用汇编语言实现，然后通过链接器将实现与声明中的函数名链接到一起。
  - 没有函数体的函 数声明是更高级的话题。 

看到这里，可能会问：同为声明，为啥函数声明与之前学过的变量声明在形式上差距这 么大呢? 变量声明中的变量名、类型名和初值与上面的函数声明是怎么对应的呢？ 

#### 函数的变量声明形式

为了更好地理解函数声明，这里就横向对比一下，把上面的函数声明**等价转换为变量声明的形式**看看：

![image-20211228163322743](go_function.assets/image-20211228163322743.png)

转换后的代码不仅和之前的函数声明是等价的，而且这也是完全合乎 Go 语法规则的代码。

对照一下这两张图，是不是有一种豁然开朗的感觉呢？这不就是在声明一个类型为 函数类型的变量吗！

函数声明中的函数名其实就是变量名，函数声明中的 func 关键字、参数列表和 返回值列表共同构成了**函数类型**。而参数列表与返回值列表的组合也被称为**函数签名**，它 是决定两个函数类型是否相同的决定因素。

因此，函数类型也可以看成是由 func 关键字与 函数签名组合而成的。

 通常，在表述函数类型时，会**省略**函数签名参数列表中的参数名，以及返回值列表中 的返回值变量名。

比如上面 Fprintf 函数的函数类型是：

```go
func (io.Writer, string, ...interface {}) (int, error)
```

这样，如果两个函数类型的函数签名是相同的，即便参数列表中的参数名，以及返回值列 表中的返回值变量名都是不同的，那么这**两个函数类型也是相同类型**，比如下面两个函数类型：

```go
func (a int, b string) (results []string, err error)
func (c int, d string) (sl []string, err error)
```

如果把这两个函数类型的参数名与返回值变量名省略，那它们都是`func (int, string) ([]string, error)`，因此它们是相同的函数类型。 

可以得到这样一个结论：每个函数声明所定义的函数，仅仅是对应的函数类 型的一个实例，就像var a int = 13这个变量声明语句中 a 是 int 类型的一个实例一 样。 

使用复合类型字面值对结构体类型变量进行显式初始化，和用变量声明来声明函数变量的形式，都以最简化的样子表现出来，看下面代码：

```go
s := T{}        // 使用复合类型字面值对结构体类型T的变量进行显式初始化
f := func () {} // 使用变量声明形式的函数声明
```

这里，T{}被称为复合类型字面值，那么处于同样位置的 func(){}是什么呢？Go 语言也为它 准备了一个名字，叫“**函数字面值（Function Literal）**”。

可以看到，函数字面值由函数类型与函数体组成，它特别像一个没有函数名的函数声明，因此也叫它**匿名函 数**。匿名函数在 Go 中用途很广。 

可能会想：既然是等价的，那以后就用这种变量声明的形式来声明一个函数吧。万万不可！这里只是为了理解函数声明做了一个等价变换。

在 Go 中的绝大 多数情况，还是会通过**传统的函数声明来声明一个特定函数类型的实例**，也就是俗称的“定义一个函数”。 



#### 函数参数

函数参数列表中的参数，是函数声明的、用于函数体实现的局部变量。由于函数分为声明 与使用两个阶段，在不同阶段，参数的称谓也有不同。

- 在函数**声明阶段**，把参数列表中的参数叫做**形式参数**（Parameter，简称形参），在函数体中，使用的都是形参； 
- 而在函数**实际调用**时传入的参数被称为**实际参数**（Argument，简称实参）。

为了便于直观理解，绘制了这张示意图，可以参考一下：

![image-20211228164835152](go_function.assets/image-20211228164835152.png)

当实际调用函数的时候，实参会传递给函数，并和形式参数逐一绑定，编译器会根据各个形参的类型与数量，来检查传入的实参的类型与数量是否匹配。只有匹配，程序才能 继续执行函数调用，否则编译器就会报错。 

Go 语言中，函数参数传递采用是**值传递**的方式。

- 所谓“值传递”，就是将实际参数在内存中的表示**逐位拷贝（Bitwise Copy）到形式参数中**。
- 对于像**整型、数组、结构体**这类类型，它们的内存表示就是它们自身的数据内容，因此当这些类型作为实参类型时，值传递拷贝的就是它们自身，传递的开销也与它们自身的大小成正比。 
- 但是像 **string、切片、map** 这些类型就不是了，它们的内存表示对应的是它们数据内容的“描述符”。当这些类型作为实参类型时，值传递拷贝的也是它们数据内容的“描述符”，不包括数据内容本身，所以这些类型传递的开销是固定的，与数据内容大小无关。 这种只拷贝“描述符”，不拷贝实际数据内容的拷贝过程，也被称为**“浅拷贝”**。 

不过函数参数的传递也有两个例外，

- 当函数的**形参为接口类型**，或者形参是**变长参数**时， 简单的值传递就不能满足要求了，这时 Go 编译器会介入：
- 对于类型为接口类型的形参， Go 编译器会把传递的实参赋值给对应的接口类型形参；
- 对于为变长参数的形参，Go 编译 器会将零个或多个实参按一定形式转换为对应的变长形参。

那么这里，零个或多个传递给变长形式参数的实参，被 Go 编译器转换为何种形式了呢？ 通过下面示例代码来看一下：

```go
func myAppend(sl []int, elems ...int) []int {
   fmt.Printf("%T\n", elems) // []int
   if len(elems) == 0 {
      println("no elems to append")
      return sl
   }
   
   sl = append(sl, elems...)
   return sl
}

func main() {
   sl := []int{1, 2, 3}
   sl = myAppend(sl) // no elems to append
   fmt.Println(sl)   // [1 2 3]
   sl = myAppend(sl, 4, 5, 6)
   fmt.Println(sl) // [1 2 3 4 5 6]
}
```

重点看一下代码中的 myAppend 函数，这个函数基于 append，实现了向一个整型切片追加数据的功能。它支持变长参数，它的第二个形参 elems 就是一个变长参数。 

myAppend 函数通过 Printf 输出了变长参数的类型。执行这段代码，将看到变长参数 elems 的类型为[]int。 

这也就说明，在 Go 中，**变长参数实际上是通过切片来实现的**。所以，在函数体中， 就可以使用切片支持的所有操作来操作变长参数，这会大大简化了变长参数的使用复杂 度。比如 myAppend 中，使用 len 函数就可以获取到传给变长参数的实参个数。 

#### 函数支持多返回值 

和其他主流静态类型语言，比如 C、C++ 和 Java 不同，Go 函数支持多返回值。

多返回值可以让函数将更多结果信息返回给它的调用者，Go 语言的错误处理机制很大程度就是建立在多返回值的机制之上的。

函数返回值列表从形式上看主要有三种：

```go
func foo()                      // 无返回值
func foo() error                // 仅有一个返回值
func foo() (int, string, error) // 有2或2个以上返回值
```

- 如果一个函数没有显式返回值，那么可以像第一种情况那样，在函数声明中省略返回值列表。
- 而且，如果一个函数仅有一个返回值，那么通常在函数声明中，就不需要将 返回值用括号括起来，
- 如果是 2 个或 2 个以上的返回值，那还是需要用括号括起来的。 

在函数声明的返回值列表中，通常会像上面例子那样，仅列举返回值的类型，但也可以像 fmt.Fprintf 函数的返回值列表那样，为每个返回值声明变量名，这种带有名字的返回值被称为**具名返回值（Named Return Value）**。这种具名返回值变量可以像函数体 中声明的局部变量一样在函数体内使用。 

那么在日常编码中，究竟该使用普通返回值形式，还是具名返回值形式呢？ 

- Go 标准库以及大多数项目代码中的函数，都选择了**使用普通的非具名返回值形式**。

- 但在一些**特定场景**下，具名返回值也会得到应用。

  - 比如，当函数使用 defer，而且还在 defer 函数中修改外部函数返回值时，具名返回值可以让代码显得更优雅清晰。

  - 再比如，当函数的返回值个数较多时，每次显式使用 return 语句时都会接一长串返回值， 这时，用具名返回值可以让函数实现的可读性更好一些，

  - 比如下面 Go 标准库 **time 包 中的 parseNanoseconds 函数**就是这样：

  - ```go
    // time/format.go
    func parseNanoseconds(value string, nbytes int) (ns int, rangeErrString string, err error) {
       if value[0] != '.' {
          err = errBad
          return
       }
       if ns, err = atoi(value[1:nbytes]); err != nil {
          return
       }
       if ns < 0 || 1e9 <= ns {
          rangeErrString = "fractional second"
          return
       }
    
       scaleDigits := 10 - nbytes
       for i := 0; i < scaleDigits; i++ {
          ns *= 10
       }
       return
    }
    ```



### 函数是“一等公民” 

这个特点就是，函数在 Go 语言中属于“**一等公民（First-Class Citizen）**”。并不是在所有编程语言中函数都是“一等公民”。 

那么，什么是编程语言的“一等公民”呢？关于这个名词，业界和教科书都没有给出精准的定义。这里可以引用一下 wiki 发明人、C2 站点作者沃德·坎宁安 (Ward Cunningham)对“一等公民”的解释：

> 如果一门编程语言对某种语言元素的创建和使用没有限制，可以像对待值（value） 一样对待这种语法元素，那么就称这种语法元素是这门编程语言的“一等公民”。 拥有“一等公民”待遇的语法元素可以存储在变量中，可以作为参数传递给函数，可以在函数内部创建并可以作为返回值从函数返回。

基于这个解释，来看看 Go 语言的函数作为“一等公民”，表现出的各种行为特征。

#### 特征一：Go 函数可以存储在变量中

按照沃德·坎宁安对一等公民的解释，身为一等公民的语法元素是可以存储在变量中的。

其实，这点在前面理解函数声明时已经验证过了，这里再用例子简单说明一下：

```go
package main

import (
   "fmt"
   "io"
   "os"
)

var (
   myFprintf = func(w io.Writer, format string, a ...interface{}) (int, error) {
      return fmt.Fprintf(w, format, a...)
   }
)

func main() {
   fmt.Printf("%T\n", myFprintf)             // func(io.Writer, string, ...interface {}) (int, error)
   myFprintf(os.Stdout, "%s\n", "Hello, Go") // 输出 Hello，Go
}
```

在这个例子中，把新创建的一个匿名函数赋值给了一个名为 myFprintf 的变量，通过 这个变量，便可以调用刚刚定义的匿名函数。

然后再通过 Printf 输出 myFprintf 变量的类型，也会发现结果与预期的函数类型是相符的。 

#### 特征二：支持在函数内创建并通过返回值返回

Go 函数不仅可以在函数外创建，还可以在函数内创建。而且由于函数可以存储在变量中， 所以函数也可以在创建后，作为函数返回值返回。来看下面这个例子：

```go
func setup(task string) func() {
   println("do some setup stuff for", task)
   return func() {
      println("do some teardown stuff for", task)
   }
}

func main() {
   teardown := setup("demo")
   defer teardown()
   println("do some bussiness stuff")
}
```

这个例子，模拟了执行一些重要逻辑之前的**上下文建立（setup）**，以及之后的上下文拆除 （teardown）。

在一些**单元测试的代码**中，也经常会在执行某些用例之前，建立此次执行的上下文（setup），并在这些用例执行后拆除上下文（teardown），避免这次执行对后续用例执行的干扰。 

在这个例子中，在 setup 函数中创建了这次执行的上下文拆除函数，并通过返回值的形式，将这个拆除函数返回给了 setup 函数的调用者。

setup 函数的调用者，在执行完对应这次执行上下文的重要逻辑后，再调用 setup 函数返回的拆除函数，就可以完成对上下文的拆除了。 

从这段代码中也可以看到，setup 函数中创建的拆除函数也是一个匿名函数，但和前面看到的匿名函数有一个不同，这个不同就在于这个**匿名函数使用了定义它的函数 setup 的局部变量 task**，这样的匿名函数在 Go 中也被称为**闭包（Closure）**。 

闭包本质上就是一个匿名函数或叫函数字面值，它们可以引用它的包裹函数，也就是创建它们的函数中定义的**变量**。然后，这些变量在包裹函数和匿名函数之间共享，只要闭包可以被访问，这些共享的变量就会继续存在。

显然，Go 语言的闭包特性也是建立在“函数是 一等公民”特性的基础上的。

#### 特征三：作为参数传入函数

既然函数可以存储在变量中，也可以作为返回值返回，那可以理所当然地想到，把函数作为参数传入函数也是可行的。

比如在日常编码时经常使用、标准库 **time 包的 AfterFunc 函数**，就是一个接受函数类型参数的典型例子。

可以看看下面这行代码，这里通过 AfterFunc 函数设置了一个 2 秒的定时器，并传入时间到了后要执行的函数。这里传入的就是一个匿名函数：

```go
time.AfterFunc(time.Second*2, func () { println("timer fired") })
```

#### 特征四：拥有自己的类型

作为一等公民的整型值拥有自己的类型 int，而这个整型值只是类型 int 的一个实例，其他作为一等公民的字符串值、布尔值等类型也都拥有自己类型。那函数呢？ 

在讲解函数声明时，曾得到过这样一个结论：每个函数声明定义的函数仅仅是对 应的函数类型的一个实例，就像var a int = 13这个变量声明语句中的 a，只是 int 类型的一个实例一样。

换句话说，每个函数都和整型值、字符串值等一等公民一样，拥有自己的类型，也就是讲过的函数类型。

甚至可以**基于函数类型来自定义类型**，就像基于整型、字符串类型等类型来自定义类型一样。下面代码中的 HandlerFunc、visitFunc 就是 Go 标准库中，基于函数类型进行自定义的类型：

```go
// net/http/server.go
type HandlerFunc func(ResponseWriter, *Request)

// sort/genzfunc.go
type visitFunc func(ast.Node) ast.Visitor
```

可以看到，Go 函数确实表现出了沃德·坎宁安诠释中“一等公民”的所有特征：Go 函数可以存储在变量中，可以在函数内创建并通过返回值返回，可以作为参数传递给其他函数，可以拥有自己的类型。

通过这些分析，也能感受到，和 C/C++ 等语言中的函数相比，作为“一等公民”的 Go 函数拥有难得的灵活性。 那么在实际生产中，怎么才能发挥出这种灵活性的最大效用，写出更加优雅 简洁的 Go 代码呢？

### 函数“一等公民”特性的高效运用 

#### 应用一：函数类型的妙用 

Go 函数是“一等公民”，也就是说，它拥有自己的类型。

而且，整型、字符串型等所有类型都可以进行的操作，比如**显式转型**，也同样可以用在函数类型上面，也就是说，函数也可以被显式转型。并且，这样的转型在特定的领域具有奇妙的作用，一个最为典型的示例 就是**标准库 http 包中的 HandlerFunc 这个类型**。

来看一个使用了这个类型的例子：

```go
func greeting(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintf(w, "Welcome, Gopher!\n")
}
func main() {
   http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
```

这日常最常见的、用 Go 构建 Web Server 的例子。

它的工作机制也很简单，就是当用户通过浏览器，或者类似 curl 这样的命令行工具，访问 Web server 的 8080 端口时， 会收到“Welcome, Gopher!”这样的文字应答。

可以使用 http 包编写 web server 的方法，先来看一下 **http 包的函数 ListenAndServe 的源码**：

```go
// net/http/server.go
func ListenAndServe(addr string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServe()
}
```

函数 ListenAndServe 会把来自客户端的 http 请求，交给它的第二个参数 handler 处理， 而这里 handler 参数的类型 http.Handler，是一个自定义的接口类型，它的源码是这样 的：

```go
// net/http/server.go
type Handler interface {
   ServeHTTP(ResponseWriter, *Request)
}
```

接口是一组方法的集合，这个接口只有一个方法 ServeHTTP，它的函数类型是func(http.ResponseWriter, *http.Request)。这和自己定义的 http 请求处理函数 greeting 的类型是一致的， 但是**没法直接将 greeting 作为参数值传入**，否则编译器会报错：

```go
func(http.ResponseWriter, *http.Request) does not implement http.Handler (missing ServeHTTP method)
```

这里，编译器提示，函数 greeting 还没有实现接口 Handler 的方法，无法将它赋值给 Handler 类型的参数。

现在再回过头来看下代码，代码中也没有直接将greeting 传给 ListenAndServe 函数，而是将http.HandlerFunc(greeting)作为参数传给了 ListenAndServe。

那这个 http.HandlerFunc 究竟是什么呢？直接来看一下 它的源码：

```go
// net/http/server.go
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

通过它的源码看到，HandlerFunc 是一个基于函数类型定义的新类型，它的底层类型为函数类型func(ResponseWriter, *Request)。

这个类型有一个方法 ServeHTTP， 然后实现了 Handler 接口。也就是说http.HandlerFunc(greeting)这句代码的真正含义，是**将函数 greeting 显式转换为 HandlerFunc 类型，后者实现了 Handler 接口**，满足 ListenAndServe 函数第二个参数的要求。 

> 大概意思明白，但是这块的更深入的调用不明白！

另外，之所以http.HandlerFunc(greeting)这段代码可以通过编译器检查，正是因为 HandlerFunc 的底层类型是func(ResponseWriter, *Request)，与 greeting 函数的类型是一致的，这和下面**整型变量的显式转型原理**也是一样的：

```go
type MyInt int

var x int = 5
y := MyInt(x) // MyInt的底层类型为int，类比HandlerFunc的底层类型为func(ResponseWriter, *Request)
```

#### 应用二：利用闭包简化函数调用

Go 闭包是在函数内部创建的匿名函数，这个匿名函数可以访问创建它的函数的参数与局部变量。

可以利用闭包的这一特性来简化函数调用，这里看一个具 体例子：

```go
func times(x, y int) int {
   return x * y
}
```

在上面的代码中，times 函数用来进行两个整型数的乘法。使用 times 函数的时候需要传入两个实参，比如：

```go
times(2, 5) // 计算2 x 5
times(3, 5) // 计算3 x 5
times(4, 5) // 计算4 x 5
```

不过，有些场景存在一些高频使用的乘数，这个时候就没必要每次都传入这样的高频乘数了。

那怎样能省去高频乘数的传入呢? 看看下面这个新函数 partialTimes：

```go
func partialTimes(x int) func(int) int {
   return func(y int) int {
      return times(x, y)
   }
}
```

这里，partialTimes 的返回值是一个接受单一参数的函数，这个由 partialTimes 函数生成的匿名函数，使用了 partialTimes 函数的参数 x。

按照前面的定义，这个匿名函数就是一 个闭包。partialTimes 实质上就是用来生成以 x 为固定乘数的、接受另外一个乘数作为参 数的、闭包函数的函数。

当程序调用 partialTimes(2) 时，partialTimes 实际上返回了一个调用 times(2,y) 的函数，这个过程的逻辑类似于下面代码：

```go
var timesTwo = func (y int) int {
	return times(2, y)
}
```

这个时候，再看看如何使用 partialTimes，分别生成以 2、3、4 为固定高频乘数的乘 法函数，以及这些生成的乘法函数的使用方法：

```go
func main() {
   // 高级方式
   timesTwo := partialTimes(2)   // 以高频乘数2为固定乘数的乘法函数
   timesThree := partialTimes(3) // 以高频乘数3为固定乘数的乘法函数
   timesFour := partialTimes(4)  // 以高频乘数4为固定乘数的乘法函数
   fmt.Println(timesTwo(5))      // 10，等价于times(2, 5)
   fmt.Println(timesTwo(6))      // 12，等价于times(2, 6)
   fmt.Println(timesThree(5))    // 15，等价于times(3, 5)
   fmt.Println(timesThree(6))    // 18，等价于times(3, 6)
   fmt.Println(timesFour(5))     // 20，等价于times(4, 5)
   fmt.Println(timesFour(6))     // 24，等价于times(4, 6)
}
```

可以看到，通过 partialTimes，生成了三个带有固定乘数的函数。这样，在计 算乘法时，就可以减少参数的重复输入。

>partialTimes 的例子就是传说中的柯里化

看到这里可能会说，这种简化的程度十分有限啊！ 不是的。这里只是举了一个比较好理解的简单例子，在那些动辄就有 5 个以上参数的复杂函数中，减少参数的重复输入给开发人员带去的收益，可要比这个简单的例子大得多。 

### 小结 

Go 代码中的基本功能逻辑单元：函数。

函数这种语法元素的诞生，源于将大问题分解为若干小任务与代码复用。 Go 语言中定义一个函数的最常用方式就是使用函数声明。

函数声明虽然形式上与变量声明不同，但本质其实是一致的，可以通过一个等价转换，将函数声明转 换为一个以函数名为变量名、以函数字面值为初值的函数变量声明形式。这个转换是深入理解函数的关键。 

对函数字面值再进行了拆解。函数字面值是由函数类型与函数体组成的，而函数类型则是由 func 关键字 + 函数签名组成。

再拆解，函数签名又包括函数的参数列表与返回值列表。通常说函数签名时，会省去参数名与返回值变量名，只保留各自的类型信息。 函数签名相同的两个函数类型就是相同的函数类型。

而且，Go 函数采用值传递的方式进行参数传递，对于 string、切片、map 等类型参数来说，这种传递方式传递的仅是“描述符”信息，是一种“浅拷贝”，这点一定要牢记。 

Go 函数支持多返回值，Go 语言的错误处理机制就是建立在多返回值的基础上的。 

最后，与传统的 C、C++、Java 等静态编程语言中的函数相比，Go 函数的最大特点就是它属于 Go 语言的“一等公民”。Go 函数具备一切作为“一等公民”的行为特征，包括函数可以存储在变量中、支持函数内创建函数并通过返回值返回、支持作为参数传递给函数，以及拥有自己的类型等。

这些“一等公民”的特征，让 Go 函数表现出极大的灵活性。日常编码中，也可以利用这些特征进行一些巧妙的代码设计，让代码的实现更简化。 



### 错误处理的设计

多返回值是 Go 语言函数，区别于其他主流静态编程语言中函数的一个重要特点。同时，它也是 Go 语言设计者建构 Go 语言错误处理机制的基础，而错误处理设计也是函数设计的一个重要环节。 

将会从 Go 语言的错误处理机制入手，围绕 Go 语言错误处理机制的原理、Go 错误处理的常见策略，结合函数的多返回值机制进行错误处理的设计。

#### Go 语言是如何进行错误处理的？ 

采用什么错误处理方式，其实是一门编程语言在设计早期就要确定下来的基本机制，它在很大程度上影响着编程语言的语法形式、语言实现的难易程度，以及语言后续的演进方 向。

Go 语言继承了“先祖”C 语言的很多语法特性，在错误处理机制上也不例外，Go 语言错误处理机制也是在 C 语言错误处理机制基础上的再创新。 

从源头讲起，先看看前辈 C 语言的错误处理机制。在 C 语言中，通常**用一个类型为整型的函数返回值作为错误状态标识**，函数调用者会基于值比较的方式，对这一代表错误状态的返回值进行检视。

通常，这个返回值为 0，就代表函数调用成 功；如果这个返回值是其它值，那就代表函数调用出现错误。也就是说，函数调用者需要根据这个返回值代表的错误状态，来决定后续执行哪条错误处理路径上的代码。 

##### C 语言错误处理

C 语言的这种简单的、**基于错误值比较的错误处理机制**有什么优点呢？ 

- 首先，它让每个开发人员必须显式地去关注和处理每个错误，经过显式错误处理的代码会更健壮，也会让开发人员对这些代码更有信心。 
- 另外，也可以发现，这些错误就是普通的值，所以不需要用额外的语言机制去处理它们，只需利用已有的语言机制，像处理其他普通类型值一样的去处理错误就可以了，这也让代码更容易调试，更容易针对每个错误处理的决策分支进行测试覆盖。

C 语言 错误处理机制的这种简单与显式结合的特征，和 Go 语言设计哲学十分契合，于是 Go 语 言设计者决定继承 C 语言这种错误处理机制。 

不过 C 语言这种错误处理机制也有一些弊端。比如，由于 C 语言中的函数最多仅支持一个返回值，很多开发者会把这单一的返回值“一值多用”。什么意思呢？就是说，一个返回值，不仅要承载函数要返回给调用者的信息，又要承载函数调用的最终错误状态。

比如 **C 标准库中的fprintf函数的**返回值就承载了两种含义。

- 在正常情况下，它的返回值表示输出到 FILE 流中的字符数量，
- 但如果出现错误，这个返回值就变成了一个负数，代表具体的错误值：

```c
// stdio.h
int  fprintf(FILE * __restrict, const char * __restrict, ...) __printflike(2, 3);
```

特别是当返回值为其他类型，比如字符串的时候，还很难将它与错误状态融合到一起。这个时候，很多 C 开发人员要么使用输出参数，承载要返回给调用者的信息，要么自定义一个包含返回信息与错误状态的结构体，作为返回值类型。

大家做法不一，就很难形 成统一的错误处理策略。 为了避免这种情况，Go 函数增加了多返回值机制，来支持错误状态与返回信息的分离，并建议开发者把要返回给调用者的信息和错误状态标识，分别放在不同的返回值中。 

##### Go 语言错误处理

继续以上面 C 语言中的 fprintf 函数为例，Go 标准库中有一个和功能等同的 **fmt.Fprintf的函数**，这个函数就是使用一个独立的表示错误状态的返回值（如下面代码中的 err），解决了 fprintf 函数中错误状态值与返回信息耦合在一起的问题：

```go
// fmt/print.go
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
   p := newPrinter()
   p.doPrintf(format, a)
   n, err = w.Write(p.buf)
   p.free()
   return
}
```

在 fmt.Fprintf 中，返回值 n 用来表示写入 io.Writer 中的字节个数，返回值 err 表示这个函数调用的最终状态，如果成功，err 值就为 nil，不成功就为特定的错误值。 

另外还可以看到，fmt.Fprintf 函数声明中代表错误状态的变量 err 的类型，并不是一个传统使用的整数类型，而是用了一个名为 **error 的类型**。 

虽然，在 Go 语言中，依然可以像传统的 C 语言那样，用一个整型值来表示错误状态，但 Go 语言惯用法，是使用 error 这个接口类型表示错误，并且按惯例，通常将 error 类型返回值放在返回值列表的末尾，就像 fmt.Fprintf 函数声明中那样。



#### error 类型与错误值构造 

error 接口是 Go 原生内置的类型，它的定义如下：

```go
// builtin/builtin.go
type error interface {
   Error() string
}
```

任何实现了 error 的 Error 方法的类型的实例，都可以作为错误值赋值给 error 接口变量。 那这里，问题就来了：难道为了构造一个错误值，还需要自定义一个新类型来实现 error 接口吗？ 

##### 字符串错误类型

Go 语言的设计者显然也想到了这一点，在标准库中提供了两种方便 Go 开发者**构造错误值的方法： errors.New和fmt.Errorf**。

使用这两种方法，可以轻松构造出一个 满足 error 接口的错误值，就像下面代码这样：

```go
err := errors.New("your first demo error")
errWithCtx = fmt.Errorf("index %d is out of bounds", i)
```

这两种方法实际上返回的是同一个实现了 error 接口的类型的实例，这个**未导出的类型就是errors.errorString**，它的定义是这样的：

```go
// errors/errors.go
type errorString struct {
   s string
}

func (e *errorString) Error() string {
   return e.s
}
```

大多数情况下，使用这两种方法构建的错误值就可以满足需求了。但也要看 到，虽然这两种构建错误值的方法很方便，但它们给错误处理者提供的**错误上下文（Error Context）只限于以字符串形式**呈现的信息，也就是 Error 方法返回的信息。 

##### 自定义错误类型

但在一些场景下，错误处理者需要从错误值中提取出更多信息，选择错误处理路 径，显然这两种方法就不能满足了。这个时候，可以**自定义错误类型**来满足这一需 求。

比如：标准库中的 **net 包**就定义了一种**携带额外错误上下文的错误类型**：

```go
// net/net.go
type OpError struct {
   Op string
   Net string
   Source Addr
   Addr Addr
   Err error
}
```

这样，错误处理者就可以**根据这个类型的错误值提供的额外上下文信息**，比如 Op、Net、 Source 等，做出**错误处理路径的选择**，比如下面标准库中的代码：

```go
// net/http/server.go
func isCommonNetReadError(err error) bool {
   if err == io.EOF {
      return true
   }
   if neterr, ok := err.(net.Error); ok && neterr.Timeout() {
      return true
   }
   if oe, ok := err.(*net.OpError); ok && oe.Op == "read" {
      return true
   }
   return false
}
```

上面这段代码利用**类型断言**，**判断 error 类型变量 err 的动态类型是否为 *net.OpError 或 net.Error**。

如果 err 的动态类型是 *net.OpError，那么类型断言就会返回这个动态类型的值（存储在 oe 中），代码就可以通过判断它的 Op 字段是否为"read"来判断它是否为 CommonNetRead 类型的错误。 

不过这里，不用过多了解类型断言（Type Assertion）到底是什么，只需要知道**通过类型断言，可以判断接口类型的动态类型**，以及获取它动态类型的值接可以了。

#### error 类型的好处

那么，使用 error 类型，而不是传统意义上的整型或其他类型作为错误类型，有什么好处呢？至少有这三点好处：

- 第一点：**统一了错误类型**。 
  - 如果不同开发者的代码、不同项目中的代码，甚至标准库中的代码，都统一以 error 接口变量的形式呈现错误类型，就能在提升代码可读性的同时，还更容易形成统一的错误处理策略。
- 第二点：**错误是值**。 
  - 构造的错误都是值，也就是说，即便赋值给 error 这个接口类型变量，也可以像整型值那样对错误做“==”和“!=”的逻辑比较，函数调用者检视错误时的体验保持不 变。 
- 第三点：**易扩展，支持自定义错误上下文**。 
  - 虽然错误以 error 接口变量的形式统一呈现，但很容易通过自定义错误类型来扩展错误上下文，就像前面的 Go 标准库的 OpError 类型那样。 
  - error 接口是错误值的提供者与错误值的检视者之间的契约。
  - error 接口的实现者负责提供错误上下文，供负责错误处理的代码使用。这种错误具体上下文与作为错误值类型的 error 接口类型的解耦，也体现了 Go 组合设计哲学中“正交”的理念。 



### 错误处理的惯用策略

#### 策略一：透明错误处理策略 

简单来说，Go 语言中的错误处理，就是根据函数 / 方法返回的 error 类型变量中携带的错误值信息做决策，并选择后续代码执行路径的过程。 

这样，最简单的错误策略莫过于完全不关心返回错误值携带的具体上下文信息，**只要发生错误就进入唯一的错误处理执行路径**，比如下面这段代码：

```go
err := doSomething()
if err != nil {
   // 不关心err变量底层错误值所携带的具体上下文信息
   // 执行简单错误处理逻辑并返回
   // ... ...
   return err
}
```

这也是 Go 语言中最常见的错误处理策略，80% 以上的 Go 错误处理情形都可以归类到这种策略下。

在这种策略下，由于错误处理方并不关心错误值的上下文，所以错误值的构造方（如上面的函数doSomething）可以直接**使用 Go 标准库提供的两个基本错误值构造方法errors.New和fmt.Errorf来构造错误值**，就像下面这样：

```go
func doSomething(...) error {
   // ... ...
   return errors.New("some error occurred")
}
```

这样构造出的错误值代表的上下文信息，对错误处理方是透明的，因此这种策略称为“透明错误处理策略”。

在错误处理方不关心错误值上下文的前提下，透明错误处理策略能最大程度地减少错误处理方与错误值构造方之间的耦合关系。 

#### 策略二：“哨兵”错误处理策略 

当错误处理方不能只根据“透明的错误值”就做出错误处理路径选取的情况下，错误处理方会**尝试对返回的错误值进行检视**，于是就有可能出现下面代码中的反模式：

```go
data, err := b.Peek(1)
if err != nil {
   switch err.Error() {
   case "bufio: negative count":
      // ... ...
      return
   case "bufio: buffer full":
      // ... ...
      return
   case "bufio: invalid use of UnreadByte":
      // ... ...
      return
   default:
      // ... ...
      return
   }
}
```

简单来说，**反模式**就是，错误处理方以透明错误值所能提供的唯一上下文信息（描述错误的字符串），作为错误处理路径选择的依据。

但这种“反模式”会造成严重的**隐式耦合**。 这也就意味着，错误值构造方不经意间的一次错误描述字符串的改动，都会造成错误处理方处理行为的变化，并且这种通过字符串比较的方式，对错误值进行检视的性能也很差。 那这有什么办法吗？

Go 标准库采用了**定义导出的（Exported）“哨兵”错误值**的方式， 来辅助错误处理方检视（inspect）错误值并做出错误处理分支的决策，比如下面的 **bufio 包**中定义的“哨兵错误”：

```go
// bufio/bufio.go
var (
   ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
   ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
   ErrBufferFull        = errors.New("bufio: buffer full")
   ErrNegativeCount     = errors.New("bufio: negative count")
)
```

下面的代码片段利用了上面的哨兵错误，进行错误处理分支的决策：

```go
data, err := b.Peek(1)
if err != nil {
   switch err {
   case bufio.ErrNegativeCount:
      // ... ...
      return
   case bufio.ErrBufferFull:
      // ... ...
      return
   case bufio.ErrInvalidUnreadByte:
      // ... ...
      return
   default:
      // ... ...
      return
   }
}
```

可以看到，一般“哨兵”错误值变量**以 ErrXXX 格式命名**。和透明错误策略相比，“哨 兵”策略让错误处理方在有检视错误值的需求时候，可以“有的放矢”。

不过，对于 API 的开发者而言，暴露“哨兵”错误值也意味着这些错误值和包的公共函数 / 方法一起成为了 API 的一部分。一旦发布出去，开发者就要对它进行很好的维护。 而“哨兵”错误值也让使用这些值的错误处理方对它产生了依赖。 

#### errors.Is

从 Go 1.13 版本开始，**标准库 errors 包提供了 Is 函数**用于错误处理方对错误值的检视。Is 函数类似于把一个 error 类型变量与“哨兵”错误值进行比较，比如下面代码：

```go
// 类似 if err == ErrOutOfBounds{ … }
if errors.Is(err, ErrOutOfBounds) {
   // 越界的错误处理
}
```

不同的是，如果 error 类型变量的底层错误值是一个**包装错误（Wrapped Error）**， errors.Is 方法会沿着该包装错误所在**错误链（Error Chain)**，与链上所有被包装的错误 （Wrapped Error）进行比较，直至找到一个匹配的错误为止。

下面是 Is 函数应用的一个例子：

```go
package main

import (
   "errors"
   "fmt"
)

var ErrSentinel = errors.New("the underlying sentinel error")

func main() {
   err1 := fmt.Errorf("wrap sentinel: %w", ErrSentinel)
   err2 := fmt.Errorf("wrap err1: %w", err1)
   println(err2 == ErrSentinel) // false

   if errors.Is(err2, ErrSentinel) {
      println("err2 is ErrSentinel")
      return
   }

   println("err2 is not ErrSentinel")
}
```

在这个例子中，通过 fmt.Errorf 函数，并且使用 %w 创建包装错误变量 err1 和 err2，其中 err1 实现了对 ErrSentinel 这个“哨兵错误值”的包装，而 err2 又对 err1 进行了包装，这样就形成了一条错误链。

位于错误链最上层的是 err2，位于最底层的是 ErrSentinel。之后，再分别通过值比较和 errors.Is 这两种方法，判断 err2 与 ErrSentinel 的关系。运行上述代码，会看到如下结果：

```sh
false
err2 is ErrSentinel
```

通过比较操作符对 err2 与 ErrSentinel 进行比较后，发现这二者并不相同。而 errors.Is 函数则会沿着 err2 所在错误链，向下找到被包装到最底层的“哨兵”错 误值ErrSentinel。 

所以，如果使用的是 Go 1.13 及后续版本，建议**尽量使用errors.Is方法**去检视某 个错误值是否就是某个预期错误值，或者包装了某个特定的“哨兵”错误值。 

#### 策略三：错误值类型检视策略 

基于 Go 标准库提供的错误值构造方法构造的“哨兵”错误值，除了让错误处理方可以“有的放矢”的进行值比较之外，并没有提供其他有效的错误上下文信息。 

那如果遇到错误处理方需要错误值提供更多的“错误上下文”的情况，上面这些错误处理策略和错误值构造方式都无法满足。 

这种情况下，需要通过**自定义错误类型**的构造错误值的方式，来**提供更多的“错误上下文”信息**。并且，由于错误值都通过 error 接口变量统一呈现，要得到底层错误类型携带的错误上下文信息，错误处理方需要使用 **Go 提供的类型断言机制（Type Assertion） 或类型选择机制（Type Switch）**，这种错误处理方式，称之为错误值类型检视策略。 

来看一个标准库中的例子加深下理解，这个 **json 包中自定义了一个 UnmarshalTypeError 的错误类型**：

```go
// encoding/json/decode.go
type UnmarshalTypeError struct {
   Value  string
   Type   reflect.Type
   Offset int64
   Struct string
   Field  string
}
```

错误处理方可以通过错误类型检视策略，获得**更多错误值的错误上下文信息**，下面就是利用这一策略的 json 包的一个方法的实现：

```go
// encoding/json/decode.go
func (d *decodeState) addErrorContext(err error) error {
   if d.errorContext.Struct != nil || len(d.errorContext.FieldStack) > 0 {
      switch err := err.(type) {
      case *UnmarshalTypeError:
         err.Struct = d.errorContext.Struct.Name()
         err.Field = strings.Join(d.errorContext.FieldStack, ".")
         return err
      }
   }
   return err
}
```

这段代码通过类型 switch 语句得到了 err 变量代表的动态类型和值，然后在匹 配的 case 分支中利用错误上下文信息进行处理。 

这里，一般自定义导出的错误类型**以XXXError的形式命名**。和“哨兵”错误处理策略一样，错误值类型检视策略，由于暴露了自定义的错误类型给错误处理方，因此这些错误类型也和包的公共函数 / 方法一起，成为了 API 的一部分。一旦发布出去，开发者就要对它们进行很好的维护。而它们也让使用这些类型进行检视的错误处理方对其产生了依赖。 

#### errors.As

从 Go 1.13 版本开始，**标准库 errors 包提供了As函数**给错误处理方检视错误值。As函数类似于通过类型断言判断一个 error 类型变量是否为特定的自定义错误类型，如下面代码所示：

```go
// 类似 if e, ok := err.(*MyError); ok { … }
var e *MyError
if errors.As(err, &e) {
   // 如果err类型为*MyError，变量e将被设置为对应的错误值
}
```

不同的是，如果 error 类型变量的动态错误值是一个**包装错误**，errors.As函数会沿着该包装错误所在错误链，与链上所有被包装的错误的类型进行比较，直至找到一个匹配的错误类型，就像 errors.Is 函数那样。

下面是As函数应用的一个例子：

```go
package main

import (
   "errors"
   "fmt"
)

type MyError struct {
   e string
}

func (e *MyError) Error() string {
   return e.e
}

func main() {
   var err = &MyError{"MyError error demo"}
   err1 := fmt.Errorf("wrap err: %w", err)
   err2 := fmt.Errorf("wrap err1: %w", err1)
   var e *MyError

   if errors.As(err2, &e) {
      println("MyError is on the chain of err2")
      println(e == err) // true
      return
   }
   println("MyError is not on the chain of err2")
}
```

运行上述代码会得到：

```sh
MyError is on the chain of err2
true
```

errors.As 函数沿着 err2 所在错误链向下找到了被包装到最深处的错误值， 并将 err2 与其类型 * MyError 成功匹配。**匹配成功后，errors.As 会将匹配到的错误值存储到 As 函数的第二个参数中**，这也是为什么 println(e == err)输出 true 的原因。 

所以，如果使用的是 Go 1.13 及后续版本，请尽量使用errors.As方法去检视某个错误值是否是某自定义错误类型的实例。 

#### 策略四：错误行为特征检视策略 

在前面的三种策略中，其实只有第一种策略，也就是“透明错误处理策略”，有效降低了错误的构造方与错误处理方两者之间的耦合。虽然前面的策略二和策略三，都是实际编码中有效的错误处理策略，但其实使用这两种策略的代码，依然在错误的构造方与错误处理方两者之间建立了**耦合**。 

那么除了“透明错误处理策略”外，是否还有手段可以降低错误处理方与错误值构造 方的耦合呢？ 

在 Go 标准库中，发现了这样一种错误处理方式：将**某个包中的错误类型归类，统一提取出一些公共的错误行为特征，并将这些错误行为特征放入一个公开的接口类型中**。这种方式也被叫做错误行为特征检视策略。 

以**标准库中的net包**为例，它将包内的所有错误类型的**公共行为特征抽象并放入 net.Error 这个接口**中，如下面代码：

```go
// net/net.go
type Error interface {
   error
   Timeout() bool   // Is the error a timeout?
   Temporary() bool // Is the error temporary?
}
```

net.Error 接口包含两个用于判断错误行为特征的方法：

- Timeout 用来判断是否是超时（Timeout）错误，
- Temporary 用于判断是否是临时（Temporary）错误。 

而错误处理方只需要依赖这个公共接口，就可以检视具体错误值的错误行为特征信息，并根据这些信息做出后续错误处理分支选择的决策。 

这里，再看一个 **http 包**使用错误行为特征检视策略进行错误处理的例子，加深下理 解：

```go
// net/http/server.go
func (srv *Server) Serve(l net.Listener) error {
   // ...
   for {
      rw, err := l.Accept()
      if err != nil {
         select {
         case <-srv.getDoneChan():
            return ErrServerClosed
         default:
         }
         if ne, ok := err.(net.Error); ok && ne.Temporary() {
            // 注：这里对临时性(temporary)错误进行处理
            time.Sleep(tempDelay)
            continue
         }
         return err
      }
      // ...
   }
}
```

在上面代码中，Accept 方法实际上返回的错误类型为*OpError，它是 net 包中的一个自定义错误类型，它**实现了错误公共特征接口 net.Error**，如下代码所示：

```go
// net/net.go
type OpError struct {
   // ...
   // Err is the error that occurred during the operation.
   // The Error method panics if the error is nil.
   Err error
}

type temporary interface {
	Temporary() bool
}

func (e *OpError) Temporary() bool {
	// ...
	if ne, ok := e.Err.(*os.SyscallError); ok {
		t, ok := ne.Err.(temporary)
		return ok && t.Temporary()
	}
	t, ok := e.Err.(temporary)
	return ok && t.Temporary()
}
```

因此，OpError 实例可以被错误处理方通过net.Error接口的方法，判断它的行为是否满足 Temporary 或 Timeout 特征。 

### 小结 

Go 函数设计中的一个重要环节：错误处理设计。 

Go 语言继承了 C 语言的基于值比较的错误处理机制，但又在 C 语言的基础上做出了优化，也就是说，Go 函数通过支持多返回值，消除了 C 语言中将错误状态值与返回给函数调用者的信息耦合在一起的弊端。 

Go 语言还统一错误类型为 error 接口类型，并提供了多种快速构建可赋值给 error 类型的错误值的函数，包括 errors.New、fmt.Errorf 等，还讲解了使用统一 error 作为错误 类型的优点。 

基于 Go 错误处理机制、统一的错误值类型以及错误值构造方法的基础上，Go 语言形成了多种错误处理的惯用策略，包括透明错误处理策略、“哨兵”错误处理策略、错误值类型检视策略以及错误行为特征检视策略等。这些策略都有适用的场合，但没有某种单一的错误处理策略可以适合所有项目或所有场合。 

在错误处理策略选择上，一些个人的建议，可以参考一下：

- 尽量使用“透明错误”处理策略，降低错误处理方与错误值构造方之间的耦合； 
- 如果可以通过错误值类型的特征进行错误检视，那么请尽量使用“错误行为特征检视策略”; 
- 在上述两种策略无法实施的情况下，再使用“哨兵”策略和“错误值类型检视”策略； 
- Go 1.13 及后续版本中，尽量用errors.Is和errors.As函数替换原先的错误检视比较语句。



### 函数的健壮性设计

良好的函数设计不仅仅要包含错误设计，函数的健壮性与简洁优雅也是在函数设计过 程中要考虑的问题。 

- 健壮的函数意味着，无论调用者如何使用函数，函数都能以合理的方式处理调用 者的任何输入，并给调用者返回预设的、清晰的错误值。即便函数发生内部异常，函数也会尽力从异常中恢复，尽可能地不让异常蔓延到整个程序。
- 而简洁优雅，则意味着函数的实现易读、易理解、更易维护，同时简洁也意味着统计意义上的更少的 bug。

#### 健壮性的“三不要”原则 

函数的健壮性设计包括很多方面，首先就有最基本的“三不要”原则，简单来分析一 下。 

##### 原则一：不要相信任何外部输入的参数

函数的使用者可能是任何人，这些人在使用函数之前可能都没有阅读过任何手册或文档， 会向函数传入意想不到的参数。

因此，为了保证函数的健壮性，函数需要对所有输 入的参数进行合法性的检查。一旦发现问题，立即终止函数的执行，返回预设的错误值。 

##### 原则二：不要忽略任何一个错误

在函数实现中，也会调用标准库或第三方包提供的函数或方法。对于这些调用，不能假定它一定会成功，一定要显式地检查这些调用返回的错误值。

一旦发现错误，要及时终止函数执行，防止错误继续传播。 

##### 原则三：不要假定异常不会发生

这里，先要确定一个认知：异常不是错误。

- 错误是可预期的，也是经常会发生的，有对应的公开错误码和错误处理预案，但异常却是少见的、意料之外的。
- 通常意义上的异常，指的是硬件异常、操作系统异常、语言运行时异常，还有更大可能是代码中潜在 bug 导致的异常，比如代码中出现了以 0 作为分母，或者是数组越界访问等情况。 

虽然异常发生是“小众事件”，但是不能假定异常不会发生。所以，函数设计时，就需要根据函数的角色和使用场景，考虑是否要在函数内设置异常捕捉和恢复的环节。 

在这三条健壮性设计原则中，做到前两条是相对容易的，也没有太多技巧可言。但对第三条异常的处理（ Go 函数的异常处理设计），很多初学者拿捏不好。

#### Go 语言中的异常：panic 

不同编程语言表示**异常（Exception）**这个概念的语法都不相同。

在 Go 语言中，异常这个概念由 panic 表示。一些教程或文章会把它译为恐慌，这里依旧选择不译，保留 panic 的原汁原味。 

panic 指的是 Go 程序在运行时出现的一个异常情况。如果异常出现了，但没有被捕获并恢复，Go 程序的执行就会被终止，即便出现异常的位置不在主 Goroutine 中也会这样。 

在 Go 中，panic 主要有**两类来源**，

- 一类是来自 **Go 运行时**，
- 另一类则是 **Go 开发人员通 过 panic 函数主动触发**的。

无论是哪种，一旦 panic 被触发，后续 Go 程序的执行过程都是一样的，这个过程被 Go 语言称为 **panicking**。 

Go 官方文档以手工调用 panic 函数触发 panic 为例，对 panicking 这个过程进行了诠释：

- 当函数 F 调用 panic 函数时，函数 F 的执行将停止。
- 不过，函数 F 中已进行求值的 deferred 函数都会得到正常执行，执行完这些 deferred 函数后，函数 F 才会把控制权返还给其调用者。 
- 对于函数 F 的调用者而言，函数 F 之后的行为就如同调用者调用的函数是 panic 一样，该 panicking过程将继续在栈上进行下去，直到当前 Goroutine 中的所有函数都返回为止，
- 然后 Go 程序将崩溃退出。 

用一个**例子**来更直观地解释一下 panicking 这个过程：

```go
package main

func foo() {
   println("call foo")
   bar()
   println("exit foo")
}

func bar() {
   println("call bar")
   panic("panic occurs in bar")
   zoo()
   println("exit bar")
}

func zoo() {
   println("call zoo")
   println("exit zoo")
}

func main() {
   println("call main")
   foo()
   println("exit main")
}
```

上面这个例子中，从 Go 应用入口开始，函数的调用次序依次为main -> foo -> bar -> zoo。在 bar 函数中，调用 panic 函数手动触发了 panic。 

执行这个程序的输出结果是这样的：

```go
call main
call foo
call bar
panic: panic occurs in bar
```

再根据前面对 panicking 过程的诠释，理解一下这个例子。 

- 程序从入口函数 main 开始依次调用了 foo、bar 函数，在 bar 函数中，代码在调用 zoo 函数之前调用了 panic 函数触发了异常。那示例的 panicking 过程就从这开始了。 
- bar 函数调用 panic 函数之后，它自身的执行就此停止了，所以也没有看到代码继续进入 zoo 函数执行。
- 并且，bar 函数没有捕捉这个 panic，这样这个 panic 就会沿着函数调用栈向上走，来到了 bar 函数的调用者 foo 函数中。 
- 从 foo 函数的视角来看，这就好比将它对 bar 函数的调用，换成了对 panic 函数的调用一样。这样一来，foo 函数的执行也被停止了。
- 由于 foo 函数也没有捕捉 panic，于是 panic 继续沿着函数调用栈向上走，来到了 foo 函数的调用者 main 函数中。 
- 同理，从 main 函数的视角来看，这就好比将它对 foo 函数的调用，换成了对 panic 函数的调用一样。
- 结果就是，main 函数的执行也被终止了，于是整个程序异常退出，日志"exit main"也没有得到输出的机会。

#### recover 函数

不过，Go 也提供了捕捉 panic 并恢复程序正常执行秩序的方法，可以通过 **recover 函数**来实现这一点。

继续用上面这个例子分析，在触发 panic 的 bar 函数中，对 panic 进行捕捉并恢复， 直接来看恢复后，整个程序的执行情况是什么样的。

这里，列出了变更后的代码：

```go
package main

import "fmt"

func foo() {
   println("call foo")
   bar()
   println("exit foo")
}

func bar() {
   defer func() {
      if e := recover(); e != nil {
         fmt.Println("recover the panic:", e)
      }
   }()
   
   println("call bar")
   panic("panic occurs in bar")
   zoo()
   println("exit bar")
}

func zoo() {
   println("call zoo")
   println("exit zoo")
}

func main() {
   println("call main")
   foo()
   println("exit main")
}
```

在更新版的 bar 函数中，在一个 defer 匿名函数中调用 recover 函数对 panic 进行了捕捉。

**recover 是 Go 内置的专门用于恢复 panic 的函数**，它必须被放在一个 defer 函数 中才能生效。如果 recover 捕捉到 panic，它就会返回以 panic 的具体内容为错误上下文信息的错误值。

如果没有 panic 发生，那么 recover 将返回 nil。而且，如果 panic 被 recover 捕捉到，panic 引发的 panicking 过程就会停止。 

无论 bar 函数正常执行结束，还是因 panic 异常终止，在那之前设置成功的 defer 函数都会得到执行。 

执行更新后的程序，得到如下结果：

```go
call main
call foo
call bar
recover the panic: panic occurs in bar
exit foo
exit main
```

可以看到 main 函数终于得以“善终”。那这个过程中究竟发生了什么呢？ 

- 在更新后的代码中，当 bar 函数调用 panic 函数触发异常后，bar 函数的执行就会被中断。
- 但这一次，在代码执行流回到 bar 函数调用者之前，bar 函数中的、在 panic 之前就已经被设置成功的 derfer 函数就会被执行。这个匿名函数会调用 recover 把刚刚触发的 panic 恢复，这样，panic 还没等沿着函数栈向上走，就被消除了。 
- 所以，这个时候，从 foo 函数的视角来看，bar 函数与正常返回没有什么差别。foo 函数依旧继续向下执行，直至 main 函数成功返回。
- 这样，这个程序的 panic“危机”就解除 了。 

面对有如此行为特点的 panic，应该如何应对呢？是不是在所有 Go 函数或方法中， 都要用 defer 函数来捕捉和恢复 panic 呢？ 

#### 如何应对 panic

其实大可不必。 

- 一来，这样做会徒增开发人员函数实现时的心智负担。
- 二来，很多函数非常简单，根本不会出现 panic 情况，增加 panic 捕获和恢复，反倒会增加函数的复杂性。
- 同时， defer 函数也不是“免费”的，也有带来性能开销。 

那么，日常情况下应该怎么做呢？这里提供了三点经验。 

##### 第一点：评估程序对 panic 的忍受度 

首先，应该知道一个事实：不同应用对异常引起的程序崩溃退出的忍受度是不一样 的。

- 比如，一个单次运行于控制台窗口中的命令行交互类程序（CLI），和一个常驻内存的后端 HTTP 服务器程序，对异常崩溃的忍受度就是不同的。 
- 前者即便因异常崩溃，对用户来说也仅仅是再重新运行一次而已。
- 但后者一旦崩溃，就很可能导致整个网站停止服务。

所以，**针对各种应用对 panic 忍受度的差异，采取的应对 panic 的策略也应该有不同**。

- 像后端 HTTP 服务器程序这样的任务关键系统，就需要在特定位置捕捉并恢复 panic，以保证服务器整体的健壮度。
- 在这方面，Go 标准库中的 http server 就是一个典型的代表。

**Go 标准库提供的 http server** 采用的是，每个客户端连接都使用一个单独的 Goroutine 进行处理的并发处理模型。也就是说，客户端一旦与 http server 连接成功，http server 就会为这个连接新创建一个 Goroutine，并在这 Goroutine 中执行对应连接（conn）的 serve 方法，来处理这条连接上的客户端请求。 

因为，无论在哪个 Goroutine 中发生未被恢复的 panic，整个程序都将崩溃退出。所以，为了保证处理某一个客户端连接的 Goroutine 出现 panic 时，不影响到 http server 主 Goroutine 的运行，**Go 标准库在 serve 方法中加入了对 panic 的捕捉与恢复**，下面是 serve 方法的部分代码片段：

```go
// net/http/server.go
// Serve a new connection.
func (c *conn) serve(ctx context.Context) {
   c.remoteAddr = c.rwc.RemoteAddr().String()
   ctx = context.WithValue(ctx, LocalAddrContextKey, c.rwc.LocalAddr())
   defer func() {
      if err := recover(); err != nil && err != ErrAbortHandler {
         const size = 64 << 10
         buf := make([]byte, size)
         buf = buf[:runtime.Stack(buf, false)]
         c.server.logf("http: panic serving %v: %v\n%s", c.remoteAddr, err, buf)
      }
      if !c.hijacked() {
         c.close()
         c.setState(c.rwc, StateClosed, runHooks)
      }
   }()
   // ...
}
```

可以看到，serve 方法在一开始处就设置了 defer 函数，并在该函数中捕捉并恢复了可能出现的 panic。

这样，即便处理某个客户端连接的 Goroutine 出现 panic，处理其他连接 Goroutine 以及 http server 自身都不会受到影响。 

这种局部不要影响整体的异常处理策略，在很多并发程序中都有应用。并且，**捕捉和恢复 panic 的位置通常都在子 Goroutine 的起始处**，这样设置可以捕捉到后面代码中可能出现的所有 panic，就像 serve 方法中那样。 

##### 第二点：提示潜在 bug

有了对 panic 忍受度的评估，panic 是不是也没有那么“恐怖”了呢？而且，甚至可 以借助 panic 来快速找到潜在 bug。 

**C 语言**中有个很好用的**辅助函数，断言（assert 宏）**。在使用 C 编写代码时，经常在一些代码执行路径上，使用断言来表达这段执行路径上某种条件一定为真的信心。断言为真，则程序处于正确运行状态，断言为否就是出现了意料之外的问题，而这个问题很可能就是一个潜在的 bug，这时可以借助断言信息快速定位到问题所在。 

不过，**Go 语言**标准库中并没有提供断言之类的辅助函数，但可以**使用 panic**，部分模拟断言对潜在 bug 的提示功能。比如，下面就是**标准库 encoding/json包**使用 panic 指示潜在 bug 的一个例子：

```go
// encoding/json/decode.go

// 当一些本不该发生的事情导致我们结束处理时，phasePanicMsg 将被用作panic消息，它可以指示JSON解码器中的bug，或者在解码器执行时还有其他代码正在修改数据切片。
// phasePanicMsg is used as a panic message when we end up with something that shouldn't happen. It can indicate a bug in the JSON decoder, or that something is editing the data slice while the decoder executes.
const phasePanicMsg = "JSON decoder out of sync - data changing underfoot?"

func (d *decodeState) init(data []byte) *decodeState {
	d.data = data
	d.off = 0
	d.savedError = nil
	d.errorContext.Struct = nil

	// Reuse the allocated space for the FieldStack slice.
	d.errorContext.FieldStack = d.errorContext.FieldStack[:0]
	return d
}

func (d *decodeState) valueQuoted() interface{} {
	switch d.opcode {
	default:
		panic(phasePanicMsg)

	case scanBeginArray, scanBeginObject:
		d.skip()
		d.scanNext()

	case scanBeginLiteral:
		v := d.literalInterface()
		switch v.(type) {
		case nil, string:
			return v
		}
	}
	return unquotedValue{}
}
```

在valueQuoted这个方法中，如果程序执行流进入了 default 分支，那这个方法就会引发 panic，这个 panic 会提示开发人员：这里很可能是一个 bug。 

同样，在 **json 包的 encode.go** 中也有使用 panic 做潜在 bug 提示的例子：

```go
// encoding/json/encode.go
func (w *reflectWithString) resolve() error {
   // ...
   switch w.v.Kind() {
   case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
      w.s = strconv.FormatInt(w.v.Int(), 10)
      return nil
   case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
      w.s = strconv.FormatUint(w.v.Uint(), 10)
      return nil
   }
   panic("unexpected map key type")
}
```

这段代码中，resolve方法的最后一行代码就相当于一个“**代码逻辑不会走到这里”的断言**。一旦触发“断言”，这很可能就是一个潜在 bug。

去掉这行代码并不会对resolve方法的逻辑造成任何影响，但真正出现问题时，开发人员就缺少了“断言”潜在 bug 提醒的辅助支持了。在 **Go 标准库中，大多数 panic 的使用都是充当类似断言的作用的**。 

##### 第三点：不要混淆异常与错误 

在日常编码中，经常会看到一些 Go 语言初学者，尤其是一些有过 Java 语言编程经验的程序员，会因为习惯了 Java 那种基于try-catch-finally的错误处理思维，而将 Go panic 当成 Java 的“checked exception”去用，这显然是混淆了 Go 中的异常与错误， 这是 Go 错误处理的一种反模式。 

查看 **Java 标准类库**，可以看到一些 Java 已预定义好的checked exception类，比 较常见的有IOException、TimeoutException、EOFException、 FileNotFoundException，等等。

看到这里，是不是觉得这些 checked exception 和 “哨兵错误值”十分相似呢？。它们都是预定义好的、代表特定场景下的错误状态。 

那 **Java 的checked exception 和 Go 中的 panic 有啥差别**呢？ 

- **Java 的checked exception**用于一些可预见的、常会发生的错误场景，比如，针对 checked exception的所谓“异常处理”，就是针对这些场景的“错误处理预案”。
- 也 可以说对checked exception的使用、捕获、自定义等行为都是“有意而为之”的。 如果它非要和 Go 中的某种语法对应来看，它对应的也是 Go 的错误处理，也就是基于 error 值比较模型的错误处理。
- 所以，Java 中对checked exception处理的本质是错误处理，虽然它的名字用了带有“异常”的字样。 
- 而 **Go 中的 panic** 呢，更接近于 Java 的RuntimeException+Error，而不是checked exception。
- 前面提到过 Java 的checked exception是必须要被上层代码处理的，也就是要么捕获处理，要么重新抛给更上层。
- 但是在 Go 中，通常会导入大量第三方包，而对于这些第三方包 API 中是否会引发panic，是不知道的。 
- 因此上层代码，也就是 API 调用者根本不会去逐一了解 API 是否会引发panic，也没有义务去处理引发的 panic。一旦在编写的 API 中，像checked exception那样使用 panic 作为正常错误处理的手段，把引发的panic当作错误，那么就会给的API 使用者带去大麻烦！
- 因此，在 Go 中，作为 API 函数的作者，一定不要将 panic 当作错误返回给 API 调用者。 

### 函数的简洁性设计

#### 使用 defer 简化函数

对函数设计来说，如何实现简洁的目标是一个大话题。

- 可以从通用的**设计原则**去谈，比如函数要遵守单一职责，职责单一的函数肯定要比担负多种职责的函数更简单。
- 也可以从**函数实现的规模**去谈，比如函数体的规模要小，尽量控制在 80 行代码之内等。 

那么，Go 中是否有**现成的语法元素**， 可以简化 Go 函数的设计和实现。有的，它就是 defer。 

同样地，也用一个具体的**例子**来理解一下。日常开发中，经常会编写一些类似下面示例中的**伪代码**：

```go
package main

import "sync"

func doSomething() error {
   var mu sync.Mutex
   mu.Lock()

   r1, err := OpenResource1()
   if err != nil {
      mu.Unlock()
      return err
   }

   r2, err := OpenResource2()
   if err != nil {
      r1.Close()
      mu.Unlock()
      return err
   }

   r3, err := OpenResource3()
   if err != nil {
      r2.Close()
      r1.Close()
      mu.Unlock()
      return err
   }

   // 使用r1，r2, r3
   err = doWithResources()
   if err != nil {
      r3.Close()
      r2.Close()
      r1.Close()
      mu.Unlock()
      return err
   }

   r3.Close()
   r2.Close()
   r1.Close()
   mu.Unlock()
   return nil
}
```

这类代码的特点就是**在函数中会申请一些资源**，并在**函数退出前释放或关闭这些资源**，比如这里的互斥锁 mu 以及资源 r1~r3 就是这样。 

函数的实现需要确保，无论函数的执行流是按预期顺利进行，还是出现错误，这些资源在函数退出时都要被及时、正确地释放。为此，需要尤为**关注函数中的错误处理**，在错误处理时不能遗漏对资源的释放。 

但这样的要求，就导致在进行资源释放，尤其是有多个资源需要释放的时候，比如上面示例那样，会大大增加开发人员的心智负担。同时当待释放的资源个数较多时，整个代码逻辑就会变得十分复杂，程序可读性、健壮性也会随之下降。

但即便如此，如果函数实现中的某段代码逻辑抛出 panic，传统的错误处理机制依然没有办法捕获它并尝试从 panic 恢复。 

Go 语言引入 defer 的初衷，就是解决这些问题。那么，defer 具体是怎么解决这些问题的呢？或者说，defer 具体的运作机制是怎样的呢？ 

**defer 是 Go 语言提供的一种延迟调用机制**，defer 的运作离不开函数。怎么理解呢？这句话至少有以下两点含义：

- 在 Go 中，只有在函数（和方法）内部才能使用 defer； 
- defer 关键字后面只能接函数（或方法），这些函数被称为 **deferred 函数**。defer 将它们注册到其所在 Goroutine 中，用于存放 deferred 函数的栈数据结构中，这些 deferred 函数将在执行 defer 的函数退出前，按**后进先出（LIFO）的顺序**被程序调度执行（如下图所示）。

![image-20211229174423436](go_function.assets/image-20211229174423436.png)

而且，无论是执行到函数体尾部返回，还是在某个错误处理分支显式 return，又或是出现 panic，**已经存储到 deferred 函数栈中的函数，都会被调度执行**。

所以说，deferred 函数是一个可以在任何情况下为函数进行**收尾工作**的好“伙伴”。 

回到刚才的那个例子，如果**把收尾工作挪到 deferred 函数**中，那么代码将变成如 下这个样子：

```go
package deferprocessmu

import "sync"

func doSomething() error {
   var mu sync.Mutex
   mu.Lock()
   defer mu.Unlock()

   r1, err := OpenResource1()
   if err != nil {
      return err
   }
   defer r1.Close()

   r2, err := OpenResource2()
   if err != nil {
      return err
   }
   defer r2.Close()

   r3, err := OpenResource3()
   if err != nil {
      return err
   }
   defer r3.Close()

   // 使用r1，r2, r3
   return doWithResources()
}
```

使用 defer 后对函数实现逻辑的简化是显而易见的。而且，这里资源释放函数 的 defer 注册动作，紧邻着资源申请成功的动作，这样**成对出现的惯例**就极大降低了遗漏资源释放的可能性，开发人员也不用再小心翼翼地在每个错误处理分支中检查是否遗漏了某个资源的释放动作。

同时，代码的简化也意味代码可读性的提高，以及代码健壮度的增强。 

#### defer 使用的几个注意事项

大多数 Gopher 都喜欢 defer，因为它不仅可以用来捕捉和恢复 panic，还能让函数变得 更简洁和健壮。但“工欲善其事，必先利其器“，一旦要用 defer，有几个关于 defer 使用的注意事项是一定要提前了解清楚的，可以避免掉进一些不必要的“坑”。 

##### 第一点：明确哪些函数可以作为 deferred 函数 

对于自定义的函数或方法，defer 可以给与无条件的支持，但是对于有返回值的自定义函数或方法，**返回值会在 deferred 函数被调度执行的时候被自动丢弃**。

而且，Go 语言中除了自定义函数 / 方法，还有 Go 语言内置的 / 预定义的函数，这里给出了 **Go 语言内置函数的完全列表**：

```go
Functions:
append cap close complex copy delete imag len
make new panic print println real recover
```

那么，Go 语言中的内置函数是否都能作为 deferred 函数呢？看下面的示例：

```go
package main

func bar() (int, int) {
   return 1, 2
}

func foo() {
   var c chan int
   var sl []int
   var m = make(map[string]int, 10)
   m["item1"] = 1
   m["item2"] = 2
   var a = complex(1.0, -1.4)
   var sl1 []int

   defer bar()
   defer append(sl, 11)
   defer cap(sl)
   defer close(c)
   defer complex(2, -2)
   defer copy(sl1, sl)
   defer delete(m, "item2")
   defer imag(a)
   defer len(sl)
   defer make([]int, 10)
   defer new(*int)
   defer panic(1)
   defer print("hello, defer\n")
   defer println("hello, defer")
   defer real(a)
   defer recover()
}

func main() {
   foo()
}
```

运行这个示例代码，可以得到：

```go
./main.go:17:2: defer discards result of append(sl, 11)
./main.go:18:2: defer discards result of cap(sl)
./main.go:20:2: defer discards result of complex(2, -2)
./main.go:23:2: defer discards result of imag(a)
./main.go:24:2: defer discards result of len(sl)
./main.go:25:2: defer discards result of make([]int, 10)
./main.go:26:2: defer discards result of new(*int)
./main.go:30:2: defer discards result of real(a)
```

Go 编译器居然给出一组错误提示！ 

从这组错误提示中可以看到，append、cap、len、make、new、imag 等内置函数都是不能直接作为 deferred 函数的，而 close、copy、delete、print、recover 等内置函数则可以直接被 defer 设置为 deferred 函数。 

不过，对于那些不能直接作为 deferred 函数的内置函数，可以使用一个**包裹它的匿名函数来间接满足要求**，以 append 为例是这样的：

```go
// 包裹 不能直接作为 deferred 函数的内置函数 
var sl []int
defer func() {
   _ = append(sl, 11)
}()
```

##### 第二点：注意 defer 关键字后面表达式的求值时机 

这里，一定要牢记一点：defer 关键字后面的表达式，是在将 deferred 函数注册到 deferred 函数栈的时候进行求值的。

同样用一个典型的例子来说明一下 **defer 后表达式的求值时机**：

```go
package main

import "fmt"

func foo1() {
   for i := 0; i <= 3; i++ {
      defer fmt.Println(i)
   }
}

func foo2() {
   for i := 0; i <= 3; i++ {
      defer func(n int) {
         fmt.Println(n)
      }(i)
   }
}

func foo3() {
   for i := 0; i <= 3; i++ {
      defer func() {
         fmt.Println(i)
      }()
   }
}

func main() {
   fmt.Println("foo1 result:")
   foo1()
   fmt.Println("\nfoo2 result:")
   foo2()
   fmt.Println("\nfoo3 result:")
   foo3()
}
```

这里，一个个分析 foo1、foo2 和 foo3 中 defer 后的表达式的求值时机。 

**首先是 foo1**。foo1 中 defer 后面直接用的是 fmt.Println 函数，每当 defer 将 fmt.Println 注册到 deferred 函数栈的时候，都会对 Println 后面的参数进行求值。根据上述代码逻辑，依次压入 deferred 函数栈的函数是：

```go
fmt.Println(0)
fmt.Println(1)
fmt.Println(2)
fmt.Println(3)
```

因此，当 foo1 返回后，deferred 函数被调度执行时，上述压入栈的 deferred 函数将以 LIFO 次序出栈执行，这时的输出的结果为：

```sh
foo1 result:
3
2
1
0
```

**然后再看 foo2**。foo2 中 defer 后面接的是一个带有一个参数的匿名函数。每当 defer 将匿名函数注册到 deferred 函数栈的时候，都会对该匿名函数的参数进行求值。根据上述代码逻辑，依次压入 deferred 函数栈的函数是：

```go
func(0)
func(1)
func(2)
func(3)
```

因此，当 foo2 返回后，deferred 函数被调度执行时，上述压入栈的 deferred 函数将以 LIFO 次序出栈执行，因此输出的结果为：

```sh
foo2 result:
3
2
1
0
```

**最后来看 foo3**。foo3 中 defer 后面接的是一个不带参数的匿名函数。根据上述代码逻辑，依次压入 deferred 函数栈的函数是：

```go
func()
func()
func()
func()
```

所以，当 foo3 返回后，deferred 函数被调度执行时，上述压入栈的 deferred 函数将以 LIFO 次序出栈执行。匿名函数会以**闭包的方式**访问外围函数的变量 i，并通过 Println 输出 i 的值，此时 i 的值为 4，因此 foo3 的输出结果为：

```sh
foo3 result:
4
4
4
4
```

通过这些例子，可以看到，无论以何种形式将函数注册到 defer 中，deferred 函数的参数值都是在注册的时候进行求值的。

##### 第三点：知晓 defer 带来的性能损耗 

通过前面的分析，可以看到，defer 让资源释放（如文件描述符、锁）的过程变得优雅很多，也不易出错。但在**性能敏感的应用**中，defer 带来的性能负担也是必须 要知晓和权衡的问题。 

这里，用一个**性能基准测试（Benchmark）**，直观地看看 defer 究竟会带来多少性能损耗。

基于 Go 工具链，可以很方便地为 Go 源码写一个性能基准测试，只需将代码放在以“_test.go”为后缀的源文件中，然后利用 testing 包提供的“框架”就可以了，看下面代码：

```go
package main

import "testing"

func sum(max int) int {
   total := 0
   for i := 0; i < max; i++ {
      total += i
   }
   return total
}

func fooWithDefer() {
   defer func() {
      sum(10)
   }()
}

func fooWithoutDefer() {
   sum(10)
}

func BenchmarkFooWithDefer(b *testing.B) {
   for i := 0; i < b.N; i++ {
      fooWithDefer()
   }
}

func BenchmarkFooWithoutDefer(b *testing.B) {
   for i := 0; i < b.N; i++ {
      fooWithoutDefer()
   }
}
```

这个基准测试包含了两个测试用例，分别是 BenchmarkFooWithDefer 和 BenchmarkFooWithoutDefer。前者测量的是带有 defer 的函数执行的性能，后者测量的是不带有 defer 的函数的执行的性能。 

在 Go 1.13 前的版本中，defer 带来的开销还是很大的。先用 **Go 1.12.7** 版本来运行一下上述基准测试，会得到如下结果：

```sh
$go1.12.7 test -bench . defer_test.go
goos: darwin
goarch: amd64
BenchmarkFooWithDefer-8         30000000                46.7 ns/op
BenchmarkFooWithoutDefer-8      300000000                5.53 ns/op
PASS
ok      command-line-arguments  3.679s
```

从这个基准测试结果中，可以清晰地看到：使用 defer 的函数的执行时间是没有使用 defer 函数的 8 倍左右。 

但从 Go 1.13 版本开始，Go 核心团队对 defer 性能进行了多次优化，到现在的 **Go 1.17 版本**，defer 的开销已经足够小了。看看使用 Go 1.17 版本运行上述基准测试的结 果：

```sh
$go1.17.5 test -bench . defer_test.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkFooWithDefer-8         175664768                7.272 ns/op
BenchmarkFooWithoutDefer-8      207940688                5.941 ns/op
PASS
ok      command-line-arguments  4.936s
```

带有 defer 的函数执行开销，仅是不带有 defer 的函数的执行开销的 1.45 倍左 右，已经达到了几乎可以忽略不计的程度，可以放心使用。



### 小结

函数设计过程中应该考虑健壮性与简洁性方面的内容。 

在函数健壮性方面，给出了“三不要”原则，这三个原则一定要记住。这里重点 讲解了第三个原则：不要假定异常不会发生。

借此，认识了 Go 语言中表示异常的 panic，也学习了 panic 发生后的代码执行流程。基于 panic 的行为特征，给出了 Go 函数设计过程中应对 panic 的三点经验，这里要注意，“评估程序对 panic 的忍受度”是选取应对 panic 措施的前提。 

另外，对于来自像 Java 这样的、基于 Exception 进行错误处理的编程语言的 Go 初学者们，切记不要将 panic 与错误处理混淆。 

接下来，讲解了如何让函数实现更加简洁。简洁性对于函数来说意味着可读性更 好，更易于理解，也有利于代码健壮性的提升。

Go 语言层面提供的 defer 机制可用于简化函数实现，尤其是在函数申请和释放资源个数较多的情况下。 

如果要用好 defer，前提就是要了解 defer 的运作机制，这里要把握住两点：

- 函数返回前，deferred 函数是按照后入先出（LIFO）的顺序执行的； 
- defer 关键字是在注册函数时对函数的参数进行了求值。

最后，在最新 Go 版本 Go1.17 中，使用 defer 带来的开销几乎可以忽略不计了，可以放心使用。 








