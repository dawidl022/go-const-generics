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

func assertEqualAfterSingleReduction(t *testing.T, program []byte, expected string) {
	p, err := parseAndReduceOneStep(program)

	require.NoError(t, err)
	require.Equal(t, expected, p.Expression.String())
}
