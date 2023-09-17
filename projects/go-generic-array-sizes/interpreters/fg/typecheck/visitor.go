package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type typeCheckingVisitor struct {
	declarations []ast.Declaration
}

func newTypeCheckingVisitor(declarations []ast.Declaration) typeCheckingVisitor {
	return typeCheckingVisitor{declarations: declarations}
}

func (t typeCheckingVisitor) TypeCheck(v ast.Visitable) error {
	return v.Accept(t)
}

func (t typeCheckingVisitor) VisitProgram(p ast.Program) error {
	for _, decl := range p.Declarations {
		if err := t.TypeCheck(decl); err != nil {
			return fmt.Errorf("ill-typed declaration: %w", err)
		}
	}
	_, err := t.typeOf(nil, p.Expression)
	return err
}
