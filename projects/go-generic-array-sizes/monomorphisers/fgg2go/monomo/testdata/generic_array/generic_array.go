package main

type any interface {
}

type Arr[T any] [2]T

func main() {
	_ = Arr[int]{1, 2}
}
