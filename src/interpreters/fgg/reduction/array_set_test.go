package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_set/generic_receiver/generic_receiver.fgg
var arraySetGenericReceiverFgg []byte

func TestReduce_givenArraySetOnGenericReceiver_returnsReducedArrayWithNewValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, arraySetGenericReceiverFgg, `Arr[3, int]{1, 2, 4}`)
}

//go:embed testdata/array_set/out_of_bounds/out_of_bounds.fgg
var arraySetOutOfBoundsFgg []byte

func TestReduce_giveArraySetOutOfBoundsOnGenericReceiver_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, arraySetOutOfBoundsFgg,
		`array set index 3 out of bounds for array of type "Arr[3, int]"`)
}

//go:embed testdata/array_set/invalid_length_type_arg/invalid_length_type_arg.fgg
var arraySetInvalidLengthTypeArgFgg []byte

func TestReduce_givenArrayOnGenericReceiverWithInvalidLengthTypeArgument(t *testing.T) {
	assertErrorAfterSingleReduction(t, arraySetInvalidLengthTypeArgFgg,
		`badly instantiated type "Arr": "int" is not a valid constant type parameter`)
}
