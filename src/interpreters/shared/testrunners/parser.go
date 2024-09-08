package testrunners

import (
	"bytes"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-const-generics/interpreters/shared/parse"
)

func ParseProgram[T any, P antlr.Parser](code []byte, actions parse.Actionable[T, P]) T {
	input := antlr.NewIoStream(bytes.NewBuffer(code))
	lexer := actions.NewLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := actions.NewParser(stream)
	p.AddErrorListener(failingErrorListener{})

	tree := actions.Program(p)
	astBuilder := actions.NewAstBuilder(tree)
	return astBuilder.BuildAST()
}

type failingErrorListener struct {
	*antlr.DefaultErrorListener
}

func (f failingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, msg string, _ antlr.RecognitionException) {
	panic(msg)
}
