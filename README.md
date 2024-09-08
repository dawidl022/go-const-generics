# Go const generics

Theory and prototype implementation of const generics (generic array sizes) in
Go, based on [Featherweight Go](https://dl.acm.org/doi/10.1145/3428217).

This monorepo contains all the work done as part of my undergraduate final year
project, from the initial proposals, to the implementation and report.

## Documentation

The following documentation is deployed to GitHub pages:

- [report.pdf](https://dawidl022.github.io/go-const-generics/report.pdf) -
  introduces and explains in depth the theory behind const generics in Go,
  including formal definitions.
- [presentation](https://dawidl022.github.io/go-const-generics/presentation) -
  used to accompany talks introducing this project.
- [cycle-detection-summary.pdf](https://dawidl022.github.io/go-const-generics/cycle-detection-summary.pdf) - formalises cycle detection rules in Go to aid in resolving issues with cycle
   detection in the Go compiler (as of version 1.23).

## Repo structure

- [`docs`](./docs/) contains the initial project definition, presentation
  slides, and the writeup in LaTeX for the final report.
- [`docs/theory`](./docs/theory/) contains the formal rules of FGA, FGGA, and
  monomorphisation, along with some example derivations in LaTeX format.
- [`docs/proposal`](./docs/proposal/) contains drafts of proposals and comments
  [submitted on GitHub](https://github.com/golang/go/issues/65555).
- [`src/benchmarks`](./src/benchmarks/) consisting of a benchmark runner capable
  of producing plots and example programs for benchmarking.
- [`src/interpreters`](./src/interpreters/) contains the FGA (fg) and FGGA (fgg)
  interpreters.
- [`src/monomorphisers`](./src/monomorphisers/) contains the FGGA to Go
  monomorphiser.

Example programs which are tested are found in `testdata` directories in the
packages under test.

## Local development

All following commands are to be executed from the `src` directory, unless
noted otherwise.

### Prerequisites

Ensure you have Make, [Go](https://go.dev/) and [Java](https://openjdk.org/)
installed. Java is needed to run the [ANTLR](https://www.antlr.org/) parser
generator used.

### Testing

The tests are written using the standard Go [testing
package](https://pkg.go.dev/testing).

To run all the unit tests, execute:

```bash
make test-unit
```

To run all tests (which may take a couple of minutes to complete), run:

```bash
make test-all
```

### Running interpreters and monomorphiser

The interpreters and monomorphiser come with binary targets that take programs
from stdin, and output the results (final result of interpretation or
monomorphised program) to stdout. Errors and extra information (such as
individual reduction steps and type preservation checks) are output to stderr.

Use shell redirection as necessary to read/write programs from/to files.

There is a make target for each of the 3 programs:

```bash
make run-fg
```

```bash
make run-fgg
```

```bash
make run-monomo
```

By default, the interpreters run until the program termination, or until a loop
is detected (repeated main expression).

The interpreters also accept a `maxSteps` flag which sets an upper bound on the
number of reduction steps. This flag cannot be passed along with the above make
commands. Instead, execute e.g.:

```bash
cd interpreters/fg
go run cmd/main.go -maxSteps=10
```

The interpreters and monomorphiser can also be used as libraries. The
`entrypoint` package inside each module provides a facade for the most common
features. (In fact, the monomorphiser uses the FGGA interpreter as a library to
perform parsing and type checking before monomorphising a program.)

## Running benchmarks

Since arrays have a fixed size at compile time, to benchmark arrays of various
sizes a custom benchmark runner was built.

The basic usage is as follows:

```
cd benchmarks/runner
go run cmd/main.go <benchmark-package-path> <benchmarks-pattern>
```

`benchmark-package-path` should be the path that contains the benchmarks. If
arrays are being benchmarked, they should use the constant `N` for their length,
which should be defined in a file named `config.go` inside that package. This
file will be overwritten by the benchmark runner during benchmarks to benchmark
various values for `N`, so it should not contain any other code than the
constant `N`. Currently, the runner uses values from 0 to 32M (1024 * 1024 *
32), with exponential increments (2<sup>n</sup>). This can be configured if
necessary in [`runner.go`](./benchmarks/runner/runner/runner.go).

<figure>

  ```go
  package reversed

  const N = 1024
  ```
  <figcaption>Example contents of <code>config.go</code></figcaption>
</figure>

`benchmarks-pattern` should be a RegEx for the function names of benchmarks to
run. The benchmark functions should be written using the standard [testing
library](https://pkg.go.dev/testing#hdr-Benchmarks). All benchmark functions
matching the pattern will be included as part of the plot that the runner
produces, with one line for each function.

The convention for the output is that the results go into
[`runner/outputs/<benchmark-package-name>`](./benchmarks/runner/outputs/) as
tab-separated `.dat` values. In addition, a LaTeX
[PGFPlots](https://ctan.org/pkg/pgfplots) graph will be created in a
subdirectory made from the `benchmarks-pattern` (with any wildcard characters
stripped away) under `<benchmark-package-name>.tex`. This file is hard-coded to
work when imported from the [report directory](../../deliverables/2-report/).

In addition, it is possible to compare two benchmarks (calculate the relative
speedup), by passing additional CLI parameters:

```
cd benchmarks/runner
go run cmd/main.go <benchmark-package-path> <benchmarks-pattern> <target-benchmark> <base-benchmark>
```

The speedup will be calculated using `base-benchmark` as the base.

The benchmark graph in the report was generated from the following command
(output may differ slightly, as benchmarks depend on many factors, including the
host device). It may take a few minutes to complete (no progress indicator is output).

```bash
cd benchmarks/runner
go run cmd/main.go ../benchmarks/reversed 'BenchmarkReversed(?:Array|Slice)$' BenchmarkReversedArray BenchmarkReversedSlice
```
