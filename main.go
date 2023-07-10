package main

import (
	"github.com/cilium/linters/ioreadall"
	"github.com/cilium/linters/timeafter"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(timeafter.Analyzer, ioreadall.Analyzer)
}
