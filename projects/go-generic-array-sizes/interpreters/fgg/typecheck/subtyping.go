package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
)

func (t typeEnvTypeCheckingVisitor) checkIsSubtypeOf(subtype, supertype ast.Type) error {
	if subtype.Equal(supertype) {
		return nil
	}
	if supertype.Equal(ast.ConstType{}) {
		if t.isConstSubtype(subtype) {
			return nil
		}
		return fmt.Errorf("type %q is not a subtype of %q", subtype, ast.ConstType{})
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

func (t typeEnvTypeCheckingVisitor) isConstSubtype(subtype ast.Type) bool {
	if _, isIntLiteral := subtype.(ast.IntegerLiteral); isIntLiteral {
		return true
	}
	namedSubtype, isNamedSubtype := subtype.(ast.NamedType)
	if !isNamedSubtype {
		return false
	}
	bound, isInTypeEnv := t.typeEnv[ast.TypeParameter(namedSubtype.TypeName)]
	if isInTypeEnv && (bound == ast.ConstType{}) {
		return true
	}
	return false
}
