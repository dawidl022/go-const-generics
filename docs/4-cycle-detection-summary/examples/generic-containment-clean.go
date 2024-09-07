package main

import "fmt"

type Foo[T Bar] struct {
	x *T
}

type Bar struct {
	x Foo[Bar]
}

func main() {
	x := Foo[Bar]{}
	y := Bar{x: x}
	fmt.Println(x)
	fmt.Println(y)
}
