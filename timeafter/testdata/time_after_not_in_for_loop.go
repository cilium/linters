// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package testdata

import (
	"fmt"
	"time"
)

func TimeAfterNotInForLoop() {
	<-time.After(time.Millisecond)
}

func TimeAfterNotInForLoop2() {
	select {
	case <-time.After(time.Millisecond):
	}
}

func TimeAfterInForLoop() {
	for i, l := 0, 3; i < l; i++ {
		fmt.Printf("time_after_perumutation_%d", i)
		<-time.After(time.Millisecond) // want `use of time.After in a for loop is prohibited, use inctimer instead`
	}
}

func TimeAfterInForRangeLoop() {
	for _, n := range []int{0, 1, 2} {
		fmt.Printf("time_after_perumutation_%d", n)
		<-time.After(time.Millisecond) // want `use of time.After in a for loop is prohibited, use inctimer instead`
	}
}
