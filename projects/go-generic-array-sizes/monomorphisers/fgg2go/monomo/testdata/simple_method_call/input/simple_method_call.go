package main

type Foo[N const] struct {
}

func (f Foo[N]) foo() int {
	return 0
}

func main() {
	_ = Foo[2]{}.foo()
}
