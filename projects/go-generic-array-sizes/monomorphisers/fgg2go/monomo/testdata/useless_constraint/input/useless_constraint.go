package main

type Useless[N const] interface {
}

type Arr[N const, T Useless[N]] [N]T

func main() {
	_ = Arr[2, int]{1, 2}
}
