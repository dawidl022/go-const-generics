package reduction

import (
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func ReduceToValue(program ast.Program) (ast.Value, error) {
	for program.Expression.Value() == nil {
		var err error
		program, err = program.Reduce()

		if err != nil {
			return nil, err
		}
	}
	return program.Expression.Value(), nil
}
