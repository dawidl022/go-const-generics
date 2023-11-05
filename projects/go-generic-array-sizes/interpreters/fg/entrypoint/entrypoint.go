package entrypoint

import (
	"io"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/typecheck"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/loop"
)

func Interpret(program io.Reader, debugOutput io.Writer) (string, error) {
	return loop.Interpret[fgProgram, fgExpression, ast.Type](program, debugOutput, fgInterpreter{})
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
	parsedProgram, err := parseFGProgram(program)
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

func parseFGProgram(r io.Reader) (ast.Program, error) {
	input := antlr.NewIoStream(r)
	lexer := parser.NewFGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGParser(stream)
	errListener := &errorListener{}

	p.AddErrorListener(errListener)
	p.BuildParseTrees = true

	tree := p.Program()
	if len(errListener.syntaxErrors) > 0 {
		return ast.Program{}, SyntaxErr{}
	}

	astBuilder := parsetree.NewAntlrASTBuilder(tree)
	return astBuilder.BuildAST(), nil
}

type SyntaxErr struct {
}

func (s SyntaxErr) Error() string {
	return "one or more syntax errors detected"
}

type errorListener struct {
	*antlr.DefaultErrorListener
	syntaxErrors []string
}

func (f *errorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
	f.syntaxErrors = append(f.syntaxErrors, msg)
}
