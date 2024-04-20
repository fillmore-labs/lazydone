# Fillmore Labs LazyDone

[![Go Reference](https://pkg.go.dev/badge/fillmore-labs.com/lazydone.svg)](https://pkg.go.dev/fillmore-labs.com/lazydone)
[![Test](https://github.com/fillmore-labs/lazydone/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/fillmore-labs/lazydone/actions/workflows/test.yml)
[![Coverage](https://codecov.io/gh/fillmore-labs/lazydone/branch/main/graph/badge.svg?token=OWP7JMXHD2)](https://codecov.io/gh/fillmore-labs/lazydone)
[![Maintainability](https://api.codeclimate.com/v1/badges/8e2d0c51b631c9c39a1f/maintainability)](https://codeclimate.com/github/fillmore-labs/lazydone/maintainability)
[![Go Report Card](https://goreportcard.com/badge/fillmore-labs.com/lazydone)](https://goreportcard.com/report/fillmore-labs.com/lazydone)
[![License](https://img.shields.io/github/license/fillmore-labs/lazydone)](https://www.apache.org/licenses/LICENSE-2.0)

The `lazydone` package provides a lazy initialized done channel with a valid zero value in Go.

## Usage

```go
package main

import (
	"fmt"
	"time"

	"fillmore-labs.com/lazydone"
)

type Result struct {
	lazydone.Lazy
	value int
}

func main() {
	var result Result

	go func() {
		time.Sleep(100 * time.Millisecond)
		result.value = 42
		result.Close() // The result is ready.
	}()

	<-result.Done() // Wait for the result.
	fmt.Println("Value:", result.value)
}
```

The channel can be used in [`select` statements](https://go.dev/ref/spec#Select_statements), which is not possible with
[Mutexes](https://pkg.go.dev/sync#Mutex) or [WaitGroups](https://pkg.go.dev/sync#WaitGroup).
