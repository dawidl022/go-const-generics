package semidynamicarray

type Array struct {
	arr [N]int
	len int
}

func (a *Array) push(element int) {
	a.arr[a.len] = element
	a.len++
}

func (a *Array) pop() {
	a.len--
}

func (a *Array) get(i int) int {
	if i >= a.len {
		panic("index out of bounds")
	}
	return a.arr[i]
}
