package main

type Foo struct {
}

func (f Foo) Identity(x any) any {
	return x
}

func main() {
	_ = Foo{}.Identity(1)
}
