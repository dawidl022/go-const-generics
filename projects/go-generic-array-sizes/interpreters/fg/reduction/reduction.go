package reduction

import (
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func ReduceOneStep(program ast.Program) (ast.Value, error) {
	expression := program.Expression
	switch expression.(type) {
	case ast.Select:
		return reduceField(program.Declarations, expression.(ast.Select))
	case ast.ArrayIndex:
		return reduceIndex(program.Declarations, expression.(ast.ArrayIndex))
	}
	panic("unsupported expression type")
}
