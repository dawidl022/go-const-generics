package main

type Arr [2]int

type Structure struct {
	x Arr
}

func main() {
	_ = Structure{Arr{1, 2}}.x[1]
}
