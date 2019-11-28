# pt

![](https://github.com/gochore/pt/workflows/Go/badge.svg)

return pointer of basic type value.

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
