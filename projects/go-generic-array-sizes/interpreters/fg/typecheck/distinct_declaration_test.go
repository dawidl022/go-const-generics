package typecheck

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/distinct_declaration/valid/valid.go
var distinctDeclarationValidGo []byte

func TestTypeCheck_givenDistinctDeclarations_returnsNoError(t *testing.T) {
	err := parseAndTypeCheck(distinctDeclarationValidGo)
	require.NoError(t, err)
}

//go:embed testdata/distinct_declaration/int_shadow/int_shadow.go
var distinctDeclarationIntShadowGo []byte

// valid Go, but we do not allow it in FG for simplicity
func TestTypeCheck_givenIntTypeIsShadowed_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, distinctDeclarationIntShadowGo,
		`non distinct type declarations: redeclared "int"`)
}

//go:embed testdata/distinct_declaration/value_type_redeclared/value_type_redeclared.go
var distinctDeclarationValueTypeRedeclaredGo []byte

func TestTypeCheck_givenValueTypeRedeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, distinctDeclarationValueTypeRedeclaredGo,
		`non distinct type declarations: redeclared "Foo"`)
}

//go:embed testdata/distinct_declaration/interface_type_redeclared/interface_type_redeclared.go
var distinctDeclarationInterfaceTypeRedeclaredGo []byte

func TestTypeCheck_givenInterfaceTypeRedeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, distinctDeclarationInterfaceTypeRedeclaredGo,
		`non distinct type declarations: redeclared "Speaker"`)
}
