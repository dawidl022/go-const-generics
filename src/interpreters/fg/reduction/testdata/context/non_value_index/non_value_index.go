package main

type Arr [1]int

func main() {
	_ = Arr{Arr{1}[0]}[0]
}
