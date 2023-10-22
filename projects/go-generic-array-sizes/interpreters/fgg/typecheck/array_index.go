package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

// TODO check if all FG ArrayIndex tests pass

func (t typeVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Type, error) {
	namedReceiverType, arrayTypeDecl, err := t.arrayIndexReceiverType(a)
	if err != nil {
		return nil, err
	}
	indexType, err := t.arrayIndexIndexType(a)
	if err != nil {
		return nil, err
	}
	err = t.checkArrayIndexBounds(namedReceiverType.TypeName, arrayTypeDecl, indexType)
	if err != nil {
		return nil, err
	}
	return t.elementType(namedReceiverType.TypeName), nil
}

func (t typeVisitor) arrayIndexReceiverType(a ast.ArrayIndex) (ast.NamedType, *ast.ArrayTypeLiteral, error) {
	receiverType, err := t.typeOf(a.Receiver)
	if err != nil {
		return ast.NamedType{}, nil, err
	}
	namedReceiverType, isNamedReceiverType := receiverType.(ast.NamedType)
	if !isNamedReceiverType || namedReceiverType.TypeName == intTypeName {
		return ast.NamedType{}, nil, fmt.Errorf("cannot perform array index on value of primitive type %q", receiverType)
	}
	decl := t.typeDeclarationOf(namedReceiverType.TypeName)
	arrayTypeDecl, isArrayTypeLitDecl := decl.TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayTypeLitDecl {
		return ast.NamedType{}, nil, fmt.Errorf("cannot perform array index on value of non-array type %q", receiverType)
	}
	return namedReceiverType, &arrayTypeDecl, nil
}

func (t typeVisitor) arrayIndexIndexType(a ast.ArrayIndex) (ast.Type, error) {
	indexType, err := t.typeOf(a.Index)
	if err != nil {
		return nil, err
	}
	if err := t.checkIsSubtypeOf(indexType, ast.NamedType{TypeName: intTypeName}); err != nil {
		return nil, fmt.Errorf("cannot use value %q as array index argument: %w", a.Index, err)
	}
	return indexType, nil
}

func (t typeVisitor) checkArrayIndexBounds(
	receiverType ast.TypeName,
	arrayTypeDecl *ast.ArrayTypeLiteral,
	indexType ast.Type,
) error {
	intLiteral, isIntLiteral := indexType.(ast.IntegerLiteral)
	lengthLiteral, isLengthLiteral := arrayTypeDecl.Length.(ast.IntegerLiteral)
	if !isLengthLiteral {
		panic("untested path")
	}
	if isIntLiteral && (intLiteral.IntValue < 0 || intLiteral.IntValue >= lengthLiteral.IntValue) {
		return fmt.Errorf("cannot access index %d on array of type %q of size %d",
			intLiteral.IntValue, receiverType, lengthLiteral.IntValue)
	}
	return nil
}
