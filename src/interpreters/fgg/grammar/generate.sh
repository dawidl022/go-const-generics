#!/usr/bin/env sh

java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool \
	-Dlanguage=Go -visitor -package parser -o ../parser *.g4
