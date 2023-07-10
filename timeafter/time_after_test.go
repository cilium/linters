// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package timeafter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAllAnalysis(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
