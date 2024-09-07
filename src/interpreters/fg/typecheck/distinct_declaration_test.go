package typecheck

import (
	_ "embed"
	"testing"
)

//go:embed testdata/distinct_declaration/valid/valid.go
var distinctDeclarationValidGo []byte

func TestTypeCheck_givenDistinctDeclarations_returnsNoError(t *testing.T) {
	assertPassesTypeCheck(t, distinctDeclarationValidGo)
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

//go:embed testdata/distinct_declaration/method_redeclared/method_redeclared.go
var distinctDeclarationMethodRedeclaredGo []byte

func TestTypeCheck_givenMethodRedeclared_returnsError(t *testing.T) {
	assertFailsTypeCheckWithError(t, distinctDeclarationMethodRedeclaredGo,
		`non distinct method declarations: redeclared "Arr.first"`)
}
