# Go Method

> Go 语言的方法。

函数是 Go 代码中的基本功能逻辑单元，它承载了 Go 程序的所有执行逻辑。可以说，Go 程序的执行流本质上就是在函数调用栈中上下流动，从一个函数到另一个函数。 

Go 语言还有一种语法元素，方法（method），它也可以承载代码逻辑，程序也可以从一个方法流动到另外一个方法。 

系统讲解 Go 语言中的方法，将围绕方法的本质、方法 receiver 的类型选择、方法集合，以及如何实现方法的“继承”这几个主题。 

## Method

### 认识 Go 方法 

Go 语言从设计伊始，就不支持经典的面向对象语法元素，比如类、对象、继 承，等等，但 Go 语言仍保留了名为“方法（method）”的语法元素。

当然，Go 语言中 的方法和面向对象中的方法并不是一样的。Go 引入方法这一元素，并不是要支持面向对象编程范式，而是 Go 践行**组合设计哲学**的一种实现层面的需要。

#### Go 方法的一般形式

简单了解之后，就以 **Go 标准库 net/http 包**中 *Server 类型的方法 ListenAndServeTLS 为例，讲解一下 Go 方法的一般形式：

![image-20211230110657518](go_method.assets/image-20211230110657518.png)

Go 中方法的声明和函数的声明有很多相似之处，可以参照着来学习。比如，Go 的方法也是以 func 关键字修饰的，并且和函数一样，也包含方法名（对应函数名）、参数列表、返回值列表与方法体（对应函数体）。 

而且，方法中的这几个部分和函数声明中对应的部分，在形式与语义方面都是一致的，比如：方法名字首字母大小写决定该方法是否是导出方法；方法参数列表支持变长参数；方法的返回值列表也支持具名返回值等。

不过，它们也有**不同的地方**。从上面这张图可以看到，和由五个部分组成的函数声明不同，Go 方法的声明有六个组成部分，多的一个就是图中的 **receiver 部分**。

在 receiver 部分声明的参数，Go 称之为 receiver 参数，这个 receiver 参数也是方法与类型之间的纽带，也是方法与函数的最大不同。

#### receiver 参数

Go 中的方法必须是归属于一个类型的，而 receiver 参数的类型就是这个方法归属的类型，或者说这个方法就是这个类型的一个方法。

以上图中的 ListenAndServeTLS 为例，这里的 receiver 参数 srv 的类型为 *Server，那么可以说，这个方法就是 *Server 类型的方法， 注意！这里说的是 ListenAndServeTLS 是 *Server 类型的方法，而不是 Server 类型的方法。

为了方便讲解，将上面例子中的方法声明，转换为一个**方法的一般声明形式**：

```go
func (t *T或T) MethodName(参数列表) (返回值列表) {
  // 方法体
}
```

无论 receiver 参数的类型为 *T 还是 T，都把一般声明形式中的 T 叫做 receiver 参数 t 的基类型。

- 如果 t 的类型为 T，那么说这个方法是类型 T 的一个方法；
- 如果 t 的类型为 *T，那么就说这个方法是类型 *T 的一个方法。

而且，要注意的是，**每个方法只能有一个 receiver 参数**，Go 不支持在方法的 receiver 部分放置包含多个 receiver 参数的参数列表，或者变长 receiver 参数。 

那么，**receiver 参数的作用域**是什么呢？ 

关于函数 / 方法作用域的结论：方法接收器（receiver）参数、函数 / 方法参数，以及返回值变量对应的作用域范围，都是函数 / 方法体对应的显式代码块。

这就意味着，receiver 部分的参数名不能与方法参数列表中的形参名，以及具名返回值中的变量名存在冲突，必须在这个方法的作用域中具有**唯一性**。

如果这个唯一不存在，比如像下面例子中那样，Go 编译器就会报错：

```go
// 参数名与方法参数的形参名重复
type T struct{}

func (t T) M(t string) { // 编译器报错：duplicate argument t (重复声明参数t)
	// ... ...
}
```

不过，如果在方法体中，没有用到 receiver 参数，也可以**省略 receiver 的参数 名**，就像下面这样：

```go
// 省略 receiver 的参数 名
type T struct{}

func (T) M(t string) {
	// ... ...
}
```

仅当方法体中的实现不需要 receiver 参数参与时，才会省略 receiver 参数名，不过这一情况很少使用，了解一下就好了。 

#### receiver 参数的基类型

