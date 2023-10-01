package main

type Arr[N const, T any] [N]T

func main() {
	_ = Arr{1, 2}[1]
}
