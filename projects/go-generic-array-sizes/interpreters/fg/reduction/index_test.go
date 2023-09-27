package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/index/basic/basic.go
var indexBasicGo []byte

func TestReduceIndex_givenBasicValidExpression_returnsValueAtIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, indexBasicGo, "1")
}

//go:embed testdata/index/multiple_indices/multiple_indices.go
var indexMultipleIndicesGo []byte

func TestReduceIndex_givenBasicValidExpressionWithMultipleIndices_returnsValueAtIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, indexMultipleIndicesGo, "2")
}

//go:embed testdata/index/incomplete_literal/incomplete_literal.go
var indexIncompleteLiteralGo []byte

// note that an incomplete array literal is valid Go - but FG does not allow it
func TestReduceIndex_givenArrayLiteralWithLessElementsThanDeclared_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexIncompleteLiteralGo,
		"array literal missing value at index 2")
}

//go:embed testdata/index/out_of_bounds/out_of_bounds.go
var indexOutOfBoundsGo []byte

func TestReduceIndex_givenArrayIndexOutOfBounds_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexOutOfBoundsGo,
		`index 1 out of bounds for array of type "Arr"`)
}

//go:embed testdata/index/struct_receiver/struct_receiver.go
var indexNonArrayGo []byte

func TestReduceIndex_givenIndexOnStruct_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexNonArrayGo,
		`no array type named "Foo" found in declarations`)
}

//go:embed testdata/index/array_value/array_value.go
var indexArrayValueGo []byte

func TestReduceIndex_givenValidExpressionWithArrayElement_returnsArrayValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, indexArrayValueGo, "Arr{3, 4}")
}

//go:embed testdata/index/struct_value/struct_value.go
var indexStructValueGo []byte

func TestReduceIndex_givenValidExpressionWithStructElement_returnsStructValue(t *testing.T) {
	assertEqualAfterSingleReduction(t, indexStructValueGo, "Structure{3, 4}")
}

//go:embed testdata/index/non_integer_index/non_integer_index.go
var indexNonIntegerIndexGo []byte

func TestReduceIndex_givenNonIntegerIndexArgument_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexNonIntegerIndexGo,
		`non integer value "Arr{1, 2}" used as index argument`)
}

//go:embed testdata/index/integer_receiver/integer_receiver.go
var indexIntegerReceiverGo []byte

func TestReduceIndex_givenIntegerLiteralReceiver_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexIntegerReceiverGo,
		`cannot access index on primitive value 1`)
}
