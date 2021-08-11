package main

import (
	"testing"
	"time"
)

func longAdd(a, b int) int {
	time.Sleep(365 * 24 * time.Hour)
	return a + b
}

func TestSkip(t *testing.T) {
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
			name: "fail",
			args: args{a: 1, b: 2},
			want: 30,
		},
		{
			name: "normal",
			args: args{a: 1, b: 2},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if testing.Short() {
				t.Skip("skipping test in short mode")
			}

			if got := longAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Fatalf("add() = %v, want %v", got, tt.want)
			}

		})
	}
}
