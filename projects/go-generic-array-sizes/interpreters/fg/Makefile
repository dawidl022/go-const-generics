grammar/antlr-4.13.0-complete.jar:
	mkdir -p parser && \
	curl https://www.antlr.org/download/antlr-4.13.0-complete.jar \
	-o grammar/antlr-4.13.0-complete.jar

parser: grammar/antlr-4.13.0-complete.jar grammar/*
	go generate grammar/generate.go

test-all: parser
	go test ./...
.PHONY: test-all

test-unit: parser
	go test `go list ./... | grep -v examples`
.PHONY: test-unit

clean:
	rm -rf parser
.PHONY: clean

run:
	go run cmd/main.go
.PHONY: run
