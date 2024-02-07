package main

type Foo struct {
	bar Bar
}

type Bar struct {
	baz Baz
}

type Baz struct {
	bar Bar
}

func main() {
	_ = 0
}
