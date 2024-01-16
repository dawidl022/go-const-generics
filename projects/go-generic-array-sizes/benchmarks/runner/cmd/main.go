package main

import (
	"log"
	"os"

	"github.com/dawidl022/go-generic-array-sizes/benchmarks/runner/entrypoint"
)

func main() {
	err := entrypoint.Run(os.Args[1], os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
