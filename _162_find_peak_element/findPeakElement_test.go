package _162_find_peak_element

import "testing"

func TestFindPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test-01", args{[]int{1, 2, 1, 3, 5, 6, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("FindPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
