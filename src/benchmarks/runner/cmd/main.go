package main

import (
	"log"
	"os"

	"github.com/dawidl022/go-const-generics/benchmarks/runner/entrypoint"
)

func main() {
	var comparison *[2]string
	if len(os.Args) == 5 {
		comparison = &[2]string{os.Args[3], os.Args[4]}
	}

	err := entrypoint.Run(os.Args[1], os.Args[2], comparison, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
