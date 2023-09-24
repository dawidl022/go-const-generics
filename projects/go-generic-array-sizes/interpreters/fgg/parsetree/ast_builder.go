package parsetree

import (
	"github.com/antlr4-go/antlr/v4"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

type ASTBuilder interface {
	BuildAST() ast.Program
}

type AntlrASTBuilder struct {
	*antlr.BaseParseTreeVisitor
	parseTree antlr.ParseTree
}

func NewAntlrASTBuilder(parseTree antlr.ParseTree) *AntlrASTBuilder {
	return &AntlrASTBuilder{parseTree: parseTree}
}

func (a AntlrASTBuilder) BuildAST() ast.Program {
	panic("implement me")
}
