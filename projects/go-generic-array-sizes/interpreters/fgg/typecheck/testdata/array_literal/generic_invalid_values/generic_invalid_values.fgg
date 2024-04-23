package main

type any interface {
}

type foo struct {
}

type Arr[N const, T any] [N]T

func main() {
	_ = Arr[2, int]{1, foo{}}
}
