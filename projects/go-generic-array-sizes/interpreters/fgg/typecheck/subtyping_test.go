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
