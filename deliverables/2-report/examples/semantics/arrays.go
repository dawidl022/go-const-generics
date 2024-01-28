package main

import "fmt"

type Outer struct {
	inner   Inner
	someVal int
}

type Inner struct {
	matrix [2][2]int
}

func main() {
	x := Outer{Inner{[2][2]int{{1, 2}, {3, 4}}}, 5}
	y := x                            // trivial deep-copy
	fmt.Println(x == y)               // deep-comparison: prints "true"

	myMap := map[Outer]int{x: 42}     // structure can be used as map key
	fmt.Println(myMap[x])             // prints: 42
	fmt.Println(myMap[y])             // prints: 42 (y has same value as x)

	y.inner.matrix[0][1] = 10         // update y
	fmt.Println(x.inner.matrix[0][1]) // x remains unchanged: prints "2"
	fmt.Println(x == y)               // deep-comparison: prints "false"
	fmt.Println(myMap[x])             // prints: 42 (x remains unchanged)
	fmt.Println(myMap[y])             // prints: 0 (new y value not in map)
}
