package main

type any interface {
}

type Arr[N const, T any] [N]Arr[N, T]

func main() {
	_ = 0
}
