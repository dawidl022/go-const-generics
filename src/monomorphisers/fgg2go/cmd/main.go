package main

import (
	"fmt"
	"os"

	"github.com/dawidl022/go-const-generics/monomorphisers/fgg2go/entrypoint"
)

func main() {
	output, err := entrypoint.Monomorphise(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(output)
}
