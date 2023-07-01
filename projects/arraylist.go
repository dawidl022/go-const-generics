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

type ArrayList [arrLen]any

func (this ArrayList) Map(f Function) List {
	return this.MapRec(0, f)
}

func (this ArrayList) MapRec(i int, f Function) ArrayList {
	if i == len(this) {
		return this
	}
	return this.Set(i, f.Apply(this[i])).MapRec(i+1, f)
}

func (this ArrayList) Set(i int, v any) ArrayList {
	this[i] = v
	return this
}

func main() {
	var l List = ArrayList{1, 2}
	fmt.Printf("%#v\n", l.Map(negate{}))
}
