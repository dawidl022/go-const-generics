package monomo

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	fgg "github.com/dawidl022/go-const-generics/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-const-generics/monomorphisers/fgg2go/codegen"
)

//go:embed testdata/simple_array_literal/input/simple_array_literal.go
var simpleArrayLiteralInput []byte

//go:embed testdata/simple_array_literal/output/simple_array_literal.go
var simpleArrayLiteralOutput []byte

func TestMonomorphise_givenSimpleArrayLiteralGenericOnLength(t *testing.T) {
	assertMonomorphises(t, simpleArrayLiteralInput, simpleArrayLiteralOutput)
}

//go:embed testdata/dead_code_type_elim/input/dead_code_type_elim.go
var deadCodeTypeElimInput []byte

//go:embed testdata/dead_code_type_elim/output/dead_code_type_elim.go
var deadCodeTypeElimOutput []byte

func TestMonomorphise_eliminatesDeadTypeDeclarations(t *testing.T) {
	assertMonomorphises(t, deadCodeTypeElimInput, deadCodeTypeElimOutput)
}

//go:embed testdata/dead_code_generic_type_elim/input/dead_code_generic_type_elim.go
var deadCodeGenericTypeElimInput []byte

//go:embed testdata/dead_code_generic_type_elim/output/dead_code_generic_type_elim.go
var deadCodeGenericTypeElimOutput []byte

func TestMonomorphise_eliminatesDeadGenericTypeDeclarations(t *testing.T) {
	assertMonomorphises(t, deadCodeGenericTypeElimInput, deadCodeGenericTypeElimOutput)
}

//go:embed testdata/dead_code_method_elim/input/dead_code_method_elim.go
var deadCodeMethodElimInput []byte

//go:embed testdata/dead_code_method_elim/output/dead_code_method_elim.go
var deadCodeMethodElimOutput []byte

func TestMonomorphise_eliminatesMethodsForDeadTypes(t *testing.T) {
	assertMonomorphises(t, deadCodeMethodElimInput, deadCodeMethodElimOutput)
}

//go:embed testdata/dead_code_methods_not_eliminated/dead_code_methods_not_eliminated.go
var deadCodeMethodsNotEliminated []byte

func TestMonomorphise_givenTypeIsInstantiated_allMethodsAreAlsoInstantiatedEvenIfUnused(t *testing.T) {
	assertMonomorphises(t, deadCodeMethodsNotEliminated, deadCodeMethodsNotEliminated)
}

//go:embed testdata/non_generic_array/non_generic_array.go
var nonGenericArrayIdentical []byte

func TestMonomorphise_givenNonGenericArray_doesNotChangeOutput(t *testing.T) {
	assertMonomorphises(t, nonGenericArrayIdentical, nonGenericArrayIdentical)
}

//go:embed testdata/generic_array/generic_array.go
var genericArrayIdentical []byte

func TestMonomorphise_givenGenericArrayWithNoConstParams_doesNotChangeOutput(t *testing.T) {
	assertMonomorphises(t, genericArrayIdentical, genericArrayIdentical)
}

//go:embed testdata/generic_const_array/input/generic_const_array.go
var genericConstArrayInput []byte

//go:embed testdata/generic_const_array/output/generic_const_array.go
var genericConstArrayOutput []byte

func TestMonomorphise_givenGenericArrayWithConstParams_MonomorphisesOnlyConstParams(t *testing.T) {
	assertMonomorphises(t, genericConstArrayInput, genericConstArrayOutput)
}

//go:embed testdata/useless_struct/input/useless_struct.go
var uselessStructInput []byte

//go:embed testdata/useless_struct/output/useless_struct.go
var uselessStructOutput []byte

func TestMonomorphise_givenStructWithConstTypeParam(t *testing.T) {
	assertMonomorphises(t, uselessStructInput, uselessStructOutput)
}

//go:embed testdata/useless_array/input/useless_array.go
var uselessArrayInput []byte

