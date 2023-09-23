package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/subtyping/covariant_method_return_type/covariant_method_return_type.go
var subtypingCovariantMethodReturnTypeGo []byte

// a method returning a subtype does not implement an interface with that method returning a supertype
// i.e. interface methods are invariant in Go
func TestTypeCheck_givenStructWithCovariantMethodReturnTypeUsedAsStructField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingCovariantMethodReturnTypeGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as field "getter" of struct "Bar": `+
			`type "Foo" is not a subtype of "anyGetter": `+
			`missing methods: "getAny() any"`)
}

//go:embed testdata/subtyping/incompatible_method_return_type/incompatible_method_return_type.go
var subtypingIncompatibleMethodReturnTypeGo []byte

func TestTypeCheck_giveStructWithIncompatibleMethodReturnTypeUsedAsStructField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingIncompatibleMethodReturnTypeGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as field "getter" of struct "Bar": `+
			`type "Foo" is not a subtype of "anyGetter": `+
			`missing methods: "getAny() int"`)
}

//go:embed testdata/subtyping/incompatible_method_parameter/incompatible_method_parameter.go
var subtypingIncompatibleMethodParameterGo []byte

func TestTypeCheck_givenStructWithIncompatibleMethodParameterTypeUsedAsStructField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingIncompatibleMethodParameterGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as field "getter" of struct "Bar": `+
			`type "Foo" is not a subtype of "anyGetter": `+
			`missing methods: "getAny(x int, y int) int"`)

}

//go:embed testdata/subtyping/contravariant_method_parameter/contravariant_method_parameter.go
var subtypingContravariantMethodParameterGo []byte

func TestTypeCheck_givenStructWithContravariantMethodParameterTypeUsedAsStructField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingContravariantMethodParameterGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as field "getter" of struct "Bar": `+
			`type "Foo" is not a subtype of "anyGetter": `+
			`missing methods: "getAny(x int) any"`)
}

//go:embed testdata/subtyping/array_set_in_methods/array_set_in_methods.go
var subtypingArraySetInMethodsGo []byte

func TestTypeCheck_givenTypeImplementsInterfaceViaArraySetMethod_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(subtypingArraySetInMethodsGo)
	require.NoError(t, err)
}
