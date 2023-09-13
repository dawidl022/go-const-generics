package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/declaration/invalid_array_type/invalid_array_type.go
var declarationInvalidArrayTypeGo []byte

func TestTypeCheck_givenDeclarationWithInvalidArrayType_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, declarationInvalidArrayTypeGo,
		`ill-typed declaration: type "Arr": element type name not ok: "Bar"`)
}

//go:embed testdata/declaration/zero_size_array/zero_size_array.go
var declarationZeroSizeArrayGo []byte

func TestTypeCheck_givenDeclarationWithZeroSizeArray_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(declarationZeroSizeArrayGo)
	require.NoError(t, err)
}
