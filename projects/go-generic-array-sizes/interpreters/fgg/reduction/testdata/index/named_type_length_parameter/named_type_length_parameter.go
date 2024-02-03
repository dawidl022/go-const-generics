package main

type Foo struct {
}

type Arr[T any] [Foo]T

func main() {
	_ = Arr[int]{1, 2}[1]
}
