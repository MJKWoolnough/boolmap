# boolmap

[![CI](https://github.com/MJKWoolnough/boolmap/actions/workflows/go-checks.yml/badge.svg)](https://github.com/MJKWoolnough/boolmap/actions)
[![Go Reference](https://pkg.go.dev/badge/vimagination.zapto.org/boolmap.svg)](https://pkg.go.dev/vimagination.zapto.org/boolmap)
[![Go Report Card](https://goreportcard.com/badge/vimagination.zapto.org/boolmap)](https://goreportcard.com/report/vimagination.zapto.org/boolmap)

--
    import "vimagination.zapto.org/boolmap"

Package boolmap provides data structures for storing bits, crumbs, and nibbles.

## Highlights

 - Types to store and retrieve bit (1-bit, 0->1), crumb (2-bits, 0->3), and nibble (4-bits, 0->15) values.
 - Map types for sparse data.
 - Slice type for compact data.

## Usage

```go
package main

import (
	"fmt"

	"vimagination.zapto.org/boolmap"
)

func main() {
	m := boolmap.NewMap()

	m.SetBool(42, true)
	m.SetBool(100, false)

	fmt.Println(m.GetBool(42))
	fmt.Println(m.GetBool(100))
	fmt.Println(m.GetBool(101))

	// Output:
	// true
	// false
	// false
}
```

## Documentation

Full API docs can be found at:

https://pkg.go.dev/vimagination.zapto.org/boolmap
