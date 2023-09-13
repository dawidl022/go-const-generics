package typecheck

import (
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func TypeCheck(p ast.Program) error {
	if err := checkDistinctTypeDeclarations(p); err != nil {
		return err
	}
	return checkDistinctMethodDeclarations(p)
}
