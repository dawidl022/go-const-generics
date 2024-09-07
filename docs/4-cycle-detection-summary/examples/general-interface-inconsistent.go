package main

type Bar interface {
	x() Baz[Bar]
}

type Baz[T Foo] interface {
}

type BarImpl struct {
}

func (b *BarImpl) x() Baz[Bar] {
	return nil
}

// if we move Foo to the start of decls we get a cycle error
type Foo interface {
	Bar
}

func main() {
}
