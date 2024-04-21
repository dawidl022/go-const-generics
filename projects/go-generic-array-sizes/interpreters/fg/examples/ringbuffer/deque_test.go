package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDeque_givenEmptyDeque_popFrontReturnsEmptyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	assert.Equal(t, EmptyDequeFunc{}.call(), deque.PopFront())
}

func TestArrayDeque_givenEmptyDeque_popBackReturnsEmptyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	assert.Equal(t, EmptyDequeFunc{}.call(), deque.PopBack())
}

func TestArrayDeque_givenElementIsPushedFront_gettingFrontReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetFront())
}

func TestArrayDeque_givenSoleElementIsPoppedFront_cannotBeGottenAgain(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	deque = deque.PopFront()
	assert.Equal(t, 0, deque.GetFront())
}

func TestArrayDeque_givenTwoElementsPushedFront_arePoppedFromFrontInReverseOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	deque = deque.PushFront(10)

	assert.Equal(t, 10, deque.GetFront())
	deque = deque.PopFront()

	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	assert.Equal(t, 0, deque.GetFront())
}

func TestArrayDeque_givenElementIsPushedBack_gettingBackReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetBack())
}

func TestArrayDeque_givenSoleElementIsPoppedBack_cannotBeGottenAgain(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	deque = deque.PopBack()
	assert.Equal(t, 0, deque.GetBack())
}

func TestArrayDeque_givenTwoElementsPushedBack_arePoppedFromBackInReverseOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	deque = deque.PushBack(10)

	assert.Equal(t, 10, deque.GetBack())
	deque = deque.PopBack()

	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	assert.Equal(t, 0, deque.GetBack())
}

func TestArrayDeque_givenElementIsPushedFront_gettingBackReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
}

func TestArrayDeque_givenElementIsPushedBack_gettingFrontReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
}

func TestArrayDeque_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	deque = deque.PushFront(10)

	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDeque_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	deque = deque.PushBack(10)

	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDeque_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDeque_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDeque_givenAlternatingTwoPushesAndPopsOfOppositeDirections(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushFront(20)
	assert.Equal(t, 20, deque.GetBack())
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToGetFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	deque = deque.PushBack(2)
	deque = deque.PopFront()
	deque = deque.PopFront()
	assert.Equal(t, 0, deque.GetFront())
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToPopBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	deque = deque.PushFront(2)
	deque = deque.PopBack()
	deque = deque.PopBack()
	assert.Equal(t, 0, deque.GetBack())
}

func TestArrayDeque_givenDequeFullFromFront_pushFrontDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushFront(i)
	}
	assert.Equal(t, deque, deque.PushFront(10))
}

func TestArrayDeque_givenDequeFullFromFront_pushBackDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushFront(i)
	}
	assert.Equal(t, deque, deque.PushBack(10))
}

func TestArrayDeque_givenDequeFullFromBack_pushBackDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushBack(i)
	}
	assert.Equal(t, deque, deque.PushBack(10))
}

func TestArrayDeque_givenDequeFullFromBack_pushFrontDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushBack(i)
	}
	assert.Equal(t, deque, deque.PushFront(10))
}

func TestArrayDeque_givenDequeFullInMiddle_pushBackDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	deque = deque.PushBack(2)
	deque = deque.PopFront()
	deque = deque.PopFront()

	for i := 0; i < 5; i++ {
		deque = deque.PushFront(i)
	}
	assert.Equal(t, deque, deque.PushBack(10))
}

func TestArrayDeque_givenDequeFullInMiddle_pushFrontDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	deque = deque.PushFront(2)
	deque = deque.PopBack()
	deque = deque.PopBack()

	for i := 0; i < 5; i++ {
		deque = deque.PushBack(i)
	}
	assert.Equal(t, deque, deque.PushFront(10))
}
