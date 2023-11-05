package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/ast"
	"github.com/dawidl022/go-generic-array-sizes/interpreters/shared/auxiliary"
)

func (t TypeCheckingVisitor) VisitArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) error {
	err := t.typeCheckArraySetMethodDeclaration(a)
	if err != nil {
		return fmt.Errorf(`array-set method "%s.%s": %w`, a.MethodReceiver.TypeName, a.GetMethodName(), err)
	}
	return nil
}

func (t TypeCheckingVisitor) typeCheckArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) error {
	err := t.checkArraySetSignature(a)
	if err != nil {
		return err
	}
	return t.checkArraySetMethodBody(a)
}

func (t TypeCheckingVisitor) checkArraySetSignature(a ast.ArraySetMethodDeclaration) error {
	if err := auxiliary.Distinct(paramNames(a)); err != nil {
		return err
	}
	_, isArrayType := t.typeDeclarationOf(a.MethodReceiver.TypeName).TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayType {
		return fmt.Errorf("method receiver must be of array type")
	}
	if a.IndexParameter.TypeName != intTypeName {
		return fmt.Errorf(`first parameter %q must be of type "int"`, a.IndexParameter.ParameterName)
	}
	err := t.CheckIsSubtypeOf(a.ValueParameter.TypeName, t.elementType(a.MethodReceiver.TypeName))
	if err != nil {
		return fmt.Errorf("second parameter %q cannot be used as element of array type %q: %w",
			a.ValueParameter.ParameterName, a.MethodReceiver.TypeName, err)
	}
	return nil
}

func paramNames(a ast.ArraySetMethodDeclaration) []name {
	return []name{
		name(a.MethodReceiver.ParameterName),
		name(a.IndexParameter.ParameterName),
		name(a.ValueParameter.ParameterName),
	}
}

func (t TypeCheckingVisitor) checkArraySetMethodBody(a ast.ArraySetMethodDeclaration) error {
	if a.MethodReceiver.TypeName != a.ReturnType {
		return fmt.Errorf("return type must be same as receiver type %q", a.MethodReceiver.TypeName)
	}
	if a.IndexReceiverVariable != a.MethodReceiver.ParameterName {
		return fmt.Errorf("index receiver must be the same as method receiver %q", a.MethodReceiver.ParameterName)
	}
	if a.IndexVariable != a.IndexParameter.ParameterName {
		return fmt.Errorf("index argument must be the same as first parameter %q", a.IndexParameter.ParameterName)
	}
	if a.NewValueVariable != a.ValueParameter.ParameterName {
		return fmt.Errorf("new index value must be the same as second parameter %q", a.ValueParameter.ParameterName)
	}
	if a.ReturnVariable != a.MethodReceiver.ParameterName {
		return fmt.Errorf("return variable must be the same as method receiver %q", a.MethodReceiver.ParameterName)
	}
	return nil
}