除了 receiver 参数名字要保证唯一外，Go 语言对 receiver 参数的基类型也有约束，那就是 **receiver 参数的基类型**本身不能为指针类型或接口类型。

下面的例子分别演示了基类型为指针类型和接口类型时，Go 编译器报错的情况：

```go
// receiver 参数的基类型为指针类型
type MyInt *int

func (r MyInt) String() string { // r的基类型为MyInt，编译器报错：invalid receiver type MyInt (MyInt is a pointer type)
   return fmt.Sprintf("%d", *(*int)(r))
}

// receiver 参数的基类型为接口类型
type MyReader io.Reader

func (r MyReader) Read(p []byte) (int, error) { // r的基类型为MyReader，编译器报错：invalid receiver type MyReader (MyReader is an interface type)
   return r.Read(p)
}
```

#### Go 方法声明的位置

最后，Go 对方法声明的位置也是有约束的，Go 要求，**方法声明要与 receiver 参数的基类型声明放在同一个包内**。

基于这个约束，还可以得到两个推论。

第一个推论：**不能为原生类型（诸如 int、float64、map 等）添加方法**。

比如，下面的代码试图为 Go 原生类型 int 增加新方法 Foo，这样做，Go 编译器会报 错：

```go
// 不能为原生类型添加方法
func (i int) Foo() string { // 编译器报错：cannot define new methods on non-local type int
   return fmt.Sprintf("%d", i)
}
```

第二个推论：**不能跨越 Go 包为其他包的类型声明新方法**。

比如，下面的代码试图跨越包边界，为 Go 标准库中的 http.Server 类型添加新方法 Foo，这样做，Go 编译器同样会报错：

```go
// 不能跨越Go包为其他包的类型声明新方法
func (s http.Server) Foo() { // 编译器报错：cannot define new methods on non-local type http.Server
   
}
```

#### 使用 Go 方法

到这里，已经基本了解了 Go 方法的声明形式以及对 receiver 参数的相关约束。有了 这些基础后，就可以看一下如何使用这些方法（method）。 

直接还是通过一个例子理解一下。如果 receiver 参数的基类型为 T，那么说 **receiver 参数绑定在 T** 上，可以通过 *T 或 T 的变量实例调用该方法：

```go
// 使用 Go 方法
type T struct{}

func (t T) M(n int) {
   fmt.Println(n)
}

func main() {
   var t T
   t.M(1) // 通过类型T的变量实例调用方法M
   
   p := &T{}
   p.M(2) // 通过类型*T的变量实例调用方法M
}
```

这段代码中，方法 M 是类型 T 的方法，那为什么通过 *T 类型变量也可以调用 M 方法呢？

从上面这些分析中，也可以看到，和其他主流编程语言相比，Go 语言的方法，只比函数多出了一个 receiver 参数，这就大大降低了 Gopher 们学习方法这一语法元素的门槛。 

但即便如此，在使用方法时可能仍然会有一些疑惑，比如，方法的类型是什么？是否可以将方法赋值给函数类型的变量？调用方法时方法对 receiver 参数的修改是不是外部可见的？

### 方法的本质是什么？

#### 自定义方法

Go 的方法与 Go 中的类型是通过 receiver 联系在一起，可以**为任何非内置原生类型定义方法**，比如下面的类型 T：

```go
// 自定义方法
type T struct {
   a int
}

func (t T) Get() int {
   return t.a
}

func (t *T) Set(a int) int {
   t.a = a
   return t.a
}
```

可以和典型的面向对象语言 C++ 做下对比。

如果了解 C++ 语言，尤其是看过 C++ 大牛、《C++ Primer》作者 Stanley B·Lippman 的大作《深入探索 C++ 对象模型》，大约会知道，C++ 中的对象在调用方法时，编译器会自动传入指向对象自身的 **this 指针**作为方法的第一个参数。 

而 Go 方法中的原理也是相似的，只不过是**将 receiver 参数以第一个参数的身份并入 到方法的参数列表中**。按照这个原理，示例中的类型 T 和 *T 的方法，就可以分别等价转换为下面的普通函数：

```go
// Get 类型T的方法Get的等价函数
func Get(t T) int {
	return t.a
}

// Set 类型*T的方法Set的等价函数
func Set(t *T, a int) int {
	t.a = a
	return t.a
}
```

这种等价转换后的函数的类型就是方法的类型。只不过在 Go 语言中，这种等价转换是由 Go 编译器在编译和生成代码时**自动完成**的。

#### 方法表达式 （Method Expression）

