package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
	"github.com/dawidl022/go-const-generics/interpreters/shared/auxiliary"
)

func (t typeCheckingVisitor) VisitArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) error {
	err := t.typeCheckArraySetMethodDeclaration(a)
	if err != nil {
		return fmt.Errorf(`array-set method "%s.%s": %w`, a.MethodReceiver.TypeName, a.GetMethodName(), err)
	}
	return nil
}

func (t typeCheckingVisitor) typeCheckArraySetMethodDeclaration(a ast.ArraySetMethodDeclaration) error {
	err := t.checkArraySetSignature(a)
	if err != nil {
		return err
	}
	return t.checkArraySetMethodBody(a)
}

func (t typeCheckingVisitor) checkArraySetSignature(a ast.ArraySetMethodDeclaration) error {
	if err := auxiliary.Distinct(paramNames(a)); err != nil {
		return err
	}
	_, isArrayType := t.typeDeclarationOf(a.MethodReceiver.TypeName).TypeLiteral.(ast.ArrayTypeLiteral)
	if !isArrayType {
		return fmt.Errorf("method receiver must be of array type")
	}
	if !a.IndexParameter.Type.Equal(ast.NamedType{TypeName: intTypeName}) {
		return fmt.Errorf(`first parameter %q must be of type "int"`, a.IndexParameter.ParameterName)
	}
	receiverType, typeParams, err := t.getReceiverType(a.MethodReceiver)
	if err != nil {
		return err
	}
	envChecker := t.NewTypeEnvTypeCheckingVisitor(typeParams)
	if err := envChecker.checkArraySetValueParam(a); err != nil {
		return err
	}
	if !a.ReturnType.Equal(receiverType) {
		return fmt.Errorf("return type must be same as receiver type %q", receiverType)
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

func (t typeEnvTypeCheckingVisitor) checkArraySetValueParam(a ast.ArraySetMethodDeclaration) error {
	err := t.CheckIsSubtypeOf(a.ValueParameter.Type, t.elementType(a.MethodReceiver.TypeName))
	if err != nil {
		return fmt.Errorf("second parameter %q cannot be used as element of array type %q: %w",
			a.ValueParameter.ParameterName, a.MethodReceiver.TypeName, err)
	}
	return nil
}

func (t typeCheckingVisitor) checkArraySetMethodBody(a ast.ArraySetMethodDeclaration) error {
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
