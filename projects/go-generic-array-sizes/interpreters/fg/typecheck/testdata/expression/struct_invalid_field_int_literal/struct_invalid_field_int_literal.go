package main

type Foo struct {
	x Bar
}

type Bar struct {
}

func main() {
	_ = Foo{1}
}
