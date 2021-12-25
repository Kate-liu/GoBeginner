# Go Language Base Data Type

>Go 语言数据类型 之 复合数据类型

## 复合数据类型

Go 基本数据类型，主要包括数值类型与字符串类型。 但是，这些基本数据类型建立的抽象概念，还远不足以应对真实世界的各种问题。 

- 比如，要表示一组数量固定且连续的整型数，建立一个能表示书籍的抽象数据类型， 这个类型中包含书名、页数、出版信息等；
- 又或者，要建立一个学号与学生姓名的映射表等。

这些问题基本数据类型都无法解决，所以需要一类新类型来建立这些抽象， 丰富 Go 语言的表现力。

这种新类型是怎么样的呢？

- 可以通过这些例子总结出新类型的一个特点，那就是它们都是由多个同构类型（相同类型）或异构类型（不同类型）的元素的值组合而成的。
- 这类数据类型在 Go 语言中被称为复合类型，即 Go 语言的复合类型。

Go 语言原生内置了多种复合数据类型，包括**数组、切片（slice）、map、结构体**，以及像 **channel** 这类用于并发程序设计的高级复合数据类型。

## 数组

### 数组类型的逻辑定义

先来看数组类型的概念。Go 语言的数组是**一个长度固定的、由同构类型元素组成的连 续序列**。

通过这个定义，可以识别出 Go 的数组类型包含两个重要属性：**元素的类型和数组长度（元素的个数）**。这两个属性也直接构成了 Go 语言中数组类型变量的声明：

```go
// 数组声明
var arr [N]T
```

这里声明了一个数组变量 arr，它的类型为 [N]T，其中元素的类型为 T，数组的长度为 N。

这里，要注意，数组元素的类型可以为任意的 Go 原生类型或自定义类型，而且数组的长度必须在声明数组变量时提供，Go 编译器需要在编译阶段就知道数组类型的长度。所以，只能用整型数字面值或常量表达式作为 N 值。 

通过这句代码也可以看到，**如果两个数组类型的元素类型 T 与数组长度 N 都是一样的，那么这两个数组类型是等价的，如果有一个属性不同，它们就是两个不同的数组类型**。下面这个示例很好地诠释了这一点：

```go
func foo(arr [5]int) {

}

func main() {
   var arr1 [5]int
   var arr2 [6]int
   var arr3 [5]string

   foo(arr1) // ok
   foo(arr2) // 错误：[6]int与函数foo参数的类型[5]int不是同一数组类型 // cannot use arr2 (type [6]int) as type [5]int in argument to foo
   foo(arr3) // 错误：[5]string与函数foo参数的类型[5]int不是同一数组类型 // cannot use arr3 (type [5]string) as type [5]int in argument to foo
}
```

在这段代码里，arr2 与 arr3 两个变量的类型分别为 [6]int 和 [5]string，前者的长度属性与 [5]int 不一致，后者的元素类型属性与[5]int 不一致，因此这两个变量都不能作为调用函数 foo 时的实际参数。 

### 数组类型在内存中的实际表示

了解了数组类型的逻辑定义后，再来看看数组类型在内存中的实际表示是怎样的，这是数组区别于其他类型，也是区分不同数组类型的根本依据。 

数组类型不仅是逻辑上的连续序列，而且在实际内存分配时也占据着一整块内存。

Go 编译器在为数组类型的变量实际分配内存时，会为 Go 数组**分配一整块、可以容纳它所有元素的连续内存**，如下图所示：

![image-20211225155620964](go_language_compound_data_type.assets/image-20211225155620964.png)

从这个数组类型的内存表示中可以看出来，这块内存全部空间都被用来表示数组元素，所以说这块内存的大小，就等于各个数组元素的大小之和。

如果两个数组所分配的内存大小不同，那么它们肯定是不同的数组类型。Go 提供了**预定义函数 len **可以用于获取一个数组类型变量的长度，通过 **unsafe 包提供的 Sizeof 函数**，可以获得一个数组变量的总大小，如下面代码：

```go
var arr = [6]int{1, 2, 3, 4, 5, 6}
fmt.Println("数组长度: ", len(arr))           // 6
fmt.Println("数组大小: ", unsafe.Sizeof(arr)) // 48
```

数组大小就是所有元素的大小之和，这里数组元素的类型为 int。在 64 位平台上，int 类型的大小为 8，数组 arr 一共有 6 个元素，因此它的总大小为 6x8=48 个字节。

### 数组初始化

和基本数据类型一样，声明一个数组类型变量的同时，也可以显式地对它进行初始化。

如果不进行显式初始化，那么数组中的元素值就是它类型的**零值**。比如下面的数组类型变量 arr5 的各个元素值都为 0：

```go
// 数组默认初始化为零值
var arr5 [6]int
fmt.Println(arr5)  // [0 0 0 0 0 0]
```

如果要显式地对数组初始化，需要在右值中显式放置数组类型，并通过**大括号**的方式给各个元素赋值（如下面代码中的 arr6）。

当然，也可以忽略掉右值初始化表达式中数组类型的长度，用**“…”**替代，Go 编译器会根据数组元素的个数，自动计算出数组长度 （如下面代码中的 arr7）：

```go
// 大括号显示数组赋值
var arr6 = [6]int{
  11, 12, 13, 14, 15, 16,
}
fmt.Println(arr6) // [11 12 13 14 15 16]

// 自动计算数组长度
var arr7 = [...]int{
  21, 22, 23,
}
fmt.Println(arr7)      // [21 22 23]
fmt.Println(len(arr7)) // 3
```

如果要对一个长度较大的稀疏数组进行显式初始化，这样逐一赋值就太麻烦了，还有什么更好的方法吗？

可以通过使用下标赋值的方式对它进行初始化，比如下面代码中的 arr8：

```go
// 下标赋值的方式初始化
var arr8 = [...]int{
   99: 39, // 将第100个元素(下标值为99)的值赋值为39，其余元素值均为0
}
fmt.Println(arr8) // [0 0 ... 99]
```

### 数组的访问

通过数组类型变量以及下标值，可以很容易地访问到数组中的元素值，并且这种访问是十分高效的，不存在 Go 运行时带来的额外开销。

但要记住，数组的下标值是从 0 开 始的。如果下标值超出数组长度范畴，或者是负数，那么 Go 编译器会给出错误提示，防止访问溢出：

```go
// 数组的访问
var arr9 = [5]int{11, 12, 13, 14, 15}
fmt.Println(arr9[0], arr9[4])  // 11 15
fmt.Println(arr9[-1]) // invalid array index -1 (index must be non-negative) 错误：下标值不能为负数
fmt.Println(arr9[99]) // invalid array index 99 (out of bounds for 5-element array) 错误：下标值超出了arr的长度范围
```



## 多维数组

上面这些元素类型的数组都是简单的一维数组，但 Go 语言中，其实还有更复杂的数组类型，多维数组。

