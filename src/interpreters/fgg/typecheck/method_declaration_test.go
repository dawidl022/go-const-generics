package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/method_declaration/type_params/type_params.fgg
var methodDeclTypeParamsFgg []byte

func TestTypeCheck_givenMethodDeclarationWithTypeParametersUsedInSignature_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, methodDeclTypeParamsFgg)
}

//go:embed testdata/method_declaration/renamed_type_params/renamed_type_params.fgg
var methodDeclRenamedTypeParamsFgg []byte

func TestTypeCheck_givenMethodDeclarationWithRenamedTypeParametersInReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclRenamedTypeParamsFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`receiver type parameter name "X" does not match type declaration parameter name "T"`)
}

//go:embed testdata/method_declaration/int_literal_param/int_literal_param.fgg
var methodDeclIntLiteralParamFgg []byte

func TestTypeCheck_givenMethodDeclarationWithIntegerLiteralUsedAsParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclIntLiteralParamFgg,
		`ill-typed declaration: method "Foo.something": `+
			`parameter "x": const type "42" cannot be used as parameter type`)
}

//go:embed testdata/method_declaration/const_param/const_param.fgg
var methodDeclConstParamFgg []byte

func TestTypeCheck_givenMethodDeclarationWithConstTypeParameterUsedAsParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclConstParamFgg,
		`ill-typed declaration: method "Foo.something": `+
			`parameter "x": const type "N" cannot be used as parameter type`)
}

//go:embed testdata/method_declaration/const_return_type/const_return_type.fgg
var methodDeclConstReturnTypeFgg []byte

func TestTypeCheck_givenMethodDeclarationWithConstReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclConstReturnTypeFgg,
		`ill-typed declaration: method "Foo.something": `+
			`const type "N" cannot be used as return type`)
}

//go:embed testdata/method_declaration/int_literal_return_type/int_literal_return_type.fgg
var methodDeclIntLiteralReturnTypeFgg []byte

func TestTypeCheck_givenMethodDeclarationWithIntLiteralReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclIntLiteralReturnTypeFgg,
		`ill-typed declaration: method "Foo.something": `+
			`const type "42" cannot be used as return type`)
}

//go:embed testdata/method_declaration/missing_type_params/missing_type_params.fgg
var methodDeclMissingTypeParamsFgg []byte

func TestTypeCheck_givenMethodDeclarationWithMissingReceiverTypeParameters_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclMissingTypeParamsFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`expected 2 type parameters on receiver but got 0`)
}

//go:embed testdata/method_declaration/extraneous_type_params/extraneous_type_params.fgg
var methodDeclExtraneousTypeParamsFgg []byte

func TestTypeCheck_givenMethodDeclarationWithExtraneousReceiverTypeParameters_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclExtraneousTypeParamsFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`expected 2 type parameters on receiver but got 3`)
}

//go:embed testdata/method_declaration/receiver_param/receiver_param.fgg
var methodDeclReceiverParamFgg []byte

func TestTypeCheck_givenMethodDeclarationWithParameterOfReceiverType_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, methodDeclReceiverParamFgg)
}

//go:embed testdata/method_declaration/receiver_return_type/receiver_return_type.fgg
var methodDeclReceiverReturnTypeFgg []byte

func TestTypeCheck_givenMethodDeclarationWithReturnTypeOfReceiverType_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, methodDeclReceiverReturnTypeFgg)
}

//go:embed testdata/method_declaration/uninstantiated_receiver_param/uninstantiated_receiver_param.fgg
var methodDeclUninstantiatedReceiverParamFgg []byte

func TestTypeCheck_givenMethodDeclarationWithParameterOfUninstantiatedReceiverType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclUninstantiatedReceiverParamFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`parameter "param": type "Foo" badly instantiated: `+
			`expected 2 type arguments but got 0`)
}

//go:embed testdata/method_declaration/uninstantiated_receiver_return_type/uninstantiated_receiver_return_type.fgg
var methodDeclUninstantiatedReceiverReturnTypeFgg []byte

func TestTypeCheck_givenMethodDeclarationWithReturnTypeOfUninstantiatedReceiverType(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclUninstantiatedReceiverReturnTypeFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`return type "Foo" badly instantiated: `+
			`expected 2 type arguments but got 0`)
}

//go:embed testdata/method_declaration/mismatch_generic_variant_expression/mismatch_generic_variant_expression.fgg
var methodDeclMismatchGenericVariantExpressionFgg []byte

func TestTypeCheck_givenMethodDeclarationWithReturnTypeMismatchingGenericReturnType(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclMismatchGenericVariantExpressionFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`return expression of type "Foo[int, int]" is not a subtype of "Foo[T, R]"`)
}

//go:embed testdata/method_declaration/mismatch_variant_expression/mismatch_variant_expression.fgg
var methodDeclMismatchVariantExpressionFgg []byte

func TestTypeCheck_givenMethodDeclarationWithExpressionMismatchingReturnTypeInTermsOfTypeArguments_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclMismatchVariantExpressionFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`return expression of type "Foo[any, R]" is not a subtype of "Foo[int, R]"`)
}

//go:embed testdata/method_declaration/covariant_type_params/covariant_type_params.fgg
var methodDeclCovariantTypeParamsFgg []byte

func TestTypeCheck_givenMethodDeclarationWithExpressionCovariantToReturnTypeInTermsOfTypeArguments_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, methodDeclCovariantTypeParamsFgg,
		`ill-typed declaration: method "Foo.foo": `+
			`return expression of type "Foo[int, R]" is not a subtype of "Foo[any, R]"`)
}

//go:embed testdata/method_declaration/shadowed_parameter_type/shadowed_parameter_type.fgg
var methodDeclShadowedParameterType []byte

func TestTypeCheck_givenMethodDeclarationWhereTypeParameterShadowsGenericType_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, methodDeclShadowedParameterType)
}
