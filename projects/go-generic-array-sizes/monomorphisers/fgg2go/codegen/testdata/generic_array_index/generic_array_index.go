package main

type Arr[T any, N const] [N]T

func main() {
	_ = Arr[int, 2]{1, 2}[1]
}
