package main

type Empty interface {
	int
	string
}
type Uninstantiable[E Empty] struct{ x E }

func main() {}
