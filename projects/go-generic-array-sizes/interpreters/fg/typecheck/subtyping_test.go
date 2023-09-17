package typecheck

import (
	_ "embed"
	"testing"
)

// TODO double check terminology (covariant vs contravariant)

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

//go:embed testdata/subtyping/covariant_method_parameter/covariant_method_parameter.go
var subtypingCovariantMethodParameterGo []byte

func TestTypeCheck_givenStructWithCovariantMethodParameterTypeUsedAsStructField_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, subtypingCovariantMethodParameterGo,
		`ill-typed main expression: `+
			`cannot use "Foo{}" as field "getter" of struct "Bar": `+
			`type "Foo" is not a subtype of "anyGetter": `+
			`missing methods: "getAny(x int) any"`)
}
