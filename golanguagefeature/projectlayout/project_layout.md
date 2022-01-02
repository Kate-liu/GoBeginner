# Project Layout

一般编写的 Go 程序都是简单程序，一般由一个或几个 Go 源码文件组成，而且所有源码文件都在同一个目录中。但是生产环境中运行的实用程序可不会这么简 单，通常它们都有着复杂的项目结构布局。

弄清楚一个实用 Go 项目的项目布局标准是 Go 开发者走向编写复杂 Go 程序的第一步，也是必经的一步。 但 Go 官方到目前为止也没有给出一个关于 Go 项目布局标准的正式定义。

## Go 语言“创世项目”结构是怎样的？

要想了解 Go 项目的结构布局以及演化历史，全世界第一个 Go 语言项目是一个最好的切 入点。

什么是“Go 语言的创世项目”呢？其实就是 Go 语言项目自身，它是全世界第一个 Go 语言项目。但这么说也不够精确，因为 Go 语言项目从项目伊始就混杂着多种语言，而且以 C 和 Go 代码为主，Go 语言的早期版本 C 代码的比例还不小。 

### Go 1.0 版本

先用 loccount 工具对 Go 语言发布的第一个 Go 1.0 版本分析看看：

```sh
$loccount .
all SLOC=460992 (100.00%) LLOC=193045 in 2746 files
Go SLOC=256321 (55.60%) LLOC=109763 in 1983 files
C SLOC=148001 (32.10%) LLOC=73458 in 368 files
HTML SLOC=25080 (5.44%) LLOC=0 in 57 files
asm SLOC=10109 (2.19%) LLOC=0 in 133 files
... ...
```

在 1.0 版本中，Go 代码行数占据一半以上比例，但是 C 语言代码行数也占据 了 32.10% 的份额。而且在后续 Go 版本演进过程中，Go 语言代码行数占比还在逐步提 升，直到 Go 1.5 版本实现自举后，Go 语言代码行数占比将近 90%，C 语言比例下降为不 到 1%，这一比例一直延续至今。 

虽然 C 代码比例下降，Go 代码比例上升，但 Go 语言项目的布局结构却整体保留了下 来，十多年间虽然也有一些小范围变动，但整体没有本质变化。

作为 Go 语言的“创世项 目”，它的结构布局对后续 Go 社区的项目具有重要的参考价值，尤其是 Go 项目早期 src 目录下面的结构。 



### Go 1.3 版本

为了方便查看，首先下载 Go 语言创世项目源码：

```sh
$git clone https://github.com/golang/go.git
```

进入 Go 语言项目根目录后，使用 tree 命令来查看一下 Go 语言项目自身的最初源码 结构布局，以 Go 1.3 版本为例，结果是这样的：

```sh
$cd go // 进入Go语言项目根目录
$git checkout go1.3 // 切换到go 1.3版本
$tree -LF 1 ./src // 查看src目录下的结构布局
./src
├── all.bash*
├── clean.bash*
├── cmd/
├── make.bash*
├── Make.dist
├── pkg/
├── race.bash*
├── run.bash*
... ...
└── sudo.bash*
```

从上面的结果来看，src 目录下面的结构有这三个特点。 

首先，以 all.bash 为代表的代码构建的脚本源文件放在了 src 下面的顶层目录下。 

第二，src 下的二级目录 cmd 下面存放着 Go 相关可执行文件的相关目录，可以深入 查看一下 cmd 目录下的结构：

```sh
$ tree -LF 1 ./src/cmd
./src/cmd
... ...
├── 6a/
├── 6c/
├── 6g/
... ...
├── cc/
├── cgo/
├── dist/
├── fix/
├── gc/
├── go/
├── gofmt/
├── ld/
├── nm/
├── objdump/
├── pack/
└── yacc/
```

可以看到，这里的每个子目录都是一个 Go 工具链命令或子命令对应的可执行文件。 其中，6a、6c、6g 等是早期 Go 版本针对特定平台的汇编器、编译器等的特殊命名方式。 

第三个特点，会看到 src 下的二级目录 pkg 下面存放着运行时实现、标准库包实现，这 些包既可以被上面 cmd 下各程序所导入，也可以被 Go 语言项目之外的 Go 程序依赖并导 入。

下面是通过 tree 命令查看 pkg 下面结构的输出结果：

