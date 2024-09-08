package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-const-generics/interpreters/fg/entrypoint"
	fggEntrypoint "github.com/dawidl022/go-const-generics/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-const-generics/interpreters/shared/loop"
)

//go:embed deque.go
var dequeLibGo string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayDeque_givenEmptyDeque_popFrontReturnsEmptyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	assert.Equal(t, EmptyDequeFunc{}.call(), deque.PopFront())
}

func TestArrayDequeFG_givenEmptyDeque_popFrontReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 0, 0, 0, 0, 0}, Zero{}, Zero{}}",
		"EmptyDequeFunc{}.call().PopFront()")
}

func TestArrayDeque_givenEmptyDeque_popBackReturnsEmptyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	assert.Equal(t, EmptyDequeFunc{}.call(), deque.PopBack())
}

func TestArrayDequeFG_givenEmptyDeque_popBackReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 0, 0, 0, 0, 0}, Zero{}, Zero{}}",
		"EmptyDequeFunc{}.call().PopBack()")
}

func TestArrayDeque_givenElementIsPushedFront_gettingFrontReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetFront())
}

func TestArrayDequeFG_givenElementIsPushedFront_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushFront(1).GetFront()")
}

func TestArrayDeque_givenSoleElementIsPoppedFront_cannotBeGottenAgain(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	deque = deque.PopFront()
	assert.Equal(t, 0, deque.GetFront())
}

func TestArrayDequeFG_givenSoleElementIsPoppedFront_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushFront(1).PopFront().GetFront()")
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

func TestArrayDequeFG_givenTwoElementsPushedFront_arePoppedFromFrontInReverseOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).GetFront()")
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).PopFront().GetFront()")
	assertReducesTo(t, "0",
		`EmptyDequeFunc{}.call().PushFront(1).PushFront(10)
		.PopFront().PopFront().GetFront()`)
}

func TestArrayDeque_givenElementIsPushedBack_gettingBackReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetBack())
}

func TestArrayDequeFG_givenElementIsPushedBack_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushBack(1).GetBack()")
}

func TestArrayDeque_givenSoleElementIsPoppedBack_cannotBeGottenAgain(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	deque = deque.PopBack()
	assert.Equal(t, 0, deque.GetBack())
}

func TestArrayDequeFG_givenSoleElementIsPoppedBack_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushBack(1).PopBack().GetBack()")
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

func TestArrayDequeFG_givenTwoElementsPushedBack_arePoppedFromBackInReverseOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).GetBack()")
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopBack().GetBack()")
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopBack().PopBack().GetBack()")
}

func TestArrayDeque_givenElementIsPushedFront_gettingBackReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
}

func TestArrayDequeFG_givenElementIsPushedFront_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushFront(1).GetBack()")
}

func TestArrayDeque_givenElementIsPushedBack_gettingFrontReturnsThatElement(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
}

func TestArrayDequeFG_givenElementIsPushedBack_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
}

func TestArrayDeque_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushFront(1)
	deque = deque.PushFront(10)

	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDequeFG_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).PopBack().GetBack()")
}

func TestArrayDeque_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	deque := EmptyDequeFunc{}.call()
	deque = deque.PushBack(1)
	deque = deque.PushBack(10)

	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDequeFG_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopFront().GetFront()")
}

func TestArrayDeque_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDequeFG_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopBack().PushFront(10).GetBack()")
}

func TestArrayDeque_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDequeFG_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushBack(10).GetFront()")
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopFront().PushBack(10).GetBack()")
}

func TestArrayDeque_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopBack().PushFront(10).GetFront()")
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	assert.Equal(t, 1, deque.GetBack())
	deque = deque.PopBack()

	deque = deque.PushBack(10)
	assert.Equal(t, 10, deque.GetFront())
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopBack().PushBack(10).GetFront()")
}

func TestArrayDeque_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	assert.Equal(t, 1, deque.GetFront())
	deque = deque.PopFront()

	deque = deque.PushFront(10)
	assert.Equal(t, 10, deque.GetBack())
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).GetBack()")
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

func TestArrayDequeFG_givenAlternatingTwoPushesAndPopsOfOppositeDirections(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).GetBack()")
	assertReducesTo(t, "20",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).PopBack().PushFront(20).GetBack()")
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToGetFront(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushBack(1)
	deque = deque.PushBack(2)
	deque = deque.PopFront()
	deque = deque.PopFront()
	assert.Equal(t, 0, deque.GetFront())
}

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToGetFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(2).PopFront().PopFront().GetFront()")
}

func TestArrayDeque_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToPopBack(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	deque = deque.PushFront(1)
	deque = deque.PushFront(2)
	deque = deque.PopBack()
	deque = deque.PopBack()
	assert.Equal(t, 0, deque.GetBack())
}

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToPopBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(2).PopBack().PopBack().GetBack()")
}

func TestArrayDeque_givenDequeFullFromFront_pushFrontDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushFront(i)
	}
	assert.Equal(t, deque, deque.PushFront(10))
}

func TestArrayDequeFG_givenDequeFullFromFront_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 1, 2, 3, 4, 0}, Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}, Zero{}}",
		"EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushFront(10)")
}

func TestArrayDeque_givenDequeFullFromFront_pushBackDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushFront(i)
	}
	assert.Equal(t, deque, deque.PushBack(10))
}

func TestArrayDequeFG_givenDequeFullFromFront_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 1, 2, 3, 4, 0}, Succ{Succ{Succ{Succ{Succ{Zero{}}}}}}, Zero{}}",
		"EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushBack(10)")
}

func TestArrayDeque_givenDequeFullFromBack_pushBackDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushBack(i)
	}
	assert.Equal(t, deque, deque.PushBack(10))
}

func TestArrayDequeFG_givenDequeFullFromBack_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 4, 3, 2, 1, 0}, Zero{}, Succ{Zero{}}}",
		"EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushBack(10)")
}

func TestArrayDeque_givenDequeFullFromBack_pushFrontDoesNotModifyDeque(t *testing.T) {
	deque := EmptyDequeFunc{}.call()

	for i := 0; i < 5; i++ {
		deque = deque.PushBack(i)
	}
	assert.Equal(t, deque, deque.PushFront(10))
}

func TestArrayDequeFG_givenDequeFullFromBack_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{0, 4, 3, 2, 1, 0}, Zero{}, Succ{Zero{}}}",
		"EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushFront(10)")
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

func TestArrayDequeFG_givenDequeFullInMiddle_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{2, 3, 4, 0, 0, 1}, Succ{Succ{Succ{Zero{}}}}, Succ{Succ{Succ{Succ{Zero{}}}}}}",
		`EmptyDequeFunc{}.call().PushBack(1).PushBack(2)
		.PopFront().PopFront()
		.PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4)
		.PushBack(10)`)
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

func TestArrayDequeFG_givenDequeFullInMiddle_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque{Arr{1, 0, 0, 4, 3, 2}, Succ{Succ{Zero{}}}, Succ{Succ{Succ{Zero{}}}}}",
		`EmptyDequeFunc{}.call().PushFront(1).PushFront(2)
		.PopBack().PopBack()
		.PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4)
		.PushFront(10)`)
}

// tests both under FG and FGG interpreters
func assertReducesTo(t *testing.T, expected string, main string) {
	program := dequeLibGo + fmt.Sprintf(mainTemplate, main)

	res, err := entrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)

	res, err = fggEntrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
