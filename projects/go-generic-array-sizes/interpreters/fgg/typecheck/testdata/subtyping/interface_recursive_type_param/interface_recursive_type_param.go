package main

type any interface{}

type Eq[T any] interface {
	equal(other T) int
}

type selfComparable[T Eq[T]] interface {
	equal(other T) int
}

type comparableHolder[T selfComparable[T]] struct {
	x T
}

type Int struct {
}

func (i Int) equal(other Int) int {
	return 0
}

func main() {
	_ = comparableHolder[Int]{Int{}}
}
