package main

type Arr [1]Structure

type Structure struct {
	x int
	y int
}

func main() {
	_ = Arr{Structure{1, 2}}[0].y
}
