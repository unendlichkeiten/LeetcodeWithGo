package _075_sort_colors

import "testing"

func Test_sortColors(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01", args{[]int{2, 0, 2, 1, 1, 0}}},
		{"test01", args{[]int{2, 0, 1}}},
		{"test01", args{[]int{0}}},
	}
	for _, tt := range tests {
		sortColors(tt.args.nums)
		t.Logf("%v", tt.args.nums)
	}
}
