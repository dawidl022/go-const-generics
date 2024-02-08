package main

type Fooer interface {
	foo(f Fooer) Fooer
}

func main() {
	_ = 0
}
