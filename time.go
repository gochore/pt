package pt

import "time"

// Deprecated: use P with generics in go1.18+
// Time return pointer of time.Time
func Time(v time.Time) *time.Time {
	return &v
}

// Deprecated: use P with generics in go1.18+
// Duration return pointer of time.Duration
func Duration(v time.Duration) *time.Duration {
	return &v
}
