# Go Concurrent

> Go 并发编程

## 上下文 Context

上下文 `context.Context` 是Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。

上下文与 Goroutine 有比较密切的关系，是 Go 语言中独特的设计，在其他编程语言中很少见到类似的概念。

`context.Context` 是 Go 语言在 1.7 版本中引入标准库的接口，该接口定义了四个需要实现的方法，其中包括：

1. `Deadline` — 返回 `context.Context` 被取消的时间，也就是完成工作的截止日期；
2. `Done` — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 `Done` 方法会返回同一个 Channel；
3. Err— 返回`context.Context`结束的原因，它只会在Done 方法对应的 Channel 关闭时返回非空的值；
   1. 如果 `context.Context` 被取消，会返回 `Canceled` 错误；
   2. 如果 `context.Context` 超时，会返回 `DeadlineExceeded` 错误；
4. `Value` — 从 `context.Context` 中获取键对应的值，对于同一个上下文来说，多次调用 `Value` 并传入相同的 `Key` 会返回相同的结果，该方法可以用来传递请求特定的数据；

```go
// github.com/golang/go/src/context/context.go
type Context interface {
	Deadline() (deadline time.Time, ok bool)// Deadline returns the time when work done on behalf of this context should be canceled.
	Done() <-chan struct{} // Done returns a channel that's closed when work done on behalf of this context should be canceled.
  // If Done is not yet closed, Err returns nil.
	// If Done is closed, Err returns a non-nil error explaining why:
	// Canceled if the context was canceled
	// or DeadlineExceeded if the context's deadline passed.
	Err() error
  // Value returns the value associated with this context for key, or nil
	// if no value is associated with key. Successive calls to Value with
	// the same key returns the same result.
	Value(key interface{}) interface{}
}
```

`context` 包中提供的 `context.Background`、`context.TODO`、`context.WithDeadline` 和 `context.WithValue` 函数会返回实现该接口的私有结构体。

### 设计原理

在 Goroutine 构成的树形结构中对信号进行同步以减少计算资源的浪费是 `context.Context` 的最大作用。

Go 服务的每一个请求都是通过单独的 Goroutine 处理的，HTTP/RPC 请求的处理器会启动新的 Goroutine 访问数据库和其他服务。

如下图所示，可能会创建多个 Goroutine 来处理一次请求，而 `context.Context` 的作用是**在不同 Goroutine 之间**同步请求特定数据、取消信号以及处理请求的截止日期。

![golang-context-usage](go_concurrent.assets/golang-context-usage.png)

**Context 与 Goroutine 树**

每一个 `context.Context` 都会从最顶层的 Goroutine 一层一层传递到最下层。`context.Context` 可以在上层 Goroutine 执行出现错误时，将信号及时同步给下层。

![golang-without-context](go_concurrent.assets/golang-without-context.png)

**不使用 Context 同步信号**

如上图所示，当最上层的 Goroutine 因为某些原因执行失败时，下层的 Goroutine 由于没有接收到这个信号所以会继续工作；但是当正确地使用 `context.Context` 时，就可以在下层及时停掉无用的工作以减少额外资源的消耗：

![golang-with-context](go_concurrent.assets/golang-with-context.png)

**使用 Context 同步信号**

可以通过一个代码片段了解 `context.Context` 是如何**对信号进行同步**的。在这段代码中，创建了一个过期时间为 1s 的上下文，并向上下文传入 `handle` 函数，该方法会使用 500ms 的时间处理传入的请求：

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}
```

因为**过期时间大于处理时间**，所以有足够的时间处理该请求，运行上述代码会打印出下面的内容：

```go
$ go run context.go
process request with 500ms
main context deadline exceeded
```

`handle` 函数没有进入超时的 `select` 分支，但是 `main` 函数的 `select` 却会等待 `context.Context` 超时并打印出 `main context deadline exceeded`。

如果将处理请求时间增加至 1500ms，整个程序都会因为上下文的过期而被中止，：

```go
$ go run context.go
main context deadline exceeded
handle context deadline exceeded
```

相信这两个例子能够理解 `context.Context` 的使用方法和设计原理 — 多个 Goroutine 同时订阅 `ctx.Done()` 管道中的消息，一旦接收到取消信号就立刻停止当前正在执行的工作。

### 默认上下文

`context` 包中最常用的方法还是 `context.Background`、`context.TODO`，这两个方法都会返回预先初始化好的私有变量 `background` 和 `todo`，它们会**在同一个 Go 程序中被复用**：

```go
// github.com/golang/go/src/context/context.go
var (
	background = new(emptyCtx)
	todo       = new(emptyCtx)
)

