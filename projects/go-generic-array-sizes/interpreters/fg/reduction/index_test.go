package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/index/basic/basic.go
var indexBasicGo []byte

func TestReduceIndex_givenBasicValidExpression_returnsValueAtIndex(t *testing.T) {
	val, err := parseAndReduceOneStep(indexBasicGo)

	require.NoError(t, err)
	require.Equal(t, "1", val.String())
}

//go:embed testdata/index/multiple_indices/multiple_indices.go
var indexMultipleIndicesGo []byte

func TestReduceIndex_givenBasicValidExpressionWithMultipleIndices_returnsValueAtIndex(t *testing.T) {
	val, err := parseAndReduceOneStep(indexMultipleIndicesGo)

	require.NoError(t, err)
	require.Equal(t, "2", val.String())
}

//go:embed testdata/index/incomplete_literal/incomplete_literal.go
var indexIncompleteLiteralGo []byte

// note that an incomplete array literal is valid Go - but FG does not allow it
func TestReduceIndex_givenArrayLiteralWithLessElementsThanDeclared_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(indexIncompleteLiteralGo)

	require.Error(t, err)
	require.Equal(t, "array literal missing value at index 2", err.Error())
}

//go:embed testdata/index/out_of_bounds/out_of_bounds.go
var indexOutOfBoundsGo []byte

func TestReduceIndex_givenArrayIndexOutOfBounds_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(indexOutOfBoundsGo)

	require.Error(t, err)
	require.Equal(t, `index 1 out of bounds for array of type "Arr"`, err.Error())
}

//go:embed testdata/index/non_array/non_array.go
var indexNonArrayGo []byte

func TestReduceIndex_givenIndexOnStruct_returnsError(t *testing.T) {
	_, err := parseAndReduceOneStep(indexNonArrayGo)

	require.Error(t, err)
	require.Equal(t, `no array type named "Foo" found in declarations`, err.Error())
}

//go:embed testdata/index/array_value/array_value.go
var indexArrayValueGo []byte

func TestReduceIndex_givenValidExpressionWithArrayElement_returnsArrayValue(t *testing.T) {
	val, err := parseAndReduceOneStep(indexArrayValueGo)

	require.NoError(t, err)
	require.Equal(t, "Arr{3, 4}", val.String())
}

//go:embed testdata/index/struct_value/struct_value.go
var indexStructValueGo []byte

func TestReduceIndex_givenValidExpressionWithStructElement_returnsStructValue(t *testing.T) {
	val, err := parseAndReduceOneStep(indexStructValueGo)

	require.NoError(t, err)
	require.Equal(t, "Structure{3, 4}", val.String())
}
