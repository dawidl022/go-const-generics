package preprocessor

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/parser"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-const-generics/interpreters/shared/testrunners"
)

//go:embed testdata/method_declaration/shadowed_named_parameter_type/shadowed_named_parameter_type.go
var methodDeclShadowedNamedParameterType []byte

func TestIdentifyTypeParams_givenMethodDeclarationWhereShadowedGenericTypeIsInstantiated_returnsError(t *testing.T) {
	p := parseFGGProgram(methodDeclShadowedNamedParameterType)
	_, err := IdentifyTypeParams(p)
	assert.EqualError(t, err, `type parameter "T" does not accept any type arguments`)
}

func parseFGGProgram(code []byte) ast.Program {
	return testrunners.ParseProgram[ast.Program, *parser.FGGParser](code, parsetree.ParseFGGActions{})
}
