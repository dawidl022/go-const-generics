package main

import "fmt"

type Function interface {
	Apply(x any) any
}

type negate struct{}

func (negate) Apply(x any) any {
	return -x.(int)
}

type List interface {
	Map(f Function) List
}

// len of an array is a constant expression
const arrLen = len([2]int{})

type AnyArray2 [arrLen]any

func (this AnyArray2) Map(f Function) List {
	return this.MapRec(0, f)
}

func (this AnyArray2) MapRec(i int, f Function) AnyArray2 {
	if i == len(this) {
		return this
	}
	return this.Set(i, f.Apply(this[i])).MapRec(i+1, f)
}

func (this AnyArray2) Set(i int, v any) AnyArray2 {
	this[i] = v
	return this
}

func main() {
	var l List = AnyArray2{1, 2}
	fmt.Printf("%#v\n", l.Map(negate{}))
}
