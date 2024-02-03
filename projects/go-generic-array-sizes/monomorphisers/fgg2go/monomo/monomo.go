package monomo

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

type Visitor struct {

}

func (v Visitor) Monomorphise(p ast.Program) ast.Program {
	return p
}
