package ast

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
	// TODO check equality of type arguments
	return isOtherNamedType && n.TypeName == otherNamedType.TypeName
}

func (c ConstType) Equal(other Type) bool {
	_, isOtherConstType := other.(ConstType)
	return isOtherConstType
}
