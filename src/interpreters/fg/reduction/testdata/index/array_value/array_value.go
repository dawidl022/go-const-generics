package main

type Arr [2]int

type Matrix [2]Arr

func main() {
	_ = Matrix{Arr{1, 2}, Arr{3, 4}}[1]
}
