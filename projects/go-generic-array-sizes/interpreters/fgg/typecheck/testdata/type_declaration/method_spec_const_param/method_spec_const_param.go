package main

type fooer[N const] interface {
	foo(x int, y N) int
}

func main() {
	_ = 1
}