也就是说，数组类型自身也可以作为数组元素的类型，这样就会产生多维数组，比如下面的变量 mArr 的类型就是一个多维数组`[2][3][4]int`：

```go
var mArr [2][3][4]int

fmt.Println(mArr)
// [
//  [
//   [0 0 0 0]
//   [0 0 0 0]
//   [0 0 0 0]
//  ]
//  [
//   [0 0 0 0]
//   [0 0 0 0]
//   [0 0 0 0]
//   ]
// ]
```

多维数组也不难理解，以上面示例中的多维数组类型为例，从左向右逐维地去 看，这样就可以将一个多维数组分层拆解成这样：

![image-20211225163029792](go_language_compound_data_type.assets/image-20211225163029792.png)

- 从上向下看，首先将 mArr 这个数组看成是一个拥有两个元素，且元素类型都为 [3] [4]int 的数组，就像图中最上层画的那样。这样，mArr 的两个元素分别为 mArr[0]和 mArr [1]，它们的类型均为[3] [4]int，也就是说它们都是二维数组。 
- 而以 mArr[0]为例，可以将其看成一个拥有 3 个元素且元素类型为[4]int 的数组，也就是图中中间层画的那样。这样 mArr[0]的三个元素分别为 mArr[0] [0]、mArr[0] [1]以及 mArr[0] [2]，它们的类型均为[4]int，也就是说它们都是一维数组。 
- 图中的最后一层就是 mArr[0]的三个元素，以及 mArr[1]的三个元素的各自展开形式。

以此类推，会发现，无论多维数组究竟有多少维，都可以将它从左到右逐一展开，最终化为熟悉的一维数组。 

不过，虽然数组类型是 Go 语言中最基础的复合数据类型，但是在使用中它也会有一些**问题**。

数组类型变量是一个整体，这就意味着一个数组变量表示的是整个数组。这点与 C 语言完全不同，在 C 语言中，数组变量可视为指向数组第一个元素的指针。这样一来，无论是参与迭代，还是作为实际参数传给一个函数 / 方法，Go 传递数组的方式都是纯粹的**值拷贝**，这会带来较大的内存拷贝开销。 

这时，可能会想到可以使用指针的方式，来向函数传递数组。没错，这样做的确可以避免性能损耗，但这更像是 C 语言的惯用法。

其实，Go 语言为我们提供了一种更为灵活、更为地道的方式 ，**切片**，来解决这个问题。它的优秀特性让它成为了 Go 语言中最常用的同构复合类型。



## 切片

前面提到过，数组作为最基本同构类型在 Go 语言中被保留了下来，但数组在使用上确有两点不足：固定的元素个数，以及传值机制下导致的开销较大。

于是 Go 设计者们又引入了另外一种同构复合类型：**切片（slice）**，来弥补数组的这两处不足。 

### 切片变量的声明

切片和数组就像两个一母同胞的亲兄弟，长得像，但又各有各的行为特点。可以先声明并初始化一个切片变量看看：

```go
// 切片变量的声明
var nums = []int{1, 2, 3, 4, 5, 6}
fmt.Println(nums) // [1 2 3 4 5 6]
```

与数组声明相比，切片声明仅仅是少了一个“长度”属性。去掉“长度”这一束缚后，切片展现出更为灵活的特性。 

虽然不需要像数组那样在声明时指定长度，但切片也有自己的长度，只不过这个长度不是固定的，而是随着切片中元素个数的变化而变化的。

可以通过 **len 函数**获得切片类型变量的长度，比如上面那个切片变量的长度就是 6:

```go
fmt.Println(len(nums))  // 6
```

通过 Go 内置**函数 append**，可以动态地向切片中添加元素。当然，添加后切 片的长度也就随之发生了变化，如下面代码所示：

```go
// 添加切片元素
nums = append(nums, 7)

fmt.Println(nums)      // [1 2 3 4 5 6 7]
fmt.Println(len(nums)) // 7
```



### Go 是如何实现切片类型的？ 

Go 切片在运行时其实是一个三元组结构，它在 Go 运行时的表示如下：

```go
// runtime/slice.go
type slice struct {
   array unsafe.Pointer
   len   int
   cap   int
}
```

可以看到，每个切片包含三个字段：

- array: 是指向底层数组的指针； 
- len: 是切片的长度，即切片中当前元素的个数； 
- cap: 是底层数组的长度，也是切片的最大容量，cap 值永远大于等于 len 值。

如果用这个三元组结构表示切片类型变量 nums，会是这样：

![image-20211225182720278](go_language_compound_data_type.assets/image-20211225182720278.png)

可以看到，Go 编译器会自动为每个新创建的切片，建立一个底层数组，默认底层数组的长度与切片初始元素个数相同。

> 图中的底层数组长度是12与切片长度不同的原因是，append 了一个新的元素，此时就会默认执行切片的扩充，变为原来的二倍，即6的二倍为12。

### 创建切片

还可以用以下几种方法创建切片，并指定它底层数组的长度。 

#### make 函数创建切片

方法一：**通过 make 函数来创建切片，并指定底层数组的长度**。

直接看下面这行代码：

```go
//  make 函数创建切片
sl1 := make([]byte, 6, 19) // 其中19为cap值，即底层数组长度，6为切片的初始长度
fmt.Println(sl1)           // [0 0 0 0 0 0]
fmt.Println(len(sl1))      // 6
fmt.Println(cap(sl1))      // 19
```

如果没有在 make 中指定 cap 参数，那么底层数组长度 cap 就等于 len，比如：

```go
//  make 函数创建切片，默认 cap = len = 6
sl2 := make([]byte, 6) // 其中默认6为cap值，即底层数组长度，6为切片的初始长度 // cap = len = 6
fmt.Println(sl2)       // [0 0 0 0 0 0]
fmt.Println(len(sl2))  // 6
fmt.Println(cap(sl2))  // 6
```

到这里，肯定会有一个问题，为什么上面图中 nums 切片的底层数组长度为 12，而不是初始的 len 值 6 呢？

> 图中的底层数组长度是12与切片长度不同的原因是，append 了一个新的元素，此时就会默认执行切片的扩充，变为原来的二倍，即6的二倍为12。

#### 数组的切片化

方法二：**采用 array[low : high : max]语法基于一个已存在的数组创建切片**。

这种方式被称为**数组的切片化**，比如下面代码：

```go
// 数组的切片化
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sl3 := arr[3:7:9]
fmt.Println(arr)      // [1 2 3 4 5 6 7 8 9 10]
fmt.Println(sl3)      // [4 5 6 7]
fmt.Println(len(sl3)) // 4
fmt.Println(cap(sl3)) // 6
```

基于数组 arr 创建了一个切片 sl3，这个切片 sl3 在运行时中的表示是这样：

![image-20211225184146434](go_language_compound_data_type.assets/image-20211225184146434.png)

可以看到，基于数组创建的切片，它的起始元素从 low 所标识的下标值开始，切片的长度 （len）是 high - low，它的容量是 max - low。

