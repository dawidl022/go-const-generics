package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/value_literal/undeclared_value_literal_type/undeclared_value_literal_type.go
var expressionUndeclaredValueLiteralTypeGo []byte

func TestTypeCheck_givenValueLiteralTypeUndeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, expressionUndeclaredValueLiteralTypeGo,
		`ill-typed main expression: undeclared value literal type name: "Foo"`)
}
