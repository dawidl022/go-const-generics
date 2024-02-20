package main

type any interface {
}

type Arr__2[T any] [2]T

type Arr__3[T any] [3]T

type Foo__2__3 interface {
	foo(a Arr__2[int]) Arr__3[int]
}

type Foo__3__2 interface {
	foo(a Arr__3[int]) Arr__2[int]
}

type FooImpl__2__3 struct {
	x Arr__3[int]
}

type FooImpl__3__2 struct {
	x Arr__2[int]
}

func (f FooImpl__2__3) foo(a Arr__2[int]) Arr__3[int] {
	return f.x
}

func (f FooImpl__3__2) foo(a Arr__3[int]) Arr__2[int] {
	return f.x
}

func (f FooImpl__2__3) getFoo() Foo__2__3 {
	return f
}

func (f FooImpl__3__2) getFoo() Foo__3__2 {
	return f
}

func (f FooImpl__2__3) makeSwappedFoo(a Arr__2[int]) Foo__3__2 {
	return FooImpl__3__2{a}
}

func (f FooImpl__3__2) makeSwappedFoo(a Arr__3[int]) Foo__2__3 {
	return FooImpl__2__3{a}
}

func main() {
	_ = FooImpl__2__3{Arr__3[int]{1, 2, 3}}.foo(Arr__2[int]{1, 2})
}
