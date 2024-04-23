package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/method_spec_const_param/method_spec_const_param.fgg
var typeDeclMethodSpecConstParamFgg []byte

func TestTypeCheck_givenMethodSpecificationWithConstParameter_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclMethodSpecConstParamFgg,
		`ill-typed declaration: type "fooer": `+
			`method specification "foo": parameter "y": `+
			`const type "N" cannot be used as parameter type`)
}

//go:embed testdata/type_declaration/method_spec_int_literal_param/method_spec_int_literal_param.fgg
var typeDeclMethodSpecIntLiteralParamFgg []byte

func TestTypeCheck_givenMethodSpecificationWithIntLiteralParameter_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclMethodSpecIntLiteralParamFgg,
		`ill-typed declaration: type "fooer": `+
			`method specification "foo": parameter "y": `+
			`const type "1" cannot be used as parameter type`)
}

//go:embed testdata/type_declaration/method_spec_int_literal_return_type/method_spec_int_literal_return_type.fgg
var typeDeclMethodSpecIntLiteralReturnTypeFgg []byte

func TestTypeCheck_givenMethodSpecificationWithIntegerLiteralReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclMethodSpecIntLiteralReturnTypeFgg,
		`ill-typed declaration: type "fooer": `+
			`method specification "foo": const type "42" cannot be used as return type`)
}

//go:embed testdata/type_declaration/method_spec_const_return_type/method_spec_const_return_type.fgg
var typeDeclMethodSpecConstReturnTypeFgg []byte

func TestTypeCheck_givenMethodSpecificationWithConstReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclMethodSpecConstReturnTypeFgg,
		`ill-typed declaration: type "fooer": `+
			`method specification "foo": const type "N" cannot be used as return type`)
}

//go:embed testdata/type_declaration/method_spec_type_params/method_spec_type_params.fgg
var typeDeclMethodSpecTypeParamsFgg []byte

func TestTypeCheck_givenMethodSpecificationWithNonConstTypeParams_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclMethodSpecTypeParamsFgg)
}

//go:embed testdata/type_declaration/method_spec_uninstantiated_return_type/method_spec_uninstantiated_return_type.fgg
var typeDeclMethodSpecUninstantiatedReturnTypeFgg []byte

func TestTypeCheck_givenMethodSpecificationWithUninstantiatedGenericReturnType(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclMethodSpecUninstantiatedReturnTypeFgg,
		`ill-typed declaration: type "mapper": `+
			`method specification "Map": return type "mapper" badly instantiated: `+
			`expected 2 type arguments but got 0`)
}

//go:embed testdata/type_declaration/method_spec_generic_return_type/method_spec_generic_return_type.fgg
var typeDeclMethodSpecGenericReturnTypeFgg []byte

func TestTypeCheck_givenMethodSpecificationWithGenericReturnType_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclMethodSpecGenericReturnTypeFgg)
}
