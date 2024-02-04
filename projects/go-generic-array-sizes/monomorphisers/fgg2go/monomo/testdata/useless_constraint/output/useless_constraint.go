package main

type Useless__2 interface {
}

type Arr__2[T Useless__2] [2]T

func main() {
	_ = Arr__2[int]{1, 2}
}
