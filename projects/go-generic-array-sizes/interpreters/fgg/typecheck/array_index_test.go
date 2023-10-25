package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/array_index/generic/generic.go
var arrayIndexGenericFgg []byte

func TestTypeCheck_givenArrayIndexIntoGenericArrayTypeWithConcreteLength_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayIndexGenericFgg)
}

//go:embed testdata/array_index/generic_literal_index/generic_literal_index.go
var arrayIndexGenericLiteralIndexFgg []byte

func TestTypeCheck_givenArrayIndexWithIntLiteralIntoGenericArrayTypeWithConcreteLength_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayIndexGenericLiteralIndexFgg)
}

//go:embed testdata/array_index/generic_out_of_bounds/generic_literal_index.go
var arrayIndexGenericOutOfBoundsFgg []byte

func TestTypeCheck_givenArrayIndexWithOutOfBoundsIntLiteralIntoGenericArrayTypeWithConcreteLength_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexGenericOutOfBoundsFgg,
		`ill-typed declaration: method "Arr.getInt": `+
			`cannot access index 2 on array of type "Arr[2, int]" of size 2`)
}

//go:embed testdata/array_index/non_concrete_length/non_concrete_length.go
var arrayIndexNonConcreteLengthFgg []byte

func TestTypeCheck_givenArrayIndexIntoGenericArrayWithTypeParamLength_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, arrayIndexNonConcreteLengthFgg)
}

//go:embed testdata/array_index/non_concrete_length_literal_index/non_concrete_length_literal_index.go
var arrayIndexNonConcreteLengthLiteralIndexFgg []byte

func TestTypeCheck_givenArrayIndexWithIntLiteralIntoGenericArrayWithTypeParamLength_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arrayIndexNonConcreteLengthLiteralIndexFgg,
		`ill-typed declaration: method "Arr.get": `+
			`cannot use int literal "0" to index into array of type "Arr[N, T]" with non-concrete length`)
}