而且，由于切片 sl3 的底层数组就是数组 arr，**对切片 sl3 中元素的修改将直接影响数组 arr 变量**。如果将切片的第一个元素加 10，那么数组 arr 的第四个元素将变为 14：

```go
// 更改切片元素的值，会改变原数组的值
sl3[0] += 10
fmt.Println(arr)                // [1 2 3 14 5 6 7 8 9 10]
fmt.Println(sl3)                // [14 5 6 7]
fmt.Println("arr[3] =", arr[3]) // arr[3] = 14

fmt.Println(sl3[5])  // 测试：在切片在访问 大于长度 小于 cap 的元素，会报错：panic: runtime error: index out of range [5] with length 4
```

这样看来，切片好比打开了一个访问与修改数组的“窗口”，通过这个窗口，可以直接操作底层数组中的部分元素。

这有些类似于操作文件之前打开的“文件描述符”（Windows 上称为句柄），通过文件描述符可以对底层的真实文件进行相关操作。可以说，**切片之于数组就像是文件描述符之于文件**。 

在 Go 语言中，数组更多是“退居幕后”，承担的是底层存储空间的角色。切片就是数组 的“描述符”，也正是因为这一特性，切片才能在函数参数传递时避免较大性能开销。

因为传递的并不是数组本身，而是**数组的“描述符”**，而这个**描述符的大小是固定的** （见上面的三元组结构），无论底层的数组有多大，切片打开的“窗口”长度有多长，它都是不变的。

此外，在进行数组切片化的时候，**通常省略 max**，而 **max 的默认值为数组的长度**。（备注：这个默认值是数组的长度不太对吧！应该是原来数组被切片之后，从切片的起始位置到最后一个元素的个数。）



#### 数组的（多个切）片化

另外，针对一个已存在的数组，还可以**建立多个操作数组的切片**，这些切片共享同一底层数组，**切片对底层数组的操作也同样会反映到其他切片中**。

下面是为数组 arr 建立的两个切片的内存表示：

![image-20211225185001642](go_language_compound_data_type.assets/image-20211225185001642.png)

可以看到，上图中的两个切片 sl1 和 sl2 是数组 arr 的“描述符”，这样的情况下，无论通过哪个切片对数组进行的修改操作，都会反映到另一个切片中。

比如，将 sl2[2]置为 14，那么 sl1[0]也会变成 14，因为 sl2[2]直接操作的是底层数组 arr 的第四个元素 arr[3]。 

#### 切片创建切片

方法三：基于**切片创建切片**。 

不过这种切片的运行时表示原理与上面的是一样的。 

最后，回答一下前面切片变量 nums 在进行一次 append 操作后切片容量变为 12 的问题。

这里要清楚一个概念：切片与数组最大的不同，就在于其长度的不定长，这种 不定长需要 Go 运行时提供支持，这种支持就是切片的“动态扩容”。



### 切片的动态扩容 

“动态扩容”指的就是，当通过 append 操作向切片追加数据的时候，如果这时切片的 len 值和 cap 值是相等的，也就是说切片底层数组已经没有空闲空间再来存储追加的值了，Go 运行时就会对这个切片做扩容操作，来保证切片始终能存储下追加的新值。 

前面的切片变量 nums 之所以可以存储下新追加的值，就是因为 Go 对其进行了动态扩容，也就是重新分配了其底层数组，从一个长度为 6 的数组变成了一个长为 12 的数组。 

接下来，再通过一个**例子**来体会一下切片动态扩容的过程：

```go
var s []int
s = append(s, 11)
fmt.Println(len(s), cap(s)) // 1 1
s = append(s, 12)
fmt.Println(len(s), cap(s)) // 2 2
s = append(s, 13)
fmt.Println(len(s), cap(s)) // 3 4
s = append(s, 14)
fmt.Println(len(s), cap(s)) // 4 4
s = append(s, 15)
fmt.Println(len(s), cap(s)) // 5 8
```

在这个例子中，append 会根据切片对底层数组容量的需求，对底层数组进行 动态调整。

- 最开始，s 初值为零值（nil），这个时候 s 没有“绑定”底层数组。
- 先通过 append 操作向切片 s 添加一个元素 11，这个时候，append 会先分配底层数组 u1（数组长度 1），然后将 s 内部表示中的 array 指向 u1，并设置 len = 1, cap = 1; 
- 接着，通过 append 操作向切片 s 再添加第二个元素 12，这个时候 len(s) = 1， cap(s) = 1，append 判断底层数组剩余空间已经不能够满足添加新元素的要求了，于是它就创建了一个新的底层数组 u2，长度为 2（u1 数组长度的 2 倍），并把 u1 中的元素拷贝到 u2 中，最后将 s 内部表示中的 array 指向 u2，并设置 len = 2, cap = 2； 
- 然后，第三步，通过 append 操作向切片 s 添加了第三个元素 13，这时 len(s) = 2， cap(s) = 2，append 判断底层数组剩余空间不能满足添加新元素的要求了，于是又创建了 一个新的底层数组 u3，长度为 4（u2 数组长度的 2 倍），并把 u2 中的元素拷贝到 u3 中，最后把 s 内部表示中的 array 指向 u3，并设置 len = 3, cap 为 u3 数组长度，也就是 4 ；
- 第四步，依然通过 append 操作向切片 s 添加第四个元素 14，此时 len(s) = 3, cap(s) = 4，append 判断底层数组剩余空间可以满足添加新元素的要求，所以就把 14 放在下一个元素的位置 (数组 u3 末尾），并把 s 内部表示中的 len 加 1，变为 4； 
- 但第五步又通过 append 操作，向切片 s 添加最后一个元素 15，这时 len(s) = 4， cap(s) = 4，append 判断底层数组剩余空间又不够了，于是创建了一个新的底层数组 u4，长度为 8（u3 数组长度的 2 倍），并将 u3 中的元素拷贝到 u4 中，最后将 s 内部表示中的 array 指向 u4，并设置 len = 5, cap 为 u4 数组长度，也就是 8。 

到这里，这个动态扩容的过程就结束了。

可以看到，append 会根据切片的需要，在当前底层数组容量无法满足的情况下，**动态分配新的数组**，新数组长度会按一定规律扩展。

在上面这段代码中，针对元素是 int 型的数组，新数组的容量是当前数组的 2 倍。新数组建立后，append 会把旧数组中的数据拷贝到新数组中，之后新数组便成为了切片的底层数组，旧数组会被垃圾回收掉。 



### 动态扩容导致解除绑定问题

不过 append 操作的这种自动扩容行为，有些时候会给开发者带来一些**困惑**，比如基于一个已有数组建立的切片，一旦追加的数据操作触碰到切片的容量上限（实质上也是数组容量的上界)，切片就会和原数组**解除“绑定”**，后续对切片的任何修改都不会反映到原数组中了。再来看这段代码：

