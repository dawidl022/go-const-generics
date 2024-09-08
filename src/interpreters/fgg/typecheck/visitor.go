package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

type typeCheckingVisitor struct {
	declarations []ast.Declaration
}

func NewTypeCheckingVisitor(declarations []ast.Declaration) typeCheckingVisitor {
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
	_, err := t.TypeOf(nil, nil, p.Expression)
	if err != nil {
		return fmt.Errorf("ill-typed main expression: %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) TypeOf(
	typeEnv map[ast.TypeParameter]ast.Bound,
	variableEnv map[string]ast.Type,
	expression ast.Expression,
) (ast.Type, error) {
	return t.newTypeVisitor(typeEnv, variableEnv).typeOf(expression)
}

type typeVisitor struct {
	typeEnvTypeCheckingVisitor
	variableEnv map[string]ast.Type
}

func (t typeCheckingVisitor) newTypeVisitor(
	typeEnv map[ast.TypeParameter]ast.Bound,
	variableEnv map[string]ast.Type,
) typeVisitor {
	return typeVisitor{
		typeEnvTypeCheckingVisitor: typeEnvTypeCheckingVisitor{
			typeCheckingVisitor: t,
			typeEnv:             typeEnv,
		},
		variableEnv: variableEnv,
	}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.Type, error) {
	return expression.AcceptTypeVisitor(t)
}
