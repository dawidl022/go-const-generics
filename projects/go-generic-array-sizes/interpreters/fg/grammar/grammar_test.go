package grammar

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
)

//go:embed testdata/hello.go
var helloGo []byte

func TestGrammarRecognisesHelloGoProgram(t *testing.T) {
	input := antlr.NewIoStream(bytes.NewBuffer(helloGo))
	lexer := parser.NewFGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGParser(stream)
	p.AddErrorListener(failingErrorListener{t: t})

	p.Program()
}

type failingErrorListener struct {
	*antlr.DefaultErrorListener
	t *testing.T
}

func (f failingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, _, _ int, _ string, _ antlr.RecognitionException) {
	f.t.FailNow()
}
