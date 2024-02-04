package main

type any interface {
}

type Arr[N const, T any] [N]T

func main() {
	_ = Arr[2, Arr[3, int]]{Arr[3, int]{1, 2, 3}, Arr[3, int]{3, 5, 6}}
}
