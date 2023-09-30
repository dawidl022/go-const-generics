package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

type BindingVisitor struct {
	variables map[string]ast.Expression
}

func newBindingVisitor(variables map[string]ast.Expression) BindingVisitor {
	return BindingVisitor{variables: variables}
}

func (b BindingVisitor) bind(e ast.Expression) (ast.Expression, error) {
	return e.Accept(b)
}

func (b BindingVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) (ast.Expression, error) {
	return i, nil
}

func (b BindingVisitor) VisitVariable(v ast.Variable) (ast.Expression, error) {
	if val, isBound := b.variables[v.Id]; isBound {
		return val, nil
	}
	return nil, fmt.Errorf("unbound variable %q", v.Id)
}

func (b BindingVisitor) VisitMethodCall(m ast.MethodCall) (ast.Expression, error) {
	boundReceiver, err := b.bind(m.Receiver)
	if err != nil {
		return nil, err
	}
	boundArgs := []ast.Expression{}
	for _, arg := range m.Arguments {
		boundArg, err := b.bind(arg)
		if err != nil {
			return nil, err
		}
		boundArgs = append(boundArgs, boundArg)
	}
	return ast.MethodCall{
		Receiver:   boundReceiver,
		MethodName: m.MethodName,
		Arguments:  boundArgs,
	}, nil
}

func (b BindingVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Expression, error) {
	boundValues := []ast.Expression{}
	for _, val := range v.Values {
		boundVal, err := b.bind(val)
		if err != nil {
			return nil, err
		}
		boundValues = append(boundValues, boundVal)
	}
	return ast.ValueLiteral{Type: v.Type, Values: boundValues}, nil
}

func (b BindingVisitor) VisitSelect(s ast.Select) (ast.Expression, error) {
	boundReceiver, err := b.bind(s.Receiver)
	return ast.Select{FieldName: s.FieldName, Receiver: boundReceiver}, err
}

func (b BindingVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Expression, error) {
	boundReceiver, err := b.bind(a.Receiver)
	if err != nil {
		return nil, err
	}
	boundIndex, err := b.bind(a.Index)
	return ast.ArrayIndex{Receiver: boundReceiver, Index: boundIndex}, err
}
