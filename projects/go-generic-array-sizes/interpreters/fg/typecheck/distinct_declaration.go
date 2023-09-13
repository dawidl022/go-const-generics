package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
)

func checkDistinctTypeDeclarations(p ast.Program) error {
	typeDecl := typeDeclarations(p.Declarations)
	typeDecl = append(typeDecl, "int")

	err := distinct(typeDecl)
	if err != nil {
		return fmt.Errorf("non distinct type declarations: %w", err)
	}
	return nil
}

func typeDeclarations(decl []ast.Declaration) []string {
	res := []string{}
	for _, d := range decl {
		if typeDecl, isTypeDecl := d.(ast.TypeDeclaration); isTypeDecl {
			res = append(res, typeDecl.TypeName)
		}
	}
	return res
}

func distinct(names []string) error {
	seenNames := make(map[string]struct{})

	for _, name := range names {
		if _, seen := seenNames[name]; seen {
			return fmt.Errorf("redeclared %q", name)
		}
		seenNames[name] = struct{}{}
	}
	return nil
}
