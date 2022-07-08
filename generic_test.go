package pt_test

import (
	"testing"
	"time"

	"github.com/gochore/pt"
)

func TestP(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "int",
			args: args{
				v: 1,
			},
		},
		{
			name: "float64",
			args: args{
				v: 1.1,
			},
		},
		{
			name: "time",
			args: args{
				v: time.Now(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pt.P(tt.args.v); *got != tt.args.v {
				t.Errorf("P() = %v, want %v", got, tt.args.v)
			}
		})
	}
}
