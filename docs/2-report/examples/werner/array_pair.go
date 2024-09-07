package werner

// Example 1: syntactically same constraints
type ArrayPair[T any, A [...]T, B [...]T] struct {
	left: A,
	right: B,
}
