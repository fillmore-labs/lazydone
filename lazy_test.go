// Copyright 2024 Oliver Eikemeier. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package lazydone_test

import (
	"strconv"
	"sync"
	"testing"

	"fillmore-labs.com/lazydone"
)

func TestDone(t *testing.T) {
	t.Parallel()

	for i := range 1_000 {
		t.Run("run"+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			var lazy lazydone.Lazy
			var wg sync.WaitGroup
			for range 1_000 {
				wg.Add(1)
				go func() {
					defer wg.Done()
					<-lazy.Done()
				}()
			}
			lazy.Close()
			wg.Wait()
		})
	}
}

func TestClosed(t *testing.T) {
	t.Parallel()
	var lazy lazydone.Lazy
	if lazy.Closed() {
		t.Error("Expected null lazy not to be closed")
	}

	select {
	case <-lazy.Done():
		t.Error("Expected null lazy not to be done")
	default:
	}

	if lazy.Closed() {
		t.Error("Expected lazy still not to be closed")
	}

	lazy.Close()

	if !lazy.Closed() {
		t.Error("Expected lazy to be closed after Close()")
	}
}

func TestClosedConcurrency(t *testing.T) {
	t.Parallel()

	var wg sync.WaitGroup
	for range 100 {
		var lazy lazydone.Lazy

		wg.Add(3)
		go func() {
			<-lazy.Done()
			wg.Done()
		}()
		go func() {
			for !lazy.Closed() { //nolint:revive
				// Spin, we want to hit the “select on closed channel” branch
			}
			wg.Done()
		}()
		go func() {
			lazy.Close()
			wg.Done()
		}()
	}

	wg.Wait()
}
