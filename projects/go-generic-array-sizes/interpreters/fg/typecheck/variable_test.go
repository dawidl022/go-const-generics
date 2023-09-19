package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/variable/unbound_variable/unbound_variable.go
var expressionUnboundVariableGo []byte

func TestTypeCheck_givenVariableInMainFunc_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionUnboundVariableGo,
		`ill-typed main expression: unbound variable: "x"`)
}
