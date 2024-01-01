package main

type any interface {
}

type Eq[T Eq[T]] interface {
	equal(other T) int
}

func main() {
	_ = 0
}
