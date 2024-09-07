package main

import (
	_ "embed"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
)

//go:embed testdata/deque.fgg
var dequeLibGo string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayDequeFG_givenEmptyDeque_popFrontReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Deque[6]{Arr[6]{0, 0, 0, 0, 0, 0}, Zero[6]{}, Zero[6]{}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PopFront()")
}

func TestArrayDequeFG_givenEmptyDeque_popBackReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Deque[6]{Arr[6]{0, 0, 0, 0, 0, 0}, Zero[6]{}, Zero[6]{}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PopBack()")
}

func TestArrayDequeFG_givenElementIsPushedFront_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushFront(1).GetFront()")
}

func TestArrayDequeFG_givenSoleElementIsPoppedFront_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushFront(1).PopFront().GetFront()")
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

func TestArrayDequeFG_givenElementIsPushedBack_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushBack(1).GetBack()")
}

func TestArrayDequeFG_givenSoleElementIsPoppedBack_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushBack(1).PopBack().GetBack()")
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

func TestArrayDequeFG_givenElementIsPushedFront_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushFront(1).GetBack()")
}

func TestArrayDequeFG_givenElementIsPushedBack_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1", "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
}

func TestArrayDequeFG_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(10).PopBack().GetBack()")
}

func TestArrayDequeFG_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopFront().GetFront()")
}

func TestArrayDequeFG_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopBack().PushFront(10).GetBack()")
}

func TestArrayDequeFG_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushBack(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopFront().PushBack(10).GetBack()")
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopBack().PushFront(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushFront(1).PopBack().PushBack(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "1",
		"EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertReducesTo(t, "10",
		"EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).GetBack()")
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

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToGetFront(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushBack(1).PushBack(2).PopFront().PopFront().GetFront()")
}

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToPopBack(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "0",
		"EmptyDequeFunc{}.call().PushFront(1).PushFront(2).PopBack().PopBack().GetBack()")
}

func TestArrayDequeFG_givenDequeFullFromFront_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Deque[6]{Arr[6]{0, 1, 2, 3, 4, 0}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}, "+
			"Zero[6]{}, Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushFront(10)")
}

func TestArrayDequeFG_givenDequeFullFromFront_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque[6]{Arr[6]{0, 1, 2, 3, 4, 0}, "+
		"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}, "+
		"Zero[6]{}, Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushBack(10)")
}

func TestArrayDequeFG_givenDequeFullFromBack_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque[6]{Arr[6]{0, 4, 3, 2, 1, 0}, Zero[6]{}, "+
		"Succ[6]{Zero[6]{}}, Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushBack(10)")
}

func TestArrayDequeFG_givenDequeFullFromBack_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t, "Deque[6]{Arr[6]{0, 4, 3, 2, 1, 0}, Zero[6]{}, "+
		"Succ[6]{Zero[6]{}}, Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		"EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushFront(10)")
}

func TestArrayDequeFG_givenDequeFullInMiddle_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Deque[6]{Arr[6]{2, 3, 4, 0, 0, 1}, "+
			"Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		`EmptyDequeFunc{}.call().PushBack(1).PushBack(2)
		.PopFront().PopFront()
		.PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4)
		.PushBack(10)`)
}

func TestArrayDequeFG_givenDequeFullInMiddle_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertReducesTo(t,
		"Deque[6]{Arr[6]{1, 0, 0, 4, 3, 2}, "+
			"Succ[6]{Succ[6]{Zero[6]{}}}, Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}, "+
			"Succ[6]{Succ[6]{Succ[6]{Succ[6]{Succ[6]{Zero[6]{}}}}}}}",
		`EmptyDequeFunc{}.call().PushFront(1).PushFront(2)
		.PopBack().PopBack()
		.PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4)
		.PushFront(10)`)
}

func assertReducesTo(t *testing.T, expected string, main string) {
	program := dequeLibGo + fmt.Sprintf(mainTemplate, main)
	res, err := entrypoint.Interpret(strings.NewReader(program), io.Discard, loop.UnboundedSteps)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}
