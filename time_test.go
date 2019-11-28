package pt_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gochore/pt"
)

func TestTime(t *testing.T) {
	now := time.Now()

	type args struct {
		v time.Time
	}
	tests := []struct {
		args args
		want time.Time
	}{
		{
			args: args{
				v: time.Time{},
			},
			want: time.Time{},
		},
		{
			args: args{
				v: now,
			},
			want: now,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Time(tt.args.v); !got.Equal(tt.want) {
				t.Errorf("Time() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDuration(t *testing.T) {
	type args struct {
		v time.Duration
	}
	tests := []struct {
		args args
		want time.Duration
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: time.Hour,
			},
			want: time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Duration(tt.args.v); *got != tt.want {
				t.Errorf("Duration() = %v, want %v", got, tt.want)
			}
		})
	}
}
