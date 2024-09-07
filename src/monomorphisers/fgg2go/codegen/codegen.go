package codegen

import (
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func GenerateSourceCode(p ast.Program) string {
	return p.String()
}
