package main

type any interface {
}

type Eq[T any] interface {
	equal(other T) int
}


type Foo[T Eq[Int]] struct {
}

func main() {
	_ = 0
}
