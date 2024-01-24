package index

func first[T any, N const[1:]](arr [N]T) T {
	// allowed & safe: constraint guarantees array has at least 1 element
	return arr[0]
}
