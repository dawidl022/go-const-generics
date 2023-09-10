package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

//go:embed testdata/value/int/int.go
var valueIntGo []byte

func TestValue_givenInt_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueIntGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, ast.IntegerLiteral{IntValue: 1}, p.Expression.Value())
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

func assertEqualValueAndFailsToReduce(t *testing.T, program []byte, expectedValue string) {
	p := parseFGProgram(program)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, expectedValue, p.Expression.Value().String())
}
