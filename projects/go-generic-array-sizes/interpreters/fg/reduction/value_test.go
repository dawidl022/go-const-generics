package reduction

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

//go:embed testdata/value/int/int.go
var valueIntGo []byte

func TestValue_givenInt_returnsValueAndFailsToReduce(t *testing.T) {
	p := parseFGProgram(valueIntGo)

	require.Panics(t, func() { p.Reduce() })
	require.Equal(t, ast.IntegerLiteral{IntValue: 1}, p.Expression.Value())
}
