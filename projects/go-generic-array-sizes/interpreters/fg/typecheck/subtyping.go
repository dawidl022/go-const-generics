package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) checkIsSubtypeOf(subtype ast.Type, supertype ast.TypeName) error {
	if subtype == supertype {
		return nil
	}
	if _, isIntLiteral := subtype.(ast.IntegerLiteral); isIntLiteral && supertype == intTypeName {
		return nil
	}
	if !t.isInterfaceTypeName(supertype) {
		return fmt.Errorf("type %q is not a subtype of %q", subtype, supertype)
	}
	missingMethods := methodDifference(t.methods(supertype), t.methods(subtype))
	if len(missingMethods) > 0 {
		return fmt.Errorf("type %q is not a subtype of %q: missing methods: %s",
			subtype, supertype, missingMethods)
	}
	return nil
}
