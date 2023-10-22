package typecheck

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"

func (t typeCheckingVisitor) elementType(typeName ast.TypeName) ast.Type {
	return t.typeDeclarationOf(typeName).TypeLiteral.(ast.ArrayTypeLiteral).ElementType
}