```go
// 自动扩容问题：切片与数组解除绑定
// 定义数组
u := [...]int{11, 12, 13, 14, 15}
fmt.Println("array:", u) // [11, 12, 13, 14, 15]
// 开始切片
s := u[1:3]
fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
s = append(s, 24)
fmt.Println("after append 24, array:", u)
fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
s = append(s, 25)
fmt.Println("after append 25, array:", u)
fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
// 切片和原数组解除绑定
s = append(s, 26)
fmt.Println("after append 26, array:", u)
fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
// 测试是否真的解除绑定
s[0] = 22
fmt.Println("after reassign 1st elem of slice, array:", u)
fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
```

运行这段代码，得到这样的结果：

```sh
array: [11 12 13 14 15]
slice(len=2, cap=4): [12 13]
after append 24, array: [11 12 13 24 15]
after append 24, slice(len=3, cap=4): [12 13 24]
after append 25, array: [11 12 13 24 25]
after append 25, slice(len=4, cap=4): [12 13 24 25]
after append 26, array: [11 12 13 24 25]
after append 26, slice(len=5, cap=8): [12 13 24 25 26]
after reassign 1st elem of slice, array: [11 12 13 24 25]
after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]
```

这里，在 append 25 之后，切片的元素已经触碰到了底层数组 u 的边界了。然后再 append 26 之后，append 发现底层数组已经无法满足 append 的要求，于是**新创建了一个底层数组**（数组长度为 cap(s) 的 2 倍，即 8），并将 slice 的元素拷贝到新数组中了。 

在这之后，即便再修改切片的第一个元素值，原数组 u 的元素也不会发生改变了，因 为这个时候切片 s 与数组 u 已经解除了“绑定关系”，s 已经不再是数组 u 的“描述符”了。

这种因切片的自动扩容而导致的“绑定关系”解除，有时候会成为实践道路上的一个小陷阱。



## 小结 

最常使用的两种同构复合数据类型：数组和切片。 

- **数组**是一个固定长度的、由同构类型元素组成的连续序列。
  - 这种连续不仅仅是逻辑上的， Go 编译器为数组类型变量分配的也是一整块可以容纳其所有元素的连续内存。
  - 而且，Go 编译器为数组变量的初始化也提供了很多便利。
  - 当数组元素的类型也是数组类型时，会出现多维数组。只需要按照变量声明从左到右、按维度分层拆解，直到出现一元数组就好了。 
  - 但是，Go 值传递的机制让数组在各个函数间传递起来比较“笨重”，开销较大，且开销随数组长度的增加而增加。
  - 为了解决这个问题，Go 引入了切片这一不定长同构数据类型。
- **切片**可以看成是数组的“描述符”，为数组打开了一个访问与修改的“窗口”。
  - 切片在 Go 运行时中被实现为一个“三元组（array, len, cap）”，其中的 array 是指向底层数组的指针，真正的数据都存储在这个底层数组中；len 表示切片的长度；而 cap 则是切片底层数组的容量。
  - 可以为一个数组建立多个切片，这些切片由于共享同一个底层数组，因此通过任一个切片对数组的修改都会反映到其他切片中。 
  - 切片是不定长同构复合类型，这个不定长体现在 Go 运行时对它提供的动态扩容的支撑。 
  - 当切片的 cap 值与 len 值相等时，如果再向切片追加数据，Go 运行时会自动对切片的底层数组进行扩容，追加数据的操作不会失败。 
  - 在大多数场合，都会使用切片以替代数组，
    - 原因之一是切片作为数组“描述符”的轻量性，无论它绑定的底层数组有多大，传递这个切片花费的开销都是恒定可控的；
    - 另外一个原因是切片相较于数组指针也是有优势的，切片可以提供比指针更为强大的功能，比如下标访问、边界溢出校验、动态扩容等。
    - 而且，指针本身在 Go 语言中的功能也受到的限制，比如不支持指针算术运算。



## 思考题

请描述一下下面这两个切片变量 sl1 与 sl2 的差异。

```go
var sl1 []int
var sl2 = []int{}
```

- s1是声明，还没初始化，是nil值，和nil比较返回true，底层没有分配内存空间。
- s2初始化为 empty slice，不是nil值，和nil比较返回false，底层分配了内存空间，有地址。



## map

Go 语言中最常用的两个复合类型：数组与切片。它们代表一组连续存储的同构类型元素集合。不同的是，数组的长度是确定的，而切片，可以理解为 一种“动态数组”，它的长度在运行时是可变的。

另外一种日常 Go 编码中比较常用的复合类型， 这种类型可以**将一个值（Value）唯一关联到一个特定的键（Key）上**，可以用于**实现特定键值的快速查找与更新**，这个复合数据类型就是 map。

很多中文 Go 编程语言类技术书籍都会将它翻译为映射、哈希表或字典，为了保持原汁原味，直接使用它的英文名，map。

map 是继切片之后，学到的第二个由 Go 编译器与运行时联合实现的复合数据类型， 它有着复杂的内部实现，但却提供了十分简单友好的开发者使用接口。

### 什么是 map 类型？ 

map 是 Go 语言提供的一种抽象数据类型，它表示一组无序的键值对。在后面的讲解中， 会直接使用 key 和 value 分别代表 map 的键和值。

而且，map 集合中每个 key 都是唯一的：

![image-20211225214240723](go_language_compound_data_type.assets/image-20211225214240723.png)

和切片类似，作为复合类型的 map，它在 Go 中的类型表示也是由 key 类型与 value 类型组成的，就像下面代码：

```go
map[key_type]value_type
```

key 与 value 的类型可以相同，也可以不同：

```go
map[string]string // key与value元素的类型相同
map[int]string // key与value元素的类型不同
```

如果两个 map 类型的 key 元素类型相同，value 元素类型也相同，那么可以说它们是同一个 map 类型，否则就是不同的 map 类型。 

这里，要注意，map 类型对 value 的类型没有限制，但是对 key 的类型却有严格要求，因为 map 类型要保证 key 的唯一性。Go 语言中要求，**key 的类型必须支持“==”和“!=”两种比较操作符**。 

但是，在 Go 语言中，**函数类型、map 类型自身，以及切片只支持与 nil 的比较**，而不支持同类型两个变量的比较。

如果像下面代码这样，进行这些类型的比较，Go 编译器将会报错：

```go
// == 的比较操作
s1 := make([]int, 1)
s2 := make([]int, 2)
f1 := func() {}
f2 := func() {}
m1 := make(map[int]string)
m2 := make(map[int]string)

println(s1 == s2) // 错误：invalid operation: s1 == s2 (slice can only be compared to nil)
println(f1 == f2) // 错误：invalid operation: f1 == f2 (func can only be compared to nil)
println(m1 == m2) // 错误：invalid operation: m1 == m2 (map can only be compared to nil)
```

因此在这里，一定要注意：**函数类型、map 类型自身，以及切片类型是不能作为 map 的 key 类型的**。 



### map 变量的声明 

可以这样声明一个 map 变量：

```go
// map 的声明
var m map[string]int // 一个map[string]int 类型的变量
```

