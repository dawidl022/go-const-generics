package testconf

import (
	"flag"
	"slices"
	"strings"
)

var langsFlag = flag.String("langs", "fg", "comma-separated list of languages to run tests for")

type TestConf struct {
	langFlags []string
}

func ParseTestConf() *TestConf {
	flag.Parse()
	return &TestConf{langFlags: strings.Split(*langsFlag, ",")}
}

func (c TestConf) EnabledFG() bool {
	return slices.Contains(c.langFlags, "fg")
}

func (c TestConf) EnabledFGG() bool {
	return slices.Contains(c.langFlags, "fgg")
}
