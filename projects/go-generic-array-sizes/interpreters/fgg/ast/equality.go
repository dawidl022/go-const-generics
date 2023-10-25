package ast

import "slices"

func (t TypeParameter) Equal(other Type) bool {
	otherTypeParam, isOtherTypeParam := other.(TypeParameter)
	return isOtherTypeParam && t == otherTypeParam
}

func (i IntegerLiteral) Equal(other Type) bool {
	otherIntLiteral, isOtherIntLiteral := other.(IntegerLiteral)
	return isOtherIntLiteral && i == otherIntLiteral
}

func (n NamedType) Equal(other Type) bool {
	otherNamedType, isOtherNamedType := other.(NamedType)
	return isOtherNamedType && n.TypeName == otherNamedType.TypeName && n.typeParamsEqual(otherNamedType)
}

func (n NamedType) typeParamsEqual(other NamedType) bool {
	return slices.EqualFunc(n.TypeArguments, other.TypeArguments, func(t1 Type, t2 Type) bool {
		return t1.Equal(t2)
	})
}

func (c ConstType) Equal(other Type) bool {
	_, isOtherConstType := other.(ConstType)
	return isOtherConstType
}
