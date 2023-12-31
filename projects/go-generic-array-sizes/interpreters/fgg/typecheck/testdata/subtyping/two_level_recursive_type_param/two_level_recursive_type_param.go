package main

type any interface{}

type Eq[T any] interface {
	equal(other T) int
}

type comparableHolder[T Eq[T]] struct {
}

func (c comparableHolder[T]) equal(other comparableHolder[T]) int {
	return 0
}

type Int struct {
}

func (i Int) equal(other Int) int {
	return 0
}

func main() {
	_ = comparableHolder[comparableHolder[Int]]{}
}
