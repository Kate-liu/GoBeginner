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



## 同步原语与锁

Go 语言作为一个原生支持用户态进程（Goroutine）的语言，当提到并发编程、多线程编程时，往往都离不开锁这一概念。

**锁**是一种并发编程中的同步原语（Synchronization Primitives），它能保证多个 Goroutine 在访问同一片内存时不会出现竞争条件（Race condition）等问题。

Go 语言中常见的同步原语 `sync.Mutex`、`sync.RWMutex`、`sync.WaitGroup`、`sync.Once` 和 `sync.Cond` 以及扩展原语 `golang/sync/errgroup.Group`、`golang/sync/semaphore.Weighted` 和 `golang/sync/singleflight.Group` 的实现原理，同时也会涉及互斥锁、信号量等并发编程中的常见概念。

## 基本原语

Go 语言在 `sync` 包中提供了用于同步的一些基本原语，包括常见的 `sync.Mutex`、`sync.RWMutex`、`sync.WaitGroup`、`sync.Once` 和 `sync.Cond`：

![golang-basic-sync-primitives](go_concurrent.assets/2020-01-23-15797104327981-golang-basic-sync-primitives.png)

**基本同步原语**

这些基本原语提供了较为基础的同步功能，但是它们是一种**相对原始的同步机制**，在多数情况下，都应该使用抽象层级更高的 Channel 实现同步。

### Mutex

Go 语言的 `sync.Mutex` 由两个字段 `state` 和 `sema` 组成。其中 `state` 表示当前互斥锁的状态，而 `sema` 是用于控制锁状态的信号量。

```go
// github.com/golang/go/src/sync/mutex.go
type Mutex struct {
	state int32
	sema  uint32
}
```

上述两个字段加起来只**占 8 字节空间**的结构体表示了 Go 语言中的互斥锁。

#### 状态

互斥锁的状态比较复杂，如下图所示，最低三位分别表示 `mutexLocked`、`mutexWoken` 和 `mutexStarving`，剩下的位置用来表示当前有多少个 Goroutine 在等待互斥锁的释放：

![golang-mutex-state](go_concurrent.assets/2020-01-23-15797104328010-golang-mutex-state.png)

**互斥锁的状态**

在默认情况下，互斥锁的所有状态位都是 0，`int32` 中的不同位分别表示了不同的状态：

- `mutexLocked` — 表示互斥锁的锁定状态；
- `mutexWoken` — 表示从正常模式被唤醒；
- `mutexStarving` — 当前的互斥锁进入饥饿状态；
- `waitersCount` — 当前互斥锁上等待的 Goroutine 个数；

```go
// github.com/golang/go/src/sync/mutex.go
const (
   mutexLocked = 1 << iota // mutex is locked
   mutexWoken
   mutexStarving
   mutexWaiterShift = iota
   starvationThresholdNs = 1e6
)
```

#### 正常模式和饥饿模式

`sync.Mutex` 有两种模式 — 正常模式和饥饿模式。需要在这里先了解正常模式和饥饿模式都是什么以及它们有什么样的关系。

在正常模式下，锁的等待者会按照**先进先出的顺序获取锁**。

但是刚被唤起的 Goroutine 与新创建的 Goroutine 竞争时，大概率会获取不到锁，为了减少这种情况的出现，**一旦 Goroutine 超过 1ms 没有获取到锁**，它就会将当前互斥锁切换饥饿模式，防止部分 Goroutine 被『饿死』。

![golang-mutex-mode](go_concurrent.assets/2020-01-23-15797104328020-golang-mutex-mode.png)

**互斥锁的正常模式与饥饿模式**

