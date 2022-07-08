//go:build go1.18

package pt

// P return pointer of v
func P[V any](v V) *V {
	return &v
}
