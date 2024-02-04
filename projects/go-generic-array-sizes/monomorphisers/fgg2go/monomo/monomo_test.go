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