Go 语言规范中还提供了**方法表达式 （Method Expression）**的概念，可以更充分地理解上面的等价转换。 

还以上面类型 T 以及它的方法为例，结合前面说过的 Go 方法的调用方式，可以得到下面代码：

```go
var t T
t.Get()
t.Set(1)
```

可以用另一种方式，把上面的方法调用做一个等价替换：

```go
var t T
T.Get(t)
(*T).Set(&t, 1)
```

这种直接以类型名 T 调用方法的表达方式，被称为 Method Expression。通过 Method Expression 这种形式，类型 T 只能调用 T 的**方法集合（Method Set）**中的方法，同理类型 *T 也只能调用 *T 的方法集合中的方法。

Method Expression 有些类似于 **C++ 中的静态方法（Static Method）**， 

- C++ 中的静态方法在使用时，以该 C++ 类的某个对象实例作为第一个参数，
- 而 Go 语言 的 Method Expression 在使用时，同样以 receiver 参数所代表的类型实例作为第一个参数。 

这种通过 Method Expression 对方法进行调用的方式，与之前所做的方法到函数的等价转换是如出一辙的。

#### 方法的本质是函数

所以，Go 语言中的**方法的本质就是，一个以方法的 receiver 参数作为第一个参数的普通函数。** 

而且，Method Expression 就是 Go 方法本质的最好体现，因为方法自身的类型就是一个普通函数的类型，甚至可以将它作为右值，**赋值给一个函数类型的变量**，比如下面示 例：

```go
func main() {
   var t T
   f1 := (*T).Set                           // f1的类型，也是T类型Set方法的类型：func (t *T, int)int
   f2 := T.Get                              // f2的类型，也是T类型Get方法的类型：func(t T)int
   fmt.Printf("the type of f1 is %T\n", f1) // the type of f1 is func(*main.T, int) int
   fmt.Printf("the type of f2 is %T\n", f2) // the type of f2 is func(main.T) int
   f1(&t, 3)
   fmt.Println(f2(t)) // 3
}
```

既然**方法本质上也是函数**，可能会问：知道方法的本质是函数又怎么样呢？它对在实际编码工作有什么帮助吗？ 

下面就以一个实际例子来看看，如何基于对方法本质的深入理解，来分析解决实际编码工作中遇到的真实问题。

### 巧解难题 

#### 问题

这个例子是来一次真实的读者咨询，问题代码是这样的：

```go
package main

import (
   "fmt"
   "time"
)

type field struct {
   name string
}

func (p *field) print() {
   fmt.Println(p.name)
}

func main() {
   data1 := []*field{{"one"}, {"two"}, {"three"}}
   for _, v := range data1 {
      go v.print()
   }
  
   data2 := []field{{"four"}, {"five"}, {"six"}}
   for _, v := range data2 {
      go v.print()
   }
  
   time.Sleep(3 * time.Second)
}
```

这段代码在多核 macOS 上的运行结果是这样（由于 **Goroutine 调度顺序不同**，运行结果中的行序可能与下面的有差异）：

```sh
one
two
three
six
six
six
```

这位读者的问题显然是：为什么对 data2 迭代输出的结果是三个“six”，而不是 four、 five、six？ 那就来分析一下。 

#### 问题的本质

首先，根据 **Go 方法的本质**，也就是一个以方法的 receiver 参数作为第一个参数的普通函数，**对这个程序做个等价变换**。这里利用 Method Expression 方式，等价变换后的源码如下：

```go
package main

import (
   "fmt"
   "time"
)

type field struct {
   name string
}

func (p *field) print() {
   fmt.Println(p.name)
}

func main() {
   data1 := []*field{{"one"}, {"two"}, {"three"}}
   for _, v := range data1 {
      go (*field).print(v)
   }

   data2 := []field{{"four"}, {"five"}, {"six"}}
   for _, v := range data2 {
      go (*field).print(&v)
   }

   time.Sleep(3 * time.Second)
}
```

这段代码中，把对 field 的方法 print 的调用，替换为 Method Expression 形式，替换前后的程序输出结果是一致的。

但变换后，问题是不是豁然开朗了！可以很清楚地看到**使用 go 关键字启动一个新 Goroutine 时**，method expression 形式的 print 函数是如何绑定参数的：

