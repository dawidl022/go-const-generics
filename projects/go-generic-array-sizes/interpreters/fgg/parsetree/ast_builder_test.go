package parsetree

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/parser"
)

//go:embed testdata/fg/hello.go
var helloFg []byte

func TestAntlrAstBuilder_givenHelloFgProgram_buildsAst(t *testing.T) {
	input := antlr.NewIoStream(bytes.NewBuffer(helloFg))
	lexer := parser.NewFGGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGGParser(stream)
	p.BuildParseTrees = true

	tree := p.Program()
	astBuilder := NewAntlrASTBuilder(tree)
	actualAst := astBuilder.BuildAST()

	expectedAst := ast.Program{
		Declarations: []ast.Declaration{
			ast.TypeDeclaration{
				TypeName:    "any",
				TypeLiteral: ast.InterfaceTypeLiteral{},
			},
			ast.TypeDeclaration{
				TypeName: "AnyArray2",
				TypeLiteral: ast.ArrayTypeLiteral{
					Length:      ast.IntegerLiteral{IntValue: 2},
					ElementType: ast.NamedType{TypeName: "any"},
				},
			},
			ast.TypeDeclaration{
				TypeName: "Foo",
				TypeLiteral: ast.StructTypeLiteral{
					Fields: []ast.Field{
						{
							Name: "x",
							Type: ast.NamedType{TypeName: "int"},
						},
						{
							Name: "y",
							Type: ast.NamedType{TypeName: "string"},
						},
					},
				},
			},
			ast.TypeDeclaration{
				TypeName: "Fooer",
				TypeLiteral: ast.InterfaceTypeLiteral{
					MethodSpecifications: []ast.MethodSpecification{
						{
							MethodName: "foo",
							MethodSignature: ast.MethodSignature{
								MethodParameters: []ast.MethodParameter{
									{
										ParameterName: "x",
										Type:          ast.NamedType{TypeName: "Foo"},
									},
								},
								ReturnType: ast.NamedType{TypeName: "any"},
							},
						},
					},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodReceiver{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "First",
					MethodSignature: ast.MethodSignature{
						ReturnType: ast.NamedType{TypeName: "any"},
					},
				},
				ReturnExpression: ast.ArrayIndex{
					Receiver: ast.Variable{Id: "this"},
					Index:    ast.IntegerLiteral{IntValue: 0},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodReceiver{
					ParameterName: "_1",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "Nth",
					MethodSignature: ast.MethodSignature{
						MethodParameters: []ast.MethodParameter{
							{
								ParameterName: "_n",
								Type:          ast.NamedType{TypeName: "int"},
							},
						},
						ReturnType: ast.NamedType{TypeName: "any"},
					},
				},
				ReturnExpression: ast.ArrayIndex{
					Receiver: ast.Variable{Id: "_1"},
					Index:    ast.Variable{Id: "_n"},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodReceiver{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "foo",
					MethodSignature: ast.MethodSignature{
						MethodParameters: []ast.MethodParameter{
							{
								ParameterName: "foo",
								Type:          ast.NamedType{TypeName: "Foo"},
							},
						},
						ReturnType: ast.NamedType{TypeName: "string"},
					},
				},
				ReturnExpression: ast.Select{
					Receiver:  ast.Variable{Id: "foo"},
					FieldName: "y",
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodReceiver{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "Length",
					MethodSignature: ast.MethodSignature{
						ReturnType: ast.NamedType{TypeName: "int"},
					},
				},
				ReturnExpression: ast.IntegerLiteral{IntValue: 98765432101},
			},
			ast.ArraySetMethodDeclaration{
				MethodReceiver: ast.MethodReceiver{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodName: "Set",
				IndexParameter: ast.MethodParameter{
					ParameterName: "i",
					Type:          ast.NamedType{TypeName: "int"},
				},
				ValueParameter: ast.MethodParameter{
					ParameterName: "x",
					Type:          ast.NamedType{TypeName: "any"},
				},
				ReturnType:            ast.NamedType{TypeName: "AnyArray2"},
				IndexReceiverVariable: "this",
				IndexVariable:         "i",
				NewValueVariable:      "x",
				ReturnVariable:        "this",
			},
		},
		Expression: ast.MethodCall{
			Receiver: ast.MethodCall{
				Receiver: ast.ValueLiteral{
					Type: ast.NamedType{TypeName: "AnyArray2"},
					Values: []ast.Expression{
						ast.IntegerLiteral{IntValue: 1},
						ast.IntegerLiteral{IntValue: 2},
					},
				},
				MethodName: "Set",
				Arguments: []ast.Expression{
					ast.IntegerLiteral{IntValue: 0},
					ast.IntegerLiteral{IntValue: 3},
				},
			},
			MethodName: "First",
		},
	}

	require.Equal(t, expectedAst, actualAst)
}
