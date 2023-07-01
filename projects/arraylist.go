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

type ArrayList struct {
	array [2]any
}

func (this ArrayList) Map(f Function) List {
	return this.MapRec(0, f)
}

func (this ArrayList) MapRec(i int, f Function) ArrayList {
	if i == len(this.array) {
		return this
	}
	return this.Set(i, f.Apply(this.array[i])).MapRec(i+1, f)
}

func (this ArrayList) Set(i int, v any) ArrayList {
	this.array[i] = v
	return this
}

// len of an array is a constant expression
const arrLen = len([2]int{})

func main() {
	var l List = ArrayList{[arrLen]any{1, 2}}
	fmt.Printf("%#v\n", l.Map(negate{}))
}
