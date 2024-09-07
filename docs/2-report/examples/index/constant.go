package index

func first[T any, N const](arr [N]T) T {
	// not allowed: constant index must be valid for all array lengths
	return arr[0]
}
