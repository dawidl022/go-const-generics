package main

type any interface {
}

type Box[T any] interface {
}

type Arr [N const] [N]int

type Foo[T Box[Box[Arr[2]]]] struct{
}

func main() {
	_ = Foo[int]{}
}
