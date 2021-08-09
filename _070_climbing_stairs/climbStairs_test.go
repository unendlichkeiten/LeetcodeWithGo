package _070_climbing_stairs

import "testing"

func TestClimbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{1}, 1},
		{"test-02", args{2}, 2},
		{"test-03", args{3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairsv2(tt.args.n); got != tt.want {
				t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
