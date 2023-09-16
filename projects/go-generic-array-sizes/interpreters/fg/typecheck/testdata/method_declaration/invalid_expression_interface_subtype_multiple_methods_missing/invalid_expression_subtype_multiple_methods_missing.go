package main


type Calculator interface {
	Max(x int, y int) int
	Add(x int, y int) int
	Square(x int) int
	Zero() int
}

type Foo struct {
}

func (f Foo) Max(x int, y int) int {
	return x
}

func (f Foo) GetCalculator() Calculator {
	return f
}

func main() {
	_ = 0
}
