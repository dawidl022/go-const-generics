package main

type IntGetter interface {
	GetInt(x int) int
}

type Foo struct {
}

func (f Foo) GetInt(x int) int {
	return x
}

func (f Foo) GetIntGetter() IntGetter {
	return f
}

func main() {
	_ = 0
}
