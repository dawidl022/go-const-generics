package main

type any interface {
}

type Box[T any] interface {
}

type Arr__2 [2]int

type Foo[T Box[Box[Arr__2]]] struct {
}

func main() {
	_ = Foo[int]{}
}
