package auxiliary

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func Fields(declarations []ast.Declaration, typ ast.NamedType) ([]ast.Field, error) {
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)

		if isTypeDecl {
			structTypeLit, isStructLit := typeDecl.TypeLiteral.(ast.StructTypeLiteral)
			if isStructLit && typeDecl.TypeName == typ.TypeName {
				return structTypeLit.Fields, nil
			}
		}
	}
	return nil, fmt.Errorf("no struct type named %q found in declarations", typ.TypeName)
}
