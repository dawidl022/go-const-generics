package testrunners

type ASTBuilder[T any] interface {
	BuildAST() T
}
