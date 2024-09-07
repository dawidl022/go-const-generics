package main

type Foo struct {
	bar Bar
}

type Bar struct {
	baz Baz
}

type Baz struct {
	foo Foo
}

func main() {
	_ = 0
}
