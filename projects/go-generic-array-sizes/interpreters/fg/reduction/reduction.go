package reduction

import (
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func ReduceOneStep(program ast.Program) (ast.Value, error) {
	return reduceField(program.Declarations, program.Expression.(ast.Select))
}
