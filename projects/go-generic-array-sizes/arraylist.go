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

type AnyArray2 [2]any

func (this AnyArray2) Map(f Function) List {
	return this.MapRec(AnyArray2{}, 0, f)
}

func (this AnyArray2) MapRec(target AnyArray2, i int, f Function) AnyArray2 {
	if i == 2 {
		return target
	}
	return this.MapRec(target.Set(i, f.Apply(this[i])), i+1, f)
}

func (this AnyArray2) Set(i int, v any) AnyArray2 {
	this[i] = v
	return this
}

func main() {
	var l List = AnyArray2{1, 2}
	fmt.Printf("%#v\n", l.Map(negate{}))
}
