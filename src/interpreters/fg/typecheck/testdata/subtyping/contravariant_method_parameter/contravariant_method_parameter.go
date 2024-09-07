package main

type any interface {
}

type anyGetter interface {
	getAny(x int) any
}

type Foo struct {
}

func (f Foo) getAny(x any) any {
	return 42
}

type Bar struct {
	getter anyGetter
}

func main() {
	_ = Bar{Foo{}}
}