- 迭代 data1 时，由于 data1 中的元素类型是 field 指针 (\*field)，因此赋值后 v 就是元素地址，与 print 的 receiver 参数类型相同，每次调用 (*field).print 函数时直接传入 v 即可，实际上传入的也是各个 field 元素的地址；
- 迭代 data2 时，由于 data2 中的元素类型是 field（非指针），与 print 的 receiver 参数类型不同，因此需要将其取地址后再传入 (*field).print 函数。这样每次传入的 &v 实际上是变量 v 的地址，而不是切片 data2 中各元素的地址。

在学习 for range 使用时应注意的几个问题，其中**循环变量复用**是关键的一个。这里的 v 在整个 for range 过程中只有一个，因此 data2 迭代完成之后，v 是元素“six”的拷贝。 

这样，一旦启动的各个子 goroutine 在 main goroutine 执行到 Sleep 时才被调度执行， 

- 那么最后的三个 goroutine 在打印 &v 时，实际打印的也就是在 v 中存放的值“six”。
- 而前三个子 goroutine 各自传入的是元素“one”、“two”和“three”的地址，所以打印 的就是“one”、“two”和“three”了。 

#### 解决问题

那么**原程序要如何修改**，才能让它按期望，输 出“one”、“two”、“three”、“four”、 “five”、“six”呢？ 

其实，只需要将 field 类型 print 方法的 receiver 类型由 *field 改为 field 就可以了。 直接来看一下**修改后的代码**：

```go
package main

import (
   "fmt"
   "time"
)

type field struct {
   name string
}

func (p field) print() {
   fmt.Println(p.name)
}

func main() {
   data1 := []*field{{"one"}, {"two"}, {"three"}}
   for _, v := range data1 {
      go v.print()
   }

   data2 := []field{{"four"}, {"five"}, {"six"}}
   for _, v := range data2 {
      go v.print()
   }
   
   time.Sleep(3 * time.Second)
}
```

修改后的程序的输出结果是这样的（因 Goroutine 调度顺序不同，在结果输出顺序可能会有不同）：

```sh
one
six
three
two
five
four
```

为什么这回就可以输出预期的值了呢？



### 小结 

Go 语言中除函数之外的、另一种可承载代码执行逻辑的语法元素：方法（method）。 

Go 提供方法这种语法，并非出自对经典面向对象编程范式支持的考虑，而是出自 Go 的组合设计哲学下类型系统实现层面上的需要。 

Go 方法在声明形式上相较于 Go 函数多了一个 receiver 组成部分，这个部分是方法与类型之间联系的纽带。可以在 receiver 部分声明 receiver 参数。但 Go 对 receiver 参数有诸多限制，比如只能有一个、参数名唯一、不能是变长参数等等。 

除此之外，Go 对 receiver 参数的基类型也是有约束的，即基类型本身不能是指针类型或接口类型。Go 方法声明的位置也受到了 Go 规范的约束，方法声明必须与 receiver 参数的基类型在同一个包中。 

Go 方法本质上其实是一个函数，这个函数以方法的 receiver 参数作为第一个参数，Go 编译器会在进行方法调用时协助进行这样的转换。牢记并理解方法的这个本质可以在实际编码中解决一些奇怪的问题。



### 思考题 

在“巧解难题”部分，为啥只需要将 field 类型 print 方法的 receiver 类型，由 *field 改为 field 就可以输出预期的结果了呢？

- 由 \*field 改为 field结果正确的原因是， \*field的方法的第一个参数是*field， 这个对于[]\*field数组直接传入成员就可以了， 而对于[]field数组， 则是要取地址，也就是指针。 但是这个指针指的是for range 循环的局部变量的地址， 这个地址在for 循环中是不变的， 在for循环结束后这个地址就指向了最后一个元素， goroutine 真正实行打印的引用的地址是局部变量的地址， 自然只会打印最后一个元素了

- 使用 field 的方法， 不涉及引用， 传参都是拷贝复制

- 基于方法的本质，进行解决问题后的源码转换，得：

- ```go
  package main
  
  import (
     "fmt"
     "time"
  )
  
  type field struct {
     name string
  }
  
  func (p field) print() {
     fmt.Println(p.name)
  }
  
  func main() {
     data1 := []*field{{"one"}, {"two"}, {"three"}}
     for _, v := range data1 {
        go field.print(*v)
     }
  
     data2 := []field{{"four"}, {"five"}, {"six"}}
     for _, v := range data2 {
        go field.print(v)
     }
  
     time.Sleep(3 * time.Second)
  }
  ```

  

### Go 方法设计

在 Go 语言中，方法本质上就是函数，所以之前讲解的、关于函数设计的内容对方法也同样适用，比如错误处理设计、针对异常的处理策略、使用 defer 提升简洁性，等等。 

