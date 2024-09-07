package expressions

func newArrStrange[T any, N const, M const](n int) [N + M]T {
	if n == 0 {
		return [N + M]T{}
	}
	return newArrStrange[T, N + 1, M - 1]()
}
