package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
)

//go:embed testdata/addition/simple-literal/simple-literal.go
var additionSimpleLiteralGo []byte

func TestTypeCheck_givenSimpleAdditionExpression_typesAsSumOfValues(t *testing.T) {
	assertPassesTypeCheck(t, additionSimpleLiteralGo)

	p := parseFGProgram(additionSimpleLiteralGo)
	typ, _ := NewTypeCheckingVisitor(p.Declarations).TypeOf(nil, p.Expression)
	assert.Equal(t, ast.IntegerLiteral{IntValue: 2}, typ)
}

//go:embed testdata/addition/chained-literal/chained-literal.go
var additionChainedLiteralGo []byte

func TestTypeCheck_givenChainedAdditionExpression_typesAsSumOfValues(t *testing.T) {
	assertPassesTypeCheck(t, additionChainedLiteralGo)

	p := parseFGProgram(additionChainedLiteralGo)
	typ, _ := NewTypeCheckingVisitor(p.Declarations).TypeOf(nil, p.Expression)
	assert.Equal(t, ast.IntegerLiteral{IntValue: 6}, typ)
}

//go:embed testdata/addition/variable-left/variable-left.go
var additionVariableLeftGo []byte

func TestTypeCheck_givenVariableOnLeftOfAddition_typesToInt(t *testing.T) {
	assertPassesTypeCheck(t, additionVariableLeftGo)

	p := parseFGProgram(additionVariableLeftGo)
	typ, _ := NewTypeCheckingVisitor(p.Declarations).TypeOf(nil, p.Expression)
	assert.Equal(t, intTypeName, typ)
}

//go:embed testdata/addition/variable-right/variable-right.go
var additionVariableRightGo []byte

func TestTypeCheck_givenVariableOnRightOfAddition_typesToInt(t *testing.T) {
	assertPassesTypeCheck(t, additionVariableRightGo)

	p := parseFGProgram(additionVariableRightGo)
	typ, _ := NewTypeCheckingVisitor(p.Declarations).TypeOf(nil, p.Expression)
	assert.Equal(t, intTypeName, typ)
}

//go:embed testdata/addition/untyped-left/untyped-left.go
var additionUntypedLeftGo []byte

func TestTypeCheck_givenUndeclaredVariableOnLeftOfAddition_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, additionUntypedLeftGo,
		`ill-typed main expression: unbound variable: "x"`)
}

//go:embed testdata/addition/untyped-right/untyped-right.go
var additionUntypedRightGo []byte

func TestTypeCheck_givenUndeclaredVariableOnRightOfAddition_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, additionUntypedRightGo,
		`ill-typed main expression: unbound variable: "x"`)
}
