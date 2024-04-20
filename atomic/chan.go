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

package atomic

import (
	"sync/atomic"
	"unsafe"
)

// A Chan is an atomic channel of type chan T. The zero value is a nil channel.
type Chan[T any] struct {
	// Mention *T in a field to disallow conversion between Chan types.
	// See go.dev/issue/56603 for more details.
	// Use *T, not T, to avoid spurious recursive type definition errors.
	_ [0]*T
	_ noCopy
	v unsafe.Pointer
}

// Load atomically loads and returns the value stored in c.
func (c *Chan[T]) Load() chan T { return ptr2Ch[T](atomic.LoadPointer(&c.v)) }

// Store atomically stores ch into c.
func (c *Chan[T]) Store(ch chan T) { atomic.StorePointer(&c.v, ch2Ptr(ch)) }

// Swap atomically stores new into c and returns the previous value.
func (c *Chan[T]) Swap(new chan T) (old chan T) {
	return ptr2Ch[T](atomic.SwapPointer(&c.v, ch2Ptr(new)))
}

// CompareAndSwap executes the compare-and-swap operation for c.
func (c *Chan[T]) CompareAndSwap(old, new chan T) (swapped bool) {
	return atomic.CompareAndSwapPointer(&c.v, ch2Ptr(old), ch2Ptr(new))
}

// ch2Ptr casts from a channel to a pointer.
func ch2Ptr[T any](ch chan T) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&ch))
}

// ptr2Ch casts from a pointer to a channel.
func ptr2Ch[T any](ptr unsafe.Pointer) chan T {
	return *(*chan T)(unsafe.Pointer(&ptr))
}
