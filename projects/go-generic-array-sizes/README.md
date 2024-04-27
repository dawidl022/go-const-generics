# Go Generic Array Sizes

## General project structure

The project consists of the following components:

- [`benchmarks`](./benchmarks/) consisting of a benchmark runner capable of
  producing plots and example programs for benchmarking.
- [`interpreters`](./interpreters/) contains the FGA (fg) and FGGA (fgg)
  interpreters
- [`monomorphisers`](./monomorphisers/) contains the FGGA to Go monomorphiser
- [`proposal`](./proposal/) contains drafts of proposals and comments [submitted
  on GitHub](https://github.com/golang/go/issues/65555)
- [`theory`](./theory/) contains the formal rules of FGA, FGGA, and
  monomorphisation, along with some example derivations in LaTeX format

Example programs which are tested are found in `testdata` directories in the
packages under test.

## Local development

### Prerequisites

Ensure you have Make and [Go](https://go.dev/) installed.

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
