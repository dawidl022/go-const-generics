package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dawidl022/go-const-generics/interpreters/fg/entrypoint"
	"github.com/dawidl022/go-const-generics/interpreters/shared/loop"
)

func main() {
	maxSteps := flag.Int("maxSteps", loop.UnboundedSteps, loop.MaxStepsHelp)
	flag.Parse()

	output, err := entrypoint.Interpret(os.Stdin, os.Stderr, *maxSteps)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(output)
}
