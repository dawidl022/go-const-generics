package ringbuffer

// TODO unit test this code

type ArrayDeque struct {
	arr   [N]int
	front int
	back  int
}

func (a *ArrayDeque) PushFront(element int) {
	a.arr[a.front] = element
	a.front--
	if a.front < 0 {
		a.front = N - 1
	}
}

func (a *ArrayDeque) PopFront() int {
	a.front++
	if a.front >= N {
		a.front = 0
	}
	return a.arr[a.front]
}

func (a *ArrayDeque) PushBack(element int) {
	a.arr[a.back] = element
	a.back++
	if a.back >= N {
		a.back = 0
	}
}

func (a *ArrayDeque) PopBack() int {
	a.back--
	if a.back < 0 {
		a.back = N - 1
	}
	return a.arr[a.back]
}
