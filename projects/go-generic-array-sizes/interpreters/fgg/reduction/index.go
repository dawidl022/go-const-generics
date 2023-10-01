package reduction

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
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

	namedReceiverType, isNamedReceiverType := receiver.Type.(ast.NamedType)
	if !isNamedReceiverType {
		return nil, fmt.Errorf("type %q is not a valid value literal type", receiver.Type)
	}

	withinBounds, err := inIndexBounds(r.declarations, namedReceiverType, index)
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
				switch arrayTypeLit.Length.(type) {
				case ast.IntegerLiteral:
					return n < arrayTypeLit.Length.(ast.IntegerLiteral).IntValue, nil
				case ast.NamedType:
					typeParam := ast.TypeParameter(arrayTypeLit.Length.(ast.NamedType).TypeName)
					if len(arrayType.TypeArguments) < len(typeDecl.TypeParameters) {
						return false, fmt.Errorf("badly instantiated type %q: "+
							"expected %d type arguments but got %d",
							arrayType.TypeName, len(typeDecl.TypeParameters), len(arrayType.TypeArguments))
					}
					for i, param := range typeDecl.TypeParameters {
						if param.TypeParameter == typeParam {
							typeArg := arrayType.TypeArguments[i]
							size, isIntSize := typeArg.(ast.IntegerLiteral)
							if !isIntSize {
								return false, fmt.Errorf("badly instantiated type %q: "+
									"%q is not a valid constant type parameter", arrayType.TypeName, typeArg)
							}
							return n < size.IntValue, nil
						}
					}
					panic("untested path")
				}
			}
		}
	}
	return false, fmt.Errorf("no array type named %q found in declarations", arrayType.TypeName)
}