func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

这两个私有变量都是通过 `new(emptyCtx)` 语句初始化的，它们是指向私有结构体 `context.emptyCtx` 的指针，这是最简单、最常用的上下文类型：

```go
// github.com/golang/go/src/context/context.go
// An emptyCtx is never canceled, has no values, and has no deadline. It is not
// struct{}, since vars of this type must have distinct addresses.
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```

从上述代码中，不难发现 `context.emptyCtx` 通过空方法实现了 `context.Context` 接口中的所有方法，它没有任何功能。

![golang-context-hierarchy](go_concurrent.assets/golang-context-hierarchy.png)

**Context 层级关系**

从源代码来看，`context.Background` 和 `context.TODO` 也只是互为别名，没有太大的差别，只是在使用和语义上稍有不同：

- `context.Background` 是**上下文的默认值**，**所有其他的上下文都应该从它衍生出来**；
- `context.TODO` 应该**仅在不确定应该使用哪种上下文时使用**；

在多数情况下，如果当前函数没有上下文作为入参，都会使用 `context.Background` 作为起始的上下文向下传递。

### 取消信号

#### context.WithCancel

`context.WithCancel` 函数能够从 `context.Context` 中衍生出一个新的子上下文并返回用于取消该上下文的函数。

一旦执行返回的取消函数，**当前上下文以及它的子上下文都会被取消**，所有的 Goroutine 都会同步收到这一取消信号。

![golang-parent-cancel-context](go_concurrent.assets/2020-01-20-15795072700927-golang-parent-cancel-context.png)

**Context 子树的取消**

直接从 `context.WithCancel` 函数的实现来看它到底做了什么：

```go
// github.com/golang/go/src/context/context.go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
	c := newCancelCtx(parent)
	propagateCancel(parent, &c)
	return &c, func() { c.cancel(true, Canceled) }
}

// newCancelCtx returns an initialized cancelCtx.
func newCancelCtx(parent Context) cancelCtx {
	return cancelCtx{Context: parent}
}
```

- `context.newCancelCtx` 将传入的上下文包装成私有结构体 `context.cancelCtx`；
- `context.propagateCancel` 会构建父子上下文之间的关联，当父上下文被取消时，子上下文也会被取消：

```go
// github.com/golang/go/src/context/context.go
// propagateCancel arranges for child to be canceled when parent is.
func propagateCancel(parent Context, child canceler) {
	done := parent.Done()
	if done == nil {
    // parent is never canceled
		return // 父上下文不会触发取消信号
	}
	select {
	case <-done:
    // parent is already canceled
		child.cancel(false, parent.Err()) // 父上下文已经被取消
		return
	default:
	}

	if p, ok := parentCancelCtx(parent); ok {
		p.mu.Lock()
		if p.err != nil {
			child.cancel(false, p.err)
		} else {
			p.children[child] = struct{}{} // child 加入 parent 的 children 列表
		}
		p.mu.Unlock()
	} else {  // 开发者自定义的父上下文
		go func() {
			select {
			case <-parent.Done():
				child.cancel(false, parent.Err())
			case <-child.Done():
			}
		}()
	}
}
```

上述函数总共与父上下文相关的三种不同的情况：

1. 当 `parent.Done() == nil`，也就是 `parent` 不会触发取消事件时，当前函数会直接返回；

2. 当 child 的继承链包含可以取消的上下文时，会判断 parent 是否已经触发了取消信号；

   - 如果已经被取消，`child` 会立刻被取消；
   - 如果没有被取消，`child` 会被加入 `parent` 的 `children` 列表中，等待 `parent` 释放取消信号；

3. 当父上下文是开发者自定义的类型、实现了`context.Context` 接口并在Done() 方法中返回了非空的管道时；

   1. 运行一个新的 Goroutine 同时监听 `parent.Done()` 和 `child.Done()` 两个 Channel；
   2. 在 `parent.Done()` 关闭时调用 `child.cancel` 取消子上下文；

`context.propagateCancel` 的作用是在 `parent` 和 `child` 之间同步取消和结束的信号，保证在 `parent` 被取消时，`child` 也会收到对应的信号，不会出现状态不一致的情况。

`context.cancelCtx` 实现的几个接口方法也没有太多值得分析的地方，该结构体最重要的方法是 `context.cancelCtx.cancel`，该方法会关闭上下文中的 Channel 并向所有的子上下文同步取消信号：

