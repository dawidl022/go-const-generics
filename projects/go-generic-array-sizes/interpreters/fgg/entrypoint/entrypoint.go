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
	return loop.Interpret[fggProgram, ast.Expression, ast.Type](program, debugOutput, fggInterpreter{}, maxSteps)
}

type fggProgram struct {
	program ast.Program
}

func (f fggProgram) Expression() ast.Expression {
	return f.program.Expression
}

type fggInterpreter struct {
}

func (f fggInterpreter) ParseProgram(program io.Reader) (fggProgram, error) {
	parsedProgram, err := parse.Program[ast.Program, *parser.FGGParser](program, parsetree.ParseFGGActions{})
	if err != nil {
		return fggProgram{}, err
	}
	preprocessedProgram := preprocessor.IdentifyTypeParams(parsedProgram)
	return fggProgram{preprocessedProgram}, nil
}

func (f fggInterpreter) TypeCheck(program fggProgram) error {
	return typecheck.TypeCheck(program.program)
}

func (f fggInterpreter) TypeOf(program fggProgram) (ast.Type, error) {
	return typecheck.NewTypeCheckingVisitor(program.program.Declarations).
		TypeOf(nil, nil, program.program.Expression)
}

func (f fggInterpreter) CheckIsSubtypeOf(program fggProgram, subtype, supertype ast.Type) error {
	return typecheck.NewTypeCheckingVisitor(program.program.Declarations).
		NewTypeEnvTypeCheckingVisitor(nil).
		CheckIsSubtypeOf(subtype, supertype)
}

func (f fggInterpreter) Reduce(program fggProgram) (fggProgram, error) {
	newProgram, err := reduction.NewProgramReducer().Reduce(program.program)
	return fggProgram{newProgram}, err
}

func (f fggInterpreter) ProgramWithExpression(program fggProgram, expression ast.Expression) fggProgram {
	return fggProgram{program: ast.Program{
		Declarations: program.program.Declarations,
		Expression:   expression,
	}}
}
