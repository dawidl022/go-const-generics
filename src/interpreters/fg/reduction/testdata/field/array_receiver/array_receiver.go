package main

type Foo [1]int

func main() {
	_ = Foo{1}.x
}
