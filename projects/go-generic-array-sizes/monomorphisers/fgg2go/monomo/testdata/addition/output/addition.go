package main

type Arr__2 [2]int

type Arr__3 [3]int

func main() {
	_ = Arr__2{1, 2}[0] + Arr__3{1, 2, 3}[1]
}
