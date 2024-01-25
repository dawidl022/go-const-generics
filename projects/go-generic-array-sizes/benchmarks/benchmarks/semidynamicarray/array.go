package semidynamicarray

type Array struct {
	arr [N]int
	len int
}

func (a *Array) Push(element int) {
	if a.len >= N {
		panic("array is full")
	}
	a.arr[a.len] = element
	a.len++
}

func (a *Array) Pop() int {
	if a.len == 0 {
		panic("array is empty")
	}
	a.len--
	return a.arr[a.len]
}

func (a *Array) Get(i int) int {
	if i >= a.len {
		panic("index out of bounds")
	}
	return a.arr[i]
}

func (a *Array) Len() int {
	return a.len
}
