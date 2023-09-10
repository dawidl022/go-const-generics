package main

type Arr [2]int

func main() {
	_ = Arr{Arr{1, 2}[0], Arr{1, 2}[1]}
}
