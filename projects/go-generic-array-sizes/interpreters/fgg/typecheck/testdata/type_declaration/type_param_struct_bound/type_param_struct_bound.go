package main

type Foo struct {
}

type Arr[T Foo] [2]T

func main() {
	_ = 1
}