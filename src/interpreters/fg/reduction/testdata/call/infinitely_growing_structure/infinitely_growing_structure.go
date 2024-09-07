package main

type List interface {
}

type Node struct {
	val  int
	next List
}

type Null struct {
}

func (n Node) grow() Node {
	return Node{0, Node{0, Null{}}.grow()}
}

func main() {
	_ = Node{0, Null{}}.grow()
}
