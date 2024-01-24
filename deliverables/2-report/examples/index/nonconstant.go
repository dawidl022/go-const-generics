package index

func first[T any, N const](arr [N]T) T {
	i := 0
	// allowed: non-constant index bounds checks are performed at runtime
	return arr[i]
}
