package _088_merge_sorted_array

import "testing"

func TestMerge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name string
		args args
	}{
		{"test01", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{4, 5, 6}, 3}},
		{"test02", args{[]int{4, 5, 6, 0, 0, 0}, 3, []int{1, 2, 3}, 3}},
	}
	for _, tt := range tests {
		merge(tt.args.nums1, tt.args.n, tt.args.nums2, tt.args.m)
		t.Logf("%#v", tt.args.nums1)
	}
}
