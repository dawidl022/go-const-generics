package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) typeOf(variableEnv map[string]ast.TypeName, expression ast.Expression) (ast.TypeName, error) {
	return newTypeVisitor(variableEnv).typeOf(expression)
}

type typeVisitor struct {
	variableEnv map[string]ast.TypeName
}

func newTypeVisitor(variableEnv map[string]ast.TypeName) typeVisitor {
	return typeVisitor{variableEnv: variableEnv}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.TypeName, error) {
	return expression.Accept(t)
}

func (t typeVisitor) VisitVariable(v ast.Variable) (ast.TypeName, error) {
	if varType, isVarInEnv := t.variableEnv[v.Id]; isVarInEnv {
		return varType, nil
	}
	panic("untested branch")
}

func (t typeCheckingVisitor) checkIsSubtypeOf(subtype ast.TypeName, supertype ast.TypeName) error {
	if subtype == supertype {
		return nil
	}
	// TODO integer literal is subtype of int
	for _, decl := range t.declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if !isTypeDecl {
			continue
		}
		_, isInterfaceDecl := typeDecl.TypeLiteral.(ast.InterfaceTypeLiteral)
		if !isInterfaceDecl {
			continue
		}
		fmt.Println(subtype)
		missingMethods := methodDifference(t.methods(supertype), t.methods(subtype))
		if len(missingMethods) > 0 {
			return fmt.Errorf("type %q is not a subtype of %q: missing methods: %s", subtype, supertype, missingMethods) // TODO include methods that are missing in error message
		}
		return nil
	}
	return fmt.Errorf("type %q is not a subtype of %q", subtype, supertype)
}

func methodDifference(super MethodSet, sub MethodSet) MethodSet {
	missingMethods := MethodSet{}
supLoop:
	for _, supM := range super {
	subLoop:
		for _, subM := range sub {
			if supM.MethodName != subM.MethodName {
				continue
			}
			if supM.MethodSignature.ReturnTypeName != subM.MethodSignature.ReturnTypeName {
				continue
			}
			for i, param := range supM.MethodSignature.MethodParameters {
				if param.ParameterName != subM.MethodSignature.MethodParameters[i].ParameterName {
					continue subLoop
				}
				if param.TypeName != subM.MethodSignature.MethodParameters[i].TypeName {
					continue subLoop
				}
				continue supLoop
			}
		}
		missingMethods = append(missingMethods, supM)
	}
	return missingMethods
}

type MethodSet []ast.MethodSpecification

func (m MethodSet) String() string {
	s := ""

	for i, method := range m {
		if i > 0 {
			s += ", "
		}
		s += fmt.Sprintf("%q", method)
	}

	return s
}

func (t typeCheckingVisitor) methods(typeName ast.TypeName) MethodSet {
	// TODO extract int into constant "intTypeName"

	typeDecl := t.typeDeclarationOf(typeName)
	switch typeDecl.TypeLiteral.(type) {
	case ast.StructTypeLiteral:
		return t.valueTypeMethods(typeName)
	case ast.ArrayTypeLiteral:
		return t.valueTypeMethods(typeName)
	case ast.InterfaceTypeLiteral:
		return typeDecl.TypeLiteral.(ast.InterfaceTypeLiteral).MethodSpecifications
	default:
		panic("unhandled type literal type")
	}
}

func (t typeCheckingVisitor) typeDeclarationOf(typeName ast.TypeName) ast.TypeDeclaration {
	for _, decl := range t.declarations {
		if typeDecl, isTypeDecl := decl.(ast.TypeDeclaration); isTypeDecl && typeDecl.TypeName == typeName {
			return typeDecl
		}
	}
	panic("could not find declaration for typename")
}

func (t typeCheckingVisitor) valueTypeMethods(typeName ast.TypeName) []ast.MethodSpecification {
	res := []ast.MethodSpecification{}

	for _, decl := range t.declarations {
		methodDecl, isMethodDecl := decl.(ast.MethodDeclaration)
		if isMethodDecl {
			res = append(res, methodDecl.MethodSpecification)
			continue
		}
		_, isArraySetMethodDecl := decl.(ast.ArraySetMethodDeclaration)
		if isArraySetMethodDecl {
			panic("untested path")
		}
	}
	return res
}
