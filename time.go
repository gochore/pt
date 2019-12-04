package pt

import "time"

// Time return pointer of time.Time
func Time(v time.Time) *time.Time {
	return &v
}

// Duration return pointer of time.Duration
func Duration(v time.Duration) *time.Duration {
	return &v
}
