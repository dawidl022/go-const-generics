grammar/antlr-4.13.0-complete.jar:
	mkdir -p parser && \
	curl https://www.antlr.org/download/antlr-4.13.0-complete.jar \
	-o grammar/antlr-4.13.0-complete.jar

parser: grammar/antlr-4.13.0-complete.jar grammar/*
	go generate grammar/generate.go

clean:
	rm -rf parser
.PHONY: clean
