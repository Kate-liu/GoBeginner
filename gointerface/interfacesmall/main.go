package main

// 会飞的
type Flyable interface {
	Fly()
}

// 会游泳的
type Swimable interface {
	Swim()
}

// 会飞且会游泳的
type FlySwimable interface {
	Flyable
	Swimable
}
