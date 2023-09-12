package main

import (
	"fmt"
	"os"

	"github.com/dawidl022/go-generic-array-sizes/interpreters/fg/entrypoint"
)

func main() {
	output, err := entrypoint.Interpret(os.Stdin, os.Stderr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(output)
}
