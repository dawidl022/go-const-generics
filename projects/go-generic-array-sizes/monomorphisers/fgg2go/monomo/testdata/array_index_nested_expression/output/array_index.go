package main

type Arr__2 [2]int

type Arr__1 [1]int

func main() {
	_ = Arr__2{1, 2}[Arr__1{1}[0]]
}
