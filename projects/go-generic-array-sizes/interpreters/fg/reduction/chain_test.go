package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/chains/index_of_field/index_of_field.go
var chainsIndexOfFieldGo []byte

func TestReduce_givenIndexOfField_reducesField(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsIndexOfFieldGo, "Arr{1, 2}[1]")
}

//go:embed testdata/chains/field_of_index/field_of_index.go
var chainsFieldOfIndexGo []byte

func TestReduce_givenFieldOfIndex_reducesIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsFieldOfIndexGo, "Structure{1, 2}.y")
}

//go:embed testdata/chains/index_of_index/index_of_index.go
var chainsIndexOfIndexGo []byte

func TestReduce_givenIndexOfIndex_reducesInnerIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsIndexOfIndexGo, "Arr{3, 4}[0]")
}

//go:embed testdata/chains/field_of_field/field_of_field.go
var chainsFieldOfFieldGo []byte

func TestReduce_givenFieldOfField_reducesInnerField(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsFieldOfFieldGo, "Foo{1}.y")
}

//go:embed testdata/chains/non_value_field/non_value_field.go
var chainsNonValueFieldGo []byte

func TestReduce_givenFieldOfNonValueField_reducesToFieldOfStructLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsNonValueFieldGo, "Foo{1}.x")
}

//go:embed testdata/chains/non_value_index/non_value_index.go
var chainsNonValueIndexGo []byte

func TestReduce_givenIndexOfNonValueIndex_reducesToIndexOfArrayLiteral(t *testing.T) {
	assertEqualAfterSingleReduction(t, chainsNonValueIndexGo, "Arr{1}[0]")
}

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	p, err := parseAndReduceOneStep(program)

	require.NoError(t, err)
	require.Equal(t, expected, p.Expression.String())
}
