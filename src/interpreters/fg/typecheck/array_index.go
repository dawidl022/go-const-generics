package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
)

func (t typeVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Type, error) {
	namedReceiverType, arrayTypeDecl, err := t.arrayIndexReceiverType(a)
	if err != nil {
		return nil, err
	}
	indexType, err := t.arrayIndexIndexType(a)
	if err != nil {
		return nil, err
	}
	err = t.checkArrayIndexBounds(namedReceiverType, arrayTypeDecl, indexType)
	if err != nil {
		return nil, err
	}
	return t.elementType(namedReceiverType), nil
}

func (t typeVisitor) arrayIndexReceiverType(a ast.ArrayIndex) (ast.TypeName, *ast.ArrayTypeLiteral, error) {
	receiverType, err := t.typeOf(a.Receiver)
	if err != nil {
		return "", nil, err
	}
	namedReceiverType, isNamedReceiverType := receiverType.(ast.TypeName)
	if !isNamedReceiverType || namedReceiverType == intTypeName {
		return "", nil, fmt.Errorf("cannot perform array index on value of primitive type %q", receiverType)
	}
	decl := t.typeDeclarationOf(namedReceiverType)
	arrayTypeDecl, isArrayTypeLitDecl := decl.TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayTypeLitDecl {
		return "", nil, fmt.Errorf("cannot perform array index on value of non-array type %q", receiverType)
	}
	return namedReceiverType, &arrayTypeDecl, nil
}

func (t typeVisitor) arrayIndexIndexType(a ast.ArrayIndex) (ast.Type, error) {
	indexType, err := t.typeOf(a.Index)
	if err != nil {
		return nil, err
	}
	if err := t.CheckIsSubtypeOf(indexType, intTypeName); err != nil {
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
	if isIntLiteral && (intLiteral.IntValue < 0 || intLiteral.IntValue >= arrayTypeDecl.Length) {
		return fmt.Errorf("cannot access index %d on array of type %q of size %d",
			intLiteral.IntValue, receiverType, arrayTypeDecl.Length)
	}
	return nil
}
