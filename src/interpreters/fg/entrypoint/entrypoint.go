package entrypoint

import (
	"io"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/typecheck"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/parse"
)

func Interpret(program io.Reader, debugOutput io.Writer, maxSteps int) (string, error) {
	return loop.Interpret[fgProgram, fgExpression, ast.Type](program, debugOutput, fgInterpreter{}, maxSteps)
}

type fgProgram struct {
	program ast.Program
}

func (f fgProgram) Expression() fgExpression {
	return fgExpression{f.program.Expression}
}

type fgExpression struct {
	ast.Expression
}

func (f fgExpression) IsValue() bool {
	return f.Value() != nil
}

type fgInterpreter struct {
}

func (f fgInterpreter) ParseProgram(program io.Reader) (fgProgram, error) {
	parsedProgram, err := parse.Program[ast.Program, *parser.FGParser](program, parsetree.ParseFGActions{})
	return fgProgram{parsedProgram}, err
}

func (f fgInterpreter) TypeCheck(program fgProgram) error {
	return typecheck.TypeCheck(program.program)
}

func (f fgInterpreter) TypeOf(program fgProgram) (ast.Type, error) {
	return typecheck.NewTypeCheckingVisitor(program.program.Declarations).
		TypeOf(nil, program.program.Expression)
}

func (f fgInterpreter) CheckIsSubtypeOf(program fgProgram, subtype, supertype ast.Type) error {
	return typecheck.NewTypeCheckingVisitor(program.program.Declarations).
		CheckIsSubtypeOf(subtype, supertype)
}

func (f fgInterpreter) Reduce(program fgProgram) (fgProgram, error) {
	newProgram, err := program.program.Reduce()
	return fgProgram{newProgram}, err
}

func (f fgInterpreter) ProgramWithExpression(program fgProgram, expression fgExpression) fgProgram {
	return fgProgram{program: ast.Program{
		Declarations: program.program.Declarations,
		Expression:   expression.Expression,
	}}
}
