package entrypoint

import (
	"fmt"
	"io"

	fgg "github.com/dawidl022/go-generic-array-sizes/interpreters/fgg/entrypoint"
	"github.com/dawidl022/go-generic-array-sizes/monomorphisers/fgg2go/codegen"
	"github.com/dawidl022/go-generic-array-sizes/monomorphisers/fgg2go/monomo"
)

func Monomorphise(program io.Reader) (string, error) {
	p, err := fgg.Interpreter{}.ParseProgram(program)
	if err != nil {
		return "", fmt.Errorf("failed to parse program: %w", err)
	}

	err = fgg.Interpreter{}.TypeCheck(p)
	if err != nil {
		return "", fmt.Errorf("type error: %w", err)
	}

	monomoP := monomo.Monomorphise(p.Program)
	return codegen.GenerateSourceCode(monomoP), nil
}
