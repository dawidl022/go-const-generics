package typecheck

import (
	"reflect"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

type MethodSet []ast.MethodSpecification

func (t typeCheckingVisitor) methods(astType ast.Type) MethodSet {
	return t.newMethodVisitor().methodsOf(astType)
}

type methodVisitor struct {
	typeCheckingVisitor
}

func (t typeCheckingVisitor) newMethodVisitor() methodVisitor {
	return methodVisitor{typeCheckingVisitor: t}
}

func (v methodVisitor) methodsOf(astType ast.Type) MethodSet {
	return astType.AcceptMethodVisitor(v)
}

func (v methodVisitor) VisitTypeName(typeName ast.TypeName) []ast.MethodSpecification {
	if typeName == intTypeName {
		return nil
	}
	typeDecl := v.typeDeclarationOf(typeName)
	switch typeDecl.TypeLiteral.(type) {
	case ast.StructTypeLiteral:
		return v.valueTypeMethods(typeName)
	case ast.ArrayTypeLiteral:
		return v.valueTypeMethods(typeName)
	case ast.InterfaceTypeLiteral:
		return typeDecl.TypeLiteral.(ast.InterfaceTypeLiteral).MethodSpecifications
	default:
		panic("unhandled type literal type")
	}
}

func (v methodVisitor) valueTypeMethods(typeName ast.TypeName) []ast.MethodSpecification {
	res := []ast.MethodSpecification{}

	for _, decl := range v.declarations {
		methodDecl, isMethodDecl := decl.(ast.MethodDeclaration)
		if isMethodDecl && methodDecl.MethodReceiver.TypeName == typeName {
			res = append(res, methodDecl.MethodSpecification)
			continue
		}
		_, isArraySetMethodDecl := decl.(ast.ArraySetMethodDeclaration)
		if isArraySetMethodDecl {
			// TODO
			panic("untested path")
		}
	}
	return res
}

func (v methodVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) []ast.MethodSpecification {
	return nil
}

func methodDifference(super MethodSet, sub MethodSet) MethodSet {
	missingMethods := MethodSet{}
	for _, method := range super {
		if !hasMethod(sub, method) {
			missingMethods = append(missingMethods, method)
		}
	}
	return missingMethods
}

func hasMethod(methodSet MethodSet, method ast.MethodSpecification) bool {
	for _, m := range methodSet {
		if reflect.DeepEqual(m, method) {
			return true
		}
	}
	return false
}