//go:embed testdata/useless_array/output/useless_array.go
var uselessArrayOutput []byte

func TestMonomorphise_givenArrayWithUnusedConstTypeParam_MonomorphisesAnyways(t *testing.T) {
	assertMonomorphises(t, uselessArrayInput, uselessArrayOutput)
}

//go:embed testdata/nested_array_literal/input/nested_array_literal.go
var nestedArrayInput []byte

//go:embed testdata/nested_array_literal/output/nested_array_literal.go
var nestedArrayOutput []byte

func TestMonomorphise_givenNestedArrayLiteral(t *testing.T) {
	assertMonomorphises(t, nestedArrayInput, nestedArrayOutput)
}

//go:embed testdata/matrix/input/matrix.go
var matrixInput []byte

//go:embed testdata/matrix/output/matrix.go
var matrixOutput []byte

func TestMonomorphise_givenMatrixWithTwoConstParameters(t *testing.T) {
	assertMonomorphises(t, matrixInput, matrixOutput)
}

//go:embed testdata/matrix_bound/input/matrix_bound.go
var matrixBoundInput []byte

//go:embed testdata/matrix_bound/output/matrix_bound.go
var matrixBoundOutput []byte

func TestMonomorphise_givenMatrixWithInnerArrayAsBound(t *testing.T) {
	assertMonomorphises(t, matrixBoundInput, matrixBoundOutput)
}

//go:embed testdata/struct_field/input/struct_field.go
var structFieldInput []byte

//go:embed testdata/struct_field/output/struct_field.go
var structFieldOutput []byte

func TestMonomorphise_givenStructFieldOfConstGenericArrayType(t *testing.T) {
	assertMonomorphises(t, structFieldInput, structFieldOutput)
}

//go:embed testdata/struct_field_select/input/struct_field.go
var structFieldSelectInput []byte

//go:embed testdata/struct_field_select/output/struct_field.go
var structFieldSelectOutput []byte

func TestMonomorphise_givenStructFieldSelectExpression(t *testing.T) {
	assertMonomorphises(t, structFieldSelectInput, structFieldSelectOutput)
}

//go:embed testdata/array_index/input/array_index.go
var arrayIndexInput []byte

//go:embed testdata/array_index/output/array_index.go
var arrayIndexOutput []byte

func TestMonomorphise_givenArrayIndexExpression(t *testing.T) {
	assertMonomorphises(t, arrayIndexInput, arrayIndexOutput)
}

//go:embed testdata/array_index_nested_expression/input/array_index.go
var arrayIndexNestedExpressionInput []byte

//go:embed testdata/array_index_nested_expression/output/array_index.go
var arrayIndexNestedExpressionOutput []byte

func TestMonomorphise_givenArrayIndexExpressionWithGenericArrayWithinIndex(t *testing.T) {
	assertMonomorphises(t, arrayIndexNestedExpressionInput, arrayIndexNestedExpressionOutput)
}

//go:embed testdata/useless_constraint/input/useless_constraint.go
var uselessConstraintInput []byte

//go:embed testdata/useless_constraint/output/useless_constraint.go
var uselessConstrainOutput []byte

func TestMonomorphise_givenConstGenericEmptyInterfaceAsConstraint(t *testing.T) {
	assertMonomorphises(t, uselessConstraintInput, uselessConstrainOutput)
}

//go:embed testdata/non_generic_method_call/non_generic_method_call.go
var nonGenericMethodCall []byte

func TestMonomorphise_givenNonGenericMethodCall_doesNotChangeOutput(t *testing.T) {
	assertMonomorphises(t, nonGenericMethodCall, nonGenericMethodCall)
}

//go:embed testdata/simple_method_call/input/simple_method_call.go
var simpleMethodCallInput []byte

//go:embed testdata/simple_method_call/output/simple_method_call.go
var simpleMethodCallOutput []byte