但关于 Go 方法中独有的 receiver 组成部分，却没有现成的、可供我参考的内容。初学者在学习 Go 方法时，最头疼的一个问题恰恰就是如何选择 receiver 参数的 类型。

### receiver 参数类型

#### receiver 参数类型对 Go 方法的影响 

要想为 receiver 参数选出合理的类型，先要了解不同的 receiver 参数类型会对 Go 方法产生怎样的影响。

**Go 方法的本质，是以方法的 receiver 参数作为第一个参数的普通函数**。 

对于函数参数类型对函数的影响，是很熟悉的。那么能不能将方法等价转换为对应的函数，再通过分析 receiver 参数类型对函数的影响，从而**间接**得出它对 Go 方法的影响呢？ 

可以基于这个思路试试看。直接来看下面例子中的两个 Go 方法，以及它们等价转换后的函数：

```go
func (t T) M1() <=> F1(t T)
func (t *T) M2() <=> F2(t *T)
```

这个例子中有方法 M1 和 M2。M1 方法是 receiver 参数类型为 T 的一类方法的代表，而 M2 方法则代表了 receiver 参数类型为 *T 的另一类。下面分别来看看不同的 receiver 参数类型对 M1 和 M2 的影响。

- 首先，当 receiver 参数的类型为 T 时：
  - 当选择以 T 作为 receiver 参数类型时，M1 方法等价转换为F1(t T)。
  - Go 函数的参数采用的是**值拷贝传递**，也就是说，F1 函数体中的 t 是 T 类型实例的一个**副本**。这样，在 F1 函数的实现中对参数 t 做任何修改，都只会影响副本，而不会影响到原 T 类型实例。

据此可以得出结论：当方法 M1 采用类型为 T 的 receiver 参数时，代表 T 类型实例的 receiver 参数以值传递方式传递到 M1 方法体中的，实际上是 T 类型实例的副本，M1 方法体中对副本的任何修改操作，都不会影响到原 T 类型实例。

- 第二，当 receiver 参数的类型为 *T 时：
  - 当选择以 *T 作为 receiver 参数类型时，M2 方法等价转换为F2(t *T)。
  - 同上面分析，传递给 F2 函数的 t 是 T 类型实例的**地址**，这样 F2 函数体中对参数 t 做的任何修改，都会反映到原 T 类型实例上。

据此也可以得出结论：当方法 M2 采用类型为 *T 的 receiver 参数时，代表 *T 类型实例的 receiver 参数以值传递方式传递到 M2 方法体中的，实际上是 T 类型实例的地址，M2 方法体通过该地址可以对原 T 类型实例进行任何修改操作。 

#### receiver 类型对原类型实例的影响

再通过一个更直观的例子，证明一下上面这个分析结果，看一下 Go 方法选择不同的 receiver 类型**对原类型实例的影响**：

```go
package main

type T struct {
   a int
}

func (t T) M1() {
   t.a = 10
}

func (t *T) M2() {
   t.a = 11
}

func main() {
   var t T
   println(t.a) // 0

   t.M1()
   println(t.a) // 0
   
   p := &t
   p.M2()
   println(t.a) // 11
}
```

在这个示例中，为基类型 T 定义了两个方法 M1 和 M2，其中 M1 的 receiver 参数类型为 T，而 M2 的 receiver 参数类型为 *T。M1 和 M2 方法体都通过 receiver 参数 t 对 t 的字段 a 进行了修改。

但运行这个示例程序后，

- 方法 M1 由于使用了 T 作为 receiver 参数类型，它在 方法体中修改的仅仅是 T 类型实例 t 的副本，原实例并没有受到影响。因此 M1 调用后， 输出 t.a 的值仍为 0。 
- 而方法 M2 呢，由于使用了 *T 作为 receiver 参数类型，它在方法体中通过 t 修改的是实例本身，因此 M2 调用后，t.a 的值变为了 11，这些输出结果与前面的分析是一致 的。 

了解了不同类型的 receiver 参数对 Go 方法的影响后，就可以总结一下，日常编码中 选择 receiver 的参数类型的时候，可以参考哪些原则。 

#### 选择 receiver 参数类型的第一个原则 

基于上面的影响分析，可以得到选择 receiver 参数类型的第一个原则：**如果 Go 方法要把对 receiver 参数代表的类型实例的修改，反映到原类型实例上，那么应该选择 *T 作为 receiver 参数的类型**。 

