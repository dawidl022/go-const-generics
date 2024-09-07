package main

type any interface {}

type Eq[T any] interface {
	equal(other T) int
}

type comparableHolder[T Eq[T]] struct {
}

type Int struct {
}

func (i Int) equal(other Int) int {
	return 0
}

type Str struct {
}

func (s Str) equal(other Int) int {
	return 0
}

func main() {
	_ = comparableHolder[Str]{}
}