func TestMonomorphise_givenMethodCallOnGenericReceiver(t *testing.T) {
	assertMonomorphises(t, simpleMethodCallInput, simpleMethodCallOutput)
}

//go:embed testdata/generic_method_call/input/generic_method_call.go
var genericMethodCallInput []byte

//go:embed testdata/generic_method_call/output/generic_method_call.go
var genericMethodCallOutput []byte

func TestMonomorphise_givenMethodOnCallOnGenericType(t *testing.T) {
	assertMonomorphises(t, genericMethodCallInput, genericMethodCallOutput)
}

//go:embed testdata/method_instantiations/input/method_instantiations.go
var methodInstantiationsInput []byte

//go:embed testdata/method_instantiations/output/method_instantiations.go
var methodInstantiationsOutput []byte

func TestMonomorphise_instantiatesTypeFromMethodParamsReturnTypeAndBody(t *testing.T) {
	assertMonomorphises(t, methodInstantiationsInput, methodInstantiationsOutput)
}

//go:embed testdata/non_generic_array_set_method/non_generic_array_set_method.go
var nonGenericArraySetMethod []byte

func TestMonomorphise_givenNonGenericArraySetMethod_doesNotChangeOutput(t *testing.T) {
	assertMonomorphises(t, nonGenericArraySetMethod, nonGenericArraySetMethod)
}

//go:embed testdata/simple_array_set_method/input/simple_array_set_method.go
var simpleArraySetMethodInput []byte

//go:embed testdata/simple_array_set_method/output/simple_array_set_method.go
var simpleArraySetMethodOutput []byte

func TestMonomorphise_givenSimpleGenericArraySetMethod(t *testing.T) {
	assertMonomorphises(t, simpleArraySetMethodInput, simpleArraySetMethodOutput)
}

//go:embed testdata/nested_array_set_method/input/nested_array_set_method.go
var nestedArraySetMethodInput []byte

//go:embed testdata/nested_array_set_method/output/nested_array_set_method.go
var nestedArraySetMethodOutput []byte

func TestMonomorphise_givenNestedArraySetMethod(t *testing.T) {
	assertMonomorphises(t, nestedArraySetMethodInput, nestedArraySetMethodOutput)
}

//go:embed testdata/generic_interface/input/generic_interface.go
var genericInterfaceInput []byte

//go:embed testdata/generic_interface/output/generic_interface.go
var genericInterfaceOutput []byte

func TestMonomorphise_givenGenericInterface(t *testing.T) {
	assertMonomorphises(t, genericInterfaceInput, genericInterfaceOutput)
}

//go:embed testdata/nested_type_param_constraint/input/nested_type_param_constraint.go
var nestedTypeParamConstraintInput []byte

//go:embed testdata/nested_type_param_constraint/output/nested_type_param_constraint.go
var nestedTypeParamConstraintOutput []byte

func TestMonomorphise_givenNestedTypeParameterConstraint(t *testing.T) {
	assertMonomorphises(t, nestedTypeParamConstraintInput, nestedTypeParamConstraintOutput)
}

//go:embed testdata/addition/input/addition.go
var additionInput []byte

//go:embed testdata/addition/output/addition.go
var additionOutput []byte

func TestMonomorphise_givenAdditionExpression(t *testing.T) {
	assertMonomorphises(t, additionInput, additionOutput)
}

func assertMonomorphises(t *testing.T, input []byte, expected []byte) {
	p := parseProgram(input)
	output := codegen.GenerateSourceCode(Monomorphise(p))
	assert.Equal(t, string(expected), output)
}

func parseProgram(code []byte) ast.Program {
	parsedProgram, err := fgg.Interpreter{}.ParseProgram(bytes.NewBuffer(code))
	if err != nil {
		panic(err)
	}
	err = fgg.Interpreter{}.TypeCheck(parsedProgram)
	if err != nil {
		// ensure all tested programs are well-typed
		panic(err)
	}
	return parsedProgram.Program
}
