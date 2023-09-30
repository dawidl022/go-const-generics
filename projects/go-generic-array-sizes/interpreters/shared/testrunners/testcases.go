package testrunners

import "fmt"

type TestCase interface {
	Name() string
}

type ReductionTestCase interface {
	TestCase
	ParseAndReduce(program []byte) (fmt.Stringer, error)
	IsValue(program []byte) bool
}

type ValueTestCase interface {
	ReductionTestCase
	ParseAndValue(program []byte) fmt.Stringer
}
