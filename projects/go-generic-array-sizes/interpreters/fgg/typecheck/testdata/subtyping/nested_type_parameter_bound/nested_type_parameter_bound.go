package main

type any interface {
}

type Eq[T any] interface {
	equal(other T) int
}

type Int struct {
}

func (t Int) equal(other Int) int {
	return 0
}

type Foo[T Eq[Int]] struct {
}

func main() {
	_ = Foo[Int]{}
}
