package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/method_declaration/invalid_receiver_type/invalid_receiver_type.go
var declarationInvalidReceiverTypeGo []byte

func TestTypeCheck_givenInvalidMethodReceiverType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidReceiverTypeGo,
		`ill-typed declaration: method "Foo.something": receiver type name not declared: "Foo"`)
}

//go:embed testdata/method_declaration/receiver_and_param_names_duplicated/receiver_and_param_names_duplicated.go
var declarationMethodReceiverAndParamNamesDuplicatedGo []byte

func TestTypeCheck_givenMethodReceiverAndParamNamesAreDuplicated_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationMethodReceiverAndParamNamesDuplicatedGo,
		`ill-typed declaration: method "Foo.something": parameter redeclared "f"`)
}

//go:embed testdata/method_declaration/param_names_duplicated/param_names_duplicated.go
var declarationParameterNamesDuplicatedGo []byte

func TestTypeCheck_givenMethodParameterNamesDuplicate_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationParameterNamesDuplicatedGo,
		`ill-typed declaration: method "Foo.something": parameter redeclared "x"`)
}

//go:embed testdata/method_declaration/invalid_parameter_type/invalid_parameter_type.go
var declarationInvalidParameterTypeGo []byte

func TestTypeCheck_givenInvalidMethodParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidParameterTypeGo,
		`ill-typed declaration: method "Foo.something": parameter "y": type name not declared: "Bar"`)
}

//go:embed testdata/method_declaration/invalid_return_type/invalid_return_type.go
var declarationInvalidReturnTypeGo []byte

func TestTypeCheck_givenInvalidMethodReturnType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidReturnTypeGo,
		`ill-typed declaration: method "Foo.something": return type name not declared: "Bar"`)
}

//go:embed testdata/method_declaration/invalid_expression_concrete_subtype/invalid_expression_concrete_subtype.go
var declarationInvalidExpressionConcreteSubtypeGo []byte

func TestTypeCheck_givenInvalidMethodExpressionConcreteSubtype_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidExpressionConcreteSubtypeGo,
		`ill-typed declaration: method "Foo.something": `+
			`return expression of type "int" is not a subtype of "Foo"`)
}

//go:embed testdata/method_declaration/invalid_expression_interface_subtype/invalid_expression_interface_subtype.go
var declarationInvalidExpressionInterfaceSubtypeGo []byte

func TestTypeCheck_givenInvalidMethodExpressionInterfaceSubtype_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidExpressionInterfaceSubtypeGo,
		`ill-typed declaration: method "Foo.GetIntGetter": `+
			`return expression of type "Foo" is not a subtype of "IntGetter": `+
			`missing methods: "GetInt(x int) int"`)
}

//go:embed testdata/method_declaration/invalid_expression_interface_subtype_multiple_methods_missing/invalid_expression_subtype_multiple_methods_missing.go
var declarationInvalidExpressionInterfaceSubtypeMultipleMethodsMissingGo []byte

func TestTypeCheck_givenInvalidMethodExpressionInterfaceSubtypeWithMultipleMethodsMissing_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidExpressionInterfaceSubtypeMultipleMethodsMissingGo,
		`ill-typed declaration: method "Foo.GetCalculator": `+
			`return expression of type "Foo" is not a subtype of "Calculator": `+
			`missing methods: "Add(x int, y int) int", "Square(x int) int", "Zero() int"`)
}

//go:embed testdata/method_declaration/invalid_expression_interface_subtype_different_param_names/invalid_expression_interface_subtype_different_param_names.go
var declarationInvalidExpressionInterfaceSubtypeDifferentParamNamesGo []byte

func TestTypeCheck_givenInvalidMethodExpressionInterfaceSubtypeWithDifferentParamNames_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidExpressionInterfaceSubtypeDifferentParamNamesGo,
		`ill-typed declaration: method "Foo.GetIntGetter": `+
			`return expression of type "Foo" is not a subtype of "IntGetter": `+
			`missing methods: "GetInt(x int) int"`)
}

//go:embed testdata/method_declaration/valid_expression_subtype/valid_expression_subtype.go
var declarationValidExpressionSubtypeGo []byte

func TestTypeCheck_givenValidMethodExpressionSubtype_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(declarationValidExpressionSubtypeGo)
	require.NoError(t, err)
}

//go:embed testdata/method_declaration/basic/basic.go
var declarationBasicGo []byte

func TestTypeCheck_givenValidMethodDeclarationWithBasicReturnExpressionType_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(declarationBasicGo)
	require.NoError(t, err)
}
