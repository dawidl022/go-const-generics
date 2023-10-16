package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
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
	_, err := t.typeOf(nil, nil, p.Expression)
	if err != nil {
		return fmt.Errorf("ill-typed main expression: %w", err)
	}
	return nil
}

func (t typeCheckingVisitor) typeOf(
	typeEnv map[ast.TypeParameter]ast.Bound,
	variableEnv map[string]ast.TypeName,
	expression ast.Expression,
) (ast.Type, error) {
	return t.newTypeVisitor(typeEnv, variableEnv).typeOf(expression)
}

type typeVisitor struct {
	typeCheckingVisitor
	typeEnv     map[ast.TypeParameter]ast.Bound
	variableEnv map[string]ast.TypeName
}

func (t typeCheckingVisitor) newTypeVisitor(
	typeEnv map[ast.TypeParameter]ast.Bound,
	variableEnv map[string]ast.TypeName,
) typeVisitor {
	return typeVisitor{typeCheckingVisitor: t, typeEnv: typeEnv, variableEnv: variableEnv}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.Type, error) {
	return expression.AcceptTypeVisitor(t)
}