这个原则似乎很好掌握，不过这个时候，可能会有个疑问：如果选择了 *T 作为 Go 方法 receiver 参数的类型，那么是不是只能通过 *T 类型变量调用该方法，而不能通过 T 类型变量调用了呢？

这个问题恰恰也是遗留的一个问题。改造一下上面例子看一下：

```go
package main

type T struct {
   a int
}

func (t T) M1() {
   t.a = 10
}

func (t *T) M2() {
   t.a = 11
}

func main() {
   var t1 T
   println(t1.a) // 0
   t1.M1()
   println(t1.a) // 0
   t1.M2()
   println(t1.a) // 11

   var t2 = &T{}
   println(t2.a) // 0
   t2.M1()
   println(t2.a) // 0
   t2.M2()
   println(t2.a) // 11
}
```

先来看看类型为 T 的实例 t1。

- 看到它不仅可以调用 receiver 参数类型为 T 的方 法 M1，它还可以直接调用 receiver 参数类型为 *T 的方法 M2，并且调用完 M2 方法 后，t1.a 的值被修改为 11 了。 
- 其实，T 类型的实例 t1 之所以可以调用 receiver 参数类型为 *T 的方法 M2，都是 Go 编译器在背后自动进行转换的结果。
- 或者说，t1.M2() 这种用法是 Go 提供的“语法糖”： Go 判断 t1 的类型为 T，也就是与方法 M2 的 receiver 参数类型 *T 不一致后，会自动将 t1.M2()转换为(&t1).M2()。 

同理，类型为 *T 的实例 t2，

- 它不仅可以调用 receiver 参数类型为 *T 的方法 M2，还可以调用 receiver 参数类型为 T 的方法 M1，这同样是因为 Go 编译器在背后做了转换。
- 也就是，Go 判断 t2 的类型为 \*T，与方法 M1 的 receiver 参数类型 T 不一致，就会自动将 t2.M1()转换为(*t2).M1()。 

通过这个实例，知道了这样一个结论：无论是 T 类型实例，还是 *T 类型实例，都既可以调用 receiver 为 T 类型的方法，也可以调用 receiver 为 *T 类型的方法。

这样，在为方法选择 receiver 参数的类型的时候，就不需要担心这个方法不能被与 receiver 参数类型不一致的类型实例调用了。

#### 选择 receiver 参数类型的第二个原则 

前面第一个原则说的是，在方法中对 receiver 参数代表的类型实例进行修改，那就要为 receiver 参数选择 *T 类型，但是如果不需要在方法中对类型实例进行修改呢？

这个时候是为 receiver 参数选择 T 类型还是 *T 类型呢？ 这也得分情况。

一般情况下，通常会**为 receiver 参数选择 T 类型**，因为这样可以缩窄外部修改类型实例内部状态的“接触面”，也就是**尽量少暴露可以修改类型内部状态的方法**。 

不过也有一个例外需要特别注意。

- 考虑到 Go 方法调用时，receiver 参数是以值拷贝的形式传入方法中的。
- 那么，如果 receiver 参数类型的 size 较大，以值拷贝形式传入就会导致**较大的性能开销**，这时选择 *T 作为 receiver 类型可能更好些。 

以上这些可以作为选择 receiver 参数类型的第二个原则。 

#### 方法集合（Method Set）

先了解一个基本概念：方法集合（Method Set）， 它是理解第三条原则的前提。 

##### 方法集合解决的问题

在了解方法集合是什么之前，先通过一个示例，直观了解一下为什么要有方法集合， 它主要用来解决什么问题：

```go
package main

type Interface interface {
   M1()
   M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
   var t T
   var pt *T
   var i Interface
   i = pt
   i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
}
```

在这个例子中，

- 定义了一个接口类型 Interface 以及一个自定义类型 T。
- Interface 接口类型包含了两个方法 M1 和 M2，它们的基类型都是 T，但它们的 receiver 参数类型不同，一个为 T，另一个为 *T。
- 在 main 函数中，分别将 T 类型实例 t 和 *T 类型实例 pt 赋值给 Interface 类型变量 i。 

运行一下这个示例程序，在i = t这一行会得到 Go 编译器的错误提示，**Go 编译器提示：T 没有实现 Interface 类型方法列表中的 M2，因此类型 T 的实例 t 不能赋值给 Interface 变量**。 

为什么 *T 类型的 pt 可以被正常赋值给 Interface 类型变量 i，而 T 类型 的 t 就不行呢？如果说 T 类型是因为只实现了 M1 方法，未实现 M2 方法而不满足 Interface 类型的要求，那么 *T 类型也只是实现了 M2 方法，并没有实现 M1 方法啊？ 

