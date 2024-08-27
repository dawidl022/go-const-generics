package main

type Foo[F Foo[F]] interface{ m(f F) F }
type Bar[B Bar[B]] interface{ m(b B) B }

type E[F Foo[F], B Bar[B]] struct{}

type T struct{}

func (t T) m(t2 T) T { return t }

func main() {
	_ = E[T, T]{}
}
