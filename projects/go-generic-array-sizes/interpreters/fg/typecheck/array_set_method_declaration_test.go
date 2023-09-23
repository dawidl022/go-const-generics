package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/array_set_method_declaration/basic/basic.go
var arraySetBasicGo []byte

func TestTypeCheck_givenBasicValidArraySetMethodDeclaration_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(arraySetBasicGo)
	require.NoError(t, err)
}

//go:embed testdata/array_set_method_declaration/subtype/subtype.go
var arraySetSubtypeGo []byte

func TestTypeCheck_givenValidArraySetMethodDeclarationWithValueParameterSubtypeOfArrayElementType_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(arraySetSubtypeGo)
	require.NoError(t, err)
}

//go:embed testdata/array_set_method_declaration/invalid_index_parameter_type/invalid_index_parameter_type.go
var arraySetInvalidIndexParameterType []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithInvalidIndexParameterType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetInvalidIndexParameterType,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`first parameter "x" must be of type "int"`)
}

//go:embed testdata/array_set_method_declaration/invalid_value_parameter_type/invalid_value_parameter_type.go
var arraySetInvalidValueParameterType []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithInvalidValueType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetInvalidValueParameterType,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`second parameter "y" cannot be used as element of array type "Arr": `+
			`type "any" is not a subtype of "int"`)
}

//go:embed testdata/array_set_method_declaration/invalid_return_type/invalid_return_type.go
var arraySetInvalidReturnTypeGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithReturnTypeDifferentFromReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetInvalidReturnTypeGo,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`return type must be same as receiver type "Arr"`)
}

//go:embed testdata/array_set_method_declaration/wrong_index_receiver_variable/wrong_index_receiver_variable.go
var arraySetWrongIndexReceiverVariableGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithWrongIndexReceiverVariableInBody_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetWrongIndexReceiverVariableGo,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`index receiver must be the same as method receiver "a"`)
}

//go:embed testdata/array_set_method_declaration/wrong_index_argument_variable/wrong_index_argument_variable.go
var arraySetWrongIndexArgumentVariableGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithWrongIndexArgumentVariableInBody_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetWrongIndexArgumentVariableGo,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`index argument must be the same as first parameter "x"`)
}

//go:embed testdata/array_set_method_declaration/wrong_value_variable/wrong_value_variable.go
var arraySetWrongValueVariableGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithWrongValueVariableInBody_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetWrongValueVariableGo,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`new index value must be the same as second parameter "y"`)
}

//go:embed testdata/array_set_method_declaration/wrong_return_variable/wrong_return_variable.go
var arraySetWrongReturnVariableGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithWrongReturnVariableInBody_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetWrongReturnVariableGo,
		`ill-typed declaration: array-set method "Arr.invalidSet": `+
			`return variable must be the same as method receiver "a"`)
}

//go:embed testdata/array_set_method_declaration/non_array_receiver/non_array_receiver.go
var arraySetNonArrayReceiverGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithNonArrayTypeReceiver_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetNonArrayReceiverGo,
		`ill-typed declaration: array-set method "Foo.set": method receiver must be of array type`)
}

//go:embed testdata/array_set_method_declaration/redeclared_receiver_name/redeclared_receiver_name.go
var arraySetRedeclaredReceiverNameGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithReceiverNameRedeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetRedeclaredReceiverNameGo,
		`ill-typed declaration: array-set method "Arr.set": redeclared "a"`)
}

//go:embed testdata/array_set_method_declaration/redeclared_parameter_name/redeclared_parameter_name.go
var arraySetRedeclaredParameterNameGo []byte

func TestTypeCheck_givenArraySetMethodDeclarationWithParameterNameRedeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, arraySetRedeclaredParameterNameGo,
		`ill-typed declaration: array-set method "Arr.set": redeclared "x"`)
}
