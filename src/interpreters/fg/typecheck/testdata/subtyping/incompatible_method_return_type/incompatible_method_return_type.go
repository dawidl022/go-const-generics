package main

type any interface {
}

type anyGetter interface {
	getAny() int
}

type Foo struct {
}

func (f Foo) getAny() Foo {
	return f
}

type Bar struct {
	getter anyGetter
}

func main() {
	_ = Bar{Foo{}}
}
