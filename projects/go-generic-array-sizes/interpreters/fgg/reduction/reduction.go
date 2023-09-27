package reduction

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

type ProgramReducer struct {
}

func NewProgramReducer() ProgramReducer {
	return ProgramReducer{}
}

func (r ProgramReducer) Reduce(program ast.Program) (ast.Program, error) {
	return program, nil
}

type ReducingVisitor struct {
	declarations []ast.Declaration
}

func NewReducingVisitor(declarations []ast.Declaration) ReducingVisitor {
	return ReducingVisitor{declarations: declarations}
}

func (r ReducingVisitor) Reduce(e ast.Expression) (ast.Expression, error) {
	return e, nil
}
