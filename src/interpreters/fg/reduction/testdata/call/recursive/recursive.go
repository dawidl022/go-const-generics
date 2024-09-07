package main

type Foo struct {
}

func (f Foo) recurse() int {
	return f.recurse()
}

func main() {
	_ = Foo{}.recurse()
}