和切片类型变量一样，如果没有显式地赋予 map 变量初值，map 类型变量的**默认值 为 nil**。 

不过切片变量和 map 变量在这里也有些不同。

- 初值为零值 nil 的切片类型变量，可以借助**内置的 append 的函数进行操作**，这种在 Go 语言中被称为“零值可用”。定义“零值可用”的类型，可以提升开发者的使用体验，不用再担心变量的初始状态是否有效。 

- 但 map 类型，因为它内部实现的复杂性，**无法“零值可用”**。所以，如果对处于零值状态的 map 变量直接进行操作，就会导致运行时异常（panic），从而导致程序进程异常退出：

  - ```go
    // map 的声明
    var m map[string]int // m = nil
    m["key"] = 1         // 发生运行时异常：panic: assignment to entry in nil map
    fmt.Println(m)       // map[]
    ```

所以，必须对 map 类型变量进行**显式初始化后才能使用**。

### map 变量的初始化 

和切片一样，为 map 类型变量显式赋值有两种方式：

- 一种是使用复合字面值；
- 另外一种是使用 make 这个预声明的内置函数。 

#### 复合字面值初始化

方法一：使用**复合字面值初始化 map 类型变量**。 

先来看这句代码：

```go
// 复合字面值初始化 map 类型变量
n := map[int]string{}
```

这里，显式初始化了 map 类型变量 n。不过，要注意，虽然此时 map 类型变量 n 中没有任何键值对，但变量 n 也不等同于初值为 nil 的 map 变量。

这个时候，对 n 进行键值对的插入操作，不会引发运行时异常。 

```go
// 复合字面值初始化 map 类型变量
n := map[int]string{}
n[1] = "liu"
fmt.Println(n) // map[1:liu]
```

这里再看看怎么通过稍微**复杂一些的复合字面值**，对 map 类型变量进行初始化：

```go
// 复杂字面值初始化
m1 := map[int][]string{
   1: []string{"val1_1", "val1_2"},
   3: []string{"val3_1", "val3_2", "val3_3"},
   7: []string{"val7_1"},
}

type Position struct {
   x float64
   y float64
}

m2 := map[Position]string{
   Position{29.935523, 52.568915}:  "school",
   Position{25.352594, 113.304361}: "shopping-mall",
   Position{73.224455, 111.804306}: "hospital",
}
fmt.Println(m1) // map[1:[val1_1 val1_2] 3:[val3_1 val3_2 val3_3] 7:[val7_1]]
fmt.Println(m2) // map[{25.352594 113.304361}:shopping-mall {29.935523 52.568915}:school {73.224455 111.804306}:hospital]
```

上面代码虽然完成了对两个 map 类型变量 m1 和 m2 的显式初始化，但有一个问题，作为初值的字面值似乎**有些“臃肿”**。

作为初值的字面值采用了复合类型的元素类型，而且在编写字面值时还带上了各自的元素类型，比如作为 map[int] []string 值类型的[]string，以及作为 map[Position]string 的 key 类型的 Position。 

针对这种情况，Go 提供了**“语法糖”**。

这种情况下，Go 允许**省略字面值中的元素类型**。因为 map 类型表示中包含了 key 和 value 的元素类型，Go 编译器已经有足够的信息，来推导出字面值中各个值的类型了。

以 m2 为例，这里的显式初始化代码和上面变量 m2 的初始化代码是等价的：

```go
// 省略字面值中的元素类型
m3 := map[Position]string{
   {29.935523, 52.568915}:  "school",
   {25.352594, 113.304361}: "shopping-mall",
   {73.224455, 111.804306}: "hospital",
}
fmt.Println(m3) // map[{25.352594 113.304361}:shopping-mall {29.935523 52.568915}:school {73.224455 111.804306}:hospital]
```

以后在无特殊说明的情况下，都将使用这种简化后的字面值初始化方式。 

#### make 初始化

方法二：**使用 make 为 map 类型变量进行显式初始化**。

和切片通过 make 进行初始化一样，通过 make 的初始化方式，可以为 map 类型变量指定键值对的初始容量，但无法进行具体的键值对赋值，就像下面代码这样：

```go
// make 初始化
m4 := make(map[int]string)    // 未指定初始容量
m5 := make(map[int]string, 8) // 指定初始容量为8
```

不过，map 类型的容量不会受限于它的初始容量值，当其中的键值对数量超过初始容量后，Go 运行时会自动增加 map 类型的容量，保证后续键值对的正常插入。 

### map 的基本操作 

针对一个 map 类型变量，可以进行诸如插入新键值对、获取当前键值对数量、查找特定键和读取对应值、删除键值对，以及遍历键值等操作。

#### 插入键值对

操作一：插入新键值对。 

面对一个非 nil 的 map 类型变量，可以在其中插入符合 map 类型定义的任意新键值对。插入新键值对的方式很简单，只需要把 value 赋值给 map 中对应的 key 就可以了：

```go
// 插入操作
m := make(map[int]string)
m[1] = "value1"
m[2] = "value2"
m[3] = "value3"
fmt.Println(m)  // map[1:value1 2:value2 3:value3]
```

而且，不需要自己判断数据有没有插入成功，因为 Go 会保证插入总是成功的。

这里，Go 运行时会负责 map 变量内部的内存管理，因此除非是系统内存耗尽，可以不用担心向 map 中插入新数据的数量和执行结果。

不过，如果插入新键值对的时候，某个 key 已经存在于 map 中了，那插入操 作就会**用新值覆盖旧值**：

```go
// 插入操作 之 新值覆盖旧值
m1 := map[string]int{
   "key1": 1,
   "key2": 2,
}
fmt.Println(m1) // map[key1:1 key2:2]
m1["key1"] = 11 // 11会覆盖掉"key1"对应的旧值1
m1["key3"] = 3  // 此时m1为map[key1:11 key2:2 key3:3]
fmt.Println(m1) // map[key1:11 key2:2 key3:3]
```

从这段代码中可以看到，map 类型变量 m1 在声明的同时就做了初始化，它的内部建立了两个键值对，其中就包含键 key1。

所以后面再给键 key1 进行赋值时，Go 不会重新创建 key1 键，而是会用新值 (11) 把 key1 键对应的旧值 (1) 替换掉。 

#### 获取键值对数量

操作二：获取键值对数量。 

如果在编码中，想知道当前 map 类型变量中已经建立了多少个键值对，那可以怎么做呢？

和切片一样，map 类型也可以通过**内置函数 len**，获取当前变量已经存储的键值对数量：

```go
// 获取键值对数量
m2 := map[string]int{
   "key1": 1,
   "key2": 2,
}
fmt.Println(len(m2)) // 2
m2["key3"] = 3
fmt.Println(len(m2)) // 3
```

不过，这里要注意的是不能对 map 类型变量调用 cap，来获取当前容量，这是 map 类型与切片类型的一个不同点。 

#### 查找和数据读取

操作三：查找和数据读取

