package main

type Arr [2]int

func (a Arr) First() int {
	return a[0]
}
func main() {
	_ = Arr{1, 2}.First()
}
