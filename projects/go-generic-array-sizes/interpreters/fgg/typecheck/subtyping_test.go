package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/subtyping/instantiated_type_implements/instantiated_type_implements.go
var subtypingInstantiatedTypeImplementsFgg []byte

func TestTypeCheck_givenInstantiatedTypeImplementsRequiredInterface_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingInstantiatedTypeImplementsFgg)
}

//go:embed testdata/subtyping/instantiated_type_not_implements/instantiated_type_not_implements.go
var subtypingInstantiatedTypeNotImplementsFgg []byte

func TestTypeCheck_givenInstantiatedTypeDoesNotImplementRequiredInterface_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingInstantiatedTypeNotImplementsFgg,
		`ill-typed declaration: method "Bar.getInterface": `+
			`return expression of type "Foo[Bar]" is not a subtype of "myInterface": `+
			`missing methods: "f(x int) int"`)
}

//go:embed testdata/subtyping/const_instantiated_type_implements/const_instantiated_type_implements.go
var subtypingConstInstantiatedTypeImplementsFgg []byte

func TestTypeCheck_givenConstInstantiatedTypeImplementsRequiredInterface_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingConstInstantiatedTypeImplementsFgg)
}

//go:embed testdata/subtyping/const_instantiated_type_not_implements/const_instantiated_type_not_implements.go
var subtypingConstInstantiatedTypeNotImplementsFgg []byte

func TestTypeCheck_givenConstInstantiatedTypeDoesNotImplementRequiredInterface_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingConstInstantiatedTypeNotImplementsFgg,
		`ill-typed declaration: method "Bar.getInterface": `+
			`return expression of type "Foo[5]" is not a subtype of "myInterface": `+
			`missing methods: "f(x IntArray[2]) IntArray[2]"`)
}

//go:embed testdata/subtyping/recursive_type_param/recursive_type_param.go
var subtypingRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingRecursiveTypeParameterFgg)
}

//go:embed testdata/subtyping/interface_recursive_type_param/interface_recursive_type_param.go
var subtypingInterfaceRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidInterfaceWithRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingInterfaceRecursiveTypeParameterFgg)
}

//go:embed testdata/subtyping/invalid_recursive_type_param/invalid_recursive_type_param.go
var subtypingInvalidRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenInvalidRecursiveTypeParameter_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingInvalidRecursiveTypeParameterFgg,
		`ill-typed main expression: type "comparableHolder" badly instantiated: `+
			`type "Str" is not a subtype of "Eq[Str]": `+
			`missing methods: "equal(other Str) int"`)
}

//go:embed testdata/subtyping/two_level_recursive_type_param/two_level_recursive_type_param.go
var subtypingTwoLevelRecursiveTypeParameterFgg []byte

func TestTypeCheck_givenValidTwoLevelRecursiveTypeParameter_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingTwoLevelRecursiveTypeParameterFgg)
}

// TODO consider if the following tests belong to the subtyping file

//go:embed testdata/subtyping/nested_type_parameter_bound/nested_type_parameter_bound.go
var subtypingNestedTypeParameterBoundFgg []byte

func TestTypeCheck_givenNestedNonRecursiveTypeParameterBound_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, subtypingNestedTypeParameterBoundFgg)
}

//go:embed testdata/subtyping/undefined_nested_type_parameter_bound/undefined_nested_type_parameter_bound.go
var subtypingUndefinedNestedTypeParameterBoundFgg []byte

func TestTypeCheck_givenUndefinedNestedTypeParameterBound_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingUndefinedNestedTypeParameterBoundFgg,
		`ill-typed declaration: type "Foo": `+
			`illegal bound of type parameter "T": `+
			`type "Eq" badly instantiated: type name not declared: "Int"`)
}

//go:embed testdata/subtyping/recursive_bound_type/recursive_bound_type.go
var subtypingRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenBoundReferencingTypeBeingDeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingRecursiveBoundTypeFgg,
		``)
}

//go:embed testdata/subtyping/nested_recursive_bound_type/recursive_bound_type.go
var subtypingNestedRecursiveBoundTypeFGG []byte

func TestTypeCheck_givenNestedBoundReferencingTypeBeingDeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingNestedRecursiveBoundTypeFGG,
		``)
}

//go:embed testdata/subtyping/indirect_recursive_bound_type/indirect_recursive_bound_type.go
var subtypingIndirectRecursiveBoundTypeFgg []byte

func TestTypeCheck_givenTypeDeclarationWithCircularlyDefinedBounds_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingIndirectRecursiveBoundTypeFgg,
		``)
}
