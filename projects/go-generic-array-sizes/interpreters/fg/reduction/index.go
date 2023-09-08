package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func reduceIndex(declarations []ast.Declaration, idx ast.ArrayIndex) (ast.Value, error) {
	receiver := idx.Receiver.(ast.ValueLiteral)
	arrTypeName := receiver.TypeName
	i := idx.Index.(ast.IntegerLiteral).Value

	withinBounds, err := inIndexBounds(declarations, arrTypeName, i)
	if err != nil {
		return nil, err
	}
	if !withinBounds {
		return nil, fmt.Errorf("index %d out of bounds for array of type %q", i, arrTypeName)
	}
	if len(receiver.Values) <= i {
		return nil, fmt.Errorf("array literal missing value at index %d", i)
	}
	return receiver.Values[i].(ast.Value), nil
}

func inIndexBounds(declarations []ast.Declaration, arrayTypeName string, n int) (bool, error) {
	if n < 0 {
		return false, nil
	}
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)

		if isTypeDecl {
			arrayTypeLit, isArrayTypeLit := typeDecl.TypeLiteral.(ast.ArrayTypeLiteral)
			if isArrayTypeLit && typeDecl.TypeName == arrayTypeName {
				return n < arrayTypeLit.Length, nil
			}
		}
	}
	return false, fmt.Errorf("no array type named %q found in declarations", arrayTypeName)
}
