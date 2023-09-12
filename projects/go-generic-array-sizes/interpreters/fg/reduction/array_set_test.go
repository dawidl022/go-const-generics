package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/array_set/basic/basic.go
var arraySetBasicGo []byte

func TestReduceArraySet_reducesToArrayLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetBasicGo, "Arr{1, 3}")
}

//go:embed testdata/array_set/out_of_bounds/out_of_bounds.go
var arraySetOutOfBoundsGo []byte

func TestReduceArraySet_withIndexOutOfBounds_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(arraySetOutOfBoundsGo)

	require.Error(t, err)
	require.Equal(t, `array set index 2 out of bounds for array of type "Arr"`, err.Error())
}

//go:embed testdata/array_set/non_integer_index/non_integer_index.go
var arraySetNonIntegerIndexGo []byte

func TestReduceArraySet_withNonIntegerValue_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(arraySetNonIntegerIndexGo)

	require.Error(t, err)
	require.Equal(t, `non-integer index "Arr{1, 2}" in array set method call: Arr.Set`, err.Error())
}

//go:embed testdata/array_set/insufficient_arguments/insufficient_arguments.go
var arraySetInsufficientArgumentsGo []byte

func TestReduceArraySet_withLessArgumentsThanNecessary_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(arraySetInsufficientArgumentsGo)

	require.Error(t, err)
	require.Equal(t, `expected 2 arguments in call to "Arr.Set", but got 0`, err.Error())
}

//go:embed testdata/array_set/extraneous_arguments/extraneous_arguments.go
var arraySetExtraneousArgumentsGo []byte

func TestReduceArraySet_withMoreArgumentsThanNecessary_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(arraySetExtraneousArgumentsGo)

	require.Error(t, err)
	require.Equal(t, `expected 2 arguments in call to "Arr.Set", but got 3`, err.Error())
}

//go:embed testdata/array_set/expression_index/expression_index.go
var arraySetExpressionIndexGo []byte

func TestReduceArraySet_withNonValueIndexArgument_reducesArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetExpressionIndexGo, "Arr{1, 2}.Set(1, 3)")
}

//go:embed testdata/array_set/expression_value/expression_value.go
var arraySetExpressionValueGo []byte

func TestReduceArraySet_withNonValueValueArgument_reducesArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetExpressionValueGo, "Arr{1, 2}.Set(0, 2)")
}

//go:embed testdata/array_set/expressions/expressions.go
var arraySetExpressionsGo []byte

func TestReduceArraySet_withBothArgumentsNonValues_reducesFirstArgument(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetExpressionsGo, "Arr{1, 2}.Set(1, Arr{1, 2}[0])")
}

//go:embed testdata/array_set/expression_receiver/expression_receiver.go
var arraySetExpressionReceiverGo []byte

func TestReduceArraySet_withNonValueReceiver_reducesReceiver(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetExpressionReceiverGo, "Arr{1, 2}.Set(Arr{1, 2}[0], 3)")
}
