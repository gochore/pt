//go:build go1.18

package pt

// P return pointer of v
func P[T any](v T) *T {
	return &v
}

// V return value of p
func V[T any](p *T) T {
	if p == nil {
		return *new(T)
	}
	return *p
}
