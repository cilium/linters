package main

import (
	"github.com/cilium/customvet/ioreadall"
	"github.com/cilium/customvet/timeafter"

	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(timeafter.Analyzer, ioreadall.Analyzer)
}
