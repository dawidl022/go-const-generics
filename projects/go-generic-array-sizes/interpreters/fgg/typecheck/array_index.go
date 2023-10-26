package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Type, error) {
	namedReceiverType, err := t.arrayIndexReceiverType(a)
	if err != nil {
		return nil, err
	}
	indexType, err := t.arrayIndexIndexType(a)
	if err != nil {
		return nil, err
	}
	err = t.checkArrayIndexBounds(namedReceiverType, indexType)
	if err != nil {
		return nil, err
	}
	typeDecl := t.typeDeclarationOf(namedReceiverType.TypeName)
	substituter, err := newTypeParamSubstituter(namedReceiverType.TypeArguments, typeDecl.TypeParameters)
	if err != nil {
		return nil, err
	}
	elementType := t.elementType(namedReceiverType.TypeName)
	elementTypeWithParms := t.identifyTypeParams(elementType)
	return substituter.substituteTypeParams(elementTypeWithParms).(ast.Type), nil
}

func (t typeVisitor) arrayIndexReceiverType(a ast.ArrayIndex) (ast.NamedType, error) {
	receiverType, err := t.typeOf(a.Receiver)
	if err != nil {
		return ast.NamedType{}, err
	}
	namedReceiverType, isNamedReceiverType := receiverType.(ast.NamedType)
	if !isNamedReceiverType || namedReceiverType.TypeName == intTypeName {
		return ast.NamedType{}, fmt.Errorf("cannot perform array index on value of primitive type %q", receiverType)
	}
	decl := t.typeDeclarationOf(namedReceiverType.TypeName)
	_, isArrayTypeLitDecl := decl.TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayTypeLitDecl {
		return ast.NamedType{}, fmt.Errorf("cannot perform array index on value of non-array type %q", receiverType)
	}
	return namedReceiverType, nil
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
	namedReceiverType ast.NamedType,
	indexType ast.Type,
) error {
	expectedLen, hasDefinedLen := t.len(namedReceiverType).(ast.IntegerLiteral)
	intLiteral, isIntLiteral := indexType.(ast.IntegerLiteral)
	if isIntLiteral && !hasDefinedLen {
		return fmt.Errorf("cannot use int literal %q to index into array of type %q with non-concrete length",
			intLiteral, namedReceiverType)
	}
	if isIntLiteral && (intLiteral.IntValue < 0 || intLiteral.IntValue >= expectedLen.IntValue) {
		return fmt.Errorf("cannot access index %d on array of type %q of size %d",
			intLiteral.IntValue, namedReceiverType, expectedLen.IntValue)
	}
	return nil
}
