package testrunners

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func runTests[T TestCase](t *testing.T, tests []T, testFunc func(test T) func(*testing.T)) {
	if len(tests) == 1 {
		testFunc(tests[0])(t)
		return
	}
	for _, test := range tests {
		t.Run(test.Name(), testFunc(test))
	}
}

func AssertEqualAfterSingleReduction(t *testing.T, program []byte, expected string, reductionTests []ReductionTestCase) {
	runTests(t, reductionTests, func(test ReductionTestCase) func(*testing.T) {
		return func(t *testing.T) {
			expr, err := test.ParseAndReduce(program)

			require.NoError(t, err)
			require.Equal(t, expected, expr.String())
			require.False(t, test.IsValue(program))
		}
	})
}

func AssertErrorAfterSingleReduction(t *testing.T, program []byte, expectedErrMsg string, reductionTests []ReductionTestCase) {
	runTests(t, reductionTests, func(test ReductionTestCase) func(t *testing.T) {
		return func(t *testing.T) {
			_, err := test.ParseAndReduce(program)

			require.Error(t, err)
			require.Equal(t, expectedErrMsg, err.Error())
			require.False(t, test.IsValue(program))
		}
	})
}

func AssertEqualValueAndFailsToReduce(t *testing.T, program []byte, expectedValue string, valueTests []ValueTestCase) {
	runTests(t, valueTests, func(test ValueTestCase) func(t *testing.T) {
		return func(t *testing.T) {
			require.Panics(t, func() { test.ParseAndReduce(program) })

			require.True(t, test.IsValue(program))
			require.Equal(t, expectedValue, test.ParseAndValue(program).String())
		}
	})
}

func AssertFailsTypeCheckWithError(t *testing.T, program []byte, expectedErrMsg string, typeTests []TypeTestCase) {
	runTests(t, typeTests, func(test TypeTestCase) func(*testing.T) {
		return func(t *testing.T) {
			err := test.ParseAndTypeCheck(program)

			require.Error(t, err)
			require.Equal(t, expectedErrMsg, err.Error())
		}
	})
}

func AssertPassesTypeCheck(t *testing.T, program []byte, typeTests []TypeTestCase) {
	runTests(t, typeTests, func(test TypeTestCase) func(*testing.T) {
		return func(t *testing.T) {
			err := test.ParseAndTypeCheck(program)

			require.NoError(t, err)
		}
	})
}