和写入相比，map 类型更多用在查找和数据读取场合。所谓查找，就是判断某个 key 是否存在于某个 map 中。

有了向 map 插入键值对的基础，可能自然而然地想到，可以用下面代码去查找一个键并获得该键对应的值：

```go
// 查找
m3 := make(map[string]int)
v := m3["key1"]
fmt.Println(v) // 0

m3["key1"] = 666
v2 := m3["key1"]
fmt.Println(v2) // 666
```

乍一看，第二行代码在语法上好像并没有什么不当之处，但其实通过这行语句，还是无法确定键 key1 是否真实存在于 map 中。

这是因为，当尝试去获取一个键对应的值的时候，如果这个键在 map 中并不存在，也会得到一个值，这个值是 value 元素类型的**零值**。 

以上面这个代码为例，如果键 key1 在 map 中并不存在，那么 v 的值就会被赋予 value 元素类型 int 的零值，也就是 0。所以无法通过 v 值判断出，究竟是因为 key1 不存在返回的零值，还是因为 key1 本身对应的 value 就是 0。 

那么在 map 中查找 key 的正确姿势是什么呢？Go 语言的 map 类型支持通过用一种名 为“**comma ok**”的惯用法，进行对某个 key 的查询。

接下来就用“comma ok”惯用法改造一下上面的代码：

```go
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
```

这里通过了一个布尔类型变量 ok，来判断键“key1”是否存在于 map 中。如果存在，变量 v4 就会被正确地赋值为键“key1”对应的 value。 

不过，如果并不关心某个键对应的 value，而**只关心某个键是否在于 map 中**，可以使用空标识符替代变量 v，忽略可能返回的 value：

```go
// 查找 之 comma ok 手法 之 空标识符
m5 := make(map[string]int)
_, ok1 := m5["key1"]
// ... ...
fmt.Println(ok1)  // false
```

因此，一定要记住：在 Go 语言中，请使用“comma ok”惯用法对 map 进行键查找和键值读取操作。

#### 删除数据

操作四：删除数据。 

在 Go 中，需要借助**内置函数 delete 来从 map 中删除数据**。使用 delete 函数的情况下，传入的第一个参数是 map 类型变量，第二个参数就是想要删除的键。

可以看看这个代码示例：

```go
// 删除操作
m6 := map[string]int{
   "key1": 1,
   "key2": 2,
}
fmt.Println(m6)    // map[key1:1 key2:2]
delete(m6, "key2") // 删除"key2"
fmt.Println(m6)    // map[key1:1]
```

这里要注意的是，**delete 函数是从 map 中删除键的唯一方法**。

即便传给 delete 的键在 map 中并不存在，delete 函数的执行也不会失败，更不会抛出运行时的异常。 

#### 遍历 map 中的键值数据

操作五：遍历 map 中的键值数据 

最后，来说一下如何遍历 map 中的键值数据。这一点虽然不像查询和读取操作那么常见，但日常开发中还是有这个需求的。

在 Go 中，遍历 map 的键值对只有一种方法， 那就是像对待切片那样**通过 for range 语句对 map 数据进行遍历**。来看一个例子：

```go
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
```

通过 for range 遍历 map 变量 m7，每次迭代都会返回一个键值对，其中键存在于变量 k 中，它对应的值存储在变量 v 中。

可以运行一下这段代码，可以得到符合预期的结果：

```go
{ [1, 11] [2, 12] [3, 13] }
```

如果只关心每次迭代的键，可以使用下面的方式对 map 进行遍历：

```go
// 只关心键
for k, _ := range m7 {
   // 只使用 k
   fmt.Printf("key: %d\n", k)
}
```

更地道的方式是这样的：

```go
// 只关心键 更地道方式
for k := range m7 {
   // 只使用 k
   fmt.Printf("key: %d\n", k)
}
```

如果只关心每次迭代返回的键所对应的 value，同样可以通过空标识符替代变量 k，就像下面这样：

```go
// 只关心值
for _, v := range m7 {
   // 只使用 k
   fmt.Printf("value: %d\n", v)
}
```

不过，前面 map 遍历的输出结果都非常理想，表象好像是迭代器按照 map 中元素的插入次序逐一遍历。那事实是不是这样呢？

再来试试，多遍历几次这个 map 看看。 先来改造一下代码：

```go
package main

// map 遍历
import "fmt"

func doIteration(m map[int]int) {
   fmt.Printf("{ ")
   for k, v := range m {
      fmt.Printf("[%d, %d] ", k, v)
   }
   fmt.Printf("}\n")
}

func main() {
   m := map[int]int{
      1: 11,
      2: 12,
      3: 13,
   }

   for i := 0; i < 3; i++ {
      doIteration(m)
   }
}
```

运行一下上述代码，可以得到这样结果：

```go
{ [2, 12] [3, 13] [1, 11] }
{ [1, 11] [2, 12] [3, 13] }
{ [1, 11] [2, 12] [3, 13] }
```

可以看到，对同一map 做多次遍历的时候，**每次遍历元素的次序都不相同**。

这是 Go 语言 map 类型的一个重要特点，也是很容易让 Go 初学者掉入坑中的一个地方。所以这里一定要记住：**程序逻辑千万不要依赖遍历 map 所得到的元素次序**。 



### map 变量的传递开销 

其实不用担心开销的问题。 和切片类型一样，**map 也是引用类型**。

这就意味着 map 类型变量作为参数被传递给函数或方法的时候，实质上**传递的只是一个“描述符”**，而不是整个 map 的数据拷贝，所以这个传递的开销是固定的，而且也很小。 

并且，当 map 变量被传递到函数或方法内部后，在函数内部对 map 类型参数的修改在函数外部也是可见的。比如从这个示例中就可以看到，函数 foo 中对 map 类型变量 m 进行了修改，而这些修改在 foo 函数外也可见。

```go
package main

import "fmt"

func foo(m map[string]int) {
   m["key1"] = 11
   m["key2"] = 12
}

func main() {
   m := map[string]int{
      "key1": 1,
      "key2": 2,
   }

   fmt.Println(m) // map[key1:1 key2:2]
   foo(m)
   fmt.Println(m) // map[key1:11 key2:12]
}
```



### map 的内部实现 

和切片相比，map 类型的内部实现要更加**复杂**。

Go 运行时使用一张**哈希表**来实现抽象的 map 类型。运行时实现了 map 类型操作的所有功能，包括查找、插入、删除等。

在编译阶段，Go 编译器会将 Go 语法层面的 map 操作，重写成运行时对应的函数调用。大致的对应关系是这样的：

```go
// 创建map类型变量实例
m := make(map[keyType]valType, capacityhint) → m := runtime.makemap(maptype, capacityhint, m)

// 插入新键值对或给键重新赋值
m["key"] = "value" → v := runtime.mapassign(maptype, m, "key") // v是用于后续存储value的空间的地址

// 获取某键的值
v := m["key"] → v := runtime.mapaccess1(maptype, m, "key")
v, ok := m["key"] → v, ok := runtime.mapaccess2(maptype, m, "key")

// 删除某键
delete(m, "key") → runtime.mapdelete(maptype, m, "key")
```

