package main

type fooer interface {
	foo() int
}

type Arr[N const, T fooer] [N]T

func main() {
	_ = Arr[2, int]{1, 2}
}
