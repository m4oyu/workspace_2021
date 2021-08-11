package main

import (
	"testing"
	"time"
)

func parallelAdd(a, b int) int {
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}

func TestParallel(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"normal_1", args{1, 2}, 3},
		{"normal_1", args{2, 3}, 5},
		{"normal_1", args{3, 4}, 7},
	}

	for _, tt := range tests {
		// 変数のシャドウイング
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// testの並列化
			t.Parallel()
			if got := parallelAdd(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("parallelAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
