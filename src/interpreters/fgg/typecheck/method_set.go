package typecheck

import (
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

type MethodSet []ast.MethodSpecification

func (m MethodSet) get(methodName string) *ast.MethodSpecification {
	for _, method := range m {
		if method.MethodName == methodName {
			return &method
		}
	}
	return nil
}

func (t typeEnvTypeCheckingVisitor) methods(astType ast.Type) MethodSet {
	return t.newMethodVisitor().methodsOf(astType)
}

type methodVisitor struct {
	typeEnvTypeCheckingVisitor
}

func (t typeEnvTypeCheckingVisitor) newMethodVisitor() methodVisitor {
	return methodVisitor{typeEnvTypeCheckingVisitor: t}
}

func (v methodVisitor) methodsOf(astType ast.Type) MethodSet {
	return astType.AcceptMethodVisitor(v)
}

func (v methodVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) []ast.MethodSpecification {
	return nil
}

func (v methodVisitor) VisitConstType(c ast.ConstType) []ast.MethodSpecification {
	panic("const type should not be used in context where method set of type is required")
}

func (v methodVisitor) VisitNamedType(n ast.NamedType) []ast.MethodSpecification {
	if n.TypeName == intTypeName {
		return nil
	}
	typeDecl := v.typeDeclarationOf(n.TypeName)
	baseMethods := v.baseMethods(n, typeDecl)
	substituter, err := newTypeParamSubstituter(n.TypeArguments, typeDecl.TypeParameters)
	if err != nil {
		return nil
	}
	substitutedMethods := make([]ast.MethodSpecification, 0, len(baseMethods))
	for _, method := range baseMethods {
		subsitutedMethod := substituter.substituteTypeParams(method).(ast.MethodSpecification)
		substitutedMethods = append(substitutedMethods, subsitutedMethod)
	}
	return substitutedMethods
}

func (v methodVisitor) baseMethods(n ast.NamedType, typeDecl ast.TypeDeclaration) []ast.MethodSpecification {
	switch typeDecl.TypeLiteral.(type) {
	case ast.StructTypeLiteral:
		return v.valueTypeMethods(n.TypeName)
	case ast.ArrayTypeLiteral:
		return v.valueTypeMethods(n.TypeName)
	case ast.InterfaceTypeLiteral:
		return typeDecl.TypeLiteral.(ast.InterfaceTypeLiteral).MethodSpecifications
	default:
		panic("unhandled type literal type")
	}
}

func (v methodVisitor) VisitTypeParameter(t ast.TypeParameter) []ast.MethodSpecification {
	// a type parameter's method set is equal to its bound, if the bound is an interface type
	// apart from "const" (which has no methods), no other type of bound is allowed in FGG
	return v.methodsOf(v.typeEnv[t])
}

func (v methodVisitor) valueTypeMethods(typeName ast.TypeName) []ast.MethodSpecification {
	res := []ast.MethodSpecification{}

	for _, decl := range v.declarations {
		methodDecl, isMethodDecl := decl.(ast.MethodDeclaration)
		if isMethodDecl && methodDecl.MethodReceiver.TypeName == typeName {
			res = append(res, methodDecl.MethodSpecification)
			continue
		}
		arraySetMethodDecl, isArraySetMethodDecl := decl.(ast.ArraySetMethodDeclaration)
		if isArraySetMethodDecl && arraySetMethodDecl.MethodReceiver.TypeName == typeName {
			res = append(res, arraySetMethodDecl.MethodSpecification())
		}
	}
	return res
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
		if cmp.Equal(m, method, cmpopts.EquateEmpty()) {
			return true
		}
	}
	return false
}
