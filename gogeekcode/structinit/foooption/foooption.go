package foooption

type Foo struct {
	name string
	id   int
	age  int
	db   interface{}
}

// FooOption 代表可选参数
type FooOption func(foo *Foo)

// WithName 代表Name为可选参数
func WithName(name string) FooOption {
	return func(foo *Foo) {
		foo.name = name
	}
}

// WithAge 代表age为可选参数
func WithAge(age int) FooOption {
	return func(foo *Foo) {
		foo.age = age
	}
}

// WithDB 代表db为可选参数
func WithDB(db interface{}) FooOption {
	return func(foo *Foo) {
		foo.db = db
	}
}

// NewFoo 代表初始化
func NewFoo(id int, options ...FooOption) *Foo {
	foo := &Foo{
		name: "default",
		id:   id,
		age:  10,
		db:   nil,
	}
	for _, option := range options {
		option(foo)
	}
	return foo
}
