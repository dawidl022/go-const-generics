package ringbuffer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDeque_givenEmptyDeque_popFrontPanics(t *testing.T) {
	deque := ArrayDeque{}
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopFront() })
}

func TestArrayDeque_givenEmptyDeque_popBackPanics(t *testing.T) {
	deque := ArrayDeque{}
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopBack() })
}

func TestArrayDeque_givenElementIsPushedFront_poppingFrontReturnsThatElement(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	assert.Equal(t, 1, deque.PopFront())
}

func TestArrayDeque_givenSoleElementIsPoppedFront_cannotBePoppedAgain(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	deque.PopFront()
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopFront() })
}

func TestArrayDeque_givenTwoElementsPushedFront_arePoppedFromFrontInReverseOrder(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	deque.PushFront(10)
	assert.Equal(t, 10, deque.PopFront())
	assert.Equal(t, 1, deque.PopFront())
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopFront() })
}

func TestArrayDeque_givenElementIsPushedBack_poppingBackReturnsThatElement(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopBack())
}

func TestArrayDeque_givenSoleElementIsPoppedBack_cannotBePoppedAgain(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	deque.PopBack()
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopBack() })
}

func TestArrayDeque_givenTwoElementsPushedBack_arePoppedFromBackInReverseOrder(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	deque.PushBack(10)
	assert.Equal(t, 10, deque.PopBack())
	assert.Equal(t, 1, deque.PopBack())
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopBack() })
}

func TestArrayDeque_givenElementIsPushedFront_poppingBackReturnsThatElement(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	assert.Equal(t, 1, deque.PopBack())
}

func TestArrayDeque_givenElementIsPushedBack_poppingFrontReturnsThatElement(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopBack())
}

func TestArrayDeque_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	deque.PushFront(10)
	assert.Equal(t, 1, deque.PopBack())
	assert.Equal(t, 10, deque.PopBack())
}

func TestArrayDeque_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	deque.PushBack(10)
	assert.Equal(t, 1, deque.PopFront())
	assert.Equal(t, 10, deque.PopFront())
}

func TestArrayDeque_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	assert.Equal(t, 1, deque.PopBack())
	deque.PushFront(10)
	assert.Equal(t, 10, deque.PopBack())
}

func TestArrayDeque_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopFront())
	deque.PushBack(10)
	assert.Equal(t, 10, deque.PopFront())
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	assert.Equal(t, 1, deque.PopFront())
	deque.PushBack(10)
	assert.Equal(t, 10, deque.PopBack())
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopBack())
	deque.PushFront(10)
	assert.Equal(t, 10, deque.PopFront())
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	assert.Equal(t, 1, deque.PopBack())
	deque.PushBack(10)
	assert.Equal(t, 10, deque.PopFront())
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopFront())
	deque.PushFront(10)
	assert.Equal(t, 10, deque.PopBack())
}

func TestArrayDeque_givenAlternatingTwoPushesAndPopsOfOppositeDirections(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	assert.Equal(t, 1, deque.PopFront())
	deque.PushFront(10)
	assert.Equal(t, 10, deque.PopBack())
	deque.PushFront(20)
	assert.Equal(t, 20, deque.PopBack())
	deque.PushBack(30)
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_panicsWhenTryingToPopFront(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PopFront()
	deque.PopFront()
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopFront() })
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_panicsWhenTryingToPopBack(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PopBack()
	deque.PopBack()
	assert.PanicsWithValue(t, "deque is empty", func() { deque.PopBack() })
}

func TestArrayDeque_givenDequeFullFromFront_pushFrontPanics(t *testing.T) {
	deque := ArrayDeque{}

	for i := 0; i < N; i++ {
		deque.PushFront(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushFront(0) })
}

func TestArrayDeque_givenDequeFullFromFront_pushBackPanics(t *testing.T) {
	deque := ArrayDeque{}

	for i := 0; i < N; i++ {
		deque.PushFront(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushBack(0) })
}

func TestArrayDeque_givenDequeFullFromBack_pushBackPanics(t *testing.T) {
	deque := ArrayDeque{}

	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushBack(0) })
}

func TestArrayDeque_givenDequeFullFromBack_pushFrontPanics(t *testing.T) {
	deque := ArrayDeque{}

	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushFront(0) })
}

func TestArrayDeque_givenDequeFullInMiddle_pushBackPanics(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushBack(1)
	deque.PushBack(2)
	deque.PopFront()
	deque.PopFront()

	for i := 0; i < N; i++ {
		deque.PushFront(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushFront(0) })
}

func TestArrayDeque_givenDequeFullInMiddle_pushFrontPanics(t *testing.T) {
	deque := ArrayDeque{}
	deque.PushFront(1)
	deque.PushFront(2)
	deque.PopBack()
	deque.PopBack()

	for i := 0; i < N; i++ {
		deque.PushBack(i)
	}
	assert.PanicsWithValue(t, "deque is full", func() { deque.PushFront(0) })
}
