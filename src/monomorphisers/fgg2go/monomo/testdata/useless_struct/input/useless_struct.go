package main

type Foo[N const] struct {
}

func main() {
	_ = Foo[5]{}
}
