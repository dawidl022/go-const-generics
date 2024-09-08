package codegen

import (
	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func GenerateSourceCode(p ast.Program) string {
	return p.String()
}
