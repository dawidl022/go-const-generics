package parsetree

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/parse"
)

type ParseFGGActions struct {
}

func (ParseFGGActions) NewLexer(input antlr.CharStream) antlr.Lexer {
	return parser.NewFGGLexer(input)
}

func (ParseFGGActions) NewParser(input antlr.TokenStream) *parser.FGGParser {
	return parser.NewFGGParser(input)
}

func (ParseFGGActions) Program(p *parser.FGGParser) antlr.ParseTree {
	return p.Program()
}

func (ParseFGGActions) NewAstBuilder(tree antlr.ParseTree) parse.ASTBuilder[ast.Program] {
	return NewAntlrASTBuilder(tree)
}