有些事情并不是表面看起来这个样子的。了解方法集合后，这个问题就迎刃而解了。同时，**方法集合也是用来判断一个类型是否实现了某接口类型的唯一手段**，可以说，“方法集合决定了接口实现”。

那么，什么是类型的方法集合呢？ 

- Go 中任何一个类型都有属于自己的方法集合，或者说方法集合是 Go 类型的一个“属性”。但不是所有类型都有自己的方法呀，比如 int 类型就没有。
- 所以，对于没有定义方法的 Go 类型，称其拥有**空方法集合**。 
- 接口类型相对特殊，它只会列出代表接口的方法列表，不会具体定义某个方法，它的方法集合就是它的方法列表中的所有方法，可以一目了然地看到。
- 因此，下面重点讲解的是非接口类型的方法集合。 

##### dumpMethodSet

为了方便查看一个**非接口类型的方法集合**，这里提供了一个函数 dumpMethodSet，用于输出一个非接口类型的方法集合：

```go
func dumpMethodSet(i interface{}) {
   dynTyp := reflect.TypeOf(i)
   if dynTyp == nil {
      fmt.Printf("there is no dynamic type\n")
      return
   }
   n := dynTyp.NumMethod()
   if n == 0 {
      fmt.Printf("%s's method set is empty!\n", dynTyp)
      return
   }
   fmt.Printf("%s's method set:\n", dynTyp)
   for j := 0; j < n; j++ {
      fmt.Println("-", dynTyp.Method(j).Name)
   }
   fmt.Printf("\n")
}
```

下面利用这个函数，试着输出一下 **Go 原生类型以及自定义类型的方法集合**，看下面 代码：

```go
type T struct{}

func (T) M1()  {}
func (T) M2()  {}

func (*T) M3() {}
func (*T) M4() {}

func main() {
   var n int
   dumpMethodSet(n)
   dumpMethodSet(&n)
   
   var t T
   dumpMethodSet(t)
   dumpMethodSet(&t)
}
```

运行这段代码，得到如下结果：

```sh
int's method set is empty!
*int's method set is empty!
main.T's method set:
- M1
- M2

*main.T's method set:
- M1
- M2
- M3
- M4
```

可以看到：

- 以 int、*int 为代表的 Go 原生类型由于没有定义方法，所以它们的方法集合都是空的。
- 自定义类型 T 定义了方法 M1 和 M2，因此它的方法集合包含了 M1 和 M2，也符合预期。
- 但 \*T 的方法集合中除了预期的 M3 和 M4 之外，居然还包含了类型 T 的方 法 M1 和 M2！ 不过，这里程序的输出并没有错误。 

这是因为，Go 语言规定，***T 类型的方法集合包含所有以 *T 为 receiver 参数类型的方法，以及所有以 T 为 receiver 参数类型的方法**。这就是这个示例中为何 *T 类型的方法集合包含四个方法的原因。 

##### 验证方法集合

这个时候，是不是也找到了前面那个示例中为何i = pt没有报编译错误的原因了呢？

同样可以使用 dumpMethodSet 工具函数，输出一下那个例子中 pt 与 t 各自所属类型 的方法集合：

```go
package main

import (
   "fmt"
   "reflect"
)

// 输出方法集合，确定问题

type Interface interface {
   M1()
   M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func dumpMethodSet(i interface{}) {
   dynTyp := reflect.TypeOf(i)
   if dynTyp == nil {
      fmt.Printf("there is no dynamic type\n")
      return
   }
   n := dynTyp.NumMethod()
   if n == 0 {
      fmt.Printf("%s's method set is empty!\n", dynTyp)
      return
   }
   fmt.Printf("%s's method set:\n", dynTyp)
   for j := 0; j < n; j++ {
      fmt.Println("-", dynTyp.Method(j).Name)
   }
   fmt.Printf("\n")
}

func main() {
   var t T
   var pt *T
   dumpMethodSet(t)
   dumpMethodSet(pt)
}
```

运行上述代码，得到如下结果：

```sh
main.T's method set:
- M1

*main.T's method set:
- M1
- M2
```

通过这个输出结果，可以一目了然地看到 T、*T 各自的方法集合。 

