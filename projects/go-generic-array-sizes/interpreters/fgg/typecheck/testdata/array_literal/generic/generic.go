package main

type any interface {
}

type Arr[N const, T any] [N]T

func main() {
	_ = Arr[2, int]{1, 2}
}
