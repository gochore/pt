package pt

import "time"

func Time(v time.Time) *time.Time {
	return &v
}

func Duration(v time.Duration) *time.Duration {
	return &v
}
