package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeCheckingVisitor) elementType(typeName ast.TypeName) ast.Type {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementType
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	namedValueType, isNamedValueType := v.Type.(ast.NamedType)
	if !isNamedValueType {
		panic("untested branch")
	}
	elemType := t.elementType(namedValueType.TypeName)
	for _, val := range v.Values {
		valType, err := t.typeOf(val)
		if err != nil {
			return err
		}
		if err := t.checkIsSubtypeOf(valType, elemType); err != nil {
			return fmt.Errorf("cannot use %q as element of array of type %q: %w", val, namedValueType.TypeName, err)
		}
	}
	return nil
}
