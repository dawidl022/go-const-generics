package typecheck

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

func (t typeCheckingVisitor) elementType(typeName ast.NamedType) ast.Type {
	return t.typeDeclarationOf(typeName.TypeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementType
}
