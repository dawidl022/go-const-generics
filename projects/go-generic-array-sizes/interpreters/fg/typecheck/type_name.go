package typecheck

import "github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"

const intTypeName ast.TypeName = "int"

func isTypeNameOfType[T ast.TypeLiteral](declarations []ast.Declaration, typeName ast.TypeName) bool {
	for _, decl := range declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if !isTypeDecl {
			continue
		}
		_, isTypeDeclOfType := typeDecl.TypeLiteral.(T)
		if isTypeDeclOfType && typeDecl.TypeName == typeName {
			return true
		}
	}
	return false
}

func (t TypeCheckingVisitor) isInterfaceTypeName(typeName ast.TypeName) bool {
	return isTypeNameOfType[ast.InterfaceTypeLiteral](t.declarations, typeName)
}

func (t typeVisitor) isStructTypeName(typeName ast.TypeName) bool {
	return isTypeNameOfType[ast.StructTypeLiteral](t.declarations, typeName)
}

func (t typeVisitor) isArrayTypeName(typeName ast.TypeName) bool {
	return isTypeNameOfType[ast.ArrayTypeLiteral](t.declarations, typeName)
}
