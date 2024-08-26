package main

type Foo[F Foo[F]] interface{ m(f F) F }
type Bar[B Bar[B]] interface{ m(b B) B }

type E[F Foo[B], B Bar[F]] struct{}

func main() {}
