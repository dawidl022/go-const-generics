package reduction

import (
	_ "embed"
	"testing"
)

//go:embed testdata/index/generic_receiver/generic_receiver.fgg
var indexGenericReceiverFgg []byte

func TestReduce_givenArrayIndexOnGenericReceiver_reducesToIndex(t *testing.T) {
	assertEqualAfterSingleReduction(t, indexGenericReceiverFgg,
		`2`)
}

//go:embed testdata/index/invalid_size_type_param/invalid_size_type_param.fgg
var indexInvalidSizeTypeParamFgg []byte

func TestReduce_givenArrayIndexOnValueLiteralWithInvalidConstTypeParam_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexInvalidSizeTypeParamFgg,
		`could not check index bounds of "Arr[int, int]{1, 2}": `+
			`badly instantiated type "Arr": `+
			`"int" is not a valid constant type parameter`)
}

//go:embed testdata/index/uninstantiated_generic_type/uninstantiated_generic_type.fgg
var indexUninstantiatedGenericTypeFgg []byte

func TestReduce_givenArrayIndexOnUninstantiatedGenericValueLiteral_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexUninstantiatedGenericTypeFgg,
		`could not check index bounds of "Arr{1, 2}": `+
			`badly instantiated type "Arr": `+
			`expected 2 type arguments but got 0`)
}

//go:embed testdata/index/out_of_bounds/out_of_bounds.fgg
var indexGenericOutOfBoundsFgg []byte

func TestReduce_givenArrayIndexOutOfBoundsOnGenericType_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexGenericOutOfBoundsFgg,
		`index 2 out of bounds for array of type "Arr[2, int]"`)
}

//go:embed testdata/index/unbound_length_parameter/unbound_length_parameter.fgg
var indexUnboundLengthParameterFgg []byte

func TestReduce_givenArrayGenericTypeHasUnboundLengthParameter_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexUnboundLengthParameterFgg,
		`could not check index bounds of "Arr[int]{1, 2}": `+
			`unexpected length type parameter "Unbound" in declaration of type "Arr"; `+
			`only integer literals and type parameters may be used as length types`)
}

//go:embed testdata/index/named_type_length_parameter/named_type_length_parameter.fgg
var indexNamedTypeLengthParameterFgg []byte

func TestReduce_givenArrayGenericTypeHasNonConstLengthParameter_returnsError(t *testing.T) {
	assertErrorAfterSingleReduction(t, indexNamedTypeLengthParameterFgg,
		`could not check index bounds of "Arr[int]{1, 2}": `+
			`unexpected length type parameter "Foo" in declaration of type "Arr"; `+
			`only integer literals and type parameters may be used as length types`)
}
