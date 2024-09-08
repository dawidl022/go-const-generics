package grammar

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-const-generics/interpreters/fg/parser"
	fggParser "github.com/dawidl022/go-const-generics/interpreters/fgg/parser"
)

//go:embed testdata/hello.go
var helloGo []byte

func TestGrammarRecognisesHelloGoProgram(t *testing.T) {
	tests := []struct {
		name      string
		newLexer  func(input antlr.CharStream) antlr.Lexer
		newParser func(input antlr.TokenStream) antlr.Parser
		program   func(parser antlr.Parser)
	}{
		{
			name:      "FG",
			newLexer:  func(i antlr.CharStream) antlr.Lexer { return parser.NewFGLexer(i) },
			newParser: func(i antlr.TokenStream) antlr.Parser { return parser.NewFGParser(i) },
			program:   func(p antlr.Parser) { p.(*parser.FGParser).Program() },
		},
		{
			name:      "FGG",
			newLexer:  func(i antlr.CharStream) antlr.Lexer { return fggParser.NewFGGLexer(i) },
			newParser: func(i antlr.TokenStream) antlr.Parser { return fggParser.NewFGGParser(i) },
			program:   func(p antlr.Parser) { p.(*fggParser.FGGParser).Program() },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			input := antlr.NewIoStream(bytes.NewBuffer(helloGo))
			lexer := test.newLexer(input)
			stream := antlr.NewCommonTokenStream(lexer, 0)

			p := test.newParser(stream)
			p.AddErrorListener(failingErrorListener{t: t})

			test.program(p)
		})
	}
}

type failingErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func (f failingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, _ string, _ antlr.RecognitionException) {
	f.t.FailNow()
}
