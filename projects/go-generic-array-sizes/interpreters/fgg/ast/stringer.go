package ast

import (
	"fmt"
	"strconv"
	"strings"
)

func (i IntegerLiteral) String() string {
	return strconv.Itoa(i.IntValue)
}

func (v Variable) String() string {
	return v.Id
}

func (m MethodCall) String() string {
	str := fmt.Sprintf("%s.%s(", m.Receiver, m.MethodName)

	for i, arg := range m.Arguments {
		if i > 0 {
			str += ", "
		}
		str += arg.String()
	}

	str += ")"
	return str
}

func (m MethodSpecification) String() string {
	return fmt.Sprintf("%s%s", m.MethodName, m.MethodSignature)
}

func (m MethodSignature) String() string {
	s := "("

	for i, param := range m.MethodParameters {
		if i > 0 {
			s += ", "
		}
		s += param.String()
	}
	s += fmt.Sprintf(") %s", m.ReturnType)
	return s
}

func (p MethodParameter) String() string {
	return fmt.Sprintf("%s %s", p.ParameterName, p.Type)
}

func (v ValueLiteral) String() string {
	str := v.Type.String() + "{"

	for i, val := range v.Values {
		if i > 0 {
			str += ", "
		}
		str += val.String()
	}

	str += "}"
	return str
}

func (s Select) String() string {
	return fmt.Sprintf("%s.%s", s.Receiver, s.FieldName)
}

func (a ArrayIndex) String() string {
	return fmt.Sprintf("%s[%s]", a.Receiver, a.Index)
}

func (t TypeParameter) String() string {
	return string(t)
}

func (n NamedType) String() string {
	s := n.TypeName.String()

	if len(n.TypeArguments) > 0 {
		s += "["

		for i, arg := range n.TypeArguments {
			if i > 0 {
				s += ", "
			}
			s += arg.String()
		}
		s += "]"
	}
	return s
}

func (t TypeName) String() string {
	return string(t)
}

func (c ConstType) String() string {
	return "const"
}

func (p Program) String() string {
	declarations := make([]string, 0, len(p.Declarations))
	for _, decl := range p.Declarations {
		declarations = append(declarations, decl.String()+"\n\n")
	}

	return fmt.Sprintf(`package main

%sfunc main() {
	_ = %s
}
`, strings.Join(declarations, ""), p.Expression)
}

func (t TypeDeclaration) String() string {
	// TODO type params
	params := make([]string, 0, len(t.TypeParameters))
	for _, param := range t.TypeParameters {
		params = append(params, param.String())
	}
	paramsStr := ""
	if len(t.TypeParameters) > 0 {
		paramsStr = fmt.Sprintf("[%s]", strings.Join(params, ", "))
	}
	return fmt.Sprintf("type %s%s %s", t.TypeName, paramsStr, t.TypeLiteral)
}

func (m MethodDeclaration) String() string {
	return fmt.Sprintf(`func (%s) %s {
	return %s
}`, m.MethodReceiver, m.MethodSpecification, m.ReturnExpression)
}

func (a ArraySetMethodDeclaration) String() string {
	return fmt.Sprintf(`func (%s) %s(%s, %s) %s {
	%s[%s] = %s;
	return %s
}`, a.MethodReceiver, a.MethodName, a.IndexParameter, a.ValueParameter, a.ReturnType,
		a.IndexReceiverVariable, a.IndexVariable, a.NewValueVariable, a.ReturnVariable)
}

func (s StructTypeLiteral) String() string {
	fields := make([]string, 0, len(s.Fields))
	maxNameLen := 0
	for _, field := range s.Fields {
		if len(field.Name) > maxNameLen {
			maxNameLen = len(field.Name)
		}
	}
	for _, field := range s.Fields {
		spaces := strings.Repeat(" ", maxNameLen-len(field.Name)+1)
		fields = append(fields, fmt.Sprintf("\t%s%s%s\n", field.Name, spaces, field.Type))
	}

	return fmt.Sprintf(`struct {
%s}`, strings.Join(fields, ""))
}

func (i InterfaceTypeLiteral) String() string {
	methodSpecs := make([]string, 0, len(i.MethodSpecifications))
	for _, spec := range i.MethodSpecifications {
		methodSpecs = append(methodSpecs, fmt.Sprintf("\t%s\n", spec))
	}
	return fmt.Sprintf(`interface {
%s}`, strings.Join(methodSpecs, ""))
}

func (a ArrayTypeLiteral) String() string {
	return fmt.Sprintf("[%s]%s", a.Length, a.ElementType)
}

func (t TypeParameterConstraint) String() string {
	return fmt.Sprintf("%s %s", t.TypeParameter, t.Bound)
}

func (m MethodReceiver) String() string {
	typeParams := make([]Type, 0, len(m.TypeParameters))
	for _, param := range m.TypeParameters {
		typeParams = append(typeParams, param)
	}

	return fmt.Sprintf("%s %s", m.ParameterName, NamedType{
		TypeName:      m.TypeName,
		TypeArguments: typeParams,
	})
}

func (a Add) String() string {
	return fmt.Sprintf("%s + %s", a.Left, a.Right)
}