```go
// github.com/golang/go/src/context/context.go
// cancel closes c.done, cancels each of c's children, and, if
// removeFromParent is true, removes c from its parent's children.
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
	c.mu.Lock()
	if c.err != nil {
		c.mu.Unlock()
		return
	}
	c.err = err
	if c.done == nil {
		c.done = closedchan
	} else {
		close(c.done)
	}
	for child := range c.children {
		child.cancel(false, err)
	}
	c.children = nil
	c.mu.Unlock()

	if removeFromParent {
		removeChild(c.Context, c)
	}
}
```

#### context.WithDeadline 和 context.WithTimeout

除了 `context.WithCancel` 之外，`context` 包中的另外两个函数 `context.WithDeadline` 和 `context.WithTimeout` 也都能创建可以被取消的计时器上下文 `context.timerCtx`：

```go
// github.com/golang/go/src/context/context.go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}

func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
	if cur, ok := parent.Deadline(); ok && cur.Before(d) {
    // The current deadline is already sooner than the new one.
		return WithCancel(parent)
	}
	c := &timerCtx{
		cancelCtx: newCancelCtx(parent),
		deadline:  d,
	}
	propagateCancel(parent, c)
	dur := time.Until(d)
	if dur <= 0 {
		c.cancel(true, DeadlineExceeded) // 已经过了截止日期
		return c, func() { c.cancel(false, Canceled) }
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.err == nil {
		c.timer = time.AfterFunc(dur, func() {  // 创建定时器
			c.cancel(true, DeadlineExceeded)
		})
	}
	return c, func() { c.cancel(true, Canceled) }
}
```

`context.WithDeadline` 在创建 `context.timerCtx` 的过程中判断了父上下文的截止日期与当前日期，并通过 `time.AfterFunc` 创建定时器，当时间超过了截止日期后会调用 `context.timerCtx.cancel` 同步取消信号。

`context.timerCtx` 内部不仅通过嵌入 `context.cancelCtx` 结构体继承了相关的变量和方法，还通过持有的定时器 `timer` 和截止时间 `deadline` 实现了定时取消的功能：

```go
// github.com/golang/go/src/context/context.go
type timerCtx struct {
	cancelCtx
	timer *time.Timer // Under cancelCtx.mu.

	deadline time.Time
}

func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
	return c.deadline, true
}

func (c *timerCtx) cancel(removeFromParent bool, err error) {
	c.cancelCtx.cancel(false, err)
	if removeFromParent {
		removeChild(c.cancelCtx.Context, c)
	}
	c.mu.Lock()
	if c.timer != nil {
		c.timer.Stop()
		c.timer = nil
	}
	c.mu.Unlock()
}
```

`context.timerCtx.cancel` 方法不仅调用了 `context.cancelCtx.cancel`，还会停止持有的定时器减少不必要的资源浪费。

### 传值方法

在最后需要了解如何使用上下文传值，`context` 包中的 `context.WithValue` 能从父上下文中创建一个子上下文，传值的子上下文使用 `context.valueCtx` 类型：

```go
// github.com/golang/go/src/context/context.go
func WithValue(parent Context, key, val interface{}) Context {
	if key == nil {
		panic("nil key")
	}
	if !reflectlite.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}
	return &valueCtx{parent, key, val}
}
```

`context.valueCtx` 结构体会将除了 `Value` 之外的 `Err`、`Deadline` 等方法代理到父上下文中，它只会响应 `context.valueCtx.Value` 方法，该方法的实现也很简单：

```go
// github.com/golang/go/src/context/context.go
type valueCtx struct {
	Context
	key, val interface{}
}

func (c *valueCtx) Value(key interface{}) interface{} {
	if c.key == key {
		return c.val
	}
	return c.Context.Value(key)
}
```

如果 `context.valueCtx` 中存储的键值对与 `context.valueCtx.Value` 方法中传入的参数不匹配，就会从父上下文中查找该键对应的值，直到某个父上下文中返回 `nil` 或者查找到对应的值。

### 小结

Go 语言中的 `context.Context` 的主要作用还是在多个 Goroutine 组成的树中同步取消信号以减少对资源的消耗和占用，虽然它也有传值的功能，但是这个功能还是很少用到。

在真正使用传值的功能时也应该非常谨慎，使用 `context.Context` 传递请求的所有参数是一种非常差的设计，比较常见的使用场景是传递请求对应用户的认证令牌以及用于进行分布式追踪的请求 ID。

### 参考

- [Package context · Golang](https://golang.org/pkg/context/)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)
- [Using context cancellation in Go](https://www.sohamkamani.com/blog/golang/2018-06-17-golang-using-context-cancellation/)
- proposal: context: new package for standard library #14660 https://github.com/golang/go/issues/14660 















