package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Type, error) {
	receiverType, err := t.typeOf(a.Receiver)
	if err != nil {
		return nil, err
	}
	namedReceiverType, isNamedReceiverType := receiverType.(ast.TypeName)
	if !isNamedReceiverType || namedReceiverType == intTypeName {
		return nil, fmt.Errorf("cannot perform array index on value of primitive type %q", receiverType)
	}
	decl := t.typeDeclarationOf(namedReceiverType)
	arrayTypeDecl, isArrayTypeLitDecl := decl.TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayTypeLitDecl {
		return nil, fmt.Errorf("cannot perform array index on value of non-array type %q", receiverType)
	}
	indexType, err := t.typeOf(a.Index)
	if err != nil {
		return nil, err
	}
	if err := t.checkIsSubtypeOf(indexType, intTypeName); err != nil {
		return nil, fmt.Errorf("cannot use value %q as array index argument: %w", a.Index, err)
	}
	intLiteral, isIntLiteral := indexType.(ast.IntegerLiteral)
	if isIntLiteral && (intLiteral.IntValue < 0 || intLiteral.IntValue >= arrayTypeDecl.Length) {
		return nil, fmt.Errorf("cannot access index %d on array of type %q of size %d",
			intLiteral.IntValue, receiverType, arrayTypeDecl.Length)
	}
	return t.elementType(namedReceiverType), nil
}
