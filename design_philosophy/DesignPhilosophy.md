# DesignPhilosophy

## 显示设计哲学示例

来一段 C 程序，看看“隐式”代码的行为特征。

```c
#include <stdio.h>
int main() {
  short int a = 5;
  
  int b = 8;
  long c = 0;
  
  c = a + b;
  
  printf("%ld\n", c);
}
```

在上面这段代码中，变量 a、b 和 c 的类型均不相同，C 语言编译器在编译c = a + b这一行时，会自动将短整型变量 a 和整型变量 b，先转换为 long 类型然后相加， 并将所得结果存储在 long 类型变量 c 中。

那如果换成 Go 来实现这个计算会怎么样呢？先把上面的 C 程序转化成等价的 Go 代码：

```go
package main

import "fmt"

func main() {
	// error example
	var a int16 = 5

	var b int = 8
	var c int64

	c = a + b

	fmt.Printf("%d\n", c)
}

```

如果编译这段程序，将得到类似这样的编译器错误：“invalid operation: a + b (mismatched types int16 and int)”。

能看到 Go 与 C 语言的隐式自动类型转换不同，Go 不允许不同类型的整型变量进行混合计算，它同样也不会对其进行隐式的自动转换。 

因此，如果要使这段代码通过编译，就需要对变量 a 和 b 进行显式转型，就像下面代 码段中这样：

```go
package main

import "fmt"

func main() {
	// correct example
	var a int16 = 5

	var b int = 8
	var c int64

	c = int64(a) + int64(b)

	fmt.Printf("%d\n", c)
}
```

而这其实就是 Go 语言显式设计哲学的一个体现。 

在 Go 语言中，不同类型变量是不能在一起进行混合计算的，这是因为 Go 希望开发人员 明确知道自己在做什么，这与 C 语言的“信任程序员”原则完全不同，因此需要以显式 的方式通过转型统一参与计算各个变量的类型。 











