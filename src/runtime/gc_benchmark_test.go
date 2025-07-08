// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"testing"
)

func BenchmarkGCSmallObjects(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000000; j++ {
			_new([16]byte)
		}
		runti.GC()
	}
}
