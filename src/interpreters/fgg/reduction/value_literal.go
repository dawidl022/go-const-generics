package reduction

import "github.com/dawidl022/go-const-generics/interpreters/fgg/ast"

func (r ReducingVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Expression, error) {
	expressions := make([]ast.Expression, len(v.Values))
	copy(expressions, v.Values)

	for i, val := range v.Values {
		if !val.IsValue() {
			reducedExpr, err := r.Reduce(val)
			expressions[i] = reducedExpr
			return ast.ValueLiteral{Type: v.Type, Values: expressions}, err
		}
	}
	panic("terminal value literal cannot be reduced")
}
