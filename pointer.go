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

package lazydone

import (
	"sync/atomic"
)

// SafeLazy is a lazily initialized done channel. It is an alternate implementation of [Lazy] using
// [sync/atomic.Pointer]. The zero value for a SafeLazy is valid and can be closed.
// A SafeLazy must not be copied after first use.
type SafeLazy struct {
	done atomic.Pointer[chan struct{}]
}

// Close closes the done channel. You shouldn't close the channel twice.
func (l *SafeLazy) Close() {
	if done := l.done.Swap(&closedChan); done != nil && *done != closedChan {
		close(*done)
	}
}

// Done returns the done channel.
func (l *SafeLazy) Done() <-chan struct{} {
	if done := l.done.Load(); done != nil {
		return *done
	}

	if ch := make(chan struct{}); l.done.CompareAndSwap(nil, &ch) {
		return ch
	}

	return *l.done.Load()
}

// Closed returns true if the done channel is closed.
func (l *SafeLazy) Closed() bool {
	done := l.done.Load()
	if done == nil {
		return false
	}

	select {
	case <-*done:
		return true
	default:
		return false
	}
}

func (l *SafeLazy) String() string {
	if l.Closed() {
		return "done"
	}

	return "pending"
}
