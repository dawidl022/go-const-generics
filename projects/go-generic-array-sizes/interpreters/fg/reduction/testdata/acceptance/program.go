package acceptance

type Foo struct {
	x int
	y Arr
}

func (f Foo) getY() Arr {
	return f.y
}

type Arr [2]int

func (a Arr) set(n int, x int) Arr {
	a[n] = x;
	return a
}

func (a Arr) first() int {
	return a[0]
}

func main() {
	_ = Arr{4, 5}.set(1, 6)[Foo{3, Arr{1, 2}}.getY().first()]
}
