package main

type xGetter interface {
	getX() int
}

type Arr [2]int

func (a Arr) getX() int {
	return 42
}

func (a Arr) getGetter() xGetter {
	return a
}

func main() {
	_ = 0
}
