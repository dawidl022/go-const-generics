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

func (i IntegerLiteral) Reduce(declarations []Declaration) (Expression, error) {
	panic("terminal integer literal cannot be reduced")
}

func (i IntegerLiteral) Value() Value {
	return i
}

func (v Variable) Reduce(declarations []Declaration) (Expression, error) {
	return nil, fmt.Errorf("unbound variable %q", v)
}

func (v Variable) Value() Value {
	return nil
}

func (v Variable) String() string {
	return v.Id
}

func (m MethodCall) String() string {
	s := fmt.Sprintf("%s.%s(", m.Receiver, m.MethodName)
	for i, arg := range m.Arguments {
		s += arg.String()
		if i < len(m.Arguments)-1 {
			s += ", "
		}
	}
	s += ")"
	return s
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
