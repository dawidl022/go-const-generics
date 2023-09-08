package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func reduceField(declarations []ast.Declaration, s ast.Select) (ast.Value, error) {
	structTypeName := s.Expression.(ast.ValueLiteral).TypeName
	structFields, err := fields(declarations, structTypeName)
	if err != nil {
		return nil, err
	}
	for i, field := range structFields {
		if field.Name == s.FieldName {
			values := s.Expression.(ast.ValueLiteral).Values
			if len(values) <= i {
				return nil, fmt.Errorf("struct literal missing value at index %d", i)
			}
			return values[i].(ast.Value), nil
		}
	}
	return nil, fmt.Errorf("no field named %q found on struct of type %q", s.FieldName, structTypeName)
}

func fields(declarations []ast.Declaration, structTypeName string) ([]ast.Field, error) {
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)

		if isTypeDecl {
			structTypeLit, isStructTypeLit := typeDecl.TypeLiteral.(ast.StructTypeLiteral)
			if isStructTypeLit && typeDecl.TypeName == structTypeName {
				return structTypeLit.Fields, nil
			}
		}
	}
	return nil, fmt.Errorf("no struct type named %q found in declarations", structTypeName)
}
