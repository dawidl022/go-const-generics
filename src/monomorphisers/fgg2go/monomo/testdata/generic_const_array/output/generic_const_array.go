package main

type any interface {
}

type Arr__2[T any] [2]T

func main() {
	_ = Arr__2[int]{1, 2}
}
