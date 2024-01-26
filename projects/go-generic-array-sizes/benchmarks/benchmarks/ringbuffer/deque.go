package ringbuffer

type ArrayDeque struct {
	arr   [N + 1]int // "waste" a slot to detect fullness
	front int
	back  int
}

func (a *ArrayDeque) PushFront(element int) {
	if a.wrapped(a.front+1) == a.back { panic("deque is full") }
	a.arr[a.front] = element
	a.front = a.wrapped(a.front + 1)
}

func (a *ArrayDeque) PopFront() int {
	if a.front == a.back { panic("deque is empty") }
	a.front = a.wrapped(a.front - 1)
	return a.arr[a.front]
}

func (a *ArrayDeque) PushBack(element int) {
	if a.front == a.wrapped(a.back-1) { panic("deque is full") }
	a.back = a.wrapped(a.back - 1)
	a.arr[a.back] = element
}

func (a *ArrayDeque) PopBack() int {
	if a.front == a.back { panic("deque is empty") }
	el := a.arr[a.back]
	a.back = a.wrapped(a.back + 1)
	return el
}

func (a *ArrayDeque) wrapped(n int) int {
	if n < 0 { return N }
	if n > N { return 0 }
	return n
}
