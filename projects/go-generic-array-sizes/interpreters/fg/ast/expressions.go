package ast

import "fmt"

type Expression interface {
	Reduce(declarations []Declaration) (Expression, error)
	Value() Value
	bind(variables map[string]Expression) (Expression, error)
	fmt.Stringer
}

func (p Program) Reduce() (Program, error) {
	expr, err := p.Expression.Reduce(p.Declarations)
	return Program{
		Declarations: p.Declarations,
		Expression:   expr,
	}, err
}

func (v Variable) Reduce(declarations []Declaration) (Expression, error) {
	return nil, fmt.Errorf("unbound variable %q", v)
}

func (v Variable) Value() Value {
	return nil
}
