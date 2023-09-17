package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func (t typeCheckingVisitor) typeOf(variableEnv map[string]ast.TypeName, expression ast.Expression) (ast.Type, error) {
	return t.newTypeVisitor(variableEnv).typeOf(expression)
}

type typeVisitor struct {
	typeCheckingVisitor
	variableEnv map[string]ast.TypeName
}

func (t typeCheckingVisitor) newTypeVisitor(variableEnv map[string]ast.TypeName) typeVisitor {
	return typeVisitor{typeCheckingVisitor: t, variableEnv: variableEnv}
}

func (t typeVisitor) typeOf(expression ast.Expression) (ast.Type, error) {
	return expression.Accept(t)
}

func (t typeVisitor) VisitVariable(v ast.Variable) (ast.Type, error) {
	if varType, isVarInEnv := t.variableEnv[v.Id]; isVarInEnv {
		return varType, nil
	}
	panic("untested branch")
}

func (t typeCheckingVisitor) checkIsSubtypeOf(subtype ast.Type, supertype ast.TypeName) error {
	if subtype == supertype {
		return nil
	}
	if _, isIntLiteral := subtype.(ast.IntegerLiteral); isIntLiteral && supertype == "int" {
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
		// TODO check against right supertype name, can probably remove loop and do check inside methods call
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

func (t typeCheckingVisitor) methods(astType ast.Type) MethodSet {
	return t.newMethodVisitor().methodsOf(astType)
}

func (t typeCheckingVisitor) newMethodVisitor() methodVisitor {
	return methodVisitor{typeCheckingVisitor: t}
}

type methodVisitor struct {
	typeCheckingVisitor
}

func (v methodVisitor) VisitTypeName(typeName ast.TypeName) []ast.MethodSpecification {
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

func (v methodVisitor) VisitIntegerLiteral(i ast.IntegerLiteral) []ast.MethodSpecification {
	return nil
}

func (v methodVisitor) methodsOf(astType ast.Type) MethodSet {
	return MethodSet(astType.AcceptMethodVisitor(v))
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

func (t typeVisitor) VisitValueLiteral(v ast.ValueLiteral) (ast.Type, error) {
	if t.isStructTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckStructLiteral(v)
	}
	if t.isArrayTypeName(v.TypeName) {
		return v.TypeName, t.typeCheckArrayLiteral(v)
	}
	return nil, fmt.Errorf("undeclared value literal type name: %s", v.TypeName)
}

func (t typeVisitor) isStructTypeName(typeName ast.TypeName) bool {
	for _, decl := range t.declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if !isTypeDecl {
			continue
		}
		_, isStructTypeLit := typeDecl.TypeLiteral.(ast.StructTypeLiteral)
		if isStructTypeLit && typeDecl.TypeName == typeName {
			return true
		}
	}
	return false
}

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	fields, err := ast.Fields(t.declarations, v.TypeName)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	for i, f := range fields {
		// TODO check less values than fields
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			panic("untested branch")
		}
		err = t.checkIsSubtypeOf(fieldType, f.TypeName)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w", v.Values[i], f.Name, v.TypeName, err)
		}
	}
	return nil
}

func (t typeVisitor) isArrayTypeName(typeName ast.TypeName) bool {
	for _, decl := range t.declarations {
		typeDecl, isTypeDecl := decl.(ast.TypeDeclaration)
		if !isTypeDecl {
			continue
		}
		_, isArrayTypeLit := typeDecl.TypeLiteral.(ast.ArrayTypeLiteral)
		if isArrayTypeLit && typeDecl.TypeName == typeName {
			return true
		}
	}
	return false
}

func (t typeVisitor) typeCheckArrayLiteral(v ast.ValueLiteral) error {
	return nil
}