这是 map 类型在 Go 运行时层的实现示意图：

![image-20211225231116615](go_language_compound_data_type.assets/image-20211225231116615.png)

可以看到，和切片的运行时表示图相比，map 的实现示意图显然要复杂得多。接下来，结合这张图来简要描述一下 map 在运行时层的实现原理。

#### 初始状态 

从图中可以看到，与语法层面 map 类型变量（m）一一对应的是 runtime.hmap 的实例。

**hmap 类型**是 map 类型的头部结构（header），也就是前面在讲解 map 类型 变量传递开销时提到的 **map 类型的描述符**，它存储了后续 map 类型操作所需的所有信息，包括：

![image-20211225231434469](go_language_compound_data_type.assets/image-20211225231434469.png)

源码：

```go
// A header for a Go map.
type hmap struct {
   count     int // # live cells == size of map.  Must be first (used by len() builtin)
   flags     uint8
   B         uint8  // log_2 of # of buckets (can hold up to loadFactor * 2^B items)
   noverflow uint16 // approximate number of overflow buckets; see incrnoverflow for details
   hash0     uint32 // hash seed

   buckets    unsafe.Pointer // array of 2^B Buckets. may be nil if count==0.
   oldbuckets unsafe.Pointer // previous bucket array of half the size, non-nil only when growing
   nevacuate  uintptr        // progress counter for evacuation (buckets less than this have been evacuated)

   extra *mapextra // optional fields
}
```

真正用来存储键值对数据的是**桶，也就是 bucket**，每个 bucket 中存储的是 Hash 值低 bit 位数值相同的元素，默认的元素个数为 **BUCKETSIZE（值为 8**，在 $GOROOT/src/cmd/compile/internal/gc/reflect.go 中定义，与 runtime/map.go 中常量 bucketCnt 保持一致）。 

```go
// cmd/compile/internal/gc/reflect.go
const (
	BUCKETSIZE  = 8
	// ...
)

// runtime/map.go
const(
    bucketCntBits = 3
    bucketCnt     = 1 << bucketCntBits  // bucketCnt = 8
)
```

当某个 bucket（比如 buckets[0]) 的 8 个**空槽 slot**）都填满了，且 map 尚未达到扩容的条件的情况下，**运行时会建立 overflow bucket**，并将这个 overflow bucket 挂在上面 bucket（如 buckets[0]）末尾的 overflow 指针上，这样两个 buckets 形成了一个**链表结构**，直到下一次 map 扩容之前，这个结构都会一直存在。 

从图中可以看到，每个 bucket 由三部分组成，从上到下分别是 tophash 区域、key 存储区域和 value 存储区域。

#### tophash 区域

向 map 插入一条数据，或者是从 map 按 key 查询数据的时候，运行时都会使用哈希函数对 key 做哈希运算，并获得一个**哈希值（hashcode）**。

这个 hashcode 非常关键，运行时会**把 hashcode“一分为二”**来看待，其中**低**位区的值用于选定 bucket，**高**位区的值用于在某个 bucket 中确定 key 的位置。

把这一过程整理成了下面这张示意图， 理解起来可以更直观：

![image-20211225232502811](go_language_compound_data_type.assets/image-20211225232502811.png)

因此，每个 bucket 的 tophash 区域其实是用来快速定位 key 位置的，这样就避免了逐个 key 进行比较这种代价较大的操作。

尤其是当 key 是 size 较大的字符串类型时，好处就更突出了。这是一种**以空间换时间**的思路。

#### key 存储区域

接着，看 tophash 区域下面是一块连续的内存区域(这个连续的存储结构是数组存储形式，A map is just a hash table. The data is arranged into an array of buckets.)，存储的是这个 bucket 承载的所有 key 数据。

运行时在分配 bucket 的时候需要知道 key 的 Size。那么运行时是如何知道 key 的 size 的呢？

当声明一个 map 类型变量，比如 var m map[string]int 时，Go 运行时就会为这个变量对应的特定 map 类型，**生成一个 runtime.maptype 实例**。如果这个实例已经存在，就会直接复用。

maptype 实例的结构是这样的：

```go
// runtime/type.go
type maptype struct {
   typ    _type
   key    *_type
   elem   *_type
   bucket *_type // internal type representing a hash bucket
   // function for hashing keys (ptr to key, seed) -> hash
   hasher     func(unsafe.Pointer, uintptr) uintptr
   keysize    uint8  // size of key slot
   elemsize   uint8  // size of elem slot
   bucketsize uint16 // size of bucket
   flags      uint32
}
```

可以看到，这个实例包含了需要的 **map 类型中的所有"元信息"**。

前面提到过，编译器会把语法层面的 map 操作重写成运行时对应的函数调用，这些运行时函数都有一个共同的特点，那就是第一个参数都是 maptype 指针类型的参数。 

Go 运行时就是利用 maptype 参数中的信息确定 key 的类型和大小的。map 所用的 **hash 函数也存放在 maptype.key.alg.hash(key, hmap.hash0) 中**。

同时 maptype 的存在也让 Go 中所有 map 类型都**共享一套运行时 map 操作函数**，而不是像 C++ 那样为每种 map 类型创建一套 map 操作函数，这样就节省了对最终二进制文件空间的占用。

#### value 存储区域

再接着看 key 存储区域下方的另外一块连续的内存区域，这个区域存储的是 key 对应的 value。

和 key 一样，这个区域的创建也是得到了 maptype 中信息的帮助。Go 运行时 采用了**把 key 和 value 分开存储的方式**，而不是采用一个 kv 接着一个 kv 的 kv 紧邻方式存储，这带来的其实是算法上的复杂性，但却**减少了因内存对齐带来的内存浪费**。 

以 map[int8]int64 为例，看看下面的存储空间利用率对比图：

![image-20211226000034508](go_language_compound_data_type.assets/image-20211226000034508.png)

当前 Go 运行时使用的方案内存利用效率很高，而 kv 紧邻存储的方案在 map[int8]int64 这样的例子中内存浪费十分严重，它的内存利用率是 72/128=56.25%（或者9字节/16字节=56.25%）， 有近一半的空间都浪费掉了。 

另外，还有一点要强调一下，如果 **key 或 value 的数据长度大于一定数值**，那么运行时不会在 bucket 中直接存储数据，而是会**存储 key 或 value 数据的指针**。

目前 Go 运行时定义的**最大 key 和 value 的长度**是这样的：

```go
// runtime/map.go
const(
    maxKeySize  = 128
    maxElemSize = 128
)
```



### map 扩容 

map 会对底层使用的内存进行自动管理。因此，在使用过程中，当插入元素个数超出一定数值后，map 一定会存在自动扩容的问题，也就是怎么**扩充 bucket 的数量**，并重新在 bucket 间均衡分配数据的问题。

那么 map 在什么情况下会进行扩容呢？

