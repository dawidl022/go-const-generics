module github.com/dawidl022/go-const-generics/interpreters/fgg

go 1.21.0

require (
	github.com/antlr4-go/antlr/v4 v4.13.0
	github.com/dawidl022/go-const-generics/interpreters/shared v0.1.0
	github.com/google/go-cmp v0.6.0
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/dawidl022/go-const-generics/interpreters/shared v0.1.0 => ../shared
