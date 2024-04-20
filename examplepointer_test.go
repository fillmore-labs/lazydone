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
	"fmt"
	"time"

	"fillmore-labs.com/lazydone"
)

type SafeResult struct {
	lazydone.SafeLazy
	value int
}

func ExampleSafeLazy() {
	var result SafeResult

	go func() {
		time.Sleep(100 * time.Millisecond)
		result.value = 42
		result.Close() // The result is ready.
	}()

	fmt.Println("SafeResult:", &result.SafeLazy)

	select {
	case <-result.Done():
		fmt.Println("Already done")
	default:
		fmt.Println("Still processing...")
	}

	<-result.Done() // Wait for the result.
	fmt.Println("SafeResult:", &result.SafeLazy)
	fmt.Println("Value:", result.value)

	// Output:
	// SafeResult: pending
	// Still processing...
	// SafeResult: done
	// Value: 42
}

func ExampleSafeLazy_Done() {
	var result SafeResult

	go func() {
		time.Sleep(100 * time.Millisecond)
		result.value = 42
		result.Close() // The result is ready.
	}()

	fmt.Println("SafeResult:", &result.SafeLazy)

	<-result.Done() // Wait for the result.
	fmt.Println("SafeResult:", &result.SafeLazy)
	fmt.Println("Value:", result.value)

	// Output:
	// SafeResult: pending
	// SafeResult: done
	// Value: 42
}
