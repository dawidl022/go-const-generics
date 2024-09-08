package main

type Foo [F Foo [F]] interface { m(x F) F }
type Bar [B Bar [B]] interface { m(x B) B }

type E[F Foo [F] , B Bar [B]] struct {}

type T struct {}
func (t T) m(x T) T { return t }

func main () {
	_ = E[T, T]{}
}
