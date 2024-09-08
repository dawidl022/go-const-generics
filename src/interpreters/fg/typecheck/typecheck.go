package typecheck

import (
	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
)

func TypeCheck(p ast.Program) error {
	if err := checkDistinctTypeDeclarations(p); err != nil {
		return err
	}
	if err := checkDistinctMethodDeclarations(p); err != nil {
		return err
	}
	return NewTypeCheckingVisitor(p.Declarations).TypeCheck(p)
}
