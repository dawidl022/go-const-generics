package main

type Foo struct {
	bar Bar
}

type Bar struct {
	foo Foo
}

func main() {
	_ = 0
}
