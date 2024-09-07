package parse

import (
	"io"

	"github.com/antlr4-go/antlr/v4"
)

type Actionable[T any, P antlr.Parser] interface {
	NewLexer(input antlr.CharStream) antlr.Lexer
	NewParser(input antlr.TokenStream) P
	Program(parser P) antlr.ParseTree
	NewAstBuilder(tree antlr.ParseTree) ASTBuilder[T]
}

type ASTBuilder[T any] interface {
	BuildAST() T
}

func Program[T any, P antlr.Parser](code io.Reader, actions Actionable[T, P]) (T, error) {
	input := antlr.NewIoStream(code)
	lexer := actions.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := actions.NewParser(stream)
	errListener := &errorListener{}

	p.AddErrorListener(errListener)

	tree := actions.Program(p)
	if len(errListener.syntaxErrors) > 0 {
		var nilProgram T
		return nilProgram, SyntaxErr{}
	}

	astBuilder := actions.NewAstBuilder(tree)
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
