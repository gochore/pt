# pt

[![Go Action](https://github.com/gochore/pt/workflows/Go/badge.svg)](https://github.com/gochore/pt/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochore/pt)](https://goreportcard.com/report/github.com/gochore/pt)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gochore/pt)
![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/gochore/pt)

Return pointer of basic type value.

```go
package main

import "github.com/gochore/pt"

func main() {
	// wrong
	f(&100) // can not compile

	// bad
	v := 100
	f(&v)

	// good
	f(pt.Int(100))

}

func f(*int) {

}
```
