package ast

import "fmt"

type Expression interface {
	Reduce(declarations []Declaration) (Expression, error)
	Value() Value
	fmt.Stringer
}

func (p Program) Reduce() (Program, error) {
	expr, err := p.Expression.Reduce(p.Declarations)
	return Program{
		Declarations: p.Declarations,
		Expression:   expr,
	}, err
}

func (i IntegerLiteral) Reduce(declarations []Declaration) (Expression, error) {
	panic("terminal integer literal cannot be reduced")
}

func (i IntegerLiteral) Value() Value {
	return i
}

func (v Variable) Reduce(declarations []Declaration) (Expression, error) {
	//TODO implement me
	panic("implement me")
}

func (v Variable) Value() Value {
	//TODO implement me
	panic("implement me")
}

func (v Variable) String() string {
	return v.Id
}

func (m MethodCall) String() string {
	//TODO implement me
	panic("implement me")
}

func (v ValueLiteral) Reduce(declarations []Declaration) (Expression, error) {
	expressions := make([]Expression, len(v.Values))
	copy(expressions, v.Values)

	for i, expr := range v.Values {
		if expr.Value() == nil {
			newExpr, err := expr.Reduce(declarations)
			expressions[i] = newExpr
			return ValueLiteral{TypeName: v.TypeName, Values: expressions}, err
		}
	}
	panic("terminal value literal cannot be reduced")
}

func (v ValueLiteral) Value() Value {
	return v
}
