package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/int_literal/basic/basic.go
var intLiteralBasicGo []byte

func TestTypeCheck_givenBasicIntLiteralExpression_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(intLiteralBasicGo)
	require.NoError(t, err)
}
