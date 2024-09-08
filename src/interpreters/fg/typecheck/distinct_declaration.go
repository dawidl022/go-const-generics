package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/shared/auxiliary"
)

func checkDistinctTypeDeclarations(p ast.Program) error {
	typeDecl := typeDeclarationNames(p.Declarations)
	typeDecl = append(typeDecl, intTypeName)

	err := auxiliary.Distinct(typeDecl)
	if err != nil {
		return fmt.Errorf("non distinct type declarations: %w", err)
	}
	return nil
}

func checkDistinctMethodDeclarations(p ast.Program) error {
	methodDecl := methodDeclarationNames(p.Declarations)

	err := auxiliary.Distinct(methodDecl)
	if err != nil {
		return fmt.Errorf("non distinct method declarations: %w", err)
	}
	return nil
}

func typeDeclarationNames(decl []ast.Declaration) []ast.TypeName {
	res := []ast.TypeName{}
	for _, d := range decl {
		if typeDecl, isTypeDecl := d.(ast.TypeDeclaration); isTypeDecl {
			res = append(res, typeDecl.TypeName)
		}
	}
	return res
}

func methodDeclarationNames(decl []ast.Declaration) []methodDeclarationName {
	res := []methodDeclarationName{}
	for _, d := range decl {
		if methodDecl, isMethodDecl := d.(ast.MethodDeclaration); isMethodDecl {
			res = append(res, methodDeclarationName{
				typeName:   methodDecl.MethodReceiver.TypeName,
				methodName: methodDecl.GetMethodName(),
			})
		}
	}
	return res
}

type methodDeclarationName struct {
	typeName   ast.TypeName
	methodName string
}

func (m methodDeclarationName) String() string {
	return fmt.Sprintf("%s.%s", m.typeName, m.methodName)
}
