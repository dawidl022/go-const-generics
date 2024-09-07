package main

type any interface {
}

type Arr__2[T any] [2]T

type Foo__2 struct {
	x Arr__2[int]
}

func main() {
	_ = Foo__2{Arr__2[int]{1, 2}}.x
}
