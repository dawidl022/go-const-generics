package main

type any interface {
}

type Arr__3[T any] [3]T

type Arr__2[T any] [2]T

func main() {
	_ = Arr__2[Arr__3[int]]{Arr__3[int]{1, 2, 3}, Arr__3[int]{3, 5, 6}}
}
