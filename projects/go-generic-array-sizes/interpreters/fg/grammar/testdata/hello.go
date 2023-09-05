package main

type any interface{}

type AnyArray2 [2]any

type Foo struct {
	x int
	y string
}

type Fooer interface {
	foo(x Foo) any
}

func (this AnyArray2) First() any {
	return this[0]
}

func (_1 AnyArray2) Nth(_n int) any {
	return _1[_n]
}

func (this AnyArray2) foo(foo Foo) any {
	return foo.y
}

func (this AnyArray2) Length() int {
	return 98765432101
}

func (this AnyArray2) Set(i int, x any) AnyArray2 {
	this[i] = x;
	return this
}

func main() {
	_ = AnyArray2{1, 2}.Set(0, 3).First()
}
