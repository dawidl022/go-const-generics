package main

type any interface {
}

type mapper[T any, R any] interface {
	Map(x T) mapper[T, R]
}

func main() {
	_ = 1
}