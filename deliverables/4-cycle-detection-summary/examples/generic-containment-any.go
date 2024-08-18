package main

import "fmt"

// Let's remember Bar in the type parameter constraint is syntactic sugar for
// an anonymous type set interface: interface { Bar }. The only type
// this constraint can be instantiated with is Bar (struct type).
//
// This code compiles in this declaration order, but fails to compile if
// we reverse the order of struct declarations.
//
// The code also (rightfully so) fails to compile if we add a field of type b,
// as this would violate the notContains constraint. However e.g. a field of type
// pointer to b is fine.
//
// I.e. by introducing type set interfaces, we need to extend notContains
// checks to type parameters in generic types.
type Foo[T any] struct {
	x T
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
