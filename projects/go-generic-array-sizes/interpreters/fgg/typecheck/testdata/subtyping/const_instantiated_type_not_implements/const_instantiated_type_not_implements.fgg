package main

type IntArray[N const] [N]int

type Foo[N const] struct {
}

func (f Foo[N]) f(x IntArray[N]) IntArray[N] {
	return x
}

type myInterface interface {
	f(x IntArray[2]) IntArray[2]
}

type Bar struct {
}

func (b Bar) getInterface() myInterface {
	return Foo[5]{}
}

func main() {
	_ = Bar{}.getInterface()
}
