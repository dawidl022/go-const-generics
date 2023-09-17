package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) typeOf(variableEnv map[string]ast.TypeName, expression ast.Expression) (ast.Type, error) {
	return t.newTypeVisitor(variableEnv).typeOf(expression)
}

type typeVisitor struct {
	typeCheckingVisitor
	variableEnv map[string]ast.TypeName
}

func (t typeCheckingVisitor) newTypeVisitor(variableEnv map[string]ast.TypeName) typeVisitor {
	return typeVisitor{typeCheckingVisitor: t, variableEnv: variableEnv}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.Type, error) {
	return expression.Accept(t)
}

func (t typeVisitor) VisitVariable(v ast.Variable) (ast.Type, error) {
	if varType, isVarInEnv := t.variableEnv[v.Id]; isVarInEnv {
		return varType, nil
	}
	return nil, fmt.Errorf("unbound variable: %q", v)
}

func (t typeVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Type, error) {
	if t.isStructTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckStructLiteral(v)
	}
	if t.isArrayTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckArrayLiteral(v)
	}
	return nil, fmt.Errorf("undeclared value literal type name: %q", v.TypeName)
}

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	fields, err := ast.Fields(t.declarations, v.TypeName)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	for i, f := range fields {
		// TODO check less values than fields
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			return err
		}
		err = t.checkIsSubtypeOf(fieldType, f.TypeName)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w", v.Values[i], f.Name, v.TypeName, err)
		}
	}
	return nil
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	// TODO
	return nil
}
