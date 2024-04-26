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
func (c *Chan[T]) Load() (ch chan T) {
	*(*unsafe.Pointer)(unsafe.Pointer(&ch)) = atomic.LoadPointer(&c.v)

	return
}

// Store atomically stores ch into c.
func (c *Chan[T]) Store(ch chan T) {
	atomic.StorePointer(&c.v, *(*unsafe.Pointer)(unsafe.Pointer(&ch)))
}

// Swap atomically stores new into c and returns the previous value.
func (c *Chan[T]) Swap(new chan T) (old chan T) {
	*(*unsafe.Pointer)(unsafe.Pointer(&old)) = atomic.SwapPointer(&c.v, *(*unsafe.Pointer)(unsafe.Pointer(&new)))

	return
}

// CompareAndSwap executes the compare-and-swap operation for c.
func (c *Chan[T]) CompareAndSwap(old, new chan T) (swapped bool) {
	return atomic.CompareAndSwapPointer(&c.v,
		*(*unsafe.Pointer)(unsafe.Pointer(&old)), *(*unsafe.Pointer)(unsafe.Pointer(&new)))
}