```sh
# tree -LF 1 ./src/pkg
./src/pkg
... ...
├── flag/
├── fmt/
├── go/
├── hash/
├── html/
├── image/
├── index/
├── io/
... ...
├── net/
├── os/
├── path/
├── reflect/
├── regexp/
├── runtime/
├── sort/
├── strconv/
├── strings/
├── sync/
├── syscall/
├── testing/
├── text/
├── time/
├── unicode/
└── unsafe/
```

虽然 Go 语言的创世项目的 src 目录下的布局结构，离现在已经比较久远了，但是这样的 布局特点依然对后续很多 Go 项目的布局产生了比较大的影响，尤其是那些 Go 语言早期采纳者建立的 Go 项目。

比如，Go 调试器项目 Delve、开启云原生时代的 Go 项目 Docker，以及云原生时代的“操作系统”项目 Kubernetes 等，它们的项目布局，至今都还保持着与 Go 创世项目早期相同的风格。 

当然了，这些早期的布局结构一直在不断地演化，简单来说可以归纳为下面三个比较重要 的演进。



### Go 1.4 版本

演进一：Go 1.4 版本删除 pkg 这一中间层目录并引入 internal 目录 

```sh
$git checkout go1.4 // 切换到go 1.4版本
$tree -LF 1 ./src // 查看src目录下的结构布局
```

出于简化源码树层次的原因，Go 语言项目的 Go 1.4 版本对它原来的 src 目录下的布局做了两处调整。

- 第一处是删除了 Go 源码树中“src/pkg/xxx”中 pkg 这一层级目录而直接使用 src/xxx。这样一来，Go 语言项目的源码树深度减少一层，更便于 Go 开发者阅读和 探索 Go 项目源码。 
- 另外一处就是 Go 1.4 引入 internal 包机制，增加了 internal 目录。这个 internal 机制其实是所有 Go 项目都可以用的，Go 语言项目自身也是自 Go 1.4 版本起，就使用 internal 机制了。
  - 根据 internal 机制的定义，一个 Go 项目里的 internal 目录下的 Go 包，只可以被本项目内部的包导入。项目外部是无法导入这个 internal 目录下面的包的。
  - 可以说， internal 目录的引入，让一个 Go 项目中 Go 包的分类与用途变得更加清晰。



### Go 1.6 版本

演进二：Go 1.6 版本增加 vendor 目录 

```sh
$git checkout go1.6 // 切换到go 1.6版本
$tree -LF 1 ./src // 查看src目录下的结构布局
```

第二次的演进，其实是为了解决 Go 包依赖版本管理的问题，Go 核心团队在 **Go 1.5 版本** 中做了第一次改进。

增加了 vendor 构建机制，也就是 Go 源码的编译可以不在 GOPATH 环境变量下面搜索依赖包的路径，而在 vendor 目录下查找对应的依赖包。 

Go 语言项目自身也在 **Go 1.6 版本**中增加了 vendor 目录以支持 vendor 构建，但 vendor 目录并没有实质性缓存任何第三方包。

直到 **Go 1.7 版本**，Go 才真正在 vendor 下缓存了其依赖的外部包。这些依赖包主要是 golang.org/x 下面的包，这些包同样是由 Go 核心团队维护的，并且其更新速度不受 Go 版本发布周期的影响。 

vendor 机制与目录的引入，让 Go 项目第一次具有了可重现构建（Reproducible Build） 的能力。



### Go 1.13 版本

演进三：Go 1.13 版本引入 go.mod 和 go.sum 

第三次演进，还是为了解决 Go 包依赖版本管理的问题。

在 **Go 1.11 版本**中，Go 核心团 队做出了第二次改进尝试：引入了 Go Module 构建机制，也就是在项目引入 go.mod 以 及在 go.mod 中明确项目所依赖的第三方包和版本，项目的构建就将**摆脱 GOPATH 的束缚**，实现精准的可重现构建。 

Go 语言项目自身在 Go 1.13 版本引入 go.mod 和 go.sum 以支持 Go Module 构建机 制，下面是 Go 1.13 版本的 go.mod 文件内容：

```go
module std

go 1.12  // 这个1.2，我表示懵逼，但确实是1.12

require (
	golang.org/x/crypto v0.0.0-20190611184440-5c40567a22f8
	golang.org/x/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/sys v0.0.0-20190529130038-5219a1e1c5f8 // indirect
	golang.org/x/text v0.3.2 // indirect
)
```

