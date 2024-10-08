package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/fgg/auxiliary"
)

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	fields, err := auxiliary.Fields(t.declarations, v.Type)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	typeDecl := t.typeDeclarationOf(v.Type.TypeName)
	substituter, err := newTypeParamSubstituter(v.Type.TypeArguments, typeDecl.TypeParameters)
	if err != nil {
		return err
	}
	if len(v.Values) != len(fields) {
		return fmt.Errorf("struct literal of type %q requires %d values, but got %d",
			v.Type.TypeName, len(fields), len(v.Values))
	}

	for i, f := range fields {
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			return err
		}
		expectedFieldType := substituter.substituteTypeParams(f.Type).(ast.Type)

		err = t.CheckIsSubtypeOf(fieldType, expectedFieldType)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w",
				v.Values[i], f.Name, v.Type.TypeName, err)
		}
	}
	return nil
}
