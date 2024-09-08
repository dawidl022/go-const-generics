package typecheck

import (
	"fmt"

	"github.com/dawidl022/go-const-generics/interpreters/fgg/ast"
)

func (t typeVisitor) VisitMethodCall(m ast.MethodCall) (ast.Type, error) {
	receiverType, err := t.typeOf(m.Receiver)
	if err != nil {
		return nil, err
	}
	methodDecl := t.methods(receiverType).get(m.MethodName)
	if methodDecl == nil {
		return nil, fmt.Errorf("no method named %q on receiver of type %q", m.MethodName, receiverType)
	}
	params := methodDecl.MethodSignature.MethodParameters
	if len(m.Arguments) != len(params) {
		return nil, fmt.Errorf(`expected %d arguments in call to "%s.%s", but got %d`,
			len(params), receiverType, methodDecl.MethodName, len(m.Arguments))
	}
	for i, param := range params {
		arg := m.Arguments[i]
		if err := t.checkParameterType(param, arg); err != nil {
			return nil, fmt.Errorf(`cannot use %q as argument %q in call to "%s.%s": %w`,
				arg, param.ParameterName, receiverType, methodDecl.MethodName, err)
		}
	}
	return methodDecl.MethodSignature.ReturnType, nil
}

func (t typeVisitor) checkParameterType(param ast.MethodParameter, arg ast.Expression) error {
	argType, err := t.typeOf(arg)
	if err != nil {
		return err
	}
	return t.CheckIsSubtypeOf(argType, param.Type)
}
