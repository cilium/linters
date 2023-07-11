// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build go1.21

package slowg

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestStdlib(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "stdlib")
}
