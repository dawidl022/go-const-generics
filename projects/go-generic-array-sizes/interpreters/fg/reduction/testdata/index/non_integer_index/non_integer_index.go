package main

type Arr[2]int

func main() {
	_ = Arr{1, 2}[Arr{1, 2}]
}
