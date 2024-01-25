package main

import "fmt"

type Outer struct {
	inner   Inner
	someVal int
}

type Inner struct {
	matrix [][]int
}

func main() {
	x := Outer{Inner{[][]int{{1, 2}, {3, 4}}}, 5}
	y := x                            // non-deep copy: matrix is shared
	y.inner.matrix[0][1] = 10         // update y
	fmt.Println(x.inner.matrix[0][1]) // x also changed: prints "10"

	fmt.Println(x == y)               // invalid operation: uncomparable type
	_ = map[Outer]int{x: 42}          // invalid map key type Outer
}
