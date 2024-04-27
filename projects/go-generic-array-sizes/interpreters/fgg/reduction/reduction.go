package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

type ProgramReducer struct {
}

func NewProgramReducer() ProgramReducer {
	return ProgramReducer{}
}

func (r ProgramReducer) Reduce(program ast.Program) (ast.Program, error) {
	reducedExpr, err := NewReducingVisitor(program.Declarations).Reduce(program.Expression)
	return ast.Program{
		Declarations: program.Declarations,
		Expression:   reducedExpr,
	}, err
}

type ReducingVisitor struct {
	declarations []ast.Declaration
}

func NewReducingVisitor(declarations []ast.Declaration) ReducingVisitor {
	return ReducingVisitor{declarations: declarations}
}

func (r ReducingVisitor) Reduce(e ast.Expression) (ast.Expression, error) {
	return e.Accept(r)
}

func (r ReducingVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) (ast.Expression, error) {
	panic("terminal integer literal cannot be reduced")
}

func (r ReducingVisitor) VisitVariable(v ast.Variable) (ast.Expression, error) {
	return nil, fmt.Errorf("unbound variable %q", v)
}
