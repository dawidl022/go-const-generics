package main

type Structure struct {
	x int
	y int
}

type Arr [2]Structure

func main() {
	_ = Arr{Structure{1, 2}, Structure{3, 4}}[1]
}
