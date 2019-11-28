package pt_test

import (
	"fmt"
	"testing"

	"github.com/gochore/pt"
)

func TestBool(t *testing.T) {
	type args struct {
		v bool
	}
	tests := []struct {
		args args
		want bool
	}{
		{
			args: args{
				v: false,
			},
			want: false,
		},
		{
			args: args{
				v: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Bool(tt.args.v); *got != tt.want {
				t.Errorf("Bool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint8(t *testing.T) {
	type args struct {
		v uint8
	}
	tests := []struct {
		args args
		want uint8
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uint8(tt.args.v); *got != tt.want {
				t.Errorf("Uint8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint16(t *testing.T) {
	type args struct {
		v uint16
	}
	tests := []struct {
		args args
		want uint16
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uint16(tt.args.v); *got != tt.want {
				t.Errorf("Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	type args struct {
		v uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uint32(tt.args.v); *got != tt.want {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint64(t *testing.T) {
	type args struct {
		v uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uint64(tt.args.v); *got != tt.want {
				t.Errorf("Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8(t *testing.T) {
	type args struct {
		v int8
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Int8(tt.args.v); *got != tt.want {
				t.Errorf("Int8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16(t *testing.T) {
	type args struct {
		v int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Int16(tt.args.v); *got != tt.want {
				t.Errorf("Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	type args struct {
		v int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Int32(tt.args.v); *got != tt.want {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Int64(tt.args.v); *got != tt.want {
				t.Errorf("Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	type args struct {
		v float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Float32(tt.args.v); *got != tt.want {
				t.Errorf("Float32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat64(t *testing.T) {
	type args struct {
		v float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Float64(tt.args.v); *got != tt.want {
				t.Errorf("Float64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex64(t *testing.T) {
	type args struct {
		v complex64
	}
	tests := []struct {
		name string
		args args
		want complex64
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Complex64(tt.args.v); *got != tt.want {
				t.Errorf("Complex64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComplex128(t *testing.T) {
	type args struct {
		v complex128
	}
	tests := []struct {
		name string
		args args
		want complex128
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Complex128(tt.args.v); *got != tt.want {
				t.Errorf("Complex128() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestString(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				v: "",
			},
			want: "",
		},
		{
			args: args{
				v: "s",
			},
			want: "s",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.String(tt.args.v); *got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt(t *testing.T) {
	type args struct {
		v int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Int(tt.args.v); *got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	type args struct {
		v uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uint(tt.args.v); *got != tt.want {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUintptr(t *testing.T) {
	type args struct {
		v uintptr
	}
	tests := []struct {
		name string
		args args
		want uintptr
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Uintptr(tt.args.v); *got != tt.want {
				t.Errorf("Uintptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte(t *testing.T) {
	type args struct {
		v byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Byte(tt.args.v); *got != tt.want {
				t.Errorf("Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRune(t *testing.T) {
	type args struct {
		v rune
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			args: args{
				v: 0,
			},
			want: 0,
		},
		{
			args: args{
				v: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.want), func(t *testing.T) {
			if got := pt.Rune(tt.args.v); *got != tt.want {
				t.Errorf("Rune() = %v, want %v", got, tt.want)
			}
		})
	}
}
