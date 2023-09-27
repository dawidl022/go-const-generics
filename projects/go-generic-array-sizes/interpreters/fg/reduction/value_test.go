package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/value/int/int.go
var valueIntGo []byte

func TestValue_givenInt_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueIntGo, "1")
}

//go:embed testdata/value/array/array.go
var valueArrayGo []byte

func TestValue_givenArrayLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueArrayGo, "Arr{1, 2}")

}

//go:embed testdata/value/matrix/matrix.go
var valueMatrixGo []byte

func TestValue_givenNestedArrLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueMatrixGo, "Matrix{Arr{1, 2}, Arr{3, 4}}")
}

//go:embed testdata/value/struct/struct.go
var valueStructGo []byte

func TestValue_givenStructLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueStructGo, "Foo{1, 2}")
}

//go:embed testdata/value/nested/nested.go
var valueNestedGo []byte

func TestValue_givenNestedStructLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueNestedGo, "Bar{Foo{1, 2}, Foo{3, 4}}")
}

//go:embed testdata/value/array_of_structs/array_of_structs.go
var valueArrayOfStructsGo []byte

func TestValue_givenArrayOfStructsLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueArrayOfStructsGo, "Arr{Foo{1, 2}, Foo{3, 4}}")
}

//go:embed testdata/value/struct_of_arrays/struct_of_arrays.go
var valueStructOfArraysGo []byte

func TestValue_givenStructOfArrayLiterals_returnsValueAndFailsTo_Reduce(t *testing.T) {
	assertEqualValueAndFailsToReduce(t, valueStructOfArraysGo, "Foo{Arr{1, 2}, Arr{3, 4}}")
}

//go:embed testdata/value/non_value_field/non_value_field.go
var valueNonValueFieldGo []byte

func TestReduce_givenStructOfFieldSelect_reducesToStructLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, valueNonValueFieldGo, "Foo{1}")
}

//go:embed testdata/value/non_value_index/non_value_index.go
var valueNonValueIndexGo []byte

func TestReduce_givenArrayOfArrayIndex_reducesToArrayLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, valueNonValueIndexGo, "Arr{1, 2}")
}

//go:embed testdata/value/multiple_non_value_fields/multiple_non_value_fields.go
var valueMultipleNonValueFields []byte

func TestReduce_givenStructOfMultipleFieldSelects_reducesOnlyFirstFieldSelect(t *testing.T) {
	assertEqualAfterSingleReduction(t, valueMultipleNonValueFields, "Foo{1, Foo{1, 2}.y}")
}

//go:embed testdata/value/multiple_non_values_indices/multiple_non_value_indices.go
var valueMultipleNonValueIndices []byte

func TestReduce_givenArrayOfMultipleArrayIndices_reducesOnlyFirstArrayIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, valueMultipleNonValueIndices, "Arr{1, Arr{1, 2}[1]}")
}

//go:embed testdata/value/unbound_variable/unbound_variable.go
var valueUnboundVariable []byte

func TestReduceCall_givenUnboundIndexInMain_failsToReduceOrYieldValue(t *testing.T) {
	p := parseFGProgram(valueUnboundVariable)
	require.Nil(t, p.Expression.Value())

	assertErrorAfterSingleReduction(t, valueUnboundVariable,
		`unbound variable "x"`)
}
