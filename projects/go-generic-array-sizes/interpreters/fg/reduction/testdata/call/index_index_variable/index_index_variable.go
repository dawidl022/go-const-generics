package main

type Fib struct {
}

func (f Fib) nth(n int) int {
	return Arr{0, 1, 1, 2, 3}[n]
}

func main() {
	_ = Fib{}.nth(3)
}
