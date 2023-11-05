package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/auxiliary"
)

func (t typeVisitor) typeCheckStructLiteral(v ast.ValueLiteral) error {
	namedValueType, isNamedValueType := v.Type.(ast.NamedType)
	if !isNamedValueType {
		panic("untested branch")
	}
	fields, err := auxiliary.Fields(t.declarations, namedValueType)
	if err != nil {
		panic("type checker should not call fields on non-struct type literal")
	}
	typeDecl := t.typeDeclarationOf(namedValueType.TypeName)
	substituter, err := newTypeParamSubstituter(namedValueType.TypeArguments, typeDecl.TypeParameters)
	if err != nil {
		return err
	}
	if len(v.Values) != len(fields) {
		return fmt.Errorf("struct literal of type %q requires %d values, but got %d",
			namedValueType.TypeName, len(fields), len(v.Values))
	}
	envChecker := t.NewTypeEnvTypeCheckingVisitor(typeDecl.TypeParameters)
	for i, f := range fields {
		fieldType, err := t.typeOf(v.Values[i])
		if err != nil {
			return err
		}
		expectedFieldType := envChecker.identifyTypeParams(f.Type).(ast.Type)
		expectedFieldType = substituter.substituteTypeParams(expectedFieldType).(ast.Type)

		err = t.CheckIsSubtypeOf(fieldType, expectedFieldType)
		if err != nil {
			return fmt.Errorf("cannot use %q as field %q of struct %q: %w",
				v.Values[i], f.Name, namedValueType.TypeName, err)
		}
	}
	return nil
}
