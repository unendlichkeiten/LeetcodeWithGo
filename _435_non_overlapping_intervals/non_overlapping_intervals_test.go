package _435_non_overlapping_intervals

import "testing"

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"test01", args{[][]int{{1, 2}, {2, 4}, {1, 3}}}, 1},
		{"test02", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

// [[1,2],[2,3],[3,4],[1,3]]
