package main

type Arr[T any] [Unbound]T

func main() {
	_ = Arr[int]{1, 2}[1]
}
