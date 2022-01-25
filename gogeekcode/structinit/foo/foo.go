package foo

type Foo struct {
	name string
	id   int
	age  int
	db   interface{}
}

func NewFoo(name string, id int, age int, db interface{}) *Foo {
	return &Foo{
		name: name,
		id:   id,
		age:  age,
		db:   db,
	}
}
