package parsetree

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/require"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/parser"
)

//go:embed testdata/hello.go
var helloGo []byte

func TestAntlrAstBuilder_givenHelloGoProgram_buildsAst(t *testing.T) {
	input := antlr.NewIoStream(bytes.NewBuffer(helloGo))
	lexer := parser.NewFGLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewFGParser(stream)
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
					Length:          2,
					ElementTypeName: "any",
				},
			},
			ast.TypeDeclaration{
				TypeName: "Foo",
				TypeLiteral: ast.StructTypeLiteral{
					Fields: []ast.Field{
						{
							Name:     "x",
							TypeName: "int",
						},
						{
							Name:     "y",
							TypeName: "string",
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
										TypeName:      "Foo",
									},
								},
								ReturnTypeName: "any",
							},
						},
					},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "First",
					MethodSignature: ast.MethodSignature{
						ReturnTypeName: "any",
					},
				},
				ReturnExpression: ast.ArrayIndex{
					Receiver: ast.Variable{Id: "this"},
					Index:    ast.IntegerLiteral{IntValue: 0},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "_1",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "Nth",
					MethodSignature: ast.MethodSignature{
						MethodParameters: []ast.MethodParameter{
							{
								ParameterName: "_n",
								TypeName:      "int",
							},
						},
						ReturnTypeName: "any",
					},
				},
				ReturnExpression: ast.ArrayIndex{
					Receiver: ast.Variable{Id: "_1"},
					Index:    ast.Variable{Id: "_n"},
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "foo",
					MethodSignature: ast.MethodSignature{
						MethodParameters: []ast.MethodParameter{
							{
								ParameterName: "foo",
								TypeName:      "Foo",
							},
						},
						ReturnTypeName: "string",
					},
				},
				ReturnExpression: ast.Select{
					Receiver:  ast.Variable{Id: "foo"},
					FieldName: "y",
				},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "Length",
					MethodSignature: ast.MethodSignature{
						ReturnTypeName: "int",
					},
				},
				ReturnExpression: ast.IntegerLiteral{IntValue: 98765432101},
			},
			ast.MethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodSpecification: ast.MethodSpecification{
					MethodName: "Sum",
					MethodSignature: ast.MethodSignature{
						ReturnTypeName: "int",
					},
				},
				ReturnExpression: ast.Add{
					Left:  ast.IntegerLiteral{IntValue: 1},
					Right: ast.IntegerLiteral{IntValue: 2},
				},
			},
			ast.ArraySetMethodDeclaration{
				MethodReceiver: ast.MethodParameter{
					ParameterName: "this",
					TypeName:      "AnyArray2",
				},
				MethodName: "Set",
				IndexParameter: ast.MethodParameter{
					ParameterName: "i",
					TypeName:      "int",
				},
				ValueParameter: ast.MethodParameter{
					ParameterName: "x",
					TypeName:      "any",
				},
				ReturnType:            "AnyArray2",
				IndexReceiverVariable: "this",
				IndexVariable:         "i",
				NewValueVariable:      "x",
				ReturnVariable:        "this",
			},
		},
		Expression: ast.MethodCall{
			Receiver: ast.MethodCall{
				Receiver: ast.ValueLiteral{
					TypeName: "AnyArray2",
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
