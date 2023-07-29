package main

import "fmt"

type List interface {
	Reverse() List
}

type AnyArray2 [2]any

func (this AnyArray2) Reverse() List {
	return this.ReverseRec(AnyArray2{}, 0)
}

func (this AnyArray2) ReverseRec(target AnyArray2, i int) AnyArray2 {
	if i == 2 {
		return target
	}
	return this.ReverseRec(target.Set(2-i-1, (this[i])), i+1)
}

func (this AnyArray2) Set(i int, v any) AnyArray2 {
	this[i] = v
	return this
}

func main() {
	var l List = AnyArray2{1, 2}
	fmt.Printf("%#v\n", l.Reverse())
}
