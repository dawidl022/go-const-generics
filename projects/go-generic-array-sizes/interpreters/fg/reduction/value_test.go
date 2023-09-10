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
	p := parseFGProgram(valueArrayGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Arr{1, 2}", p.Expression.Value().String())
}

//go:embed testdata/value/matrix/matrix.go
var valueMatrixGo []byte

func TestValue_givenNestedArrLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueMatrixGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Matrix{Arr{1, 2}, Arr{3, 4}}", p.Expression.Value().String())
}

//go:embed testdata/value/struct/struct.go
var valueStructGo []byte

func TestValue_givenStructLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueStructGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Foo{1, 2}", p.Expression.Value().String())
}

//go:embed testdata/value/nested/nested.go
var valueNestedGo []byte

func TestValue_givenNestedStructLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueNestedGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Bar{Foo{1, 2}, Foo{3, 4}}", p.Expression.Value().String())
}

//go:embed testdata/value/array_of_structs/array_of_structs.go
var valueArrayOfStructsGo []byte

func TestValue_givenArrayOfStructsLiteral_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueArrayOfStructsGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Arr{Foo{1, 2}, Foo{3, 4}}", p.Expression.Value().String())
}

//go:embed testdata/value/struct_of_arrays/struct_of_arrays.go
var valueStructOfArraysGo []byte

func TestValue_givenStructOfArrayLiterals_returnsValueAndFailsTo_Reduce(t *testing.T) {
	p := parseFGProgram(valueStructOfArraysGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, "Foo{Arr{1, 2}, Arr{3, 4}}", p.Expression.Value().String())
}
