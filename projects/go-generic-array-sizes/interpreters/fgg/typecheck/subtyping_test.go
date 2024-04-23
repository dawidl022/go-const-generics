package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/subtyping/instantiated_type_implements/instantiated_type_implements.fgg
var subtypingInstantiatedTypeImplementsFgg []byte

func TestTypeCheck_givenInstantiatedTypeImplementsRequiredInterface_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingInstantiatedTypeImplementsFgg)
}

//go:embed testdata/subtyping/instantiated_type_not_implements/instantiated_type_not_implements.fgg
var subtypingInstantiatedTypeNotImplementsFgg []byte

func TestTypeCheck_givenInstantiatedTypeDoesNotImplementRequiredInterface_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingInstantiatedTypeNotImplementsFgg,
		`ill-typed declaration: method "Bar.getInterface": `+
			`return expression of type "Foo[Bar]" is not a subtype of "myInterface": `+
			`missing methods: "f(x int) int"`)
}

//go:embed testdata/subtyping/const_instantiated_type_implements/const_instantiated_type_implements.fgg
var subtypingConstInstantiatedTypeImplementsFgg []byte

func TestTypeCheck_givenConstInstantiatedTypeImplementsRequiredInterface_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingConstInstantiatedTypeImplementsFgg)
}

//go:embed testdata/subtyping/const_instantiated_type_not_implements/const_instantiated_type_not_implements.fgg
var subtypingConstInstantiatedTypeNotImplementsFgg []byte

func TestTypeCheck_givenConstInstantiatedTypeDoesNotImplementRequiredInterface_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingConstInstantiatedTypeNotImplementsFgg,
		`ill-typed declaration: method "Bar.getInterface": `+
			`return expression of type "Foo[5]" is not a subtype of "myInterface": `+
			`missing methods: "f(x IntArray[2]) IntArray[2]"`)
}

//go:embed testdata/subtyping/recursive_type_param/recursive_type_param.fgg
var subtypingRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingRecursiveTypeParameterFgg)
}

//go:embed testdata/subtyping/interface_recursive_type_param/interface_recursive_type_param.fgg
var subtypingInterfaceRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidInterfaceWithRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingInterfaceRecursiveTypeParameterFgg)
}

//go:embed testdata/subtyping/invalid_recursive_type_param/invalid_recursive_type_param.fgg
var subtypingInvalidRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenInvalidRecursiveTypeParameter_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingInvalidRecursiveTypeParameterFgg,
		`ill-typed main expression: type "comparableHolder" badly instantiated: `+
			`type "Str" is not a subtype of "Eq[Str]": `+
			`missing methods: "equal(other Str) int"`)
}

//go:embed testdata/subtyping/two_level_recursive_type_param/two_level_recursive_type_param.fgg
var subtypingTwoLevelRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidTwoLevelRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingTwoLevelRecursiveTypeParameterFgg)
}

// TODO consider if the following tests belong to the subtyping file

//go:embed testdata/subtyping/nested_type_parameter_bound/nested_type_parameter_bound.fgg
var subtypingNestedTypeParameterBoundFgg []byte

func TestTypeCheck_givenNestedNonRecursiveTypeParameterBound_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingNestedTypeParameterBoundFgg)
}

//go:embed testdata/subtyping/undefined_nested_type_parameter_bound/undefined_nested_type_parameter_bound.fgg
var subtypingUndefinedNestedTypeParameterBoundFgg []byte

func TestTypeCheck_givenUndefinedNestedTypeParameterBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingUndefinedNestedTypeParameterBoundFgg,
		`ill-typed declaration: type "Foo": `+
			`illegal bound of type parameter "T": `+
			`type "Eq" badly instantiated: type name not declared: "Int"`)
}