Go 运行时的 map 实现中引入了一个 **LoadFactor（负载因子）**，**当 count > LoadFactor * 2^B 或 overflow bucket 过多时**，运行时会自动对 map 进行扩容。

目前 Go 最新 1.17 版本 LoadFactor 设置为 6.5（loadFactorNum/ loadFactorDen）。(目前在 1.65 版本看到也是 6.5)

这里是 Go 中与 map 扩容相关的部分源码：

```go
// runtime/map.go
const(
	// Maximum average load of a bucket that triggers growth is 6.5.
	// Represent as loadFactorNum/loadFactorDen, to allow integer math.
	loadFactorNum = 13
	loadFactorDen = 2
)

func overLoadFactor(count int, B uint8) bool {
   return count > bucketCnt && uintptr(count) > loadFactorNum*(bucketShift(B)/loadFactorDen)
}

func mapassign(t *maptype, h *hmap, key unsafe.Pointer) unsafe.Pointer {
  // ...
	if !h.growing() && (overLoadFactor(h.count+1, h.B) || tooManyOverflowBuckets(h.noverflow, h.B)) {
		hashGrow(t, h)
		goto again // Growing the table invalidates everything, so try again
	}
  // ...
}
```

这两方面原因导致的扩容，在运行时的操作其实是不一样的。

- 如果是因为 overflow bucket 过多导致的“扩容”，实际上运行时会**新建一个和现有规模一样的 bucket 数组**， 然后在 assign 和 delete 时做排空和迁移。 
- 如果是因为当前数据数量超出 LoadFactor 指定水位而进行的扩容，那么运行时会**建立一个两倍于现有规模的 bucket 数组**，但真正的排空和迁移工作也是在 assign 和 delete 时逐步进行的。
  - 原 bucket 数组会挂在 hmap 的 oldbuckets 指针下面，直到原 buckets 数组中所有数据都迁移到新数组后，原 buckets 数组才会被释放。
  - 可以结合下面的 map 扩容示意图来理解这个过程，这会理解得更深刻一些：

![image-20211226003016256](go_language_compound_data_type.assets/image-20211226003016256.png)

### map 与并发 

从上面的实现原理来看，充当 map 描述符角色的 hmap 实例自身是有状态的（hmap.flags），而且对状态的读写是没有并发保护的。

所以说 map 实例不是并发写安全的，也**不支持并发读写**。如果对 map 实例进行并发读写，程序运行时就会抛出异常。

可以看看下面这个并发读写 map 的例子：

```go
package main

import (
   "fmt"
   "time"
)

func doIteration(m map[int]int) {
   for k, v := range m {
      _ = fmt.Sprintf("[%d, %d] ", k, v)
   }
}

func doWrite(m map[int]int) {
   for k, v := range m {
      m[k] = v + 1
   }
}

func main() {
   m := map[int]int{
      1: 11,
      2: 12,
      3: 13,
   }

   go func() {
      for i := 0; i < 1000; i++ {
         doIteration(m)
      }
   }()

   go func() {
      for i := 0; i < 1000; i++ {
         doWrite(m)
      }
   }()

   time.Sleep(5 * time.Second)
}
```

运行这个示例程序，会得到下面的执行错误结果：

```go
fatal error: concurrent map iteration and map write
```

不过，如果仅仅是进行并发读，map 是没有问题的。

而且，Go 1.9 版本中引入了**支持并发写安全的 sync.Map 类型**，可以用来在并发读写的场景下替换掉 map，可以查看一下 sync.Map 的手册。 

另外，要注意，考虑到 map 可以自动扩容，map 中数据元素的 value 位置可能在这一过程中发生变化，所以 **Go 不允许获取 map 中 value 的地址**，这个约束是在编译期间就生效的。

下面这段代码就展示了 Go 编译器识别出获取 map 中 value 地址的语句后，给出的编译错误：

```go
// 获取 map 中的 value 地址，编译不通过
m8 := make(map[string]int)
m8["key1"] = 678
p := &m8["key1"] // cannot take the address of m[key]
fmt.Println(p)
```



## 小结

在 Go 语言中，map 类型是一个无序的键值对的集合。它有两种类型元素，一类是键 （key），另一类是值（value）。在一个 map 中，键是唯一的，在集合中不能有两个相同的键。

Go 也是通过这两种元素类型来表示一个 map 类型，要记得这个通用的 map 类型表示：“map[key_type]value_type”。 

map 类型对 key 元素的类型是有约束的，它要求 key 元素的类型必须支 持"==“和”!="两个比较操作符。value 元素的类型可以是任意的。 

不过，map 类型变量声明后必须对它进行初始化后才能操作。

map 类型支持插入新键值对、查找和数据读取、删除键值对、遍历 map 中的键值数据等操作，Go 为开发者提供了十分简单的操作接口。这里要重点记住的是，在查找和数据读取时一定要使用“comma ok”惯用法。

此外，map 变量在函数与方法间传递的开销很小，并且在函数内部通过 map 描述符对 map 的修改会对函数外部可见。 

另外，map 的内部实现要比切片复杂得多，它是由 Go 编译器与运行时联合实现的。Go 编译器在编译阶段会将语法层面的 map 操作，重写为运行时对应的函数调用。Go 运行时则采用了高效的算法实现了 map 类型的各类操作，这里要结合 Go 项目源码来理解 map 的具体实现。 

和切片一样，map 是 Go 语言提供的重要数据类型，也是 Gopher 日常 Go 编码是最常使用的类型之一。在日常使用 map 的场合要把握住下面几个要点，不要走弯路：

- 不要依赖 map 的元素遍历顺序； 
- map 不是线程安全的，不支持并发读写； 
- 不要尝试获取 map 中元素（value）的地址。



## 思考题 

对 map 类型进行遍历所得到的键的次序是随机的，那么是否可以实现一个方法，能对 map 的进行稳定次序遍历？

- 取出map中的key，并且把把key存到有序切片中，用切片遍历；

- 也可以使用额外的链表来保存顺序，参考：java的LinkedHashMap；

- go 中可以基于container/list来实现。github上现成的功能，https://github.com/elliotchance/orderedmap

- Go 官网给出的解决方案：https://go.dev/blog/maps

  - Iteration order

  - When iterating over a map with a range loop, the iteration order is not specified and is not guaranteed to be the same from one iteration to the next. If you require a stable iteration order you must maintain a separate data structure that specifies that order. 

  - This example uses a separate sorted slice of keys to print a `map[int]string` in key order:

  - ```go
    import "sort"
    
    var m map[int]string
    var keys []int
    for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
    for _, k := range keys {
        fmt.Println("Key:", k, "Value:", m[k])
    }
    ```

- 深入研究 map 的小伙伴，可以去研究这些博客：
  - 深度解密Go语言之map：https://www.qcrao.com/2019/05/22/dive-into-go-map/
  - 深度解密Go语言之sync.map：https://qcrao.com/2020/05/06/dive-into-go-sync-map/
  - 哈希表：https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/





















