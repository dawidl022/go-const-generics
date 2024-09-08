package typecheck

import (
	"github.com/dawidl022/go-const-generics/interpreters/fg/testconf"
	"github.com/dawidl022/go-const-generics/interpreters/shared/testrunners"
)

func testCases() []testrunners.TypeTestCase {
	conf := testconf.ParseTestConf()
	tests := []testrunners.TypeTestCase{}

	if conf.EnabledFG() {
		tests = append(tests, fgTestCase{})
	}
	if conf.EnabledFGG() {
		tests = append(tests, fggTestCase{})
	}
	return tests
}

type fgTestCase struct {
}

func (f fgTestCase) Name() string {
	return "FG"
}

func (f fgTestCase) ParseAndTypeCheck(program []byte) error {
	return parseAndTypeCheckFG(program)
}

type fggTestCase struct {
}

func (f fggTestCase) Name() string {
	return "FGG"
}

func (f fggTestCase) ParseAndTypeCheck(program []byte) error {
	return parseAndTypeCheckFGG(program)
}
