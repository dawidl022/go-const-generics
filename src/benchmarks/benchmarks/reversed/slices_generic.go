package reversed

func reversedSliceGeneric[T any](s []T) []T {
	n := len(s)
	newS := make([]T, n)
	for i := 0; i < n/2; i++ {
		newS[i], newS[n-i-1] = s[n-i-1], s[i]
	}
	return newS
}
