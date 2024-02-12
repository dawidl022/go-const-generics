package main

type any interface{}

type Eq[T any] interface {
	equal(other T) int
}

type SelfEq[T Eq[T]] interface {
	equal(other T) int
}

// the simple Eq constraint seems sufficient to specify the self-eq property
type ComparableHolder[T Eq[T]] struct {
	x T
}

type Int struct {
}

func (i Int) equal(other Int) int {
	return 0
}

type String struct {
}

func (i String) equal(other Int) int {
	return 0
}

func main() {
	_ = ComparableHolder[Int]{Int{}}       // compiles
	_ = ComparableHolder[Int]{String{}}    // does not compile
	_ = ComparableHolder[String]{String{}} // does not compile
	var _ SelfEq[Int] = String{}           // compiles
	var _ SelfEq[String] = String{}        // does mot compile
}
