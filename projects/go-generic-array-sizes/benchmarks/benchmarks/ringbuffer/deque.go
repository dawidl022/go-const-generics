package ringbuffer

type Deque struct {
	arr   [N + 1]int // "waste" a slot to detect fullness
	front int
	back  int
}

func (d *Deque) PushFront(el int) {
	if d.wrapped(d.front+1) == d.back { panic("deque is full") }
	d.arr[d.front] = el
	d.front = d.wrapped(d.front + 1)
}

func (d *Deque) PopFront() int {
	if d.front == d.back { panic("deque is empty") }
	d.front = d.wrapped(d.front - 1)
	return d.arr[d.front]
}

func (d *Deque) PushBack(element int) {
	if d.front == d.wrapped(d.back-1) { panic("deque is full") }
	d.back = d.wrapped(d.back - 1)
	d.arr[d.back] = element
}

func (d *Deque) PopBack() int {
	if d.front == d.back { panic("deque is empty") }
	el := d.arr[d.back]
	d.back = d.wrapped(d.back + 1)
	return el
}

func (d *Deque) wrapped(n int) int {
	if n < 0 { return N }
	if n > N { return 0 }
	return n
}
