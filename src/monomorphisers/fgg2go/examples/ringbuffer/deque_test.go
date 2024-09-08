package main

import (
	_ "embed"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-const-generics/monomorphisers/fgg2go/monomo"
)

//go:embed testdata/deque.fgg
var dequeLibFgg string

const mainTemplate = `func main() {
	_ = %s
}`

func TestArrayDequeFG_givenEmptyDeque_popFrontReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PopFront()")
}

func TestArrayDequeFG_givenEmptyDeque_popBackReturnsEmptyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PopBack()")
}

func TestArrayDequeFG_givenElementIsPushedFront_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).GetFront()")
}

func TestArrayDequeFG_givenSoleElementIsPoppedFront_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PopFront().GetFront()")
}

func TestArrayDequeFG_givenTwoElementsPushedFront_arePoppedFromFrontInReverseOrder(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PushFront(10).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PushFront(10).PopFront().GetFront()")
	assertBisimulates(t, `EmptyDequeFunc{}.call().PushFront(1).PushFront(10)
		.PopFront().PopFront().GetFront()`)
}

func TestArrayDequeFG_givenElementIsPushedBack_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetBack()")
}

func TestArrayDequeFG_givenSoleElementIsPoppedBack_cannotBeGottenAgain(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopBack().GetBack()")
}

func TestArrayDequeFG_givenTwoElementsPushedBack_arePoppedFromBackInReverseOrder(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(10).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopBack().GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopBack().PopBack().GetBack()")
}

func TestArrayDequeFG_givenElementIsPushedFront_gettingBackReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).GetBack()")
}

func TestArrayDequeFG_givenElementIsPushedBack_gettingFrontReturnsThatElement(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
}

func TestArrayDequeFG_givenTwoElementsArePushedFront_arePoppedBackInSameOrder(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PushFront(10).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PushFront(10).PopBack().GetBack()")
}

func TestArrayDequeFG_givenTwoElementsArePushedBack_arePoppedFrontInSameOrder(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(10).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(10).PopFront().GetFront()")
}

func TestArrayDequeFG_givenAlternatingFrontPushesAndBackPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PopBack().PushFront(10).GetBack()")
}

func TestArrayDequeFG_givenAlternatingBackPushesAndFrontPops_popsPushedElements(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopFront().PushBack(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithFront(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PopFront().PushBack(10).GetBack()")
}

func TestArrayDequeFG_givenAlternatingPushedAndPops_startingWithBack(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopBack().PushFront(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithFront(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PopBack().PushBack(10).GetFront()")
}

func TestArrayDequeFG_givenAlternatingPushesAndPopsOfOppositeDirections_startingWithBack(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).GetBack()")
}

func TestArrayDequeFG_givenAlternatingTwoPushesAndPopsOfOppositeDirections(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).GetFront()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).GetBack()")
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PopFront().PushFront(10).PopBack().PushFront(20).GetBack()")
}

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToGetFront(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(1).PushBack(2).PopFront().PopFront().GetFront()")
}

func TestArrayDequeFG_givenDequeIsEmptiedAtMiddle_returnsZeroWhenTryingToPopBack(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(1).PushFront(2).PopBack().PopBack().GetBack()")
}

func TestArrayDequeFG_givenDequeFullFromFront_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushFront(10)")
}

func TestArrayDequeFG_givenDequeFullFromFront_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4).PushBack(10)")
}

func TestArrayDequeFG_givenDequeFullFromBack_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushBack(10)")
}

func TestArrayDequeFG_givenDequeFullFromBack_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, "EmptyDequeFunc{}.call().PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4).PushFront(10)")
}

func TestArrayDequeFG_givenDequeFullInMiddle_pushBackDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, `EmptyDequeFunc{}.call().PushBack(1).PushBack(2)
		.PopFront().PopFront()
		.PushFront(0).PushFront(1).PushFront(2).PushFront(3).PushFront(4)
		.PushBack(10)`)
}

func TestArrayDequeFG_givenDequeFullInMiddle_pushFrontDoesNotModifyDeque(t *testing.T) {
	t.Parallel()
	assertBisimulates(t, `EmptyDequeFunc{}.call().PushFront(1).PushFront(2)
		.PopBack().PopBack()
		.PushBack(0).PushBack(1).PushBack(2).PushBack(3).PushBack(4)
		.PushFront(10)`)
}

func assertBisimulates(t *testing.T, main string) {
	program := dequeLibFgg + fmt.Sprintf(mainTemplate, main)
	interpreter := entrypoint.Interpreter{}

	fggProg, err := interpreter.ParseProgram(strings.NewReader(program))
	require.NoError(t, err)

	fggMonoProg := entrypoint.FggProgram{monomo.Monomorphise(fggProg.Program)}

	for !fggProg.Expression().IsValue() {
		fggProg, err = interpreter.Reduce(fggProg)
		require.NoError(t, err)

		fggMonoProg, err = interpreter.Reduce(fggMonoProg)
		require.NoError(t, err)

		expectedMonoProg := monomo.Monomorphise(fggProg.Program)

		// monomorphise is called on monomorphised-line only for dead-type elimination!
		fggMonoProg = entrypoint.FggProgram{monomo.Monomorphise(fggMonoProg.Program)}

		require.Equal(t, expectedMonoProg, fggMonoProg.Program)
	}
}
