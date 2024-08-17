package main

// Let's take an example of types that cannot be instantiated, yet the current
// compiler implementation permits their declaration.
//
// If we cannot instantiate Foo, when it recursively refers to struct type Bar,
// then we would encounter the same issue with a constraint of the
// form [a Foo[a]]. The compiler currently rejects the latter case due to
// the notReferenced rule.
//
// If the goal of cycle detection is to prevent declaring non-instantiable types,
// then the rules should be amended to prevent scenarios such as the below.
type Foo[a Bar[a]] struct {
	x a
}

type Bar[b any] struct {
}

func main() {
}