饥饿模式是 Go 语言在 1.9 版本中通过提交 [sync: make Mutex more fair](https://github.com/golang/go/commit/0556e26273f704db73df9e7c4c3d2e8434dec7be) 引入的优化，引入的目的是**保证互斥锁的公平性**。

在饥饿模式中，互斥锁会直接交给等待队列最前面的 Goroutine。新的 Goroutine 在该状态下不能获取锁、也不会进入自旋状态，它们只会在队列的末尾等待。

如果一个 Goroutine 获得了互斥锁并且它在队列的末尾或者它等待的时间少于 1ms，那么当前的互斥锁就会**切换回正常模式**。

与饥饿模式相比，正常模式下的互斥锁能够提供更好地性能，饥饿模式能避免 Goroutine 由于陷入等待无法获取锁而造成的高尾延时。

#### 加锁

互斥锁的加锁和解锁过程，它们分别使用 `sync.Mutex.Lock` 和 `sync.Mutex.Unlock`方法。

互斥锁的加锁是靠 `sync.Mutex.Lock` 完成的，最新的 Go 语言源代码中已经将 `sync.Mutex.Lock` 方法进行了简化，方法的主干只保留最常见、简单的情况 — 当锁的状态是 0 时，将 `mutexLocked` 的位 置为 1：

```go
// github.com/golang/go/src/sync/mutex.go
func (m *Mutex) Lock() {
  // Fast path: grab unlocked mutex.
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
  // Slow path (outlined so that the fast path can be inlined)
	m.lockSlow()
}
```

如果互斥锁的状态不是 0 时就会调用 `sync.Mutex.lockSlow` 尝试通过自旋（Spinnig）等方式等待锁的释放，该方法的主体是一个非常大 for 循环，这里将它分成几个部分介绍获取锁的过程：

1. 判断当前 Goroutine 能否进入自旋；
2. 通过自旋等待互斥锁的释放；
3. 计算互斥锁的最新状态；
4. 更新互斥锁的状态并获取锁；

先来介绍互斥锁是如何**判断当前 Goroutine 能否进入自旋**等互斥锁的释放：

```go
// github.com/golang/go/src/sync/mutex.go
func (m *Mutex) lockSlow() {
	var waitStartTime int64
	starving := false
	awoke := false
	iter := 0
	old := m.state
	for {
		if old&(mutexLocked|mutexStarving) == mutexLocked && runtime_canSpin(iter) {
			if !awoke && old&mutexWoken == 0 && old>>mutexWaiterShift != 0 &&
				atomic.CompareAndSwapInt32(&m.state, old, old|mutexWoken) {
				awoke = true
			}
			runtime_doSpin()
			iter++
			old = m.state
			continue
		}
```

**自旋**是一种多线程同步机制，当前的进程在进入自旋的过程中会一直保持 CPU 的占用，持续检查某个条件是否为真。在多核的 CPU 上，自旋可以避免 Goroutine 的切换，使用恰当会对性能带来很大的增益，但是使用的不恰当就会拖慢整个程序，所以 **Goroutine 进入自旋的条件**非常苛刻：

1. 互斥锁只有在普通模式才能进入自旋；
2. `runtime.sync_runtime_canSpin` 需要返回true ：
   1. 运行在多 CPU 的机器上；
   2. 当前 Goroutine 为了获取该锁进入自旋的次数小于四次；
   3. 当前机器上至少存在一个正在运行的处理器 P 并且处理的运行队列为空；

一旦当前 Goroutine 能够进入自旋就会调用`runtime.sync_runtime_doSpin` 和 `runtime.procyield` 并执行 30 次的 `PAUSE` 指令，该指令只会占用 CPU 并消耗 CPU 时间：

```go
// github.com/golang/go/src/runtime/proc.go
func sync_runtime_doSpin() {
	procyield(active_spin_cnt)
}

// github.com/golang/go/src/runtime/lock_sema.go
const (
	locked uintptr = 1

	active_spin     = 4
	active_spin_cnt = 30  // 执行 30 次 PAUSE 指令
	passive_spin    = 1
)

// github.com/golang/go/src/runtime/asm_386.s
TEXT runtime·procyield(SB),NOSPLIT,$0-0
	MOVL	cycles+0(FP), AX
again:
	PAUSE
	SUBL	$1, AX
	JNZ	again
	RET
```

处理了自旋相关的特殊逻辑之后，互斥锁会**根据上下文计算当前互斥锁最新的状态**。

几个不同的条件分别会更新 `state` 字段中存储的不同信息 — `mutexLocked`、`mutexStarving`、`mutexWoken` 和 `mutexWaiterShift`：

```go
// github.com/golang/go/src/sync/mutex.go
		new := old
		if old&mutexStarving == 0 {
			new |= mutexLocked
		}
		if old&(mutexLocked|mutexStarving) != 0 {
			new += 1 << mutexWaiterShift
		}
		if starving && old&mutexLocked != 0 {
			new |= mutexStarving
		}
		if awoke {
			new &^= mutexWoken
		}
```

计算了新的互斥锁状态之后，会使用 CAS 函数 `sync/atomic.CompareAndSwapInt32` 更新状态：

```go
// github.com/golang/go/src/sync/mutex.go
		if atomic.CompareAndSwapInt32(&m.state, old, new) {
			if old&(mutexLocked|mutexStarving) == 0 {
        // locked the mutex with CAS
				break // 通过 CAS 函数获取了锁
			}
			...
      // github.com/golang/go/src/runtime/sema.go
			runtime_SemacquireMutex(&m.sema, queueLifo, 1)
			starving = starving || runtime_nanotime()-waitStartTime > starvationThresholdNs
			old = m.state
			if old&mutexStarving != 0 { // 饥饿模式
				delta := int32(mutexLocked - 1<<mutexWaiterShift)
				if !starving || old>>mutexWaiterShift == 1 {
					delta -= mutexStarving
				}
				atomic.AddInt32(&m.state, delta)
				break
			}
			awoke = true  // 正常模式：设置唤醒标记
			iter = 0
		} else {
			old = m.state
		}
	}
}
```

如果没有通过 CAS 获得锁，会调用 `runtime.sync_runtime_SemacquireMutex` 通过信号量保证资源不会被两个 Goroutine 获取。

```go
//go:linkname sync_runtime_SemacquireMutex sync.runtime_SemacquireMutex
func sync_runtime_SemacquireMutex(addr *uint32, lifo bool, skipframes int) {
   semacquire1(addr, lifo, semaBlockProfile|semaMutexProfile, skipframes)
}

func semacquire1(addr *uint32, lifo bool, profile semaProfileFlags, skipframes int) {
	...
	for {
		...
		// Any semrelease after the cansemacquire knows we're waiting
		// (we set nwait above), so go to sleep.
		root.queue(addr, s, lifo)  // 陷入休眠
		goparkunlock(&root.lock, waitReasonSemacquire, traceEvGoBlockSync, 4+skipframes)
		if s.ticket != 0 || cansemacquire(addr) {
			break
		}
	}
	if s.releasetime > 0 {
		blockevent(s.releasetime-t0, 3+skipframes)
	}
	releaseSudog(s)
}
```

`runtime.sync_runtime_SemacquireMutex` 会在方法中不断尝试获取锁并陷入休眠等待信号量的释放，一旦当前 Goroutine 可以获取信号量，它就会立刻返回，`sync.Mutex.Lock` 的剩余代码也会继续执行。

- 在正常模式下，这段代码会设置唤醒和饥饿标记、重置迭代次数并重新执行获取锁的循环；
- 在饥饿模式下，当前 Goroutine 会获得互斥锁，如果等待队列中只存在当前 Goroutine，互斥锁还会从饥饿模式中退出；

#### 解锁

互斥锁的解锁过程 `sync.Mutex.Unlock`与加锁过程相比就很简单，该过程会先使用 `sync/atomic.AddInt32` 函数快速解锁，这时会发生下面的两种情况：

- 如果该函数返回的新状态等于 0，当前 Goroutine 就成功解锁了互斥锁；
- 如果该函数返回的新状态不等于 0，这段代码会调用 `sync.Mutex.unlockSlow` 开始慢速解锁：

```go
// github.com/golang/go/src/sync/mutex.go
func (m *Mutex) Unlock() {
  // Fast path: drop lock bit.
	new := atomic.AddInt32(&m.state, -mutexLocked)
	if new != 0 {
    // Outlined slow path to allow inlining the fast path.
		// To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
		m.unlockSlow(new)
	}
}
```

`sync.Mutex.unlockSlow` 会先校验锁状态的合法性 — 如果当前互斥锁已经被解锁过了会直接抛出异常 “sync: unlock of unlocked mutex” 中止当前程序。

在正常情况下会根据当前互斥锁的状态，分别处理正常模式和饥饿模式下的互斥锁：

```go
// github.com/golang/go/src/sync/mutex.go
func (m *Mutex) unlockSlow(new int32) {
	if (new+mutexLocked)&mutexLocked == 0 {
		throw("sync: unlock of unlocked mutex")
	}
	if new&mutexStarving == 0 { // 正常模式
		old := new
		for {
			if old>>mutexWaiterShift == 0 || old&(mutexLocked|mutexWoken|mutexStarving) != 0 {
				return
			}
			new = (old - 1<<mutexWaiterShift) | mutexWoken
			if atomic.CompareAndSwapInt32(&m.state, old, new) {
				runtime_Semrelease(&m.sema, false, 1)
				return
			}
			old = m.state
		}
	} else { // 饥饿模式
		runtime_Semrelease(&m.sema, true, 1)
	}
}
```

- 在正常模式下，上述代码会使用如下所示的处理过程：
  - 如果互斥锁不存在等待者或者互斥锁的 `mutexLocked`、`mutexStarving`、`mutexWoken` 状态不都为 0，那么当前方法可以直接返回，不需要唤醒其他等待者；
  - 如果互斥锁存在等待者，会通过 `sync.runtime_Semrelease` 唤醒等待者并移交锁的所有权；
- 在饥饿模式下，上述代码会直接调用 `sync.runtime_Semrelease` 将当前锁交给下一个正在尝试获取锁的等待者，等待者被唤醒后会得到锁，在这时互斥锁还不会退出饥饿状态；

#### 小结

已经从多个方面分析了互斥锁 `sync.Mutex` 的实现原理，这里从加锁和解锁两个方面总结注意事项。

互斥锁的加锁过程比较复杂，它涉及自旋、信号量以及调度等概念：

- 如果互斥锁处于初始化状态，会通过置位 `mutexLocked` 加锁；
- 如果互斥锁处于 `mutexLocked` 状态并且在普通模式下工作，会进入自旋，执行 30 次 `PAUSE` 指令消耗 CPU 时间等待锁的释放；
- 如果当前 Goroutine 等待锁的时间超过了 1ms，互斥锁就会切换到饥饿模式；
- 互斥锁在正常情况下会通过 `runtime.sync_runtime_SemacquireMutex` 将尝试获取锁的 Goroutine 切换至休眠状态，等待锁的持有者唤醒；
- 如果当前 Goroutine 是互斥锁上的最后一个等待的协程或者等待的时间小于 1ms，那么它会将互斥锁切换回正常模式；

互斥锁的解锁过程与之相比就比较简单，其代码行数不多、逻辑清晰，也比较容易理解：

- 当互斥锁已经被解锁时，调用 `sync.Mutex.Unlock`会直接抛出异常；
- 当互斥锁处于饥饿模式时，将锁的所有权交给队列中的下一个等待者，等待者会负责设置 `mutexLocked` 标志位；
- 当互斥锁处于普通模式时，如果没有 Goroutine 等待锁的释放或者已经有被唤醒的 Goroutine 获得了锁，会直接返回；在其他情况下会通过 `sync.runtime_Semrelease` 唤醒对应的 Goroutine；



### RWMutex

读写互斥锁 `sync.RWMutex` 是细粒度的互斥锁，它不限制资源的并发读，但是读写操作和写写操作无法并行执行。

|      |  读  |  写  |
| :--: | :--: | :--: |
|  读  |  Y   |  N   |
|  写  |  N   |  N   |

**RWMutex 的读写并发**

常见服务的资源读写比例会非常高，因为大多数的读请求之间不会相互影响，所以可以**分离读写操作**，以此来提高服务的性能。

#### 结构体

`sync.RWMutex` 中总共包含以下 5 个字段：

```go
// github.com/golang/go/src/sync/rwmutex.go
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```

- `w` — 复用互斥锁提供的能力；
- `writerSem` 和 `readerSem` — 分别用于`写等待读`和`读等待写`：
- `readerCount` 存储了当前正在执行的读操作数量；
- `readerWait` 表示当写操作被阻塞时等待的读操作个数；

会依次分析获取写锁和读锁的实现原理，其中：

- 写操作使用 `sync.RWMutex.Lock` 和`sync.RWMutex.Unlock` 方法；
- 读操作使用 `sync.RWMutex.RLock` 和 `sync.RWMutex.RUnlock`方法；

#### 写锁

当资源的使用者想要**获取写锁**时，需要调用 `sync.RWMutex.Lock` 方法：

```go
// github.com/golang/go/src/sync/rwmutex.go
// Lock locks rw for writing.
// If the lock is already locked for reading or writing,
// Lock blocks until the lock is available.
func (rw *RWMutex) Lock() {
  // First, resolve competition with other writers.
	rw.w.Lock()
  // Announce to readers there is a pending writer.
	r := atomic.AddInt32(&rw.readerCount, -rwmutexMaxReaders) + rwmutexMaxReaders
  // Wait for active readers.
	if r != 0 && atomic.AddInt32(&rw.readerWait, r) != 0 {
		runtime_SemacquireMutex(&rw.writerSem, false, 0)
	}
}
```

1. 调用结构体持有的`sync.Mutex` 结构体的`sync.Mutex.Lock`阻塞后续的写操作；

   - 因为互斥锁已经被获取，其他 Goroutine 在获取写锁时会进入自旋或者休眠；

2. 调用 `sync/atomic.AddInt32` 函数阻塞后续的读操作：

3. 如果仍然有其他 Goroutine 持有互斥锁的读锁，该 Goroutine 会调用 `runtime.sync_runtime_SemacquireMutex` 进入休眠状态等待所有读锁所有者执行结束后释放 `writerSem` 信号量将当前协程唤醒；

**写锁的释放**会调用`sync.RWMutex.Unlock`：

```go
// github.com/golang/go/src/sync/rwmutex.go
// Unlock unlocks rw for writing. It is a run-time error if rw is
// not locked for writing on entry to Unlock.
func (rw *RWMutex) Unlock() {
  // Announce to readers there is no active writer.
	r := atomic.AddInt32(&rw.readerCount, rwmutexMaxReaders)
	if r >= rwmutexMaxReaders {
		throw("sync: Unlock of unlocked RWMutex")
	}
  // Unblock blocked readers, if any.
	for i := 0; i < int(r); i++ {
		runtime_Semrelease(&rw.readerSem, false, 0)
	}
  // Allow other writers to proceed.
	rw.w.Unlock()
}
```

与加锁的过程正好相反，写锁的释放分以下几个执行：

1. 调用 `sync/atomic.AddInt32` 函数将 `readerCount` 变回正数，释放读锁；
2. 通过 for 循环释放所有因为获取读锁而陷入等待的 Goroutine：
3. 调用 `sync.Mutex.Unlock`释放写锁；

获取写锁时会先阻塞写锁的获取，后阻塞读锁的获取，这种策略能够保证读操作不会被连续的写操作『饿死』。

#### 读锁

**读锁的加锁**方法 `sync.RWMutex.RLock` 很简单，该方法会通过 `sync/atomic.AddInt32` 将 `readerCount` 加一：

```go
// github.com/golang/go/src/sync/rwmutex.go
// RLock locks rw for reading.
func (rw *RWMutex) RLock() {
	if atomic.AddInt32(&rw.readerCount, 1) < 0 {
    // A writer is pending, wait for it.
		runtime_SemacquireMutex(&rw.readerSem, false, 0)
	}
}
```

1. 如果该方法返回负数 — 表示其他 Goroutine 获得了写锁，当前 Goroutine 就会调用 `runtime.sync_runtime_SemacquireMutex` 陷入休眠等待锁的释放；
2. 如果该方法的结果为非负数 — 没有 Goroutine 获得写锁，当前方法会成功返回；

当 Goroutine 想要**释放读锁**时，会调用如下所示的 `sync.RWMutex.RUnlock`方法：

```go
// github.com/golang/go/src/sync/rwmutex.go
// RUnlock undoes a single RLock call;
func (rw *RWMutex) RUnlock() {
	if r := atomic.AddInt32(&rw.readerCount, -1); r < 0 {
    // Outlined slow-path to allow the fast-path to be inlined
		rw.rUnlockSlow(r)
	}
}
```

该方法会先减少正在读资源的 `readerCount` 整数，根据 `sync/atomic.AddInt32` 的返回值不同会分别进行处理：

- 如果返回值大于等于零 — 读锁直接解锁成功；
- 如果返回值小于零 — 有一个正在执行的写操作，在这时会调用`sync.RWMutex.rUnlockSlow` 方法；

```go
// github.com/golang/go/src/sync/rwmutex.go
func (rw *RWMutex) rUnlockSlow(r int32) {
	if r+1 == 0 || r+1 == -rwmutexMaxReaders {
		throw("sync: RUnlock of unlocked RWMutex")
	}
	// A writer is pending.
	if atomic.AddInt32(&rw.readerWait, -1) == 0 {
		// The last reader unblocks the writer.
		runtime_Semrelease(&rw.writerSem, false, 1)
	}
}
```

`sync.RWMutex.rUnlockSlow` 会减少获取锁的写操作等待的读操作数 `readerWait` ，并在所有读操作都被释放之后触发写操作的信号量 `writerSem`，该信号量被触发时，调度器就会唤醒尝试获取写锁的 Goroutine。

#### 小结

虽然读写互斥锁 `sync.RWMutex` 提供的功能比较复杂，但是因为它建立在 `sync.Mutex` 上，所以实现会简单很多。总结一下读锁和写锁的关系：

- 调用`sync.RWMutex.Lock`尝试获取写锁时；

  - 每次 `sync.RWMutex.RUnlock`都会将 `readerCount` 其减一，当它归零时该 Goroutine 会获得写锁；
  - 将 `readerCount` 减少 `rwmutexMaxReaders` 个数以阻塞后续的读操作；

- 调用`sync.RWMutex.Unlock` 释放写锁时，会先通知所有的读操作，然后才会释放持有的互斥锁；

读写互斥锁在互斥锁之上提供了额外的更细粒度的控制，能够在读操作远远多于写操作时提升性能。



### WaitGroup

`sync.WaitGroup`可以**等待一组 Goroutine 的返回**，一个比较常见的使用场景是批量发出 RPC 或者 HTTP 请求：

```go
requests := []*Request{...}
wg := &sync.WaitGroup{}
wg.Add(len(requests))

for _, request := range requests {
    go func(r *Request) {
        defer wg.Done()
        // res, err := service.call(r)
    }(request)
}
wg.Wait()
```

可以通过 `sync.WaitGroup`将原本顺序执行的代码**在多个 Goroutine 中并发执行**，加快程序处理的速度。

![golang-syncgroup](go_concurrent.assets/2020-01-23-15797104328028-golang-syncgroup.png)

**WaitGroup 等待多个 Goroutine**

#### 结构体 

`sync.WaitGroup`结构体中只包含两个成员变量：

```go
// github.com/golang/go/src/sync/waitgroup.go
type WaitGroup struct {
	noCopy noCopy
  // 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit compilers do not ensure it. 
	// So we allocate 12 bytes and then use the aligned 8 bytes in them as state, and the other 4 as storage for the sema.
	state1 [3]uint32 
}
```

- `noCopy` — 保证 `sync.WaitGroup`不会被开发者通过再赋值的方式拷贝；
- `state1` — 存储着状态和信号量；

`sync.noCopy`是一个特殊的私有结构体，`tools/go/analysis/passes/copylock` 包中的**分析器**会在编译期间检查被拷贝的变量中是否包含 `sync.noCopy`或者实现了 `Lock` 和 `Unlock` 方法，如果包含该结构体或者实现了对应的方法就会报出以下错误：

```go
func main() {
	wg := sync.WaitGroup{}
	yawg := wg  // https://pkg.go.dev/golang.org/x/tools/go/analysis/passes/copylock
	fmt.Println(wg, yawg)
}

$ go vet proc.go
./prog.go:10:10: assignment copies lock value to yawg: sync.WaitGroup
./prog.go:11:14: call of fmt.Println copies lock value: sync.WaitGroup
./prog.go:11:18: call of fmt.Println copies lock value: sync.WaitGroup
```

这段代码会因为变量赋值或者调用函数时**发生值拷贝导致分析器报错**。

除了 `sync.noCopy`之外，`sync.WaitGroup`结构体中还包含一个总共**占用 12 字节的数组**，这个数组会存储当前结构体的状态，在 64 位与 32 位的机器上表现也非常不同。

![golang-waitgroup-state](go_concurrent.assets/2020-01-23-15797104328035-golang-waitgroup-state.png)

**WaitGroup 在 64 位和 32 位机器的不同状态**

`sync.WaitGroup`提供的私有方法 `sync.WaitGroup.state` 能够从 `state1` 字段中**取出它的状态和信号量**。

```go
// github.com/golang/go/src/sync/waitgroup.go
// state returns pointers to the state and sema fields stored within wg.state1.
func (wg *WaitGroup) state() (statep *uint64, semap *uint32) {
   if uintptr(unsafe.Pointer(&wg.state1))%8 == 0 {
      return (*uint64)(unsafe.Pointer(&wg.state1)), &wg.state1[2]
   } else {
      return (*uint64)(unsafe.Pointer(&wg.state1[1])), &wg.state1[0]
   }
}
```

#### 接口

`sync.WaitGroup`对外暴露了三个方法 — `sync.WaitGroup.Add`、`sync.WaitGroup.Wait` 和 `sync.WaitGroup.Done`。

因为其中的 `sync.WaitGroup.Done` 只是向 `sync.WaitGroup.Add` 方法传入了 -1，所以重点分析另外两个方法，即 `sync.WaitGroup.Add` 和 `sync.WaitGroup.Wait`：

```go
// github.com/golang/go/src/sync/waitgroup.go
func (wg *WaitGroup) Add(delta int) {
	statep, semap := wg.state()
	state := atomic.AddUint64(statep, uint64(delta)<<32)
  v := int32(state >> 32) // v: counter
	w := uint32(state)
	if v < 0 {
		panic("sync: negative WaitGroup counter")
	}
	if v > 0 || w == 0 {
		return
	}
	*statep = 0
	for ; w != 0; w-- {
		runtime_Semrelease(semap, false, 0)
	}
}
```

`sync.WaitGroup.Add` 可以更新 `sync.WaitGroup`中的计数器 `counter`。

虽然 `sync.WaitGroup.Add` 方法传入的参数可以为负数，但是**计数器只能是非负数**，一旦出现负数就会发生程序崩溃。

当调用计数器归零，即所有任务都执行完成时，才会通过 `sync.runtime_Semrelease` 唤醒处于等待状态的 Goroutine。

`sync.WaitGroup`的另一个方法 `sync.WaitGroup.Wait` 会在计数器大于 0 并且不存在等待的 Goroutine 时，调用 `runtime.sync_runtime_Semacquire` 陷入睡眠。

```go
// github.com/golang/go/src/sync/waitgroup.go
func (wg *WaitGroup) Wait() {
	statep, semap := wg.state()
	for {
		state := atomic.LoadUint64(statep)
		v := int32(state >> 32)
		if v == 0 {
			return
		}
		if atomic.CompareAndSwapUint64(statep, state, state+1) {
			runtime_Semacquire(semap)
			if +statep != 0 {
				panic("sync: WaitGroup is reused before previous Wait has returned")
			}
			return
		}
	}
}
```

当 `sync.WaitGroup`的计数器归零时，陷入睡眠状态的 Goroutine 会被唤醒，上述方法也会立刻返回。

#### 小结

通过对 `sync.WaitGroup`的分析和研究，能够得出以下结论：

- `sync.WaitGroup`必须在 `sync.WaitGroup.Wait` 方法返回之后才能被重新使用；
- `sync.WaitGroup.Done` 只是对 `sync.WaitGroup.Add` 方法的简单封装，可以向 `sync.WaitGroup.Add` 方法传入任意负数（需要保证计数器非负），快速将计数器归零以唤醒等待的 Goroutine；
- 可以同时有多个 Goroutine 等待当前 `sync.WaitGroup`计数器的归零，这些 Goroutine 会被同时唤醒；



### Once

Go 语言标准库中 `sync.Once` 可以保证在 Go 程序运行期间的某段代码只会执行一次。在运行如下所示的代码时，会看到如下所示的运行结果：

```go
func main() {
    o := &sync.Once{}
    for i := 0; i < 10; i++ {
        o.Do(func() {
            fmt.Println("only once")
        })
    }
}

$ go run main.go
only once
```

#### 结构体

每一个 `sync.Once` 结构体中都只包含一个用于标识代码块是否执行过的 `done` 以及一个互斥锁 `sync.Mutex`：

```go
// github.com/golang/go/src/sync/once.go
type Once struct {
  // done indicates whether the action has been performed.
	done uint32
	m    Mutex
}
```

#### 接口

`sync.Once.Do` 是 `sync.Once` 结构体对外唯一暴露的方法，该方法会接收一个入参为空的函数：

- 如果传入的函数已经执行过，会直接返回；
- 如果传入的函数没有执行过，会调用 `sync.Once.doSlow` 执行传入的函数：

```go
// github.com/golang/go/src/sync/once.go
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}
```

1. 为当前 Goroutine 获取互斥锁；
2. 执行传入的无入参函数；
3. 运行延迟函数调用，将成员变量 `done` 更新成 1；

`sync.Once` 会通过成员变量 `done` 确保函数不会执行第二次。

#### 小结

作为用于保证函数执行次数的 `sync.Once` 结构体，它使用互斥锁和 `sync/atomic` 包提供的方法实现了某个函数在程序运行期间只能执行一次的语义。在使用该结构体时，也需要注意以下的问题：

- `sync.Once.Do` 方法中传入的函数只会被执行一次，哪怕函数中发生了 `panic`；
- 两次调用 `sync.Once.Do` 方法传入不同的函数只会执行第一次调用传入的函数；



### Cond

Go 语言标准库中还包含条件变量 `sync.Cond`，它可以**让一组的 Goroutine 都在满足特定条件时被唤醒**。

每一个 `sync.Cond` 结构体在初始化时都需要传入一个互斥锁，可以通过下面的例子了解它的使用方法：

```go
var status int64

func main() {
	c := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go listen(c)
	}
	time.Sleep(1 * time.Second)
	go broadcast(c)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 {
		c.Wait()
	}
	fmt.Println("listen")
	c.L.Unlock()
}

$ go run main.go
listen
...
listen
```

上述代码同时运行了 11 个 Goroutine，这 11 个 Goroutine 分别做了不同事情：

- 10 个 Goroutine 通过 `sync.Cond.Wait` 等待特定条件的满足；
- 1 个 Goroutine 会调用 `sync.Cond.Broadcast` 唤醒所有陷入等待的 Goroutine；

调用 `sync.Cond.Broadcast` 方法后，上述代码会打印出 10 次 “listen” 并结束调用。

![golang-cond-broadcast](go_concurrent.assets/2020-01-23-15797104328042-golang-cond-broadcast.png)

**Cond 条件广播**

#### 结构体

`sync.Cond` 的结构体中包含以下 4 个字段：

```go
// github.com/golang/go/src/sync/cond.go
type Cond struct {
	noCopy  noCopy
	L       Locker // L is held while observing or changing the condition
	notify  notifyList
	checker copyChecker
}
```

- `noCopy` — 用于保证结构体不会在编译期间拷贝；
- `copyChecker` — 用于禁止运行期间发生的拷贝；
- `L` — 用于保护内部的 `notify` 字段，`Locker` 接口类型的变量；
- `notify` — 一个 Goroutine 的链表，它是实现同步机制的核心结构；

```go
// github.com/golang/go/src/sync/runtime2.go
type notifyList struct {
	wait uint32
	notify uint32

	lock mutex // key field of the mutex
	head *sudog
	tail *sudog
}
```

在 `sync.notifyList`结构体中，`head` 和 `tail` 分别指向的链表的头和尾，`wait` 和 `notify` 分别表示当前正在等待的和已经通知到的 Goroutine 的索引。

#### 接口

##### Goroutine 陷入休眠

`sync.Cond` 对外暴露的 `sync.Cond.Wait` 方法会将当前 Goroutine 陷入休眠状态，它的执行过程分成以下两个步骤：

1. 调用 `runtime.notifyListAdd` 将等待计数器加一并解锁；
2. 调用 `runtime.notifyListWait` 等待其他 Goroutine 的唤醒并加锁：

```go
// github.com/golang/go/src/sync/cond.go
func (c *Cond) Wait() {
	c.checker.check()
	t := runtime_notifyListAdd(&c.notify) // runtime.notifyListAdd 的链接名
	c.L.Unlock()
	runtime_notifyListWait(&c.notify, t) // runtime.notifyListWait 的链接名
	c.L.Lock()
}

// github.com/golang/go/src/runtime/sema.go
func notifyListAdd(l *notifyList) uint32 {
	return atomic.Xadd(&l.wait, 1) - 1
}
```

`runtime.notifyListWait` 会获取当前 Goroutine 并将它追加到 Goroutine 通知链表的最末端：

```go
// github.com/golang/go/src/runtime/sema.go
func notifyListWait(l *notifyList, t uint32) {
	s := acquireSudog()
	s.g = getg()
	s.ticket = t
	if l.tail == nil {
		l.head = s
	} else {
		l.tail.next = s
	}
	l.tail = s
	goparkunlock(&l.lock, waitReasonSyncCondWait, traceEvGoBlockCond, 3)
	releaseSudog(s)
}
```

除了将当前 Goroutine 追加到链表的末端之外，还会调用 `runtime.goparkunlock` 将当前 Goroutine 陷入休眠，该函数也是在 Go 语言切换 Goroutine 时经常会使用的方法，它会直接让出当前处理器的使用权并等待调度器的唤醒。

![golang-cond-notifylist](go_concurrent.assets/2020-01-23-15797104328049-golang-cond-notifylist.png)

**Cond 条件通知列表**

##### 唤醒休眠的 Goroutine

`sync.Cond.Signal`和 `sync.Cond.Broadcast` 就是用来唤醒陷入休眠的 Goroutine 的方法，它们的实现有一些细微的差别：

- `sync.Cond.Signal`方法会唤醒队列最前面的 Goroutine；
- `sync.Cond.Broadcast` 方法会唤醒队列中全部的 Goroutine；

```go
// github.com/golang/go/src/sync/cond.go
func (c *Cond) Signal() {
	c.checker.check()
	runtime_notifyListNotifyOne(&c.notify)
}

func (c *Cond) Broadcast() {
	c.checker.check()
	runtime_notifyListNotifyAll(&c.notify)
}
```

`runtime.notifyListNotifyOne` 只会从 `sync.notifyList`链表中找到满足 `sudog.ticket == l.notify` 条件的 Goroutine 并通过 `runtime.readyWithTime` 唤醒：

```go
// github.com/golang/go/src/runtime/sema.go
// notifyListNotifyOne notifies one entry in the list.
func notifyListNotifyOne(l *notifyList) {
	t := l.notify
	atomic.Store(&l.notify, t+1)

	for p, s := (*sudog)(nil), l.head; s != nil; p, s = s, s.next {
		if s.ticket == t {
			n := s.next
			if p != nil {
				p.next = n
			} else {
				l.head = n
			}
			if n == nil {
				l.tail = p
			}
			s.next = nil
			readyWithTime(s, 4)
			return
		}
	}
}
```

`runtime.notifyListNotifyAll` 会依次通过 `runtime.readyWithTime` 唤醒链表中 Goroutine：

```go
// github.com/golang/go/src/runtime/sema.go
// notifyListNotifyAll notifies all entries in the list.
func notifyListNotifyAll(l *notifyList) {
	s := l.head
	l.head = nil
	l.tail = nil

	atomic.Store(&l.notify, atomic.Load(&l.wait))

	for s != nil {
		next := s.next
		s.next = nil
		readyWithTime(s, 4)
		s = next
	}
}
```

Goroutine 的唤醒顺序也是按照**加入队列的先后顺序**，先加入的会先被唤醒，而后加入的可能 Goroutine 需要等待调度器的调度。

在一般情况下，都会先调用 `sync.Cond.Wait` 陷入休眠等待满足期望条件，当满足唤醒条件时，就可以选择使用 `sync.Cond.Signal`或者 `sync.Cond.Broadcast` 唤醒一个或者全部的 Goroutine。

#### 小结

`sync.Cond` 不是一个常用的同步机制，但是**在条件长时间无法满足时**，与使用 `for {}` 进行忙碌等待相比，`sync.Cond` 能够让出处理器的使用权，提高 CPU 的利用率。使用时也需要注意以下问题：

- `sync.Cond.Wait` 在调用之前一定要传入需要使用的互斥锁，否则会触发程序崩溃；
- `sync.Cond.Signal`唤醒的 Goroutine 都是队列最前面、等待最久的 Goroutine；
- `sync.Cond.Broadcast` 会按照一定顺序广播通知等待的全部 Goroutine；



## 扩展原语

除了标准库中提供的同步原语之外，Go 语言还在子仓库 [sync](https://github.com/golang/sync) 中提供了四种扩展原语，`golang/sync/errgroup.Group`、`golang/sync/semaphore.Weighted`、`golang/sync/singleflight.Group` 和 `golang/sync/syncmap.Map`，其中的 `golang/sync/syncmap.Map` 在 1.9 版本中被移植到了标准库中。

![golang-extension-sync-primitives](go_concurrent.assets/2020-01-23-15797104328056-golang-extension-sync-primitives.png)

**Go 扩展原语**

介绍 Go 语言在扩展包中提供的三种同步原语，也就是 `golang/sync/errgroup.Group`、`golang/sync/semaphore.Weighted` 和 `golang/sync/singleflight.Group`。

### ErrGroup

`golang/sync/errgroup.Group` 在一组 Goroutine 中提供了同步、错误传播以及上下文取消的功能，可以使用如下所示的方式**并行获取网页的数据**：

```go
var g errgroup.Group
var urls = []string{
    "http://www.golang.org/",
    "http://www.google.com/",
}
for i := range urls {
    url := urls[i]
    g.Go(func() error {
        resp, err := http.Get(url)
        if err == nil {
            resp.Body.Close()
        }
        return err
    })
}
if err := g.Wait(); err == nil {
    fmt.Println("Successfully fetched all URLs.")
}
```

`golang/sync/errgroup.Group.Go` 方法能够创建一个 Goroutine 并在其中执行传入的函数，而 `golang/sync/errgroup.Group.Wait` 会等待所有 Goroutine 全部返回，该方法的不同返回结果也有不同的含义：

- 如果返回错误 — 这一组 Goroutine 最少返回一个错误；
- 如果返回空值 — 所有 Goroutine 都成功执行；

#### 结构体

`golang/sync/errgroup.Group` 结构体同时由三个比较重要的部分组成：

1. `cancel` — 创建 `context.Context` 时返回的取消函数，用于在多个 Goroutine 之间同步取消信号；
2. `wg` — 用于等待一组 Goroutine 完成子任务的同步原语；
3. `errOnce` — 用于保证只接收一个子任务返回的错误；

```go
// github.com/golang/sync/errgroup/errgroup.go
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}
```

这些字段共同组成了 `golang/sync/errgroup.Group` 结构体并提供同步、错误传播以及上下文取消等功能。

#### 接口

通过 `golang/sync/errgroup.WithContext` 构造器创建新的 `golang/sync/errgroup.Group` 结构体：

```go
// github.com/golang/sync/errgroup/errgroup.go
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}
```

运行新的并行子任务需要使用 `golang/sync/errgroup.Group.Go` 方法，这个方法的执行过程如下：

1. 调用 `sync.WaitGroup.Add` 增加待处理的任务；
2. 创建新的 Goroutine 并运行子任务；
3. 返回错误时及时调用 `cancel` 并对 `err` 赋值，只有最早返回的错误才会被上游感知到，后续的错误都会被舍弃：

```go
// github.com/golang/sync/errgroup/errgroup.go
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel()
				}
			})
		}
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.err
}
```

另一个用于等待的 `golang/sync/errgroup.Group.Wait` 方法只是调用了 `sync.WaitGroup.Wait`，在子任务全部完成时取消 `context.Context` 并返回可能出现的错误。

#### 小结

`golang/sync/errgroup.Group` 的实现没有涉及底层和运行时包中的 API，它只是对基本同步语义进行了封装以提供更加复杂的功能。在使用时也需要注意下面几个问题：

- `golang/sync/errgroup.Group` 在出现错误或者等待结束后会调用 `context.Context` 的 `cancel` 方法同步取消信号；
- 只有第一个出现的错误才会被返回，剩余的错误会被直接丢弃；



### Semaphore

信号量是在并发编程中常见的一种同步机制，在需要**控制访问资源的进程数量时就会用到信号量**，它会保证持有的计数器在 0 到初始化的权重之间波动。

- 每次获取资源时都会将信号量中的计数器减去对应的数值，在释放时重新加回来；
- 当遇到计数器大于信号量大小时，会进入休眠等待其他线程释放信号；

Go 语言的扩展包中就提供了**带权重的信号量** `golang/sync/semaphore.Weighted`，可以按照不同的权重对资源的访问进行管理，这个结构体对外也只暴露了四个方法：

- `golang/sync/semaphore.NewWeighted` 用于创建新的信号量；
- `golang/sync/semaphore.Weighted.Acquire` 阻塞地获取指定权重的资源，如果当前没有空闲资源，会陷入休眠等待；
- `golang/sync/semaphore.Weighted.TryAcquire` 非阻塞地获取指定权重的资源，如果当前没有空闲资源，会直接返回 `false`；
- `golang/sync/semaphore.Weighted.Release` 用于释放指定权重的资源；

#### 结构体

`golang/sync/semaphore.NewWeighted` 方法能根据传入的最大权重创建一个指向 `golang/sync/semaphore.Weighted` 结构体的指针：

```go
// github.com/golang/sync/semaphore/semaphore.go
func NewWeighted(n int64) *Weighted {
	w := &Weighted{size: n}
	return w
}

type Weighted struct {
	size    int64  // 上限
	cur     int64  // 计数器
	mu      sync.Mutex
	waiters list.List
}
```

`golang/sync/semaphore.Weighted` 结构体中包含一个 `waiters` 列表，其中存储着等待获取资源的 Goroutine，除此之外它还包含当前信号量的上限`size`以及一个计数器 `cur`，这个计数器的范围就是 [0, size]：

![golang-semaphore](go_concurrent.assets/2020-01-23-15797104328063-golang-semaphore.png)

**权重信号量**

信号量中的计数器会随着用户对资源的访问和释放进行改变，引入的权重概念能够提供更细粒度的资源的访问控制，尽可能满足常见的用例。

#### 获取

`golang/sync/semaphore.Weighted.Acquire` 方法能用于获取指定权重的资源，其中包含三种不同情况：

1. 当信号量中剩余的资源大于获取的资源并且没有等待的 Goroutine 时，会直接获取信号量；
2. 当需要获取的信号量大于 `golang/sync/semaphore.Weighted` 的上限时，由于不可能满足条件会直接返回错误；
3. 遇到其他情况时会将当前 Goroutine 加入到等待列表并通过 `select` 等待调度器唤醒当前 Goroutine，Goroutine 被唤醒后会获取信号量；

```go
// github.com/golang/sync/semaphore/semaphore.go
func (s *Weighted) Acquire(ctx context.Context, n int64) error {

  // 第一种情况
  if s.size-s.cur >= n && s.waiters.Len() == 0 {
		s.cur += n
		return nil
	}

  // 第二种情况
	if n > s.size {
		// Don't make other Acquire calls block on one that's doomed to fail.
		s.mu.Unlock()
		<-ctx.Done()
		return ctx.Err()
	}
  
  // 其他情况
	ready := make(chan struct{})
	w := waiter{n: n, ready: ready}
	elem := s.waiters.PushBack(w)
	select {
	case <-ctx.Done():
		err := ctx.Err()
		select {
		case <-ready:
			err = nil
		default:
			s.waiters.Remove(elem)
		}
		return err
	case <-ready:
		return nil
	}
}
```

另一个用于获取信号量的方法 `golang/sync/semaphore.Weighted.TryAcquire` 只会非阻塞地判断当前信号量是否有充足的资源，如果有充足的资源会直接立刻返回 `true`，否则会返回 `false`：

```go
// github.com/golang/sync/semaphore/semaphore.go
func (s *Weighted) TryAcquire(n int64) bool {
	s.mu.Lock()
	success := s.size-s.cur >= n && s.waiters.Len() == 0
	if success {
		s.cur += n
	}
	s.mu.Unlock()
	return success
}
```

因为 `golang/sync/semaphore.Weighted.TryAcquire` 不会等待资源的释放，所以可能更适用于一些延时敏感、用户需要立刻感知结果的场景。

#### 释放

当要释放信号量时，`golang/sync/semaphore.Weighted.Release` 方法会从头到尾遍历 `waiters` 列表中全部的等待者，如果释放资源后的信号量有充足的剩余资源就会通过 Channel 唤起指定的 Goroutine：

```go
// github.com/golang/sync/semaphore/semaphore.go
func (s *Weighted) Release(n int64) {
	s.mu.Lock()
	s.cur -= n
	for {
		next := s.waiters.Front()
		if next == nil {
			break
		}
		w := next.Value.(waiter)
		if s.size-s.cur < w.n {
			break
		}
		s.cur += w.n
		s.waiters.Remove(next)
		close(w.ready)
	}
	s.mu.Unlock()
}
```

当然也可能会出现剩余资源无法唤起 Goroutine 的情况，在这时当前方法在释放锁后会直接返回。

通过对 `golang/sync/semaphore.Weighted.Release` 的分析能发现，如果一个信号量需要的占用的资源非常多，它可能会**长时间无法获取锁**，这也是 `golang/sync/semaphore.Weighted.Acquire` 引入上下文参数的原因，即为信号量的获取设置超时时间。

#### 小结

带权重的信号量确实有着更多的应用场景，这也是 Go 语言对外提供的唯一一种信号量实现，在使用的过程中需要注意以下的几个问题：

- `golang/sync/semaphore.Weighted.Acquire` 和 `golang/sync/semaphore.Weighted.TryAcquire` 都可以用于获取资源，前者会阻塞地获取信号量，后者会非阻塞地获取信号量；
- `golang/sync/semaphore.Weighted.Release` 方法会按照先进先出的顺序唤醒可以被唤醒的 Goroutine；
- 如果一个 Goroutine 获取了较多地资源，由于 `golang/sync/semaphore.Weighted.Release` 的释放策略可能会等待比较长的时间；



### SingleFlight

`golang/sync/singleflight.Group` 是 Go 语言扩展包中提供了另一种同步原语，它能够**在一个服务中抑制对下游的多次重复请求**。

一个比较常见的使用场景是：在使用 Redis 对数据库中的数据进行缓存，发生**缓存击穿**时，大量的流量都会打到数据库上进而影响服务的尾延时。

![golang-query-without-single-flight](go_concurrent.assets/2020-01-23-15797104328070-golang-query-without-single-flight.png)

**Redis 缓存击穿问题**

但是 `golang/sync/singleflight.Group` 能有效地解决这个问题，它能够限制对同一个键值对的多次重复请求，减少对下游的瞬时流量。

![golang-extension-single-flight](go_concurrent.assets/2020-01-23-15797104328078-golang-extension-single-flight.png)

**缓解缓存击穿问题**

在资源的获取非常昂贵时（例如：访问缓存、数据库），就很适合使用 `golang/sync/singleflight.Group` 优化服务。来了解一下它的使用方法：

```go
type service struct {
    requestGroup singleflight.Group
}

func (s *service) handleRequest(ctx context.Context, request Request) (Response, error) {
    v, err, _ := requestGroup.Do(request.Hash(), func() (interface{}, error) {
        rows, err := // select * from tables
        if err != nil {
            return nil, err
        }
        return rows, nil
    })
    if err != nil {
        return nil, err
    }
    return Response{
        rows: rows,
    }, nil
}
```

因为请求的哈希在业务上一般表示相同的请求，所以上述代码使用它作为请求的键。当然，也可以选择其他的字段作为 `golang/sync/singleflight.Group.Do` 方法的第一个参数减少重复的请求。

#### 结构体

`golang/sync/singleflight.Group` 结构体由一个互斥锁 `sync.Mutex` 和一个映射表组成，每一个 `golang/sync/singleflight.call` 结构体都保存了当前调用对应的信息：

```go
// github.com/golang/sync/singleflight/singleflight.go
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

type call struct {
	wg sync.WaitGroup

	val interface{}
	err error

	dups  int
	chans []chan<- Result
}
```

`golang/sync/singleflight.call` 结构体中的 `val` 和 `err` 字段都只会在执行传入的函数时，赋值一次并在 `sync.WaitGroup.Wait` 返回时被读取；`dups` 和 `chans` 两个字段分别存储了抑制的请求数量以及用于同步结果的 Channel。

#### 接口

`golang/sync/singleflight.Group` 提供了两个用于抑制相同请求的方法：

- `golang/sync/singleflight.Group.Do` — 同步等待的方法；
- `golang/sync/singleflight.Group.DoChan` — 返回 Channel 异步等待的方法；

这两个方法在功能上没有太多的区别，只是在接口的表现上稍有不同。

每次调用 `golang/sync/singleflight.Group.Do` 方法时都会获取互斥锁，随后判断是否已经存在键对应的 `golang/sync/singleflight.call`：

1. 当不存在对应的`golang/sync/singleflight.call` 时：

   1. 初始化一个新的 `golang/sync/singleflight.call` 指针；
   2. 增加 `sync.WaitGroup` 持有的计数器；

   3. 将 `golang/sync/singleflight.call` 指针添加到映射表；
   4. 释放持有的互斥锁；
   5. 阻塞地调用 `golang/sync/singleflight.Group.doCall` 方法等待结果的返回；

2. 当存在对应的`golang/sync/singleflight.call` 时；

   1. 增加 `dups` 计数器，它表示当前重复的调用次数；
   2. 释放持有的互斥锁；
   3. 通过 `sync.WaitGroup.Wait` 等待请求的返回；

```go
// github.com/golang/sync/singleflight/singleflight.go
func (g *Group) Do(key string, fn func() (interface{}, error)) (v interface{}, err error, shared bool) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err, true
	}
	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	g.doCall(c, key, fn)
	return c.val, c.err, c.dups > 0
}
```

因为 `val` 和 `err` 两个字段都只会在 `golang/sync/singleflight.Group.doCall` 方法中赋值，所以当 `golang/sync/singleflight.Group.doCall` 和 `sync.WaitGroup.Wait` 返回时，函数调用的结果和错误都会返回给 `golang/sync/singleflight.Group.Do` 的调用者。

```go
// github.com/golang/sync/singleflight/singleflight.go
func (g *Group) doCall(c *call, key string, fn func() (interface{}, error)) {
	c.val, c.err = fn()
	c.wg.Done()

	g.mu.Lock()
	delete(g.m, key)
	for _, ch := range c.chans {
		ch <- Result{c.val, c.err, c.dups > 0}
	}
	g.mu.Unlock()
}
```

1. 运行传入的函数 `fn`，该函数的返回值会赋值给 `c.val` 和 `c.err`；
2. 调用 `sync.WaitGroup.Done` 方法通知所有等待结果的 Goroutine — 当前函数已经执行完成，可以从 `call` 结构体中取出返回值并返回了；
3. 获取持有的互斥锁并通过管道将信息同步给使用 `golang/sync/singleflight.Group.DoChan` 方法的 Goroutine；

```go
func (g *Group) DoChan(key string, fn func() (interface{}, error)) <-chan Result {
	ch := make(chan Result, 1)
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	if c, ok := g.m[key]; ok {
		c.dups++
		c.chans = append(c.chans, ch)
		g.mu.Unlock()
		return ch
	}
	c := &call{chans: []chan<- Result{ch}}
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock()

	go g.doCall(c, key, fn)

	return ch
}
```

`golang/sync/singleflight.Group.Do` 和 `golang/sync/singleflight.Group.DoChan` 分别提供了同步和异步的调用方式，这让使用起来也更加灵活。

#### 小结

当需要减少对下游的相同请求时，可以使用 `golang/sync/singleflight.Group` 来增加吞吐量和服务质量，不过在使用的过程中也需要注意以下的几个问题：

- `golang/sync/singleflight.Group.Do` 和 `golang/sync/singleflight.Group.DoChan` 一个用于同步阻塞调用传入的函数，一个用于异步调用传入的参数并通过 Channel 接收函数的返回值；
- `golang/sync/singleflight.Group.Forget` 可以通知 `golang/sync/singleflight.Group` 在持有的映射表中删除某个键，接下来对该键的调用就不会等待前面的函数返回了；
- 一旦调用的函数返回了错误，所有在等待的 Goroutine 也都会接收到同样的错误；

## 小结

介绍了 Go 语言标准库中提供的基本原语以及扩展包中的扩展原语，这些并发编程的原语能够更好地利用 Go 语言的特性构建高吞吐量、低延时的服务、解决并发带来的问题。

在设计同步原语时，不仅要考虑 API 接口的易用、解决并发编程中可能遇到的线程竞争问题，还需要对尾延时进行优化，保证公平性，理解同步原语也是理解并发编程无法跨越的一个步骤。

## 参考

- “sync: allow inlining the Mutex.Lock fast path” https://github.com/golang/go/commit/41cb0aedffdf4c5087de82710c4d016a3634b4ac
- “sync: allow inlining the Mutex.Unlock fast path” https://github.com/golang/go/commit/4c3f26076b6a9853bcc3c7d7e43726c044ac028a#diff-daec021895d1400f2c064a3e851c0d2c
- “runtime: fall back to fair locks after repeated sleep-acquire failures” https://github.com/golang/go/issues/13086
- Go Team. May 2014. “The Go Memory Model” https://golang.org/ref/mem
- Chris. May 2017. “The X-Files: Exploring the Golang Standard Library Sub-Repositories” https://rodaine.com/2017/05/x-files-intro/
- Dmitry Vyukov, Russ Cox. Dec 13, 2016. “sync: make Mutex more fair” https://github.com/golang/go/commit/0556e26273f704db73df9e7c4c3d2e8434dec7be 
- golang/sync/syncmap: recommend sync.Map #33867 https://github.com/golang/go/issues/33867 



## 计时器

准确的时间对于任何一个正在运行的应用非常重要，但是在分布式系统中很难保证各个节点的绝对时间一致，哪怕通过 NTP 这种标准的对时协议也只能把各个节点上时间的误差控制在毫秒级，所以准确的相对时间在分布式系统中显得更为重要，本节会分析用于获取相对时间的计时器的设计与实现原理。

### 设计原理

Go 语言从实现计时器到现在经历过很多个版本的迭代，到最新的版本为止，计时器的实现分别经历了以下几个过程：

1. Go 1.9 版本之前，所有的计时器由全局唯一的四叉堆维护；
2. Go 1.10 ~ 1.13，全局使用 64 个四叉堆维护全部的计时器，每个处理器（P）创建的计时器会由对应的四叉堆维护；
3. Go 1.14 版本之后，每个处理器单独管理计时器并通过网络轮询器触发；

分别介绍计时器在不同版本的不同设计，梳理计时器实现的演进过程。

#### 全局四叉堆

Go 1.10 之前的计时器都使用**最小四叉堆**实现，所有的计时器都会存储在如下所示的结构体 `runtime.timers:093adee` 中：

```go
// github.com/golang/go/src/runtime/time.go
// go 1.9
var timers struct {
	lock         mutex
	gp           *g
	created      bool
	sleeping     bool
	rescheduling bool
	sleepUntil   int64
	waitnote     note
	t            []*timer
}
```

这个结构体中的字段 `t` 就是最小四叉堆，运行时创建的所有计时器都会加入到四叉堆中。

```go
// github.com/golang/go/src/runtime/time.go
func timerproc() {
   timers.gp = getg()
   for {
      ...
      for {
         if len(timers.t) == 0 {
            delta = -1
            break
         }
         ...
         if t.period > 0 {
            // leave in heap but adjust next time to fire
            t.when += t.period * (1 + -delta/t.period)
            siftdownTimer(0)
         } else {
            // remove from heap
            last := len(timers.t) - 1
            if last > 0 {
               timers.t[0] = timers.t[last]
               timers.t[0].i = 0
            }
            timers.t[last] = nil
            timers.t = timers.t[:last]
            ...
         }
        ...
      }
   }
}
```

`runtime.timerproc:093adee` 的Goroutine 会运行 时间驱动的事件，运行时 会在发生以下事件时唤醒计时器：

- 四叉堆中的计时器到期；
- 四叉堆中加入了触发时间更早的新计时器；

![golang-timer-quadtree](go_concurrent.assets/2020-01-25-15799218054781-golang-timer-quadtree.png)

**计时器四叉堆**

然而全局四叉堆**共用的互斥锁**对计时器的影响非常大，计时器的各种操作都需要获取全局唯一的互斥锁，这会严重影响计时器的性能。

#### 分片四叉堆

Go 1.10 将全局的四叉堆分割成了 64 个更小的四叉堆。

在理想情况下，四叉堆的数量应该等于处理器的数量，但是这需要实现动态的分配过程，所以经过权衡最终选择初始化 64 个四叉堆，以牺牲内存占用的代价换取性能的提升（空间换时间的思路）。

```go
// github.com/golang/go/src/runtime/time.go
// go 1.10
const timersLen = 64

var timers [timersLen]struct {
	timersBucket
  
  pad [sys.CacheLineSize - unsafe.Sizeof(timersBucket{})%sys.CacheLineSize]byte
}

type timersBucket struct {
	lock         mutex
	gp           *g
	created      bool
	sleeping     bool
	rescheduling bool
	sleepUntil   int64
	waitnote     note
	t            []*timer
}
```

如果当前机器上的处理器 P 的个数超过了 64，多个处理器上的计时器就可能存储在同一个桶中。每一个计时器桶都由一个运行 `runtime.timerproc:76f4fd8` 函数的 Goroutine 处理。

![golang-timer-bucket](go_concurrent.assets/2020-01-25-15799218054791-golang-timer-bucket.png)

**分片计时器桶**

将全局计时器分片的方式，虽然能够降低锁的粒度，提高计时器的性能，但是 `runtime.timerproc:76f4fd8` 造成的**处理器和线程之间频繁的上下文切换**却成为了影响计时器性能的首要因素。

#### 网络轮询器

在最新版本的实现中，计时器桶已经被移除，所有的计时器都以最小四叉堆的形式存储在处理器 `runtime.p` 中。

![golang-p-and-timers](go_concurrent.assets/2020-01-25-15799218054798-golang-p-and-timers.png)

**处理器中的最小四叉堆**

处理器 `runtime.p` 中与计时器相关的有以下字段：

- `timersLock` — 用于保护计时器的互斥锁；
- `timers` — 存储计时器的最小四叉堆；
- `numTimers` — 处理器中的计时器数量；
- `adjustTimers` — 处理器中处于 `timerModifiedEarlier` 状态的计时器数量；
- `deletedTimers` — 处理器中处于 `timerDeleted` 状态的计时器数量；

```go
// github.com/golang/go/src/runtime/runtime2.go
type p struct {
	...
	timersLock mutex
	timers []*timer

	numTimers     uint32
	adjustTimers  uint32
	deletedTimers uint32
	...
}
```

原本用于管理计时器的 `runtime.timerproc:76f4fd8` 也已经被移除，目前计时器都交**由处理器的网络轮询器和调度器触发**，这种方式能够充分利用本地性、减少上下文的切换开销，也是目前性能最好的实现方式。

### 数据结构

`runtime.timer` 是 Go 语言计时器的内部表示，每一个计时器都存储在对应处理器的最小四叉堆中，下面是运行时计时器对应的结构体：

```go
// github.com/golang/go/src/runtime/time.go
// go 1.17
type timer struct {
	pp puintptr

	when     int64
	period   int64
	f        func(interface{}, uintptr)
	arg      interface{}
	seq      uintptr
	nextwhen int64
	status   uint32
}
```

- `when` — 当前计时器被唤醒的时间；
- `period` — 两次被唤醒的间隔；
- `f` — 每当计时器被唤醒时都会调用的函数；
- `arg` — 计时器被唤醒时调用 `f` 传入的参数；
- `nextWhen` — 计时器处于 `timerModifiedXX` 状态时，用于设置 `when` 字段；
- `status` — 计时器的状态；

然而这里的 `runtime.timer` 只是计时器运行时的私有结构体，对外暴露的计时器使用 `time.Timer` 结构体：

```go
// github.com/golang/go/src/time/sleep.go
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

`time.Timer` 计时器必须通过 `time.NewTimer`、`time.AfterFunc`或者 `time.After` 函数创建。 

当计时器失效时，订阅计时器 Channel 的 Goroutine 会收到计时器失效的时间。

### 状态机

运行时使用状态机的方式处理全部的计时器，其中包括 **10 种状态**和 **7 种操作**（最新状态显示，删除了重置计数器操作）。由于 Go 语言的计时器需要同时支持增加、删除、修改和重置等操作，所以它的状态非常复杂，目前会包含以下 10 种可能：

| 状态                 |          解释          |
| :------------------- | :--------------------: |
| timerNoStatus        |     还没有设置状态     |
| timerWaiting         |        等待触发        |
| timerRunning         |     运行计时器函数     |
| timerDeleted         |         被删除         |
| timerRemoving        |       正在被删除       |
| timerRemoved         | 已经被停止并从堆中删除 |
| timerModifying       |       正在被修改       |
| timerModifiedEarlier |  被修改到了更早的时间  |
| timerModifiedLater   |  被修改到了更晚的时间  |
| timerMoving          |  已经被修改正在被移动  |

**计时器的状态**

上述表格已经展示了不同状态的含义，但是还需要展示一些重要的信息，例如状态的存在时间、计时器是否在堆上等：

- `timerRunning`、`timerRemoving`、`timerModifying` 和 `timerMoving` — 停留的时间都比较短；
- `timerWaiting`、`timerRunning`、`timerDeleted`、`timerRemoving`、`timerModifying`、`timerModifiedEarlier`、`timerModifiedLater` 和 `timerMoving` — 计时器在处理器的堆上；
- `timerNoStatus` 和 `timerRemoved` — 计时器不在堆上；
- `timerModifiedEarlier` 和 `timerModifiedLater` — 计时器虽然在堆上，但是可能位于错误的位置上，需要重新排序；

当操作计时器时，运行时会根据状态的不同而做出反应，所以在分析计时器时会将状态作为切入点分析其实现原理。

计时器的状态机中包含如下所示的 6 种不同操作，它们分别承担了不同的职责：

- `runtime.addtimer`— 向当前处理器增加新的计时器；
- `runtime.deltimer` — 将计时器标记成 `timerDeleted` 删除处理器中的计时器；
- `runtime.modtimer` — 网络轮询器会调用该函数修改计时器；
- `runtime.cleantimers` — 清除队列头中的计时器，能够提升程序创建和删除计时器的性能；
- `runtime.adjusttimers` — 调整处理器持有的计时器堆，包括移动会稍后触发的计时器、删除标记为 `timerDeleted` 的计时器；
- `runtime.runtimer` — 检查队列头中的计时器，在其准备就绪时运行该计时器；

在这里会依次分析计时器的上述 6 个不同操作。

#### 增加计时器

当调用 `time.NewTimer` 增加新的计时器时，会执行程序中的 `runtime.addtimer`函数根据以下的规则处理计时器：

- `timerNoStatus` -> `timerWaiting`
- 其他状态 -> 崩溃：不合法的状态

```go
// github.com/golang/go/src/runtime/time.go
// addtimer adds a timer to the current P.
func addtimer(t *timer) {
	if t.status != timerNoStatus {
		badTimer()
	}
	t.status = timerWaiting
	cleantimers(pp)
	doaddtimer(pp, t)
	wakeNetPoller(when)
}
```

1. 调用 `runtime.cleantimers` 清理处理器中的计时器；

2. 调用 `runtime.doaddtimer` 将当前计时器加入处理器的 timers四叉堆中；

   1. 调用 `runtime.netpollGenericInit` 函数惰性初始化网络轮询器；

3. 调用 `runtime.wakeNetPoller` 唤醒网络轮询器中休眠的线程；

   1. 调用 `runtime.netpollBreak` 函数中断正在阻塞的网络轮询；

每次增加新的计时器都会中断正在阻塞的轮询，触发调度器检查是否有计时器到期，会在后面详细介绍计时器的触发过程。

#### 删除计时器

`runtime.deltimer` 函数会标记需要删除的计时器，它会根据以下的规则处理计时器：

- `timerWaiting` -> `timerModifying` -> `timerDeleted`
- `timerModifiedEarlier` -> `timerModifying` -> `timerDeleted`
- `timerModifiedLater` -> `timerModifying` -> `timerDeleted`
- 其他状态 -> 等待状态改变或者直接返回

```go
// github.com/golang/go/src/runtime/time.go
func deltimer(t *timer) bool {
   for {
      switch s := atomic.Load(&t.status); s {
      case timerWaiting, timerModifiedLater:
         ...
         if atomic.Cas(&t.status, s, timerModifying) {
            ...
            if !atomic.Cas(&t.status, timerModifying, timerDeleted) {
               badTimer()
            }
            ...
            // Timer was not yet run.
            return true
         } else {
            releasem(mp)
         }
      case timerModifiedEarlier:
         ...
         if atomic.Cas(&t.status, s, timerModifying) {
            ...
            if !atomic.Cas(&t.status, timerModifying, timerDeleted) {
               badTimer()
            }
            ..
            // Timer was not yet run.
            return true
         } else {
            releasem(mp)
         }
      case timerDeleted, timerRemoving, timerRemoved:
         // Timer was already run.
         return false
      case timerRunning, timerMoving:
         // The timer is being run or moved, by a different P.
         // Wait for it to complete.
         osyield()
      case timerNoStatus:
         // Removing timer that was never added or
         // has already been run. Also see issue 21874.
         return false
      case timerModifying:
         // Simultaneous calls to deltimer and modtimer.
         // Wait for the other call to complete.
         osyield()
      default:
         badTimer()
      }
   }
}
```

在删除计时器的过程中，可能会遇到其他处理器的计时器，在设置中需要将计时器标记为删除状态，并由持有计时器的处理器完成清除工作。

#### 修改计时器

`runtime.modtimer` 会修改已经存在的计时器，它会根据以下的规则处理计时器：

- `timerWaiting` -> `timerModifying` -> `timerModifiedXX`
- `timerModifiedXX` -> `timerModifying` -> `timerModifiedYY`
- `timerNoStatus` -> `timerModifying` -> `timerWaiting`
- `timerRemoved` -> `timerModifying` -> `timerWaiting`
- `timerDeleted` -> `timerModifying` -> `timerWaiting`
- 其他状态 -> 等待状态改变

```go
// github.com/golang/go/src/runtime/time.go
func modtimer(t *timer, when, period int64, f func(interface{}, uintptr), arg interface{}, seq uintptr) bool {
	status := uint32(timerNoStatus)
	wasRemoved := false
loop:
	for {
		switch status = atomic.Load(&t.status); status {
			...
		}
	}

	t.period = period
	t.f = f
	t.arg = arg
	t.seq = seq

	if wasRemoved {
		t.when = when
		doaddtimer(pp, t)
		wakeNetPoller(when)
	} else {
		t.nextwhen = when
		newStatus := uint32(timerModifiedLater)
		if when < t.when {
			newStatus = timerModifiedEarlier
		}
		...
		if newStatus == timerModifiedEarlier {
			wakeNetPoller(when)
		}
	}
}
```

如果待修改的计时器已经被删除，那么该函数会调用 `runtime.doaddtimer` 创建新的计时器。在正常情况下会根据修改后的时间进行不同的处理：

- 如果修改后的时间大于或者等于修改前时间，设置计时器的状态为 `timerModifiedLater`；
- 如果修改后的时间小于修改前时间，设置计时器的状态为 `timerModifiedEarlier` 并调用 `runtime.netpollBreak` 触发调度器的重新调度；

因为修改后的时间会影响计时器的处理，所以用于修改计时器的 `runtime.modtimer` 也是状态机中最复杂的函数了。

#### 清除计时器

`runtime.cleantimers` 函数会根据状态清理处理器队列头中的计时器，该函数会遵循以下的规则修改计时器的触发时间：

- `timerDeleted` -> `timerRemoving` -> `timerRemoved`
- `timerModifiedXX` -> `timerMoving` -> `timerWaiting`

```go
// github.com/golang/go/src/runtime/time.go
func cleantimers(pp *p) bool {
	for {
		if len(pp.timers) == 0 {
			return true
		}
		t := pp.timers[0]
		switch s := atomic.Load(&t.status); s {
		case timerDeleted:
			atomic.Cas(&t.status, s, timerRemoving)
			dodeltimer0(pp)
			atomic.Cas(&t.status, timerRemoving, timerRemoved)
		case timerModifiedEarlier, timerModifiedLater:
			atomic.Cas(&t.status, s, timerMoving)

			t.when = t.nextwhen

			dodeltimer0(pp)
			doaddtimer(pp, t)
			atomic.Cas(&t.status, timerMoving, timerWaiting)
		default:
			return true
		}
	}
}
```

`runtime.cleantimers` 函数只会处理计时器状态为 `timerDeleted`、`timerModifiedEarlier` 和 `timerModifiedLater` 的情况：

- 如果计时器的状态为 timerDeleted；

  - 将计时器的状态修改成 `timerRemoving`；
  - 调用 `runtime.dodeltimer0` 删除四叉堆顶上的计时器；
  - 将计时器的状态修改成 `timerRemoved`；

- 如果计时器的状态为 timerModifiedEarlier 或者 timerModifiedLater ；

  - 将计时器的状态修改成 `timerMoving`；
  - 使用计时器下次触发的时间 `nextWhen` 覆盖 `when`；
  - 调用 `runtime.dodeltimer0` 删除四叉堆顶上的计时器；
  - 调用 `runtime.doaddtimer` 将计时器加入四叉堆中；
  - 将计时器的状态修改成 `timerWaiting`；

`runtime.cleantimers` 会删除已经标记的计时器，修改状态为 `timerModifiedXX` 的计时器。

#### 调整计时器

`runtime.adjusttimers` 与 `runtime.cleantimers` 的作用相似，它们都会删除堆中的计时器并修改状态为 `timerModifiedEarlier` 和 `timerModifiedLater` 的计时器的时间，它们也会遵循相同的规则处理计时器状态：

- `timerDeleted` -> `timerRemoving` -> `timerRemoved`
- `timerModifiedXX` -> `timerMoving` -> `timerWaiting`

```go
// github.com/golang/go/src/runtime/time.go
func adjusttimers(pp *p, now int64) {
	var moved []*timer
loop:
	for i := 0; i < len(pp.timers); i++ {
		t := pp.timers[i]
		switch s := atomic.Load(&t.status); s {
		case timerDeleted:
			// 删除堆中的计时器
		case timerModifiedEarlier, timerModifiedLater:
			// 修改计时器的时间
		case ...
		}
	}
	if len(moved) > 0 {
		addAdjustedTimers(pp, moved)
	}
}
```

与 `runtime.cleantimers` 不同的是，上述函数可能会遍历处理器堆中的全部计时器（包含退出条件），而不是只修改四叉堆顶部。

#### 运行计时器

`runtime.runtimer` 函数会检查处理器四叉堆上最顶上的计时器，该函数也会处理计时器的删除以及计时器时间的更新，它会遵循以下的规则处理计时器：

- `timerNoStatus` -> 崩溃：未初始化的计时器

- `timerWaiting`
  - -> `timerWaiting`
  - -> `timerRunning` -> `timerNoStatus`
  - -> `timerRunning` -> `timerWaiting`

- `timerModifying` -> 等待状态改变

- `timerModifiedXX` -> `timerMoving` -> `timerWaiting`

- `timerDeleted` -> `timerRemoving` -> `timerRemoved`

- `timerRunning` -> 崩溃：并发调用该函数

- `timerRemoved`、`timerRemoving`、`timerMoving` -> 崩溃：计时器堆不一致

```go
// github.com/golang/go/src/runtime/time.go
func runtimer(pp *p, now int64) int64 {
	for {
		t := pp.timers[0]
		switch s := atomic.Load(&t.status); s {
		case timerWaiting:
			if t.when > now {
				return t.when
			}
			atomic.Cas(&t.status, s, timerRunning)
			runOneTimer(pp, t, now)
			return 0
		case timerDeleted:
			// 删除计时器
		case timerModifiedEarlier, timerModifiedLater:
			// 修改计时器的时间
		case ...
		}
	}
}
```

如果处理器四叉堆顶部的计时器没有到触发时间会直接返回，否则调用 `runtime.runOneTimer` 运行堆顶的计时器：

```go
// github.com/golang/go/src/runtime/time.go
func runOneTimer(pp *p, t *timer, now int64) {
	f := t.f
	arg := t.arg
	seq := t.seq

	if t.period > 0 {
		delta := t.when - now
		t.when += t.period * (1 + -delta/t.period)
		siftdownTimer(pp.timers, 0)
		atomic.Cas(&t.status, timerRunning, timerWaiting)
		updateTimer0When(pp)
	} else {
		dodeltimer0(pp)
		atomic.Cas(&t.status, timerRunning, timerNoStatus)
	}

	unlock(&pp.timersLock)
	f(arg, seq)
	lock(&pp.timersLock)
}
```

根据计时器的 `period` 字段，上述函数会做出不同的处理：

- 如果period字段大于 0；

  - 修改计时器下一次触发的时间并更新其在堆中的位置；
  - 将计时器的状态更新至 `timerWaiting`；
  - 调用 `runtime.updateTimer0When` 函数设置处理器的 `timer0When` 字段；

- 如果period字段小于或者等于 0；

  - 调用 `runtime.dodeltimer0` 函数删除计时器；
- 将计时器的状态更新至 `timerNoStatus`；

更新计时器之后，上述函数会运行计时器中存储的函数并传入触发时间等参数。

### 触发计时器

已经分析了计时器状态机中的 10 种状态以及几种操作。这里将分析计时器的触发过程，Go 语言会在两个模块触发计时器，运行计时器中保存的函数：

- 调度器调度时会检查处理器中的计时器是否准备就绪；
- 系统监控会检查是否有未执行的到期计时器；

将依次分析上述这两个触发过程。

#### 调度器

`runtime.checkTimers` 是调度器用来运行处理器中计时器的函数，它会在发生以下情况时被调用：

- 调度器调用 `runtime.schedule` 执行调度时；
- 调度器调用 `runtime.findrunnable`获取可执行的 Goroutine 时；
- 调度器调用 `runtime.findrunnable`从其他处理器窃取计时器时；

这里不展开介绍 `runtime.schedule` 和 `runtime.findrunnable`的实现了，重点分析用于执行计时器的`runtime.checkTimers`，将该函数的实现分成调整计时器、运行计时器和删除计时器三个部分。

首先是**调整堆中计时器**的过程：

- 如果处理器中不存在需要调整的计时器；
  - 当没有需要执行的计时器时，直接返回；
  - 当下一个计时器没有到期并且需要删除的计时器较少时都会直接返回；
- 如果处理器中存在需要调整的计时器，会调用 `runtime.adjusttimers`；

```go
// github.com/golang/go/src/runtime/proc.go
func checkTimers(pp *p, now int64) (rnow, pollUntil int64, ran bool) {
	if atomic.Load(&pp.adjustTimers) == 0 {
		next := int64(atomic.Load64(&pp.timer0When))
		if next == 0 {
			return now, 0, false
		}
		if now == 0 {
			now = nanotime()
		}
		if now < next {
			if pp != getg().m.p.ptr() || int(atomic.Load(&pp.deletedTimers)) <= int(atomic.Load(&pp.numTimers)/4) {
				return now, next, false
			}
		}
	}

	lock(&pp.timersLock)
	adjusttimers(pp)
```

调整了堆中的计时器之后，会通过 `runtime.runtimer` 依次**查找堆中是否存在需要执行的计时器**：

- 如果存在，直接运行计时器；
- 如果不存在，获取最新计时器的触发时间；

```go
// github.com/golang/go/src/runtime/proc.go
  rnow = now
	if len(pp.timers) > 0 {
		if rnow == 0 {
			rnow = nanotime()
		}
		for len(pp.timers) > 0 {
			if tw := runtimer(pp, rnow); tw != 0 {
				if tw > 0 {
					pollUntil = tw
				}
				break
			}
			ran = true
		}
	}
```

在 `runtime.checkTimers` 的最后，如果当前 Goroutine 的处理器和传入的处理器相同，并且处理器中删除的计时器是堆中计时器的 1/4 以上，就会调用 `runtime.clearDeletedTimers` **删除处理器全部被标记为 `timerDeleted` 的计时器**，保证堆中靠后的计时器被删除。

```go
// github.com/golang/go/src/runtime/proc.go
	if pp == getg().m.p.ptr() && int(atomic.Load(&pp.deletedTimers)) > len(pp.timers)/4 {
		clearDeletedTimers(pp)
	}

	unlock(&pp.timersLock)
	return rnow, pollUntil, ran
}
```

`runtime.clearDeletedTimers` 能够避免堆中出现大量长时间运行的计时器，该函数和 `runtime.moveTimers` 是唯二会遍历计时器堆的函数。

#### 系统监控

系统监控函数 `runtime.sysmon` 也可能会触发函数的计时器，下面的代码片段中省略了大量与计时器无关的代码：

```go
// github.com/golang/go/src/runtime/proc.go
func sysmon() {
	...
	for {
		...
		now := nanotime()
		next, _ := timeSleepUntil()
		...
		lastpoll := int64(atomic.Load64(&sched.lastpoll))
		if netpollinited() && lastpoll != 0 && lastpoll+10*1000*1000 < now {
			atomic.Cas64(&sched.lastpoll, uint64(lastpoll), uint64(now))
			list := netpoll(0) // non-blocking - returns list of goroutines
			if !list.empty() {
				incidlelocked(-1)
				injectglist(&list)
				incidlelocked(1)
			}
		}
		if next < now {
			startm(nil, false)
		}
		...
}
```

1. 调用 `runtime.timeSleepUntil` 获取计时器的到期时间以及持有该计时器的堆；
2. 如果超过 10ms 的时间没有轮询，调用 `runtime.netpoll` 轮询网络；
3. 如果当前有应该运行的计时器没有执行，可能存在无法被抢占的处理器，这时应该启动新的线程处理计时器；

在上述过程中 `runtime.timeSleepUntil` 会遍历运行时的全部处理器并查找下一个需要执行的计时器。

### 小结

Go 语言的计时器在并发编程起到了非常重要的作用，它能够**提供比较准确的相对时间**，基于它的功能，标准库中还提供了定时器、休眠等接口能在 Go 语言程序中更好地**处理过期和超时等问题**。

标准库中的计时器在大多数情况下是能够正常工作并且高效完成任务的，但是在遇到极端情况或者性能敏感场景时，它可能没有办法胜任，而**在 10ms 的这个粒度中**，在社区中也没有找到能够使用的计时器实现，一些使用时间轮询算法的开源库也不能很好地完成这个任务。

### 历史变更

2021-01-05 更新：Go 1.15 修改并 合并了计时器处理的多个函数并改变了状态的迁移过程，这里删除了重置计数器的内容；



### 参考

- “runtime: switch to using new timer code” https://github.com/golang/go/commit/6becb033341602f2df9d7c55cc23e64b925bbee2
- jaypei/use_c_sleep.go · Gist https://gist.github.com/jaypei/5334115
- Alexander Morozov Vyacheslav Bakhmutov. Dec 4, 2016. “How Do They Do It: Timers in Go” https://blog.gopheracademy.com/advent-2016/go-timers/
- Russ Cox. January 26, 2017. “Proposal: Monotonic Elapsed Time Measurements in Go” https://go.googlesource.com/proposal/+/master/design/12914-monotonic.md

- Go 1.9 之前的计时器实现 https://github.com/golang/go/blob/093adeef4004fd029de1a8fd138802607265dc73/src/runtime/time.go 
- Aliaksandr Valialkin, Ian Lance Taylor. Jan 6, 2017. “runtime: improve timers scalability on multi-CPU systems” https://github.com/golang/go/commit/76f4fd8a5251b4f63ea14a3c1e2fe2e78eb74f81 
- Dmitry Vyukov. Apr 6, 2016. “runtime: make timers faster” https://github.com/golang/go/issues/6239#issuecomment-206361959 
- Dmitry Vyukov. Apr 6, 2016. “runtime: timer doesn’t scale on multi-CPU systems with a lot of timers #15133” https://github.com/golang/go/issues/15133#issuecomment-206376049 
- Go 1.10 ~ 1.13 的计时器实现 https://github.com/golang/go/blob/76f4fd8a5251b4f63ea14a3c1e2fe2e78eb74f81/src/runtime/time.go 
- “time: excessive CPU usage when using Ticker and Sleep”https://github.com/golang/go/issues/27707 
- Ian Lance Taylor. Apr 12, 2019. “runtime, time: remove old timer code” https://github.com/golang/go/commit/580337e268a0581bc537e67ca4005b7682be5d66 
- Ian Lance Taylor. “runtime: add new addtimer function” https://github.com/golang/go/commit/2e0aa581b4a2544249ad2f8e86e17204ca778ca7 
- Ian Lance Taylor. “runtime: add new deltimer function” https://github.com/golang/go/commit/7416315e3358b0bc2774c92f39d8f7c4b33790ad 
- Ian Lance Taylor. “runtime: add modtimer function” https://github.com/golang/go/commit/48eb79ec2197aeea0eb43597b00cad1ebcad61d2 
- Ian Lance Taylor. “runtime: add cleantimers function” https://github.com/golang/go/commit/466547014769bbdf7d5a62ca1019bf52d809dfcd 
- Ian Lance Taylor. “runtime: add adjusttimers function” https://github.com/golang/go/commit/220150ff3c03a0d2618093689ab129ab5ea7dc7b 
- Ian Lance Taylor. “runtime: add new runtimer function” https://github.com/golang/go/commit/432ca0ea83d12519004c6f7f7c1728410923987f 
- Ian Lance Taylor. “runtime: add netpollBreak” https://github.com/golang/go/commit/50f4896b72d16b6538178c8ca851b20655075b7f
- Ian Lance Taylor. “runtime: don’t panic on racy use of timers” https://github.com/golang/go/commit/98858c438016bbafd161b502a148558987aa44d5 



## Channel


























