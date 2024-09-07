package werner

type array interface {
	[2]int | [3]int
}

func foo[A array](a A) {
	// compile error: len(a) (value of type int) is not constant
	const n = len(a)
}
