package main

type xGetter interface {
	getX() int
	getX(y int) int
}

func main() {
	_ = 0
}
