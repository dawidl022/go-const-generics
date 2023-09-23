package entrypoint

import (
	"fmt"
	"io"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parsetree"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/reduction"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/typecheck"
)

func Interpret(program io.Reader, debugOutput io.Writer) (string, error) {
	p, err := parseFGProgram(program)
	if err != nil {
		return "", fmt.Errorf("failed to parse program: %w", err)
	}

	err = typecheck.TypeCheck(p)
	if err != nil {
		return "", fmt.Errorf("type error: %w", err)
	}

	debugObserver := &debugObserver{writer: debugOutput}
	typeObserver := &typeCheckingObserver{declarations: p.Declarations, writer: debugOutput}

	reducer := reduction.NewProgramReducer([]reduction.Observer{debugObserver, typeObserver})

	val, err := reducer.ReduceToValue(p)
	if err != nil {
		return "", err
	}
	return val.String(), nil
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

type debugObserver struct {
	writer     io.Writer
	stepNumber int
}

func (d *debugObserver) Notify(expression ast.Expression) error {
	d.stepNumber++
	_, err := fmt.Fprintf(d.writer, "reduction step %d: %s\n", d.stepNumber, expression)
	if err != nil {
		return fmt.Errorf("failed to write to debug output: %w", err)
	}
	return nil
}

type typeCheckingObserver struct {
	writer       io.Writer
	declarations []ast.Declaration
}

func (t *typeCheckingObserver) Notify(expression ast.Expression) error {
	err := typecheck.TypeCheck(ast.Program{
		Declarations: t.declarations,
		Expression:   expression,
	})
	if err != nil {
		return fmt.Errorf("type error: %s\n", err)
	}
	_, err = fmt.Fprint(t.writer, "program well typed\n\n")
	if err != nil {
		return fmt.Errorf("failed to write to debug output: %w", err)
	}
	return nil
}
