package main

type Edge[E any, V any] interface {
	Source() V
	Target() V
}

type Vertex[E any, V any] interface {
	Edges() []E
}

type Edge2[E Edge[E, V], V Vertex[E, V]] interface {
	Edge[E, V]
}

type Vertex2[E Edge[E, V], V Vertex[E, V]] interface {
	Vertex[E, V]
}

type graphHolder[E Edge[E, V], V Vertex[E, V]] struct {
	edges    []E
	vertices []V
}

func main() {
	_ = 0
}
