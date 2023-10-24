//go:build go1.18

package pt_test

import (
	"testing"
	"time"

	"github.com/gochore/pt"
)

func TestP(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if got := pt.P(1); *got != 1 {
			t.Errorf("P() = %v, want %v", *got, 1)
		}
	})
	t.Run("float64", func(t *testing.T) {
		if got := pt.P(1.1); *got != 1.1 {
			t.Errorf("P() = %v, want %v", *got, 1)
		}
	})
	t.Run("time", func(t *testing.T) {
		now := time.Now()
		if got := pt.P(now); *got != now {
			t.Errorf("P() = %v, want %v", *got, 1)
		}
	})
}

func TestV(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		if got := pt.V(pt.P(1)); got != 1 {
			t.Errorf("V() = %v, want %v", got, 1)
		}
	})
	t.Run("float64", func(t *testing.T) {
		if got := pt.V(pt.P(1.1)); got != 1.1 {
			t.Errorf("V() = %v, want %v", got, 1)
		}
	})
	t.Run("time", func(t *testing.T) {
		now := time.Now()
		if got := pt.V(pt.P(now)); got != now {
			t.Errorf("V() = %v, want %v", got, 1)
		}
	})
	t.Run("nil int", func(t *testing.T) {
		var i *int
		if got := pt.V(i); got != 0 {
			t.Errorf("V() = %v, want %v", got, 0)
		}
	})
	t.Run("nil string", func(t *testing.T) {
		var s *string
		if got := pt.V(s); got != "" {
			t.Errorf("V() = %v, want %v", got, "")
		}
	})
}
