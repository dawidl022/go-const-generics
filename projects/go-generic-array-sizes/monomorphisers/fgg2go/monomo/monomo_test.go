package monomo

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	fgg "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-generic-array-sizes/monomorphisers/fgg2go/codegen"
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
	return parsedProgram.Program
}
