package entrypoint

import (
	"io"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/preprocessor"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/reduction"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/typecheck"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/parse"
)

func Interpret(program io.Reader, debugOutput io.Writer, maxSteps int) (string, error) {
	return loop.Interpret[FggProgram, ast.Expression, ast.Type](program, debugOutput, Interpreter{}, maxSteps)
}

type FggProgram struct {
	Program ast.Program
}

func (f FggProgram) Expression() ast.Expression {
	return f.Program.Expression
}

type Interpreter struct {
}

func (f Interpreter) ParseProgram(program io.Reader) (FggProgram, error) {
	parsedProgram, err := parse.Program[ast.Program, *parser.FGGParser](program, parsetree.ParseFGGActions{})
	if err != nil {
		return FggProgram{}, err
	}
	preprocessedProgram, err := preprocessor.IdentifyTypeParams(parsedProgram)
	return FggProgram{preprocessedProgram}, err
}

func (f Interpreter) TypeCheck(program FggProgram) error {
	return typecheck.TypeCheck(program.Program)
}

func (f Interpreter) TypeOf(program FggProgram) (ast.Type, error) {
	return typecheck.NewTypeCheckingVisitor(program.Program.Declarations).
		TypeOf(nil, nil, program.Program.Expression)
}

func (f Interpreter) CheckIsSubtypeOf(program FggProgram, subtype, supertype ast.Type) error {
	return typecheck.NewTypeCheckingVisitor(program.Program.Declarations).
		NewTypeEnvTypeCheckingVisitor(nil).
		CheckIsSubtypeOf(subtype, supertype)
}

func (f Interpreter) Reduce(program FggProgram) (FggProgram, error) {
	newProgram, err := reduction.NewProgramReducer().Reduce(program.Program)
	return FggProgram{newProgram}, err
}

func (f Interpreter) ProgramWithExpression(program FggProgram, expression ast.Expression) FggProgram {
	return FggProgram{Program: ast.Program{
		Declarations: program.Program.Declarations,
		Expression:   expression,
	}}
}
