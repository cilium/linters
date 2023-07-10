// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package testdata

import (
	"io"
	"io/ioutil"
	"strings"
)

func IOReadAllPresent() {
	r := strings.NewReader("this is a string reader")
	io.ReadAll(r) // want `use of io.ReadAll is prohibited, use safeio.ReadAllLimit instead`
}

func IOUtilReadAllPresent() {
	r := strings.NewReader("this is a string reader")
	ioutil.ReadAll(r) // want `use of ioutil.ReadAll is prohibited, use safeio.ReadAllLimit instead`
}
