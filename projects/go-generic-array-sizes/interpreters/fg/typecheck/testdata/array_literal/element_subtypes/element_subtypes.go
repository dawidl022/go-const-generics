package main

type any interface {
}

type Arr [2]any

func main() {
	_ = Arr{1, Arr{1, 2}}
}
