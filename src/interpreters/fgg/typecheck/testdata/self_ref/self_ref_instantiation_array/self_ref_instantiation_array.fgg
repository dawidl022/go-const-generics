package main

type any interface {
}

type Arr[T any] [2]T

func main() {
	_ = Arr[Arr[int]]{Arr[int]{1, 2}, Arr[int]{3, 4}}
}