- T 类型的方法集合中只包含 M1，没有 Interface 类型方法集合中的 M2 方法， 这就是 Go 编译器认为变量 t 不能赋值给 Interface 类型变量的原因。 
- 在输出的结果中，还看到 \*T 类型的方法集合除了包含它自身定义的 M2 方法外，还包含了 T 类型定义的 M1 方法，*T 的方法集合与 Interface 接口类型的方法集合是一样的， 因此 pt 可以被赋值给 Interface 接口类型的变量 i。 

到这里，已经知道了所谓的**方法集合决定接口实现**的含义就是：

- 如果某类型 T 的方法集合与某接口类型的方法集合相同，或者类型 T 的方法集合是接口类型 I 方法集合的超集，那么就说这个类型 T 实现了接口 I。
- 或者说，方法集合这个概念在 Go 语言中的主要用途，就是用来判断某个类型是否实现了某个接口。 

有了方法集合的概念做铺垫，选择 receiver 参数类型的第三个原则已经呼之欲出了。

#### 选择 receiver 参数类型的第三个原则 

这个原则的选择依据就是 **T 类型是否需要实现某个接口**。 

- 如果 T 类型需要实现某个接口，那就要使用 T 作为 receiver 参数的类型，来满足接口类型方法集合中的所有方法。
- 如果 T 不需要实现某一接口，但 \*T 需要实现该接口，那么根据方法集合概念，*T 的方法集合是包含 T 的方法集合的，这样在确定 Go 方法的 receiver 的类型时，参考原则一 和原则二就可以了。 

如果说前面的两个原则更多聚焦于类型内部，从单个方法的实现层面考虑，那么这第三个原则是更多从全局的设计层面考虑，聚焦于这个类型与接口类型间的耦合关系。



### 小结 

Go 方法本质上也是函数。所以 Go 方法设计的多数地方，都可以借鉴函数设计的相关内容。唯独 Go 方法的 receiver 部分，是没有现成经验可循的。

如何为 Go 方法的 receiver 参数选择类型。 

- 先了解了不同类型的 receiver 参数对 Go 方法行为的影响，这是进行 receiver 参 数选型的前提。 
- 在这个前提下，提出了 receiver 参数选型的三个经验原则，实际进行 Go 方法设计时，首先应该考虑的是原则三，即 T 类型是否要实现某一接口。 
  - 如果 T 类型需要实现某一接口的全部方法，那么就需要使用 T 作为 receiver 参数的类型来满足接口类型方法集合中的所有方法。 
  - 如果 T 类型不需要实现某一接口，那么就可以参考原则一和原则二来为 receiver 参数选择类型了。
  - 也就是，如果 Go 方法要把对 receiver 参数所代表的类型实例的修改反映到 原类型实例上，那么应该选择 *T 作为 receiver 参数的类型。
  - 否则通常会为 receiver 参数选择 T 类型，这样可以减少外部修改类型实例内部状态的“渠道”。
  - 除非 receiver 参数类型的 size 较大，考虑到传值的较大性能开销，选择 *T 作为 receiver 类型可能更适合。 
- Go 语言中的一个重要概念：方法集合。
  - 它在 Go 语言中的主要用途就是判断某个类型是否实现了某个接口。
  - 方法集合像“胶水”一样，将自定义类型与接口隐式地“粘结”在一起，后面理解带有类型嵌入的类型时还会借助这个概 念。 

### 思考题 

方法集合是一个很重要也很实用的概念， 如果一个类型 T 包含两个方法 M1 和 M2：

```go
type T struct{}

func (T) M1() {}
func (T) M2() {}
```

然后，再使用类型定义语法，又基于类型 T 定义了一个新类型 S：

```go
type S T
```

那么，这个 S 类型包含哪些方法呢？*S 类型又包含哪些方法呢？

- 验证：

  - ```go
    type T struct{}
    
    func (T) M1() {}
    func (T) M2() {}
    
    type S T
    // type S =  T
    
    func main() {
       var t T
       dumpMethodSet(t)
       dumpMethodSet(&t)
       
       var s S
       dumpMethodSet(s)
       dumpMethodSet(&s)
    }
    ```

- 输出：

  - ```sh
    main.T's method set:
    - M1
    - M2
    
    *main.T's method set:
    - M1
    - M2
    
    main.S's method set is empty!
    *main.S's method set is empty!
    ```

- S 类型 和 *S 类型都没有包含方法

- 因为S是新的类型，它不能调用T的方法，必须显示转换之后 才可以调用，所以本身的S或*S类型都不具有任何的方法

- 但是如果用 type S = T 则S和*S类型都包含两个方法。





### Go 类型嵌入







































