package main

type Foo[F Foo[F]] interface{ m(f F) F }
type Bar[B Bar[B]] interface{ m(b B) B }

type T struct{}

func (t T) m(t2 T) T { return t }

type E[F Foo[F], B Bar[B]] struct{}

func main() {
	_ = E[T, T]{}
}
