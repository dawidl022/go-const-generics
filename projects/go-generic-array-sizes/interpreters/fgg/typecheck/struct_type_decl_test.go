package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/type_declaration/struct_const_field/struct_const_field.go
var typeDeclStructConstFieldFgg []byte

func TestTypeCheck_givenStructTypeDeclarationWithIntLiteralFieldType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclStructConstFieldFgg,
		`ill-typed declaration: type "Foo": cannot use const type "42" as field type`)
}

//go:embed testdata/type_declaration/struct_type_params/struct_type_params.go
var typeDeclStructTypeParams []byte

func TestTypeCheck_givenStructTypeDeclarationWithTypeParameters_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, typeDeclStructTypeParams)
}

//go:embed testdata/type_declaration/struct_invalid_field_type_args/struct_invalid_field_type_args.go
var typeDeclStructInvalidFieldTypeArgsFgg []byte

func TestTypeCheck_givenStructDeclarationWithInvalidTypeArgument_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, typeDeclStructInvalidFieldTypeArgsFgg,
		`ill-typed declaration: type "Foo": `+
			`field "arr": type "Arr" badly instantiated: `+
			`type "T" is not a subtype of "fooer": missing methods: "foo() int"`)
}
