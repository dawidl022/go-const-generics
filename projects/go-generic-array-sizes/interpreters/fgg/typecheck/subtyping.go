package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeEnvTypeCheckingVisitor) checkIsSubtypeOf(subtype, supertype ast.Type) error {
	// TODO return zero values in stub functions instead of nil so that this doesn't panic
	if subtype.Equal(supertype) {
		return nil
	}
	namedSupertype, isNamedSupertype := supertype.(ast.NamedType)
	if _, isIntLiteral := subtype.(ast.IntegerLiteral); isIntLiteral && isNamedSupertype && namedSupertype.TypeName == intTypeName {
		return nil
	}
	if !(isNamedSupertype && t.isInterfaceTypeName(namedSupertype.TypeName)) {
		return fmt.Errorf("type %q is not a subtype of %q", subtype, supertype)
	}
	missingMethods := methodDifference(t.methods(supertype), t.methods(subtype))
	if len(missingMethods) > 0 {
		return fmt.Errorf("type %q is not a subtype of %q: missing methods: %s",
			subtype, supertype, missingMethods)
	}
	return nil
}
