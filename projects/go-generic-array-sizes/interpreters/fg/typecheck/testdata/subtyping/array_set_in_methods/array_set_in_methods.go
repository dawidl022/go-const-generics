package main

type settable interface {
	set(x int, y any) Arr
}

type any interface {
}

type Arr [2]any

func (a Arr) set(x int, y any) Arr {
	a[x] = y;
	return a
}

func (a Arr) getSettable() settable {
	return a
}

func main() {
	_ = 0
}