Go 语言项目自身所依赖的包在 go.mod 中都有对应的信息，而原本这些依赖包是缓存在 vendor 目录下的。 



### 总结

总的来说，这三次演进（1.4->1.6->1.13）主要体现在简化结构布局，以及优化包依赖管理方面，起到了改善 Go 开发体验的作用。

可以说，Go 创世项目的源码布局以及演化对 Go 社区项目的布局具 有重要的启发意义，以至于在多年的 Go 社区实践后，Go 社区逐渐形成了公认的 Go 项目 的典型结构布局。



## 现在的 Go 项目的典型结构布局是怎样的？ 

一个 Go 项目通常分为可执行程序项目和库项目，现在来分析一下这两类 Go 项目 的典型结构布局分别是怎样的。 

### Go 可执行程序项目的典型结构布局

可执行程序项目是以构建可执行程序为目的的项目，Go 社区针对这类 Go 项目所形成的典 型结构布局是这样的：

```sh
$tree -F exe-layout
exe-layout
├── cmd/
│ ├── app1/
│ │ └── main.go
│ └── app2/
│ 	└── main.go
├── go.mod
├── go.sum
├── internal/
│ ├── pkga/
│ │ └── pkg_a.go
│ └── pkgb/
│ 	└── pkg_b.go
├── pkg1/
│ └── pkg1.go
├── pkg2/
│ └── pkg2.go
└── vendor/
```

这样的一个 Go 项目典型布局就是“脱胎”于 Go 创世项目的最新结构布局，解释一下这里面的几个要点。 

#### cmd 目录

cmd 目录就是存放项目要编译构建的可执行文件对应的 **main 包的源文件**。

如果项目中有多个可执行文件需要构建，每个可执行 文件的 main 包单独放在一个子目录中，比如图中的 app1、app2，cmd 目录下的各 app 的 main 包将整个项目的依赖连接在一起。 

而且通常来说，main 包应该很简洁。

在 main 包中会做一些命令行参数解析、资源初始化、日志设施初始化、数据库连接初始化等工作，之后就会将程序的执行权限交给更高级的执行控制对象。

另外，也有一些 Go 项目将 cmd 这个名字**改为 app 或其他名字**，但它的功能其实并没有变。 

#### pkgN 目录

pkgN 目录，这是一个存放项目自身要使用、同样也是可执行文件对应 **main 包所要依赖的库文件**，同时这些目录下的包还可以被外部项目引用。

#### go.mod 和 go.sum

然后是 go.mod 和 go.sum，它们是 **Go 语言包依赖管理使用的配置文件**。

Go 1.11 版本引入了 Go Module 构建机制，建议所有新项目都基于 Go Module 来进行包依赖管理，因为这是目前 Go 官方推荐的标准构建模式。 

对于还没有使用 Go Module 进行包依赖管理的遗留项目，比如之前采用 dep、glide 等作为包依赖管理工具的，建议尽快迁移到 Go Module 模式。

Go 命令支持直接将 dep 的 `Gopkg.toml/Gopkg.lock` 或 glide 的 `glide.yaml/glide.lock` 转换为 go.mod。 

#### vendor 目录

vendor 目录 是 Go 1.5 版本引入的用于在项目本地缓存特定版本依赖包的机制，在 Go Modules 机制引入前，基于 vendor 可以实现可重现构建，保证基于同一源码构建出的可执行程序是等价的。 

不过呢，这里将 vendor 目录视为一个可选目录。原因在于，Go Module 本身就支持 可再现构建，而无需使用 vendor。 

当然 Go Module 机制也保留了 vendor 目录（通过 go mod vendor 可以生成 vendor 下的依赖包，通过 `go build -mod=vendor` 可以实现 基于 vendor 的构建）。

一般仅保留项目根目录下的 vendor 目录，否则会造成不必要的依赖选择的复杂性。

当然了，有些开发者喜欢借助一些第三方的构建工具辅助构建，比如：make、bazel 等。 可以将这类外部辅助构建工具涉及的诸多脚本文件（比如 Makefile）放置在项目的顶层目录下，就像 Go 创世项目中的 all.bash 那样。 

另外，这里只要说明一下的是，Go 1.11 引入的 module 是一组同属于一个版本管理单元的包的集合。

并且 Go 支持在一个项目 / 仓库中**存在多个 module**，但这种管理方式可能要比一定比例的代码重复引入更多的复杂性。 

