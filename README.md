# pt

[![Go Reference](https://pkg.go.dev/badge/github.com/gochore/pt.svg)](https://pkg.go.dev/github.com/gochore/pt)
[![Actions](https://github.com/gochore/pt/actions/workflows/test.yaml/badge.svg)](https://github.com/gochore/pt/actions)
[![Codecov](https://codecov.io/gh/gochore/pt/branch/master/graph/badge.svg)](https://codecov.io/gh/gochore/pt)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochore/pt)](https://goreportcard.com/report/github.com/gochore/pt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gochore/pt)](https://github.com/gochore/pt/blob/master/go.mod)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/gochore/pt)](https://github.com/gochore/pt/releases)

Return pointer of basic type value.

## Example

```go
package main

import "github.com/gochore/pt"

func main() {
	// 💀 It cannot work because Go does not allow taking the address of a constant or literal.
	f(&100)

	// ☹️ It works, but it requires two lines and declares a variable that could pollute the namespace.
	v := 100
	f(&v)

	// 😊 It works. Only one line and no new variables.
	// But you have to use different functions for different types.
	// It's the only way to do it before Go1.18.
	f(pt.Int(100))

	// 🤩 It works. Only one line and no new variables, and a single function for all types.
	// It's based on generics, so it requires Go1.18 and above.
	f(pt.P(100))
}

func f(p *int) {
	// 💀 It could panic if p is nil.
	println(*p)

	// ☹️ It's safe, but it requires multiple lines and declares a variable that could pollute the namespace.
	v := 0
	if p != nil {
		v = *p
	}
	println(v)

	// 🤩 It's safe. Only one line and no new variables.
	// It's based on generics, so it requires Go1.18 and above.
	println(pt.V(p))
}
```

## Document

### Go1.18 and later

#### func P

```go
func P[T any](v T) *T
```
P returns pointer of v.
It's a short form of "Pointer" or "GetPointer".

#### func V

```go
func V[T any](p *T) T
```
V returns value of p. If p is nil, return zero value of T.
It's a short form of "Value" or "GetValue".

### Before Go1.18 (deprecated)

#### func Bool

```go
func Bool(v bool) *bool
```
Bool returns pointer of bool

#### func Byte

```go
func Byte(v byte) *byte
```
Byte returns pointer of byte

#### func Complex128

```go
func Complex128(v complex128) *complex128
```
Complex128 returns pointer of complex128

#### func Complex64

```go
func Complex64(v complex64) *complex64
```
Complex64 returns pointer of complex64

#### func Duration

```go
func Duration(v time.Duration) *time.Duration
```
Duration returns pointer of time.Duration

#### func Float32

```go
func Float32(v float32) *float32
```
Float32 returns pointer of float32

#### func Float64

```go
func Float64(v float64) *float64
```
Float64 returns pointer of float64

#### func Int

```go
func Int(v int) *int
```
Int returns pointer of int

#### func Int16

```go
func Int16(v int16) *int16
```
Int16 returns pointer of int16

#### func Int32

```go
func Int32(v int32) *int32
```
Int32 returns pointer of int32

#### func Int64

```go
func Int64(v int64) *int64
```
Int64 returns pointer of int64

#### func Int8

```go
func Int8(v int8) *int8
```
Int8 returns pointer of int8

#### func Rune

```go
func Rune(v rune) *rune
```
Rune returns pointer of rune

#### func String

```go
func String(v string) *string
```
String returns pointer of string

#### func Time

```go
func Time(v time.Time) *time.Time
```
Time returns pointer of time.Time

#### func Uint

```go
func Uint(v uint) *uint
```
Uint returns pointer of uint

#### func Uint16

```go
func Uint16(v uint16) *uint16
```
Uint16 returns pointer of uint16

#### func Uint32

```go
func Uint32(v uint32) *uint32
```
Uint32 returns pointer of uint32

#### func Uint64

```go
func Uint64(v uint64) *uint64
```
Uint64 returns pointer of uint64

#### func Uint8

```go
func Uint8(v uint8) *uint8
```
Uint8 returns pointer of uint8

#### func Uintptr

```go
func Uintptr(v uintptr) *uintptr
```
Uintptr returns pointer of uintptr
