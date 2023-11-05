package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t TypeCheckingVisitor) CheckIsSubtypeOf(subtype ast.Type, supertype ast.Type) error {
	if subtype == supertype {
		return nil
	}
	namedSupertype, isNamedSuperType := supertype.(ast.TypeName)
	if !isNamedSuperType {
		return fmt.Errorf("type %q is not a subtype of %q", subtype, supertype)
	}
	if _, isIntLiteral := subtype.(ast.IntegerLiteral); isIntLiteral && supertype == intTypeName {
		return nil
	}
	if !t.isInterfaceTypeName(namedSupertype) {
		return fmt.Errorf("type %q is not a subtype of %q", subtype, supertype)
	}
	missingMethods := methodDifference(t.methods(supertype), t.methods(subtype))
	if len(missingMethods) > 0 {
		return fmt.Errorf("type %q is not a subtype of %q: missing methods: %s",
			subtype, supertype, missingMethods)
	}
	return nil
}
