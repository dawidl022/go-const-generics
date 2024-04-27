../../interpreters/fgg/grammar/antlr-4.13.0-complete.jar:
	cd ../../interpreters/fgg && \
	mkdir -p parser && \
	curl https://www.antlr.org/download/antlr-4.13.0-complete.jar \
	-o grammar/antlr-4.13.0-complete.jar

../../interpreters/fgg/parser: ../../interpreters/fgg/grammar/antlr-4.13.0-complete.jar ../../interpreters/fgg/grammar/*
	cd ../../interpreters/fgg && \
	go generate grammar/generate.go

test-all: ../../interpreters/fgg/parser
	go test ./...
.PHONY: test-all

test-unit: ../../interpreters/fgg/parser
	go test `go list ./... | grep -v examples`
.PHONY: test-unit

run:
	go run cmd/main.go
.PHONY: run
