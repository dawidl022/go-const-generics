package main

type Eq[T Eq[T]] interface {
	equal(other T) int
}

func main() {
	_ = 0
}
