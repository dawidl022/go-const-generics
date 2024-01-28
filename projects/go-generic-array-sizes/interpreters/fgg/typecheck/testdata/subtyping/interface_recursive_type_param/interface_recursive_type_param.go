package main

type any interface{}

type Eq[T any] interface {
	equal(other T) int
}

type SelfEq[T Eq[T]] interface {
	equal(other T) int
}

type ComparableHolder[T SelfEq[T]] struct {
	x T
}

type Int struct {
}

func (i Int) equal(other Int) int {
	return 0
}

func main() {
	_ = ComparableHolder[Int]{Int{}}
}
