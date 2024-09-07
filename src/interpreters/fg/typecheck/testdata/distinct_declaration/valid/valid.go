package main

type any interface {}

type xGetter interface {
	getX() any
}

type Int struct {
	val int
}

type Foo struct {
	x int
	y int
}

func (f Foo) getX() any {
	return f.x
}

type Bar struct {
	x Foo
}

func (b Bar) getX() any {
	return b.x
}

type Arr [2]int

func (a Arr) first() int {
	return a[0]
}

type ArrOfBars [2]Bar

func main() {
	_ = 1
}
