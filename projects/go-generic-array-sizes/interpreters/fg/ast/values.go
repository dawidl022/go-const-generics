package ast

import (
	"fmt"
)

type Value interface {
	fmt.Stringer
	val()
}

func (i IntegerLiteral) Reduce(declarations []Declaration) (Expression, error) {
	panic("terminal integer literal cannot be reduced")
}

func (i IntegerLiteral) Value() Value {
	return i
}

func (i IntegerLiteral) val() {
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

func (v ValueLiteral) val() {
}
