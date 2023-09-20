package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	elemType := t.elementType(v.TypeName)
	for _, val := range v.Values {
		valType, err := t.typeOf(val)
		if err != nil {
			return err
		}
		if err := t.checkIsSubtypeOf(valType, elemType); err != nil {
			return fmt.Errorf("cannot use %q as element of array of type %q: %w", val, v.TypeName, err)
		}
	}
	return nil
}

func (t typeCheckingVisitor) elementType(typeName ast.TypeName) ast.TypeName {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementTypeName
}
