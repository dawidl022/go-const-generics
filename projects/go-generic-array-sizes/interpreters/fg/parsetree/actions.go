package parsetree

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/testrunners"
)

type ParseFGActions struct{}

func (ParseFGActions) NewLexer(input antlr.CharStream) antlr.Lexer {
	return parser.NewFGLexer(input)
}

func (ParseFGActions) NewParser(input antlr.TokenStream) *parser.FGParser {
	return parser.NewFGParser(input)
}

func (ParseFGActions) Program(p *parser.FGParser) antlr.ParseTree {
	return p.Program()
}

func (ParseFGActions) NewAstBuilder(tree antlr.ParseTree) testrunners.ASTBuilder[ast.Program] {
	return NewAntlrASTBuilder(tree)
}
