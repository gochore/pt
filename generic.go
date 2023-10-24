//go:build go1.18

package pt

// P returns pointer of v.
// It's a short form of "Pointer" or "GetPointer".
func P[T any](v T) *T {
	return &v
}

// V returns value of p. If p is nil, return zero value of T.
// It's a short form of "Value" or "GetValue".
func V[T any](p *T) T {
	if p == nil {
		return *new(T)
	}
	return *p
}
