package reduction

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func (r ReducingVisitor) VisitArrayIndex(a ast.ArrayIndex) (ast.Expression, error) {
	if !a.Receiver.IsValue() {
		return r.arrayIndexWithReducedReceiver(a)
	}
	receiver, isReceiverValue := a.Receiver.(ast.ValueLiteral)
	if !isReceiverValue {
		return nil, fmt.Errorf("cannot access index on primitive value %s", a.Receiver)
	}

	if !a.Index.IsValue() {
		return r.arrayIndexWithReducedIndex(a)
	}
	intArgument, isIntArgument := a.Index.(ast.IntegerLiteral)
	if !isIntArgument {
		return nil, fmt.Errorf("non integer value %q used as index argument", a.Index)
	}
	return r.reduceArrayIndex(receiver, intArgument)
}

func (r ReducingVisitor) arrayIndexWithReducedReceiver(a ast.ArrayIndex) (ast.Expression, error) {
	reducedReceiver, err := r.Reduce(a.Receiver)
	return ast.ArrayIndex{Index: a.Index, Receiver: reducedReceiver}, err
}

func (r ReducingVisitor) arrayIndexWithReducedIndex(a ast.ArrayIndex) (ast.Expression, error) {
	reducedIndex, err := r.Reduce(a.Index)
	return ast.ArrayIndex{Index: reducedIndex, Receiver: a.Receiver}, err
}

func (r ReducingVisitor) reduceArrayIndex(receiver ast.ValueLiteral, argument ast.IntegerLiteral) (ast.Expression, error) {
	index := argument.IntValue

	withinBounds, err := inIndexBounds(r.declarations, receiver.Type, index)
	if err != nil {
		return nil, fmt.Errorf("could not check index bounds of %q: %w", receiver, err)
	}
	if !withinBounds {
		return nil, fmt.Errorf("index %d out of bounds for array of type %q", index, receiver.Type)
	}
	if len(receiver.Values) <= index {
		return nil, fmt.Errorf("array literal missing value at index %d", index)
	}
	return receiver.Values[index], nil
}

func inIndexBounds(declarations []ast.Declaration, arrayType ast.NamedType, n int) (bool, error) {
	if n < 0 {
		return false, nil
	}
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)

		if isTypeDecl {
			arrayTypeLit, isArrayTypeLit := typeDecl.TypeLiteral.(ast.ArrayTypeLiteral)
			if isArrayTypeLit && typeDecl.TypeName == arrayType.TypeName {
				return checkIndexBounds(n, arrayType, typeDecl, arrayTypeLit)
			}
		}
	}
	return false, fmt.Errorf("no array type named %q found in declarations", arrayType.TypeName)
}

func checkIndexBounds(n int, arrayType ast.NamedType, typeDecl ast.TypeDeclaration, arrayTypeLit ast.ArrayTypeLiteral) (bool, error) {
	switch arrayTypeLit.Length.(type) {
	case ast.IntegerLiteral:
		return n < arrayTypeLit.Length.(ast.IntegerLiteral).IntValue, nil
	case ast.TypeParameter:
		length, err := getGenericArrayLength(arrayType, typeDecl, arrayTypeLit)
		return n < length, err
	case ast.NamedType:
		return false, fmt.Errorf("unexpected length type parameter %q in declaration of type %q; "+
			"only integer literals and type parameters may be used as length types",
			arrayTypeLit.Length, arrayType.TypeName)
	default:
		panic("unexpected Type type for Length")
	}
}

func getGenericArrayLength(arrayType ast.NamedType, arrayTypeDecl ast.TypeDeclaration, arrayTypeLit ast.ArrayTypeLiteral) (int, error) {
	err := checkTypeArgumentsCount(arrayType, arrayTypeDecl)
	if err != nil {
		return 0, err
	}
	lengthParam := arrayTypeLit.Length.(ast.TypeParameter)

	for i, param := range arrayTypeDecl.TypeParameters {
		if param.TypeParameter == lengthParam {
			return getArrayLengthFromTypeArgument(arrayType, i)
		}
	}
	return 0, fmt.Errorf("unbound length type parameter %q in declaration of type %q", lengthParam, arrayTypeDecl.TypeName)
}

func checkTypeArgumentsCount(instantiatedType ast.NamedType, typeDecl ast.TypeDeclaration) error {
	if len(instantiatedType.TypeArguments) < len(typeDecl.TypeParameters) {
		return fmt.Errorf("badly instantiated type %q: "+
			"expected %d type arguments but got %d",
			instantiatedType.TypeName, len(typeDecl.TypeParameters), len(instantiatedType.TypeArguments))
	}
	return nil
}

func getArrayLengthFromTypeArgument(arrayType ast.NamedType, i int) (int, error) {
	typeArg := arrayType.TypeArguments[i]
	size, isIntSize := typeArg.(ast.IntegerLiteral)
	if !isIntSize {
		return 0, fmt.Errorf("badly instantiated type %q: "+
			"%q is not a valid constant type parameter", arrayType.TypeName, typeArg)
	}
	return size.IntValue, nil
}
