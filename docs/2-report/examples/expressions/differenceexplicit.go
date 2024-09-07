package expressions

func difference[T any, N const[M:], M const](a [N]T, b [M]T) [N - M]T {
	// some presumably useful code...
}
