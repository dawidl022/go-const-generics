package main

type A interface{ M(B[T]) T }
type B[a A] struct{}

type T struct{}

func (t T) M(b B[T]) T { return t }

func main() {}
