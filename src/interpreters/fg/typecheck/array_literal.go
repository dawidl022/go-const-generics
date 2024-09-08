package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
)

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	elemType := t.elementType(v.TypeName)
	if len(v.Values) != t.len(v.TypeName) {
		return fmt.Errorf("expected %d values in array literal of type %q but got %d",
			t.len(v.TypeName), v.TypeName, len(v.Values))
	}
	for _, val := range v.Values {
		valType, err := t.typeOf(val)
		if err != nil {
			return err
		}
		if err := t.CheckIsSubtypeOf(valType, elemType); err != nil {
			return fmt.Errorf("cannot use %q as element of array of type %q: %w", val, v.TypeName, err)
		}
	}
	return nil
}

func (t TypeCheckingVisitor) elementType(typeName ast.TypeName) ast.TypeName {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementTypeName
}

func (t TypeCheckingVisitor) len(typeName ast.TypeName) int {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).Length
}
