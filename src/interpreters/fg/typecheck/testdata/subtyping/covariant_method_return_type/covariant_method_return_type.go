package main

type any interface {
}

type anyGetter interface {
	getAny() any
}

type Foo struct {
}

func (f Foo) getAny() int {
	return 42
}

type Bar struct {
	getter anyGetter
}

func main() {
	_ = Bar{Foo{}}
}
