package goutils

import "testing"

func TestMin(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a less then b",
			args: args{
				a: 1,
				b: 2,
			},
			want: 1,
		},
		{
			name: "a greater then b",
			args: args{
				a: 2,
				b: 1,
			},
			want: 1,
		},
		{
			name: "a equals to b",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a less then b",
			args: args{
				a: 1,
				b: 2,
			},
			want: 2,
		},
		{
			name: "a greater then b",
			args: args{
				a: 2,
				b: 1,
			},
			want: 2,
		},
		{
			name: "a equals to b",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxUint64(t *testing.T) {
	type args struct {
		a uint64
		b uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "a less then b",
			args: args{
				a: 1,
				b: 2,
			},
			want: 2,
		},
		{
			name: "a greater then b",
			args: args{
				a: 2,
				b: 1,
			},
			want: 2,
		},
		{
			name: "a equals to b",
			args: args{
				a: 1,
				b: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxUint64(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}
