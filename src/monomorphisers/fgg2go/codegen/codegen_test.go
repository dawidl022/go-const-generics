package codegen

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/parse"
)

//go:embed testdata/trivial/trivial.go
var testdataTrivialGo []byte

func TestGenerateSourceCode_givenTrivialProgram(t *testing.T) {
	assertGeneratesCode(t, testdataTrivialGo)
}

//go:embed testdata/int_literal/int_literal.go
var testdataIntLiteral []byte

func TestGenerateSourceCode_givenMainExpressionWithIntegerLiteral(t *testing.T) {
	assertGeneratesCode(t, testdataIntLiteral)
}

//go:embed testdata/struct_decl/struct_decl.go
var testdataStructDecl []byte

func TestGenerateSourceCode_givenStructTypeDeclaration(t *testing.T) {
	assertGeneratesCode(t, testdataStructDecl)
}

//go:embed testdata/struct_decl_with_field/struct_decl_with_field.go
var testdataStructDeclWithField []byte

func TestGenerateSourceCode_givenStructTypeDeclWithSingleField(t *testing.T) {
	assertGeneratesCode(t, testdataStructDeclWithField)
}

//go:embed testdata/struct_decl_with_multiple_fields/struct_decl_with_multiple_fields.go
var testdataStructDeclWithMultipleFields []byte

func TestGenerateSourceCode_givenStructTypeDeclWithMultipleFields(t *testing.T) {
	assertGeneratesCode(t, testdataStructDeclWithMultipleFields)
}

//go:embed testdata/array_type_literal/array_type_literal.go
var testdataArrayTypeLiteral []byte

func TestGenerateSourceCode_givenArrayTypeDecl(t *testing.T) {
	assertGeneratesCode(t, testdataArrayTypeLiteral)
}

//go:embed testdata/generic_struct_decl/generic_struct_decl.go
var testdataGenericStructDecl []byte

func TestGenerateSourceCode_givenGenericStructTypeLiteral(t *testing.T) {
	assertGeneratesCode(t, testdataGenericStructDecl)
}

//go:embed testdata/generic_array_decl/generic_array_decl.go
var testdataGenericArrayDecl []byte

func TestGenerateSourceCode_givenGenericArrayTypeLiteral(t *testing.T) {
	assertGeneratesCode(t, testdataGenericArrayDecl)
}

//go:embed testdata/interface_empty_type_literal/interface_empty_type_literal.go
var testdataInterfaceEmptyTypeLiteral []byte

func TestGenerateSourceCode_givenInterfaceWithNoMethods(t *testing.T) {
	assertGeneratesCode(t, testdataInterfaceEmptyTypeLiteral)
}

//go:embed testdata/interface_type_literal/interface_type_literal.go
var testdataInterfaceTypeLiteral []byte

func TestGenerateSourceCode_givenInterfaceTypeLiteral(t *testing.T) {
	assertGeneratesCode(t, testdataInterfaceTypeLiteral)
}

//go:embed testdata/interface_multiple_methods/interface_multiple_methods.go
var testdataInterfaceMultipleMethods []byte

func TestGenerateSourceCode_givenInterfaceWithMultipleMethodsAndGenericTypes(t *testing.T) {
	assertGeneratesCode(t, testdataInterfaceMultipleMethods)
}

//go:embed testdata/method_decl/method_decl.go
var testdataMethodDecl []byte

func TestGenerateSourceCode_givenMethodDeclaration(t *testing.T) {
	assertGeneratesCode(t, testdataMethodDecl)
}

//go:embed testdata/generic_method_decl/generic_method_decl.go
var testdataGenericMethodDecl []byte

func TestGenerateSourceCode_givenGenericMethodDeclaration(t *testing.T) {
	assertGeneratesCode(t, testdataGenericMethodDecl)
}

//go:embed testdata/array_set_method_decl/array_set_method_decl.go
var testdataArraySetMethodDecl []byte

func TestGenerateSourceCode_givenArraySetMethod(t *testing.T) {
	assertGeneratesCode(t, testdataArraySetMethodDecl)
}

//go:embed testdata/generic_array_set_method_decl/array_set_method_decl.go
var testdataGenericArraySetMethodDecl []byte

func TestGenerateSourceCode_givenGenericArraySetMethod(t *testing.T) {
	assertGeneratesCode(t, testdataGenericArraySetMethodDecl)
}

//go:embed testdata/method_call_with_variable/method_call_with_variable.go
var testdataMethodCallWithVariable []byte

func TestGenerateSourceCode_givenMethodCallWhereMethodReturnsVariable(t *testing.T) {
	assertGeneratesCode(t, testdataMethodCallWithVariable)
}

//go:embed testdata/value_literal/value_literal.go
var testdataValueLiteral []byte

func TestGenerateSourceCode_givenValueLiteralExpression(t *testing.T) {
	assertGeneratesCode(t, testdataValueLiteral)
}

//go:embed testdata/generic_value_literal/generic_value_literal.go
var testdataGenericValueLiteral []byte

func TestGenerateSourceCode_givenGenericValueLiteral(t *testing.T) {
	assertGeneratesCode(t, testdataGenericValueLiteral)
}

//go:embed testdata/field_select/field_select.go
var testdataFieldSelect []byte

func TestGenerateSourceCode_givenFieldSelect(t *testing.T) {
	assertGeneratesCode(t, testdataFieldSelect)
}

//go:embed testdata/generic_array_index/generic_array_index.go
var testdataGenericArrayIndex []byte

func TestGenerateSourceCode_givenGenericArrayIndex(t *testing.T) {
	assertGeneratesCode(t, testdataGenericArrayIndex)
}

func assertGeneratesCode(t *testing.T, code []byte) {
	p := parseProgram(code)
	output := GenerateSourceCode(p)
	assert.Equal(t, string(code), output)
}

func parseProgram(code []byte) ast.Program {
	parsedProgram, err := parse.Program[ast.Program, *parser.FGGParser](bytes.NewBuffer(code), parsetree.ParseFGGActions{})
	if err != nil {
		panic(err)
	}
	return parsedProgram
}
