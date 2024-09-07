package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeVisitor) VisitVariable(v ast.Variable) (ast.Type, error) {
	if varType, isVarInEnv := t.variableEnv[v.Id]; isVarInEnv {
		return varType, nil
	}
	return nil, fmt.Errorf("unbound variable: %q", v)
}
