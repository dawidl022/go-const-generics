package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type TypeCheckingVisitor struct {
	declarations []ast.Declaration
}

func NewTypeCheckingVisitor(declarations []ast.Declaration) TypeCheckingVisitor {
	return TypeCheckingVisitor{declarations: declarations}
}

func (t TypeCheckingVisitor) TypeCheck(v ast.Visitable) error {
	return v.Accept(t)
}

func (t TypeCheckingVisitor) VisitProgram(p ast.Program) error {
	for _, decl := range p.Declarations {
		if err := t.TypeCheck(decl); err != nil {
			return fmt.Errorf("ill-typed declaration: %w", err)
		}
	}
	_, err := t.TypeOf(nil, p.Expression)
	if err != nil {
		return fmt.Errorf("ill-typed main expression: %w", err)
	}
	return nil
}

func (t TypeCheckingVisitor) TypeOf(variableEnv map[string]ast.TypeName, expression ast.Expression) (ast.Type, error) {
	return t.newTypeVisitor(variableEnv).typeOf(expression)
}

type typeVisitor struct {
	TypeCheckingVisitor
	variableEnv map[string]ast.TypeName
}

func (t TypeCheckingVisitor) newTypeVisitor(variableEnv map[string]ast.TypeName) typeVisitor {
	return typeVisitor{TypeCheckingVisitor: t, variableEnv: variableEnv}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.Type, error) {
	return expression.Accept(t)
}
