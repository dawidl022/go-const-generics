package werner

// Example 2: same constraints using shorthand
type ArrayPair[T any, A, B [...]T] struct {
	left: A,
	right: B,
}
