# pt

[![Build Status](https://travis-ci.com/gochore/pt.svg?branch=master)](https://travis-ci.com/gochore/pt)
[![codecov](https://codecov.io/gh/gochore/pt/branch/master/graph/badge.svg)](https://codecov.io/gh/gochore/pt)
[![Go Report Card](https://goreportcard.com/badge/github.com/gochore/pt)](https://goreportcard.com/report/github.com/gochore/pt)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/gochore/pt)](https://github.com/gochore/pt/blob/master/go.mod)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/gochore/pt)](https://github.com/gochore/pt/releases)

Return pointer of basic type value.

## Example

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

## Document

#### func  Bool

```go
func Bool(v bool) *bool
```
Bool return pointer of bool

#### func  Byte

```go
func Byte(v byte) *byte
```
Byte return pointer of byte

#### func  Complex128

```go
func Complex128(v complex128) *complex128
```
Complex128 return pointer of complex128

#### func  Complex64

```go
func Complex64(v complex64) *complex64
```
Complex64 return pointer of complex64

#### func  Duration

```go
func Duration(v time.Duration) *time.Duration
```
Duration return pointer of time.Duration

#### func  Float32

```go
func Float32(v float32) *float32
```
Float32 return pointer of float32

#### func  Float64

```go
func Float64(v float64) *float64
```
Float64 return pointer of float64

#### func  Int

```go
func Int(v int) *int
```
Int return pointer of int

#### func  Int16

```go
func Int16(v int16) *int16
```
Int16 return pointer of int16

#### func  Int32

```go
func Int32(v int32) *int32
```
Int32 return pointer of int32

#### func  Int64

```go
func Int64(v int64) *int64
```
Int64 return pointer of int64

#### func  Int8

```go
func Int8(v int8) *int8
```
Int8 return pointer of int8

#### func  Rune

```go
func Rune(v rune) *rune
```
Rune return pointer of rune

#### func  String

```go
func String(v string) *string
```
String return pointer of string

#### func  Time

```go
func Time(v time.Time) *time.Time
```
Time return pointer of time.Time

#### func  Uint

```go
func Uint(v uint) *uint
```
Uint return pointer of uint

#### func  Uint16

```go
func Uint16(v uint16) *uint16
```
Uint16 return pointer of uint16

#### func  Uint32

```go
func Uint32(v uint32) *uint32
```
Uint32 return pointer of uint32

#### func  Uint64

```go
func Uint64(v uint64) *uint64
```
Uint64 return pointer of uint64

#### func  Uint8

```go
func Uint8(v uint8) *uint8
```
Uint8 return pointer of uint8

#### func  Uintptr

```go
func Uintptr(v uintptr) *uintptr
```
Uintptr return pointer of uintptr