因此，如果项目结构中存在版本管理的“分 歧”，比如：app1 和 app2 的发布版本并不总是同步的，那么建议**将项目拆分为多个项目（仓库）**，每个项目单独作为一个 module 进行单独的版本管理和演进。 

当然如果非要在一个代码仓库中存放多个 module，那么**新版 Go 命令**也提供了很好的支持。比如下面代码仓库 multi-modules 下面有三个 module：mainmodule、 module1 和 module2：

```sh
$tree multi-modules
multi-modules
├── go.mod // mainmodule
├── module1
│ └── go.mod // module1
└── module2
	└── go.mod // module2
```

可以通过 **git tag 名字**来区分不同 module 的版本。其中 vX.Y.Z 形式的 tag 名字用于 代码仓库下的 mainmodule；而 module1/vX.Y.Z 形式的 tag 名字用于指示 module1 的版本；同理，module2/vX.Y.Z 形式的 tag 名字用于指示 module2 版本。 

#### 简化项目布局

如果 Go 可执行程序项目有一个且只有一个可执行程序要构建，那就比较好办了，可以将上面项目布局进行简化：

```sh
$tree -F -L 1 single-exe-layout
single-exe-layout
├── go.mod
├── internal/
├── main.go
├── pkg1/
├── pkg2/
└── vendor/
```

可以看到，删除了 cmd 目录，将唯一的可执行程序的 main 包就放置在项目根目录 下，而其他布局元素的功用不变。 



### Go 库项目的典型结构布局

Go 库项目仅对外暴露 Go 包，这类项目的典型布局形式是这样的：

```sh
$tree -F lib-layout
lib-layout
├── go.mod
├── internal/
│ ├── pkga/
│ │ └── pkg_a.go
│ └── pkgb/
│ 	└── pkg_b.go
├── pkg1/
│ └── pkg1.go
└── pkg2/
	└── pkg2.go
```

#### go.mod

库类型项目相比于 Go 可执行程序项目的布局要简单一些。因为这类项目不需要构建可执行程序，所以去除了 cmd 目录。 

而且，vendor 也不再是可选目录了。对于库类型项目而言，并**不推荐在项目中放置 vendor 目录**去缓存库自身的第三方依赖，库项目**仅通过 go.mod 文件**明确表述出该项目依赖的 module 或包以及版本要求就可以了。 

#### internal 目录

Go 库项目的初衷是为了对外部（开源或组织内部公开）暴露 API，对于仅限项目内部使用而不想暴露到外部的包，可以放在项目顶层的 internal 目录下面。

当然 internal 也可以有多个并存在于项目结构中的任一目录层级中，关键是项目结构设计人员要明确各级 internal 包的应用层次和范围。 

#### 简化项目布局

对于有一个且仅有一个包的 Go 库项目来说，也可以将上面的布局做进一步简化，简 化的布局如下所示：

```sh
$tree -L 1 -F single-pkg-lib-layout
single-pkg-lib-layout
├── feature1.go
├── feature2.go
├── go.mod
└── internal/
```

简化后，将这唯一包的所有源文件放置在项目的顶层目录下（比如上面的 feature1.go 和 feature2.go），其他布局元素位置和功用不变。 



## 注意早期 Go 可执行程序项目的典型布局

很多早期接纳 Go 语言的开发者所建立的 Go 可执行程序项目，深受 Go 创世项目 **1.4 版本之前的布局影响**，这些项目将所有可暴露到外面的 Go 包聚合在 pkg 目录下，就像前面 Go 1.3 版本中的布局那样，它们的典型布局结构是这样的：

```sh
$tree -L 3 -F early-project-layout
early-project-layout
└── exe-layout/
  ├── cmd/
  │ ├── app1/
  │ └── app2/
  ├── go.mod
  ├── internal/
  │ ├── pkga/
  │ └── pkgb/
  ├── pkg/
  │ ├── pkg1/
  │ └── pkg2/
  └── vendor/
```

原本放在项目顶层目录下的 pkg1 和 pkg2 公共包被统一聚合到 pkg 目录下 了。而且，这种早期 Go 可执行程序项目的典型布局在 Go 社区内部也不乏受众，很多新 建的 Go 项目依然采用这样的项目布局。 

所以，当看到这样的布局也不要奇怪，应该明确在这样的布局下 pkg 目录所起到的“聚类”的作用。

建议在创建新的 Go 项目时，优先采用前面的标准项目布局。









