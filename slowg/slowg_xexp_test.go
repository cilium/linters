// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package slowg

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestExp(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, Analyzer, "xexp")
}
